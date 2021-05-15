PACKAGE_NAME=github.com/saurabh3460/cli_cobra_1

BUILD_DIR=build/
BUILD_FILE=$(BUILD_DIR)$(PACKAGE_NAME)

make: compile run

compile:
	go build -o $(BUILD_FILE)

clean: 
	rm -f $(BUILD_FILE)

run: 
	$(BUILD_FILE)