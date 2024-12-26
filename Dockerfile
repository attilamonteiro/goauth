FROM golang:1.23

# Enable CGO
ENV CGO_ENABLED=1

# Install GCC
RUN apt-get update && apt-get install -y gcc

WORKDIR /goauth

COPY . .

RUN go get -d -v ./...

RUN go build -o goauth .

EXPOSE 8080

CMD ["./goauth"]