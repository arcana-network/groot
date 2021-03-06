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

# Tests needs to run separate, refer logger_test.go
test-acceptance:
	@echo 'Running acceptance tests and reporting to coverage.txt...'
	@go test ./... -v -run TestNewZapLogger
	@go test ./... -v -run TestGlobalNewZapLogger
	@go test ./... -v -run TestNewZapLoggerEmptyService
	@go test ./... -v -run TestNewGlobalZapLoggerEmptyService
	@go test ./... -v -run TestSinkRepeat
	@go test ./... -v -run TestGlobalSinkRepeat
	@go test ./... -v -run TestFileCreation
	@go test ./... -v -run TestGlobalFileCreation
	@go test ./... -v -run TestFileContent
	@go test ./... -v -run TestGlobalFileContent
	@go test ./... -v -run TestFatalLogs
	@go test ./... -v -run TestGlobalFatalLogs
	@go test ./... -v -run TestPanicLogs