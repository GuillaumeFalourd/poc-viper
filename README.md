# poc-viper

This repository is a POC using [Viper Golang Library](https://github.com/spf13/viper).

## Setup

- Golang installed
- Viper library installed: `go get github.com/spf13/viper`

## Run

At the repository `root` directory run: `go run ./viper.go`

_Note: This will **read** fields from the `test.yaml` file at the repository `root`._

Example:

```yaml
name: Test Yaml
id: 123
inputs:
  - run: ls -lha
  - run: mkdir new
```

Will return:

```bash
Test Yaml
123
[map[run:ls -lha] map[run:mkdir new]]
map[id:123 inputs:[map[run:ls -lha] map[run:mkdir new]] name:Test Yaml] # All Settings
```
