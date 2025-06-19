# golangci-lint
$(go env GOPATH)/bin/golangci-lint run --issues-exit-code=1

$(go env GOPATH)/bin/golangci-lint run -c .golangci.yml

`run.skip-files`/ `issues.exclude-files` and `run.skip-dirs`/`issues.exclude-dirs` have been merged into `linters.exclusions.paths`.
