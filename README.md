# checkers-backend
Server for SpaceCheckers game. Written entirely using only Go's standard library.

## Requirements

- Golang 1.21 or newer. Get it from [official site](https://go.dev/dl/)


## How to build

- Simply open up your terminal (or CMD) at this project root directory and run the following command.

```bash
    go build --ldflags="-s -w" .
```
- Set ENV variable for `GAME_PORT` which the server will listen to.
- If port not set, it will be default listen at 9876
- How to determine a point (x,y) is inside circle with radius r and center (a,b) see [link](https://study.com/skill/learn/determining-if-a-point-lies-inside-outside-or-on-a-circle-given-the-center-point-a-radius-explanation.html)


### Relation between Pieces on Checkboard

- PLAYER 1 is always RED 🔴. PLAYER 2 IS always BLACK ⚫️
- When a piece is on EVEN row of cells. The number denotes the cell_index _delta_ between two pieces.

isEven : true, deltaForward = 4 , deltaBehindEnemy  =5
```
4         3
 \       /   
  \     /
 EVEN_ROW   
  /     \
 /       \   
4         5

```

- When a piece is on ODD row of cells.
```
5         4
 \       /   
  \     /
  ODD_ROW   
  /     \
 /       \   
3         4

```
