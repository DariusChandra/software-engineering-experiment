# software-engineering-experiment

generate coverage file
```go test ./... -coverprofile=coverage.out```

we can ask for the coverage to be broken down by function, although that’s not going to illuminate much in this case since there’s only one function:
```go tool cover -func=coverage.out```

A much more interesting way to see the data is to get an HTML presentation of the source code decorated with coverage information. This display is invoked by the -html flag:
```go tool cover -html=coverage.out```

uploading coverage for golang 

bash <(curl -Ls https://coverage.codacy.com/get.sh) report \ --force-coverage-parser go -r coverage.out