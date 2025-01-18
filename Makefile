TARGET := go-api-example
DOCKER_PORT := 443

SRC_DIR := src
BIN_DIR := bin
DIST_DIR := dist
RESOURCES_DIR := resources

### LOCAL ###

#
# This builds and runs the project locally.
#
.PHONY: run
run:
	@echo "== Running project..."
	API_ENV=test go run ./src/

#
# This builds the project so that you can run it locally.
#
.PHONY: build
build:
	@echo "== Building project..."
	mkdir -p $(BIN_DIR)
	go build -o ./$(BIN_DIR)/$(TARGET) ./src

### PACKAGE ###

#
# This creates the directory ./dist and copies over the required files.
# You can then either run it using the provided Dockerfile, or you can
# run it manually as, the server binary is entirely statically linked.
# There is provided a Makefile in ./dist to help you build and run the
# docker container. The rules for docker-build and docker-run in this
# Makefile uses the makefile in ./dist
#
.PHONY: package
package:
	@echo "== Packaging up project into $(DIST_DIR)..."
	mkdir -p $(DIST_DIR)
	cp -rv ./$(RESOURCES_DIR) ./$(DIST_DIR)/
	cp -v Dockerfile.dist ./$(DIST_DIR)/Dockerfile
	cp -v Makefile.dist ./$(DIST_DIR)/Makefile
	# Build the server binary and copy it out.
	sudo docker build -t $(TARGET)_build .
	sudo docker create --name $(TARGET)_dummy $(TARGET)_build
	sudo docker cp $(TARGET)_dummy:/app/server ./$(DIST_DIR)
	sudo docker rm -f $(TARGET)_dummy

### DOCKER ###

#
# This builds the docker container for the packaged version of the project.
# Will package the project before running.
#
.PHONY: docker-build
docker-build: package
	$(MAKE) docker-build -C ./$(DIST_DIR)

#
# This runs the docker container for the packaged version of the project.
# Will build the docker container before running.
#
.PHONY: docker-run
docker-run: docker-build
	$(MAKE) docker-run -C ./$(DIST_DIR)
