# hangouts-wrapped

## Compilation

Go is a compiled language. All compilation commands should be run from the `go` directory, and use the [official Go Compiler](https://go.dev/dl/).

### Native

To compile a native binary, run `go build -o hangouts-wrapped` to create an executable called hangouts-wrapped

To compile Go code for web: `GOOS=js GOARCH=wasm go build -o ../services/analysis/analysis.wasm`

## Running on a local server

There is a hard-coded assumption that the web version is running at a path prefixed with `https://hostname/hangouts-wrapped/`. Set up your local server accordingly.
