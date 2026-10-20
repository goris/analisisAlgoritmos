BIN_DIR := bin
EXERCISES := $(notdir $(shell find . -mindepth 1 -maxdepth 1 -type d ! -name '.git' ! -name '$(BIN_DIR)' ! -name 'docs' ! -name '.*'))

.PHONY: build test bench run clean list help notebook-env

help:
	@echo "make build        build every exercise into $(BIN_DIR)/"
	@echo "make test         run all tests"
	@echo "make bench        run all benchmarks"
	@echo "make run EX=name  run a single exercise, e.g. make run EX=expoBySqr"
	@echo "make list         list available exercises"
	@echo "make clean        remove $(BIN_DIR)/"
	@echo "make notebook-env set up the Python environment for the notebooks (.venv/)"

build:
	@mkdir -p $(BIN_DIR)
	@for ex in $(EXERCISES); do \
		echo "Building $$ex..."; \
		go build -o $(BIN_DIR)/$$ex ./$$ex || exit 1; \
	done

test:
	go test ./...

bench:
	go test -bench=. -run=^$$ ./...

run:
	@if [ -z "$(EX)" ]; then \
		echo "Usage: make run EX=<exercise>"; \
		exit 1; \
	fi
	go run ./$(EX)

list:
	@for ex in $(EXERCISES); do echo $$ex; done

clean:
	rm -rf $(BIN_DIR)

notebook-env:
	uv sync
