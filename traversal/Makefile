bench:
	go test -bench='BenchmarkTraverse' -tags=cols -count=5 > cols.txt
	go test -bench='BenchmarkTraverse' -tags=rows -count=5 > rows.txt
	benchstat cols.txt rows.txt

build:
	GOOS=linux GOARCH=amd64 go build -o traversal