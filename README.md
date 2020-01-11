## dependency
execute the command and set IDEA to vendor mode then
```bash
go mod vendor
```

## Road map
- [ ] golang syntax
- [x] [project layout](https://github.com/golang-standards/project-layout)
- [x] [gin](https://gin-gonic.com/docs/examples/)
- [x] lint
- [x] unit test
- [x] swagger integration
- [x] orm
- [ ] metrics

### lint
```bash
golangci-lint run
```

### unit test
```bash
# test all
go test ./...

# coverage
go test ./folder -cover

# visual coverage
go test ./folder -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### swagger integration
```bash
# create or update swag docs
# the docs will be served at /swagger/index.html then
swag init
```

