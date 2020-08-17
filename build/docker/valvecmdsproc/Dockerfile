FROM gcr.io/distroless/base-debian10 as prod
CMD ["/app"]

FROM golang:1.15-rc as build
WORKDIR /app
RUN mkdir -p bin 

COPY go.mod .
RUN go mod download \
    && go mod verify

COPY cmd/valvecmdsproc cmd/valvecmdsproc
COPY internal internal
COPY pkg/valvepb pkg/valvepb

RUN go build -o bin/valvecmdsproc cmd/valvecmdsproc/main.go

FROM prod 
COPY --from=build /app/bin/valvecmdsproc /app