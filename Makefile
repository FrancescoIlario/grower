valve: proto/valve.proto
	mkdir -p pkg/valvepb
	protoc \
		-I proto \
		--go_out=plugins=grpc:pkg/valvepb \
		--go_opt=paths=source_relative \
		valve.proto

.PHONY: pinctl
pinctl:
	@mkdir -p bin/
	env GOOS=linux GOARCH=arm GOARM=6 go build -o bin/pinctl cmd/pinctl/main.go

.PHONY: valvectl
valvectl:
	@mkdir -p bin/
	env GOOS=linux GOARCH=arm GOARM=6 go build -o bin/valvectl cmd/valvectl/main.go

.PHONY: valvecmdrctl
valvecmdrctl:
	@mkdir -p bin/
	env GOOS=linux GOARCH=arm GOARM=6 go build -o bin/valvecmdctl cmd/valvecmdr/client/main.go

.PHONY: valvecmdrsvr
valvecmdrsvr:
	@mkdir -p bin/
	env GOOS=linux GOARCH=arm GOARM=6 go build -o bin/valvecmdsvr cmd/valvecmdr/server/main.go

.PHONY: all
all: pinctl valvectl valvecmdrctl valvecmdrsvr