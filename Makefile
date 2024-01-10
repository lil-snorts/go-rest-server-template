build:
	@templ generate
	@go build -o maxtrackr -v
run:
	@templ generate
	@go build -o maxtrackr -v
	@./maxtrackr
