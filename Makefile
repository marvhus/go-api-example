TARGET := go-api-example
DOCKER_PORT := 443

SRC_DIR := src
BIN_DIR := bin
DIST_DIR := dist
RESOURCES_DIR := resources

### LOCAL ###

.PHONY: run
run:
	@echo "== Running project..."
	API_ENV=test go run ./src/

.PHONY: build
build:
	@echo "== Building project..."
	mkdir -p $(BIN_DIR)
	go build -o ./$(BIN_DIR)/$(TARGET) ./src

### DOCKER ###

.PHONY: docker-build
docker-build: package
	@echo "== Building docker container..."
	sudo docker build -t $(TARGET) ./$(DIST_DIR)

.PHONY: docker-run
docker-run: docker-build
	@echo "== Running docker container..."
	sudo docker run --rm -p $(DOCKER_PORT):$(DOCKER_PORT) $(TARGET)

### PACKAGE ###

.PHONY: package
package:
	@echo "== Packaging up project into $(DIST_DIR)..."
	mkdir -p $(DIST_DIR)
	cp -rv ./$(RESOURCES_DIR) ./$(DIST_DIR)/
	cp -rv ./$(SRC_DIR) ./$(DIST_DIR)/
	cp -v go.mod ./$(DIST_DIR)/
	cp -v go.sum ./$(DIST_DIR)/
	cp -v Dockerfile ./$(DIST_DIR)/
