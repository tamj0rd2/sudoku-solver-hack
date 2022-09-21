include .bingo/Variables.mk

.PHONY: test
.DEFAULT_GOAL := run

run:
	docker-compose build app
	docker-compose run --rm app

run2:
	go run ./cmd/cli2/...

setup:
	git config core.hooksPath .hooks
	go install github.com/bwplotka/bingo@latest
	bingo get
	bingo get -l github.com/bwplotka/bingo

t: test
test: lint
	$(GOTEST) --race --count=1 ./...

ci:
	git pull -r
	make test
	git push

lint:
	$(GOLANGCI_LINT) run --timeout=5m ./...

fix-imports:
	@$(GCI) write -s standard -s default -s "prefix(github.com/tamj0rd2/sudoku-solver-hack)" $$(find . -type f -name '*.go' -not -path "./vendor")

lf: lintfix
lintfix:
	@$(GOLANGCI_LINT) run ./... --fix
	@$(MAKE) fix-imports
