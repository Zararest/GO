gogen:
	go generate ./...

gobuild: 
	go build -o bin/sfss cmd/sfss/main.go cmd/sfss/cli.go 
	
gotest:
	go test -v --cover ./...

godocs:
	go run github.com/swaggo/swag/cmd/swag init -g  cmd/sfss/main.go