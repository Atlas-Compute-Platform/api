help:
	@printf "Atlas Make Options\n"
	@printf	"build: Builds repository\nclean: Cleans repository\n"
build:
	@go fmt
	@go build
clean:
	@go clean
docker:build
	@docker build