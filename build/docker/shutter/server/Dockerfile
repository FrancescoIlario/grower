FROM gcr.io/distroless/base-debian10 as prod
CMD ["/app"]

FROM golang:1.15-rc as build
WORKDIR /app
RUN mkdir -p bin 

COPY go.mod .
RUN go mod download \
    && go mod verify

COPY cmd/shutter/server cmd/shutter/server
COPY internal internal
COPY pkg/valvepb pkg/valvepb
COPY pkg/shutterpb pkg/shutterpb

RUN go build -o bin/shutter/server cmd/shutter/server/main.go

FROM prod 
COPY --from=build /app/bin/shutter/server /app