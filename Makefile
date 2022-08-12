include .env

build:
	go build -o ./bin/tg-proxy

run:
	./bin/tg-proxy -a=$(ADDRESS) -t=$(BOT_TOKEN)