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
+---------------------------------------------------+
|     |  32  |     |  31  |     |  30  |     |  29  |  Row 7 (Top)
+---------------------------------------------------+
|  28 |     |  27  |     |  26  |     |  25  |      |  Row 6
+---------------------------------------------------+
|     |  24  |     |  23  |     |  22  |     |  21  |  Row 5
+---------------------------------------------------+
|  20 |     |  19  |     |  18  |     |  17  |      |  Row 4
+---------------------------------------------------+
|     |  16  |     |  15  |     |  14  |     |  13  |  Row 3
+---------------------------------------------------+
|  12 |     |  11  |     |  10  |     |  9   |      |  Row 2
+---------------------------------------------------+
|     |   8  |     |   7  |     |   6  |     |   5  |  Row 1
+---------------------------------------------------+
|   4 |     |   3  |     |   2  |     |   1  |      |  Row 0 (Bottom)
+---------------------------------------------------+
Col 0  Col 1  Col 2  Col 3  Col 4  Col 5  Col 6  Col 7
```
