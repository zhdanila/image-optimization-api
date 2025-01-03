up:
	go run cmd/server/main.go

install-easyjson:
	go get -u github.com/mailru/easyjson/...

# Generate easyjson code for internal/app/service/garage/objects.go
geneasyj: install-easyjson
	easyjson -all -omit_empty internal/service/*/objects.go