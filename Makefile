#Run build and run to create and use the project

information:
	echo "Zera Generator by Luca Becci"

build:
	go build -o bin/main cmd/main.go

run:
	bin/main

