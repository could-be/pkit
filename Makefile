genstatic:
	go build -o bin/genstatic cmd/genstatic/main.go
	mv bin/genstatic $(GOPATH)/bin

install: genstatic
	genstatic
	go build -o bin/pkit cmd/pkit/main.go
	mv bin/pkit $(GOPATH)/bin

