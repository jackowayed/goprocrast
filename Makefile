build:
	go build

install: build
	sudo cp goprocrast /usr/local/bin/
	sudo chmod a+s /usr/local/bin/goprocrast