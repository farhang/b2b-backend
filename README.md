# Wallex B2B

[![Intergration tests](https://github.com/farhang/b2b-backend/actions/workflows/intergration-test.yml/badge.svg)](https://github.com/farhang/b2b-backend/actions/workflows/intergration-test.yml)
[![Build](https://github.com/farhang/b2b-backend/actions/workflows/build.yml/badge.svg)](https://github.com/farhang/b2b-backend/actions/workflows/build.yml)
## Todo 

- [x] Input Validations 
- [x] Organize DTOs 
- [x] Error handling 
- [x] Find a structure for responses 
- [x] Decide where we should put http middlewares 
- [x] Move the secret key to environment variables
- [x] Remove redundant codes
- [x] Add swagger 
- [x] Use a library for environment variables
- [ ] Log management
- [ ] Install swagger & ginkgo locally 
- [ ] Add expiration time to jwt
- [x] Fix docker compose up postgresDB
- [ ] Split main.go into modules
- [ ] Add benchmarks 

## Getting started
```bash
docker compose up postgresDB 
```
```bash
make run
```
## Testing
#### run intergration tests
```
go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo
```
```
ginkgo -v test/...
```
## Documentation
#### generate swagger
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```
```bash
make swagger
```
#### the swagger-ui is served in $BASE_URL/swagger/index.html
## Contributing guidelines 
1. Please treat git with  [Conventional Commits](https://www.conventionalcommits.org/).
2. Please do not push directly in main branch, instead of that create a pull request.
3. Make sure all actions are passed while creating pull request.
4. For now Issues should be closed only by maintainers.


