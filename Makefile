TEMP_DIR=./dist
BIN_NAME=gedis
BIN_PATH=$(TEMP_DIR)/$(BIN_NAME)

clean:
	rm -rf ./dist/

build: clean
	@mkdir $(TEMP_DIR)
	go build -o $(BIN_PATH) *.go

run: clean build
	exec $(BIN_PATH)

test-integration: clean build
	go test -v -count=1 ./tests/integration/... 
