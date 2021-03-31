check:
	go vet ./src

run:	
	go run ./src/cryptogo.go
%:
	echo "Hola"

test:
	go test ./tests/cryptogo_test.go