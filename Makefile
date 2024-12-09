all: clean build docker

clean:
	rm openmhz-to-discord

build:
	go build -v

docker:
	docker build -t jbuchbinder/openmhz-to-discord .
	docker push jbuchbinder/openmhz-to-discord
