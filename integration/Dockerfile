FROM golang:1.15-rc
WORKDIR /app

# RUN apt-get update \
#     && apt-get install -y tee

COPY go.mod .
RUN go mod download -x \
    && go mod verify

COPY . .
CMD [ "bash" ]