all: clean build docker

clean:
	rm openmhz-to-discord

build:
	go build -v

docker:
	wget https://johnvansickle.com/ffmpeg/releases/ffmpeg-release-amd64-static.tar.xz
	tar Jxvf ffmpeg-release-amd64-static.tar.xz
	mv ffmpeg-*-amd64-static/ffmpeg
	rm ffmpeg-*-amd64-static ffmpeg-release-amd64-static.tar.xz -Rf
	docker build -t jbuchbinder/openmhz-to-discord .
	docker push jbuchbinder/openmhz-to-discord
