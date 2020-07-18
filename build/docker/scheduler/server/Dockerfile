FROM gcr.io/distroless/base-debian10 as prod
CMD ["/app"]

FROM golang:1.15-rc as build
WORKDIR /app
RUN mkdir -p bin 

COPY go.mod .
RUN go mod download \
    && go mod verify

COPY cmd/scheduler/server cmd/scheduler/server
COPY internal internal
COPY pkg/valvepb pkg/valvepb
COPY pkg/schedulerpb pkg/schedulerpb

RUN go build -o bin/scheduler/server cmd/scheduler/server/main.go

FROM prod 
COPY --from=build /app/bin/scheduler/server /app