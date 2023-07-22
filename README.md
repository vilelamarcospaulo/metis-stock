## Metis - The Wise Stock Advisor

#### Introduction
Metis, named after the ancient Greek goddess of wisdom, is my personal evaluation software designed to provide insightful guidance on my investiment decisions. </br></br>
Taking inspiration from the goddess Metis' wise counsel, this tool employs a carefully curated set of rules to delve deep into the financial health and intrinsic value of stocks. 
By leveraging fundamental analysis techniques, Metis empowers us with the knowledge we need to make informed decisions.


### Proto

#### Go

Install:

```bash
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

Generate:
```bash
protoc --proto_path=./proto --go_out=./generated/go --go-grpc_out=./generated/go ./proto/stock_picker.proto
```

It should generate a few `*.pb.go` at `/generated/go/*`.


### Usage

Since we're establishing a connection, it's important to create our own network :)

```bash
docker network create grpc-network
```

#### Server

Build:

```bash
docker build . -t server --progress=plain
```

Run:

```bash
docker run -d --name server --network grpc-network -p 50051:50051 server
```

### Client

Build:

```bash
docker build . -t client --progress=plain
```

Run:

```bash
docker run --network grpc-network client
```
