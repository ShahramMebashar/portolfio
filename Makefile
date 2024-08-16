.PHONY: watch
# watch with air (auto reload)
watch:
	air --build.cmd "go build -o bin/web cmd/main.go" --build.bin "./bin/web"