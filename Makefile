help:
	@echo "lint - runs linter across the project using golangci-lint"
	@echo "upgrade - upgrades all go dependencies "
	@echo "build - installs all depdencies and builds go binary in GOPATH"
	@echo "test - runs the test file across the project"

lint: 
	@golangci-lint run  ./...

upgrade:
	@echo "Upgrading dependencies..."
	@go get -u
	@go mod tidy

# Build won't build binary to run as main file is not present in project root. 
# But it can be used to check whether packages can build properly.
build: 
	@echo 'Building binary...'
	@go install ./... 

test:
	@echo 'Running tests and reporting to coverage.txt...'
	@go test -race -coverprofile=coverage.txt ./...