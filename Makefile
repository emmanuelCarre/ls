PROJECT_NAME := ls
PKG := gitlab.com/emmanuel.b.carre/ls
OUTPUT_DIR := output

GOTEST_PKGS=$(shell go list ./... | sed 's/gitlab.com\/emmanuel\.b\.carre\/ls/./' | grep -v mocks)

OUTPUT = ${OUTPUT_DIR}/${GOOS}-${GOARCH}/${PROJECT_NAME}

.PHONY: all
all: test build-linux-amd64 build-linux-386 build-osx-amd64 build-osx-386

.PHONY: test
test:
	@echo "Running tests:"
	@go test $(GOTEST_PKGS) -cover

.PHONY: mock_gen
mock_gen:
	mockgen -source=./io/ioutil/ioutil.go -destination=mocks/ioutil.go
	mockgen -source=${GOROOT}/src/os/types.go -destination=mocks/os/types.go

.PHONY: build-linux-amd64
build-linux-amd64: 
	$(eval export GOOS=linux)
	$(eval export GOARCH=amd64)
	@echo "Build ${PROJECT_NAME} - ${GOOS} ${GOARCH}"
	@go build -o ${OUTPUT}

.PHONY: build-linux-386
build-linux-386: 
	$(eval export GOOS=linux)
	$(eval export GOARCH=386)
	@echo "Build ${PROJECT_NAME} - ${GOOS} ${GOARCH}"
	@go build -o ${OUTPUT}

.PHONY: build-osx-amd64
build-osx-amd64: 
	$(eval export GOOS=darwin)
	$(eval export GOARCH=amd64)
	@echo "Build ${PROJECT_NAME} - ${GOOS} ${GOARCH}"
	@go build -o ${OUTPUT}

.PHONY: build-osx-386
build-osx-386: 
	$(eval export GOOS=darwin)
	$(eval export GOARCH=386)
	@echo "Build ${PROJECT_NAME} - ${GOOS} ${GOARCH}"
	@go build -o ${OUTPUT}

.PHONY: clean
clean:
	rm -rf ${OUTPUT_DIR}