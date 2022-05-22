GO = go
COMPOSE = docker-compose
DOCKER_DIR = ./docker

# Run API without building it
# example: make run
run:
	$(GO) run ./cmd/api/main.go

# Docker compose commands for Apple M1
# example: make m1-up
# example: make m1-down
m1-up:
	$(COMPOSE) -f $(DOCKER_DIR)/docker-compose-m1.yaml up
m1-down:
	$(COMPOSE) -f $(DOCKER_DIR)/docker-compose-m1.yaml down

# Docker compose commands for Intel
# example: make intel-up
# example: make intel-down
intel-up:
	$(COMPOSE) -f $(DOCKER_DIR)/docker-compose-intel.yaml up
intel-down:
	$(COMPOSE) -f $(DOCKER_DIR)/docker-compose-intel.yaml down
