run:
	@templ generate
	@go build -o maxtrackr -v
	@./maxtrackr