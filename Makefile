all: clean  yaccdata build test

test:
	go test -v
build:
	go build
clean:
	rm -f ./yaccpragmaids.y.go

yaccdata:
	goyacc  -o ./yaccpragmaids.y.go  -p "YaccPragmaIds"  ./yaccpragmaids.y
