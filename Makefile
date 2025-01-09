TARGET := go-api-example
PORT := 5050

### LOCAL ###

.PHONY: run
run:
	API_ENV=test go run ./src/

.PHONY: build
build:
	mkdir -p bin
	go build -o bin/$(TARGET) ./src

### DOCKER ###

.PHONY: docker-build
docker-build: Dockerfile
	@echo "Building docker container..."
	sudo docker build -t $(TARGET) .

.PHONY: docker-run
docker-run: docker-build
	@echo "Running docker container..."
	sudo docker run --rm -p $(PORT):$(PORT) $(TARGET)

### HELPERS ###

Dockerfile: Dockerfile.in
	@echo "Building Dockerfile..."
	cat $< \
		| sed 's/{{TARGET}}/'$(TARGET)'/g' \
		| sed 's/{{PORT}}/'$(PORT)'/g' \
		| tee $@
