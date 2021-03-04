# golic
license appender

## 1. Add generator marks into project main package
```go
package main

//go:generate bash -c "cd $(mktemp -d) && GO111MODULE=on go get github.com/kuritka/golic@v0.0.2"

//go:generate bash -c "${GOPATH}/bin/golic run -l ./.licignore"
```
## 2. Add .licignore into project root
.licignore works the very same way as `.gitignore` but says which files will be skipped when license injected.
i.e.:
```gitignore
# Ignore everything
*

# But not these files...
!CODEOWNERS
!LICENSE
!README.md
!.gitignore
!.licignore
!.golangci
!*.yaml
!Makefile
!*.go
!go.mod
!go.sum

# ...even if they are in subdirectories
!*/

# except gitignore
/**/.gitignore
```

# 3. run generator
```shell
go generate
```

