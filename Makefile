TEMP_DIR=./dist
BIN_NAME=gedis
BIN_PATH=$(TEMP_DIR)/$(BIN_NAME)
clean:
	rm -rf ./dist/

run: clean
	@mkdir $(TEMP_DIR)
	go build -o $(BIN_PATH) *.go
	exec $(BIN_PATH)
