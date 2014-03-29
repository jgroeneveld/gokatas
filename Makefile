default: test

package = github.com/jgroeneveld/gokatas

.PHONY: default test

test:
	go test $(package)/...

format:
	goimports -w .
