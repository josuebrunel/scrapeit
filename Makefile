NAME=scrapeit
BIN_SERVER="./bin/${NAME}-server"
BIN_CLIENT="./bin/${NAME}-client"
BIN_CONTRACT="./bin/${NAME}-contract"
MAIN_SERVER=./cmd/server/main.go
MAIN_CLIENT=./cmd/client/main.go
MAIN_CONTRACT=./cmd/contract/main.go

deps:
	apt update && apt upgrade -y && apt install protobuf-compiler -y
	go get "github.com/labstack/echo/v5@v5.0.0-20230722203903-ec5b858dab61"
	go install github.com/a-h/templ/cmd/templ@latest
	go get github.com/a-h/templ@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go get "github.com/ethereum/go-ethereum"

debug.build: build
	go build -gcflags "all=-N -l" -ldflags="-compressdwarf=false" -o ${BIN} ${MAIN}

debug: debug.build
	dlv --listen=:4000 --headless=true --log=true --accept-multiclient --api-version=2 exec ${BIN}

tidy:
	go mod tidy

templ:
	templ generate

build.osx:
	GOOS=darwin GOARCH=amd64 go build -o "./bin/scrapeit-client-darwin-amd64" ${MAIN_CLIENT}

build.windows:
	GOOS=windows GOARCH=amd64 go build -o "./bin/scrapeit-client-windows-amd64.exe" ${MAIN_CLIENT}

build.others: build.osx build.windows

build: tidy templ
	go build -o ${BIN_SERVER} ${MAIN_SERVER}
	go build -o ${BIN_CLIENT} ${MAIN_CLIENT}
	go build -o ${BIN_CONTRACT} ${MAIN_CONTRACT}

build.contract:
	solc --bin --abi --overwrite contract/contract.sol -o contract/
	abigen --bin=contract/Scraper.bin --abi=contract/Scraper.abi --pkg=contract --out=contract/scraper.go

test.int:
	python3 test-int.py

run: build
	./${BIN_SERVER} "scraper"
