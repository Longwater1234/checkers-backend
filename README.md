# checkers-backend
Server for SpaceCheckers game. Written entirely using only Go's standard library.

## Requirements

- Golang 1.21 or newer. Get it from [official site](https://go.dev/dl/)


## How to build

- Simply open up your terminal (or CMD) at this project root directory and run the following command.

```bash
    go build --ldflags="-s -w" .
```
- Set ENV variable for `PORT` which the server will listen to.
- If port not set, it will be default listen at 9876
- How to determine a point (x,y) is inside circle with radius r and center (a,b) see [link](https://study.com/skill/learn/determining-if-a-point-lies-inside-outside-or-on-a-circle-given-the-center-point-a-radius-explanation.html)


