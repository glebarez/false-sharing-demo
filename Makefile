bench:
	@>base.txt 
	@for i in `seq 1 5`; do go test -bench=. -cpu=2 >> base.txt; done

	@>padded.txt
	@for i in `seq 1 5`; do go test -bench=. -cpu=2 -tags padded >> padded.txt; done

	benchstat base.txt padded.txt

bench-atomic:
	@>base.txt # erase file
	@for i in `seq 1 5`; do go test -bench=. -cpu=2 -tags atomic >> base.txt; done

	@>padded.txt
	@for i in `seq 1 5`; do go test -bench=. -cpu=2 -tags atomic,padded >> padded.txt; done

	benchstat base.txt padded.txt

build:
	GOOS=linux GOARCH=amd64 go build -o test
	GOOS=linux GOARCH=amd64 go build -tags padded -o test-padded
