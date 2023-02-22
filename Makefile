.PHONY: clean test

clean:
	rm -rf cover.out

test:
	go test -v -race -gcflags="all=-N -l" -coverprofile="cover.out" -coverpkg="./..." ./... && go tool cover -html="cover.out"