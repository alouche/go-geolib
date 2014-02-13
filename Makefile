bootstrap:
	@bash --norc -i ./scripts/bootstrap.sh

cov: bootstrap
	gocov test ./... | gocov report

test: bootstrap
	go test -v ./...

annotate:
	FILENAME=$(shell uuidgen)
	gocov test > /tmp/--go-test-server-coverage.json
	gocov annotate /tmp/--go-test-server-coverage.json $(fn)

.PNONY: test cov annotate
