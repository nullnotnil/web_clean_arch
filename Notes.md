# golangci-lint
$(go env GOPATH)/bin/golangci-lint run --issues-exit-code=1

$(go env GOPATH)/bin/golangci-lint run -c .golangci.yml

`run.skip-files`/ `issues.exclude-files` and `run.skip-dirs`/`issues.exclude-dirs` have been merged into `linters.exclusions.paths`.



# Taskfile

https://taskfile.dev/installation/#install-script

## Generate a Takfile
`task --init`

## Forwarding CLI arguments to commands
```yml
version: '3'

tasks:
  yarn:
    cmds:
      - yarn {{.CLI_ARGS}}
```

# Goreleaser
https://goreleaser.com/install/


goreleaser init


goreleaser check
