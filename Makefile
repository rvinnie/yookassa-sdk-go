test:
	go test -v -count=1 ./test/...

test100:
	go test -v -count=100 ./test/...

race:
	go test -v -race -count=1 ./test/...

.PHONY: test test100 race cover