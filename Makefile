up:
	go run cmd/server/main.go

install-easyjson:
	go get -u github.com/mailru/easyjson/...

geneasyj: install-easyjson
	easyjson -all -omit_empty internal/service/*/objects.go