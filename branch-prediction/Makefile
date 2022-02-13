bench:
	go test -bench=.  -count=5 -tags=rand > random.txt
	go test -bench=.  -count=5            > partitioned.txt
	benchstat random.txt partitioned.txt

build:
	GOOS=linux GOARCH=amd64 go build -o partitioned
	GOOS=linux GOARCH=amd64 go build -tags=rand -o rand