# help
.PHONY: help
help:
	@echo "make run"
	@echo "       run the web server"
	@echo "make watch"
	@echo "       watch with air (auto reload)"



.PHONY: run
# run the web server
run:
	go run cmd/main.go



.PHONY: watch
# watch with air (auto reload)
watch:
	air --build.cmd "go build -o bin/web cmd/main.go" --build.bin "./bin/web"