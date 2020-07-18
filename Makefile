.PHONY: protos valvepb schedulerpb
protos: valvepb schedulerpb

valvepb: proto/valve.proto
	mkdir -p pkg/valvepb
	protoc \
		-I proto \
		--go_out=plugins=grpc:pkg/valvepb \
		--go_opt=paths=source_relative \
		valve.proto

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

.PHONY: valvecmdrsvr valvecmdrctl
valvecmdrsvr: bindir
	env GOOS=linux GOARCH=arm GOARM=6 go build -o bin/valvecmdsvr cmd/valvecmdr/server/main.go

valvecmdrctl: bindir
	env GOOS=linux GOARCH=arm GOARM=6 go build -o bin/valvecmdctl cmd/valvecmdr/client/main.go

.PHONY: schedulersvr schedulerctl
schedulersvr: bindir
	env GOOS=linux GOARCH=arm GOARM=6 go build -o bin/schedulersvr cmd/scheduler/server/main.go

schedulerctl: bindir
	env GOOS=linux GOARCH=arm GOARM=6 go build -o bin/schedulerctl cmd/scheduler/client/main.go

.PHONY: utest scheduler-itest
utest:
	go run ./... -short

scheduler-itest:
	eval integration/scheduler/mongostore/runtest.sh