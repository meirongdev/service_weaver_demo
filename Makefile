# Build service weaver
build:
	@echo "Building service weaver"
	@weaver generate .
	@go build .
deploy: build
	@echo "Deploying service weaver"
	@weaver multi deploy config.toml