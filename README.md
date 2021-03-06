# golic
license generator
```
golic inject -c="2021 SuperPower Group Limited" -l=.licignore --dry
```
![Screenshot 2021-03-08 at 11 42 52](https://user-images.githubusercontent.com/7195836/110310942-6d2f3680-8003-11eb-9540-b2e21b4f2b87.png)


## Running from commandline

create `.golicignore`
```shell
# Ignore everything
*

# But not these files...
!Makefile
!*.go

# ...even if they are in subdirectories
!*/
````
And run **GOLIC**
```shell
GO111MODULE=on go get github.com/kuritka/golic@v0.1.0
$(GOBIN)/golic inject -c="2021 MyCompany Group Limited" -l=.golicignore
```


## Usage
```shell
Usage:
   inject [flags]

Flags:
  -u, --config-url string   config URL (default "https://raw.githubusercontent.com/kuritka/golic/main/config.yaml")
  -c, --copyright string    company initials entered into license (default "2021 MyCompany")
  -d, --dry                 dry run
  -h, --help                help for inject
  -l, --licignore string    .licignore path
  -t, --template string     license key (default "apache2")

Global Flags:
  -v, --verbose   verbose output
```

## Configuration
For more details see: [default configuration](https://raw.githubusercontent.com/kuritka/golic/main/config.yaml). 
Use `-u` flag to run against custom configuration or create PR. 

