.PHONY: all build clean

all: build

build:
	gomobile bind -v -ldflags='-s -w' -target=android -androidapi 19 -o ../app/libs/server.aar

clean:
	rm -f ../app/libs/server.aar
