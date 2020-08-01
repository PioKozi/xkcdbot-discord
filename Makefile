build:
	go build -v -o bin/xkcdbot-discord ./cmd/main

clean:
	rm -f bin/xkcdbot-discord
