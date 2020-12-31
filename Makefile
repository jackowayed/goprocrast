build:
	go build

release: build
	sudo cp goprocrast /usr/local/bin/
	sudo chmod a+s /usr/local/bin/goprocrast