build:
	 go build -o bin/ldat cmd/ldat/main.go

clean:
	rm -f bin/ldat

.PHONY: clean