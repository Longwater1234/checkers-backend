
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