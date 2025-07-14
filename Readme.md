## Criteria to select programming language: I chose Go as language.

## Criteria to select framework: use Gin since fast and minimalistic 

## steps to create api command and dependency
go mod init github.com/akshayverma91/rydeapi
go get github.com/gin-gonic/gin
go get go.mongodb.org/mongo-driver/mongo
go get github.com/swaggo/gin-swagger
go get github.com/swaggo/files
go install github.com/swaggo/swag/cmd/swag@latest


## Initialize Swagger
use command to generate swagger using swag init
add or update controller run again swag init
Run project and go to http://localhost:8080/swagger/index.html 

## Run Docker Container
docker build -t rydeapi .
docker run -p 8080:8080 rydeapi

for docker compose
to build: docker-compose build
to start: docker-compose up
to stop: docker-compose down


## To run test
go test ./tests