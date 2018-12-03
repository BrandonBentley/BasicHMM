run: build
	./hmm
no: build
	./hmm 100
mil: build
	./hmm 1000000
clean:
	rm dataset.json
build:
	go build -o hmm main.go dataset.go