### Relation between Pieces on Checkboard

- PLAYER 1 is always RED 🔴. PLAYER 2 IS always BLACK ⚫️
- When a piece is on EVEN row of cells. The number denotes the cell*index \_delta* between two pieces.

isEven : true, deltaForward = 4 , deltaBehindEnemy =5

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

## Board Layout:

Here is layout of the game board as shown on the game client. The empty boxes are light-colored, non-playable. The ones with numbers are dark-cells, which are are playable cells.

```txt

0    1    2    3    4    5    6    7
  +---------------------------------------+
7 |    | 32 |    | 31 |    | 30 |    | 29 |
  |---------------------------------------|
6 | 28 |    | 27 |    | 26 |    | 25 |    |
  |---------------------------------------|
5 |    | 24 |    | 23 |    | 22 |    | 21 |
  |---------------------------------------|
4 | 20 |    | 19 |    | 18 |    | 17 |    |
  |---------------------------------------|
3 |    | 16 |    | 15 |    | 14 |    | 13 |
  |---------------------------------------|
2 | 12 |    | 11 |    | 10 |    |  9 |    |
  |---------------------------------------|
1 |    |  8 |    |  7 |    |  6 |    |  5 |
  |---------------------------------------|
0 |  4 |    |  3 |    |  2 |    |  1 |    |
  +---------------------------------------+

```