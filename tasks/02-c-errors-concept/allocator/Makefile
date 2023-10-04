SRC = main.c
BIN = ./main.out

.PHONY: run check
.DEFAULT_GOAL := run

build:
	@docker run \
		--rm \
		--platform linux/amd64 \
		--workdir $(HOME) \
		-v $(PWD):$(HOME) \
		--entrypoint gcc \
		 karek/valgrind -Wall -g -o $(BIN) $(SRC)

run: build
	@docker rm -f valgrind >/dev/null 2>&1 || true
	@docker run \
		  --rm \
		  --name valgrind \
		  -i \
		  --platform linux/amd64 \
		  --workdir $(HOME) \
		  -v $(PWD):$(HOME) \
		  --entrypoint valgrind \
		  karek/valgrind --tool=memcheck --suppressions=allocator.suppr --leak-check=full $(BIN)

check: build
	@docker rm -f valgrind >/dev/null 2>&1 || true
	@for input in ./testdata/*.txt ; do \
		docker run \
		  --rm \
		  --name valgrind \
		  -i \
		  --platform linux/amd64 \
		  --workdir $(HOME) \
		  -v $(PWD):$(HOME) \
		  --entrypoint valgrind \
		  karek/valgrind -q --tool=memcheck --suppressions=allocator.suppr --leak-check=full $(BIN) < $$input ; \
  	done
