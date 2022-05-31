build:
	go build -o cmd ./

test: 
	go test -timeout 99999s
