# checkers-backend

Server for [SpaceCheckers](https://longwater1234.itch.io/spacecheckers) , a multiplayer checkers game for Windows and MacOS, available on itch.io. Written entirely using only Go's standard library and protobuf lib.

## Requirements

- Golang 1.21 or newer. Get it from [official site](https://go.dev/dl/)

## How to build

- Simply open up your terminal (or CMD) at this project root directory and run the following command.

```bash
    go build --ldflags="-s -w" .
```

- Set ENV variable for `PORT` which the server will listen to.
- If port not set, it will be default listen at 9876
