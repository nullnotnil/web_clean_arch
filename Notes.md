# golangci-lint
$(go env GOPATH)/bin/golangci-lint run --issues-exit-code=1

$(go env GOPATH)/bin/golangci-lint run -c .golangci.yml
