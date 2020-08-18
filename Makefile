.PHONY: protos valvepb schedulerpb
protos: valvepb schedulerpb shutterpb valveespb valvecqrspb valvesharedpb

valvepb: proto/valve.grpc.proto valvesharedpb
	mkdir -p pkg/valvepb/grpc
	protoc \
		-I proto \
		--go_out=plugins=grpc:pkg/valvepb/grpc \
		--go_opt=paths=source_relative \
		valve.grpc.proto

valveespb: proto/valve.es.proto valvesharedpb
	mkdir -p pkg/valvepb/es
	protoc \
		-I proto \
		--gofast_out=plugins=grpc:pkg/valvepb/es \
		--gofast_opt=paths=source_relative \
		valve.es.proto

valvecqrspb: proto/valve.cqrs.proto valvesharedpb
	mkdir -p pkg/valvepb/cqrs
	protoc \
		-I proto \
		--gofast_out=plugins=grpc:pkg/valvepb/cqrs  \
		--gofast_opt=paths=source_relative \
		valve.cqrs.proto

valvesharedpb: proto/valve.shared.proto
	mkdir -p pkg/valvepb/shared
	protoc \
		-I proto \
		--go_out=plugins=grpc:pkg/valvepb/shared \
		--go_opt=paths=source_relative \
		valve.shared.proto

schedulerpb: proto/scheduler.proto
	mkdir -p pkg/schedulerpb

	protoc  \
		--proto_path=${GOPATH}/src \
		--proto_path=${GOPATH}/src/github.com/google/protobuf/src \
		--proto_path=proto \
		--go_out=plugins=grpc:pkg/schedulerpb \
		--govalidators_out=pkg/schedulerpb \
		--govalidators_opt=paths=source_relative \
		--go_opt=paths=source_relative \
		scheduler.proto

shutterpb: proto/shutter.proto
	mkdir -p pkg/shutterpb

	protoc  \
		--proto_path=${GOPATH}/src \
		--proto_path=${GOPATH}/src/github.com/google/protobuf/src \
		--proto_path=proto \
		--go_out=plugins=grpc:pkg/shutterpb \
		--govalidators_out=pkg/shutterpb \
		--govalidators_opt=paths=source_relative \
		--go_opt=paths=source_relative \
		shutter.proto

.PHONY: all init
all: pinctl valvectl valvecmdrctl valvecmdrsvr scheduler
init:
	@chmod +x init.sh
	bash init.sh

bindir: 
	@mkdir -p bin/

.PHONY: pinctl
pinctl: bindir
	env GOOS=linux GOARCH=arm GOARM=6 go build -o bin/pinctl cmd/pinctl/main.go

.PHONY: valvectl
valvectl: bindir
	env GOOS=linux GOARCH=arm GOARM=6 go build -o bin/valvectl cmd/valvectl/main.go

.PHONY: shutter
shutter: bindir
	env GOOS=linux GOARCH=arm GOARM=6 go build -o bin/shutter cmd/shutter/main.go

.PHONY: valvecmdrsvr valvecmdrctl
valvecmdrsvr: bindir
	env GOOS=linux GOARCH=arm GOARM=6 go build -o bin/valvecmdsvr cmd/valvegrpc/server/main.go

valvecmdrctl: bindir
	env GOOS=linux GOARCH=arm GOARM=6 go build -o bin/valvecmdctl cmd/valvegrpc/client/main.go

.PHONY: schedulersvr schedulerctl
schedulersvr: bindir
	env GOOS=linux GOARCH=arm GOARM=6 go build -o bin/schedulersvr cmd/scheduler/server/main.go

schedulerctl: bindir
	env GOOS=linux GOARCH=arm GOARM=6 go build -o bin/schedulerctl cmd/scheduler/client/main.go

.PHONY: utest scheduler-itest
utest:
	go test ./... -short

scheduler-itest:
	eval integration/scheduler/mongostore/runtest.sh