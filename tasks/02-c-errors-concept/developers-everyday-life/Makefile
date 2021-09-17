SRC = main.c
BIN = ./main.out

.PHONY: run check
.DEFAULT_GOAL := run

build:
	@docker run \
		--platform linux/amd64 \
		--workdir $(HOME) \
		-v $(PWD):$(HOME) \
		--entrypoint gcc \
		 mooreryan/valgrind -o $(BIN) $(SRC)

run: build
	@docker run \
		  -i \
		  --platform linux/amd64 \
		  --workdir $(HOME) \
		  -v $(PWD):$(HOME) \
		  mooreryan/valgrind --tool=memcheck --leak-check=full $(BIN)

check: build
	@for user_request in ./testdata/*.json ; do \
		docker run \
		  -i \
		  --platform linux/amd64 \
		  --workdir $(HOME) \
		  -v $(PWD):$(HOME) \
		  mooreryan/valgrind -q --tool=memcheck --leak-check=full $(BIN) < $$user_request ; \
  	done
