package mongostore

import (
	"context"

	"github.com/FrancescoIlario/grower/internal/scheduler"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConfiguration ...
type MongoConfiguration struct {
	ConnectionString string
	Collection       string
	Database         string
	Username         string
	Password         string
}

type mongoStore struct {
	Client        *mongo.Client
	Collection    *mongo.Collection
	Configuration MongoConfiguration
}

//NewMongoRepo ...
func NewMongoRepo(conf MongoConfiguration) (scheduler.PairStore, error) {
	// Set client options
	clientOptions := options.Client().
		ApplyURI(conf.ConnectionString)

	if conf.Username != "" {
		clientOptions = clientOptions.
			SetAuth(options.Credential{
				Username: conf.Username,
				Password: conf.Password,
			})
	}

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	collection := client.
		Database(conf.Database).
		Collection(conf.Collection)

	// build database
	stconf := conf
	stconf.Password = ""

	db := mongoStore{
		Client:        client,
		Collection:    collection,
		Configuration: stconf,
	}
	return &db, nil
}

func (s *mongoStore) Store(ctx context.Context, pair scheduler.Pair) (*uuid.UUID, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	pair.ID = id.String()

	if _, err := s.Collection.InsertOne(ctx, pair); err != nil {
		return nil, err
	}

	return &id, nil
}

func (s *mongoStore) Read(ctx context.Context, id uuid.UUID) (*scheduler.Pair, error) {
	filter := bson.M{"_id": id.String()}

	sr := s.Collection.FindOne(ctx, filter)
	if err := sr.Err(); err != nil {
		return nil, err
	}

	var res scheduler.Pair
	if err := sr.Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *mongoStore) List(ctx context.Context) ([]scheduler.Pair, error) {
	c, err := s.Collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer c.Close(ctx)

	var res []scheduler.Pair
	if err := c.All(ctx, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *mongoStore) Delete(ctx context.Context, id uuid.UUID) error {
	filter := bson.M{"_id": id.String()}

	dr, err := s.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if dr.DeletedCount == 0 {
		return scheduler.ErrNotFound
	}

	return nil
}
