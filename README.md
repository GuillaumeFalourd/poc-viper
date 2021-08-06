# poc-viper

This repository is a POC using [Viper Golang Library](https://github.com/spf13/viper).

References:

- [golang viper tutorial](https://www.youtube.com/watch?v=rU9o3BlLJO8&ab_channel=GeneralistProgrammer)

## Setup

- Golang installed
- Viper library installed: `go get github.com/spf13/viper`

## Usage

This implementation will **read** fields from a `.yaml` file located at the repository `root` according to the following format

To set the `*.yaml` to read, update the [viper.go file](https://github.com/GuillaumeFalourd/poc-viper/blob/main/viper.go), line 11: `vi.SetConfigFile("{name}.yaml")`

_The default value for the yaml file is `test-rest.yaml`._

### For REST API

```yaml
name: Test Viper REST
id: 123
inputs:
  request:
    api: REST
    method: "POST"
    url: "http://localhost:8080"
    headers:  # You can use any custom-fields inside this headers map
      authorization: "Bearer token"
      accept: "application/json"
    body:  # You can use any custom-fields inside this body map
      document: "123456YZ"
      email: "test@example.com"
```

### For gRPC API

```yaml
name: Test Viper gRPC
id: 456
inputs:
  request1:
  request:
    api: gRPC
    server-address: "http://localhost:9090"
    service: "LoginRequest"
    fields: # You can use any custom-fields inside this fields map
      document: "78901AB"
      email: "abcde@test.com"
```

## Run

At the repository `root` directory run: `go run ./viper.go`

### For the REST API yaml

```bash
Name: Test Viper REST
Id: 123
Method: POST
Url: http://localhost:8080
Headers: map[accept:application/json authorization:Bearer token]
Body: map[document:123456YZ email:test@example.com]
```

### For the gRPC API yaml

```bash
Name: Test Viper gRPC
Id: 456
Server Address: http://localhost:9090
Service: LoginRequest
Fields: map[document:78901AB email:abcde@test.com]
```
