.PHONY: run
run: main.go testlib.go
	go run $^
