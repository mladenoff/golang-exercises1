## Eight Queens

Solve the eight queens problem. I created two structs:

1. `position`, which consists of `row` and `col`.
2. `state`, which contained:
  * `colsOccupied`, an `[8]bool`.
  * `upDiagonalsOccupied`, a `[15]bool`.
  * `downDiagonalsOccupied`, a `[15]bool`.
  * `queenPositions`, a `[8]position`.
  * `numQueensPlaced`, an `int`.

I assumed that I would try to fill the queens in from row zero up to row
seven, so I didn't need a `rowsOccupied`.

I wrote a method called `placeQueens(s *state, row int) bool`. This
tries to place a queen in the row. It iterates through the columns,
checking if there are any conflicts in the occupied state variables.

If not, it tries to place a queen, updating the various occupied
properties and placing the queen position in queenPositions and
incrementing `numQueensPlaced`.

If we just placed the eighth queen, I return true. Else, I recursively
call `placeQueen`, but for the next row.

If the recursive call returns true, you can keep on returning true back
up the stack. But if false, you need to undo your updates to the
occupied variables, and decrement `numQueensPlaced`. Continue with the
next column.
