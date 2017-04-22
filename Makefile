cov:
	gocov test ./... | gocov report

test:
	go test -v ./...

annotate:
	FILENAME=$(shell uuidgen)
	gocov test > /tmp/--go-test-server-coverage.json
	gocov annotate /tmp/--go-test-server-coverage.json $(fn)

.PNONY: test cov annotate
