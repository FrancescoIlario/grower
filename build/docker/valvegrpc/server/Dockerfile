FROM gcr.io/distroless/base-debian10 as prod
CMD ["/app"]

FROM golang:1.15-rc as build
WORKDIR /app
RUN mkdir -p bin 

COPY go.mod .
RUN go mod download \
    && go mod verify

COPY cmd/valvegrpc/server cmd/valvegrpc/server
COPY internal internal
COPY pkg/valvepb pkg/valvepb

RUN go build -o bin/valvegrpc/server cmd/valvegrpc/server/main.go

FROM prod 
COPY --from=build /app/bin/valvegrpc/server /app