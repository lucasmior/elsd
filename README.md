# ELSd

Entity Locator Service

## Building

```
$ docker-compose build
```

### Modifying the code

Generating gRPC client and server interfaces.

```
$ protoc -I pkg/api/ pkg/api/els.proto --go_out=plugins=grpc:pkg/api
```

Updating dependencies.
```
$ dep ensure -update
```

To update a dependency to a new version, you might run

```
$ dep ensure github.com/pkg/errors@^0.8.0
```

## Running

```
$ docker-compose up
```

## Testing

You can build the client, add some routing keys and get them

```
$ go install  github.com/hpcwp/elsd/cmd/elscli
$ elscli -grpc.addr localhost:8082  -method Add   123 http://localhost:8072 rw
$ elscli -grpc.addr localhost:8082  -method Add   123 http://localhost:8080 r
$ elscli -grpc.addr localhost:8082  -method Add   124 http://localhost:8072 rw
$ elscli -grpc.addr localhost:8082  -method Add   125 http://localhost:8080 rw
```

```
$ elscli -grpc.addr localhost:8082  -method Get   125
```

## Prometheus Metrics

```
http://localhost:8080/metrics
```

## Using AWS CLI 

```
$ export AWS_ACCESS_KEY_ID=123
$ export AWS_SECRET_ACCESS_KEY=123
$ aws dynamodb list-tables --endpoint-url http://localhost:8000
 ```