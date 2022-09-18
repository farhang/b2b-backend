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
- [x] Fix docker compose up postgresDB
- [ ] Enhance log management messages (we use log forwarding in digital ocean)
- [ ] Install swagger & ginkgo locally 
- [ ] Add expiration time to jwt
- [ ] Split main.go into modules
- [ ] Use A Library for [Dependency injection](https://github.com/avelino/awesome-go#dependency-injection)
- [ ] Add benchmarks 
- [ ] Remove [Kavenegar](https://kavenegar.com/) token from .env.dev

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


