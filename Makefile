tpl:
	@templ generate
build: tpl
	@go build -o ff cmd/factsfood/server.go
run: build
	@./ff
