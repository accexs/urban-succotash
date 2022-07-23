## Readme

Go to quickstart section for a fast start.

Used gin-gonic framework as core of the API, gorm as ORM with postgres as DB and SQLite as memory database for 
integration tests. API is secured using JWT and tests uses testify package.

Tests mainly cover "happy" path but good coverage was achieved. There are integration and unit tests, mocking techniques
were used to isolate tests, service and repository layers and use of structs and interfaces were used to try to achieve
a clean architecture. To run all tests: `docker-compose exec me_wallet_service go test ./...`  You could run them
locally as well.

Docker compose file is supplied with air for hot reload and a postgres database (no volumes for reboot persistence).
Build and server with air (hot reload) should be ready to use with docker-compose.

Godotenv package is installed, and for security reasons a `ACCESS_SECRET` could be used. DB credentials are set by
default on server creation to work with docker however the `.env` file could be used to provide something more specific.

Running `docker-compose up` should be enough to have everything setup.

On `/swagger/index.html` is the API documentation and playground, everything should be usable but consider the 
target user id for the transfer. The API needs an access token that is provided by the login method, on the swagger ui, 
just execute the example after running the seed.

Run seed to create test users: `docker-compose exec me_wallet_service go run setup/seed.go`

On `src` directory run `swag init -g utils/server.go --output ../docs --parseDependency` to generate 
swagger/open API v2 specs

### Quickstart

```
docker-compose up
docker-compose exec me_wallet_service go run setup/seed.go
swag init -g utils/server.go --output ../docs --parseDependency