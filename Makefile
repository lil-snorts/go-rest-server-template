run:
	@templ generate
	@go build -o bin/application application.go
	@go run application.go