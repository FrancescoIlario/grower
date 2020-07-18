FROM gcr.io/distroless/base-debian10 as prod
CMD ["/app"]

FROM golang:1.15-rc as build
WORKDIR /app
RUN mkdir -p bin 

COPY go.mod .
RUN go mod download \
    && go mod verify

COPY cmd/scheduler/client cmd/scheduler/client
COPY internal internal
COPY pkg/schedulerpb pkg/schedulerpb

RUN go build -o bin/scheduler/client cmd/scheduler/client/main.go

FROM prod 
COPY --from=build /app/bin/scheduler/client /app