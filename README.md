# checkers-backend

Server for my game SpaceCheckers, a multiplayer Checkers game for Windows, MacOS and Linux, [now available on itch.io](https://longwater1234.itch.io/spacecheckers). The game is free to play, and gets updated regularly.

You can self-host this server locally, or in the Cloud, for use with the game client. The game client is [available on this repo](https://github.com/Longwater1234/space-checkers)

## Requirements

- Golang 1.24 or newer. Get it from [official site](https://go.dev/dl/)

## How to build

- Simply open up your terminal (or Powershell) at this project root directory and run the following command.

```bash
    go build --ldflags="-s -w" .
    ./checkers-backend
```

- Set ENV variable for `PORT` which the server will listen to.
- If port not set, it will be default listen at 9876.
- Alternatively, you can use the provided Dockerfile.

## License

[MIT License](LICENSE) &copy; 2024, Davis T.
