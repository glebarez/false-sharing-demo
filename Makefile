bench:
	go test -bench=. -cpu=2 -count=5              > base.txt
	go test -bench=. -cpu=2 -count=5 -tags padded > padded.txt
	benchstat base.txt padded.txt

bench-atomic:
	go test -bench=. -cpu=2 -count=5 -tags atomic        > base.txt
	go test -bench=. -cpu=2 -count=5 -tags atomic,padded > padded.txt
	benchstat base.txt padded.txt

build:
	GOOS=linux GOARCH=amd64 go build -o test
	GOOS=linux GOARCH=amd64 go build -tags padded -o test-padded
