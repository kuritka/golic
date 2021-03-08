# golic
license generator

## Running from commandline

create `.golicignore`
```shell
cat .golicignore <<EOF
# Ignore everything
*

# But not these files...
!Makefile
!*.go

# ...even if they are in subdirectories
!*/
EOF
````
And run **GOLIC**
```shell
GO111MODULE=on go get github.com/kuritka/golic@v0.1.0
$(GOBIN)/golic inject -c="2021 MyCompany Group Limited" -l=.golicignore
```
