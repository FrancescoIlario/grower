valve: proto/valve.proto
	mkdir -p pkg/valvepb
	protoc \
		-I proto \
		--go_out=plugins=grpc:pkg/valvepb \
		--go_opt=paths=source_relative \
		valve.proto

.PHONY: pinctl
pinctl:
	@mkdir -p bin/pinctl/
	env GOOS=linux GOARCH=arm GOARM=6 go build -o bin/pinctl/pinctl cmd/pinctl/main.go

.PHONY: valvectl
valvectl:
	@mkdir -p bin/valvectl
	env GOOS=linux GOARCH=arm GOARM=6 go build -o bin/valvectl/valvectl cmd/valvectl/main.go