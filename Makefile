all: gobuild

webbuild:
	@cd webui && make build

gobuild:
	@cd gosrc && make build

run:
	@cd gosrc && make run
