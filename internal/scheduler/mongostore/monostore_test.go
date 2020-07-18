package mongostore_test

import (
	"os"
	"testing"

	"github.com/FrancescoIlario/grower/internal/scheduler"
	"github.com/FrancescoIlario/grower/internal/scheduler/mongostore"
	"github.com/google/uuid"
)

// TODO: current integration tests are not implement
// in the best way, but it's the easier and agilest one
// for the moment I could came up! :)

var (
	mongoConnStr string
	mongoDb      string
	mongoPswd    string
	mongoUser    string
)

func init() {
	mongoConnStr = os.Getenv("MONGO_CONNSTR")
	mongoDb = os.Getenv("MONGO_DATABASE")
	mongoPswd = os.Getenv("MONGO_PASSWORD")
	mongoUser = os.Getenv("MONGO_USERNAME")
}

func baseMongoConfiguration(collection string) *mongostore.MongoConfiguration {
	return &mongostore.MongoConfiguration{
		ConnectionString: mongoConnStr,
		Database:         mongoDb,
		Password:         mongoPswd,
		Username:         mongoUser,
		Collection:       collection,
	}
}

func generateNewStoreOrFatal(t *testing.T) scheduler.PairStore {
	collection := generateCollectionNameOrFatal(t)
	conf := baseMongoConfiguration(collection)
	store, err := mongostore.NewMongoRepo(*conf)
	if err != nil {
		t.Fatalf("error creating store: %v", err)
	}
	return store
}

func generateCollectionNameOrFatal(t *testing.T) string {
	collection, err := uuid.NewUUID()
	if err != nil {
		t.Fatalf("error creating uuid for collection name: %v", err)
		return ""
	}
	return collection.String()
}
