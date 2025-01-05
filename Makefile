up:
	go run cmd/server/main.go

install-easyjson:
	go get -u github.com/mailru/easyjson/...

geneasyj: install-easyjson
	easyjson -all -omit_empty internal/service/*/objects.go

swagger:
	swag init --parseDependency --parseInternal -g cmd/server/main.go

dbuild:
	docker build -t image-optimization-api .

dcomposebuild:
	docker-compose down
	docker-compose up --build