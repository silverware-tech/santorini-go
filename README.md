
# Santorini

Remake of Santorini board game in Go language.

## Rules
The youngest player is the Start Player.

Players take turns, starting with the Start Player, who first placed their Workers. On your turn, select one of your Workers. You must **move** and then **build** with the selected Worker.

### Move

- Move your selected Worker into one of the (up to) eight neighboring spaces.
- A Worker may move up a maximum of one level higher, move down any number of levels lower, or move along the same level.
- A Worker may not move up more than one level.
- The space your Worker moves into must be unoccupied (not containing a Worker or Dome).

### Build

- Build a block or dome on an unoccupied space neighboring the moved Worker.
- You can build onto a level of any height, but you must choose the correct shape of block or dome for the level being built (See the schema below).
- A tower with 3 blocks and a dome is considered a “Complete Tower”.

``` txt
     ________
     | Dome |
    ___________
    | Level 3 | ---|
    ___________    |
    | Level 2 |    |--> Blocks
    ___________    |
    | Level 1 | ---|
------ Ground ------
```

### Winning the game

- If one of your Workers moves up on top of level 3 during your turn, you instantly win!
- You must always perform a move then build on your turn. If you are unable to, you lose.
