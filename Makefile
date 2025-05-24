NAME=scrapeit
BIN_SERVER="./bin/${NAME}-server"
BIN_CLIENT="./bin/${NAME}-client"
MAIN_SERVER=./cmd/server/main.go
MAIN_CLIENT=./cmd/client/main.go

debug.build: build
	go build -gcflags "all=-N -l" -ldflags="-compressdwarf=false" -o ${BIN} ${MAIN}

debug: debug.build
	dlv --listen=:4000 --headless=true --log=true --accept-multiclient --api-version=2 exec ${BIN}

tidy:
	go mod tidy

build: 
	go build -o ${BIN_SERVER} ${MAIN_SERVER}
	go build -o ${BIN_CLIENT} ${MAIN_CLIENT}
	

run: build
	./${BIN_SERVER} "scraper"
