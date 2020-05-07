## GNOME Chess and UCI
GNOME Chess communicates with chess engines over the Universal Chess Interface 
(UCI) and the Chess Engine Communication Protocol (CECP). The spec for UCI is
included in [engine-interface.txt](./engine-interface.txt). Engines are
configured for GNOME Chess using the file `/etch/gnome-chess/engines.conf`. 
Engine binaries must be on the path, typically in `/usr/bin`.

## Forsyth-Edwards Notation
[Notes on FEN](https://en.wikipedia.org/wiki/Forsyth%E2%80%93Edwards_Notation)

## Portable Game Notation
[Notes on PGN](https://en.wikipedia.org/wiki/Portable_Game_Notation)

## Useful libraries
### notnil/chess
Useful chess library in Go found [here](https://github.com/notnil/chess).

###  dylhunn/dragontoothmg
Another Go chess library found [here](https://github.com/dylhunn/dragontoothmg).

## Other Resources
* [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
* [go-makefile-example](https://github.com/azer/go-makefile-example)
* [OpenChess Forum](https://www.open-chess.org/index.php)
* [Crafty chess engine](http://craftychess.com/)

## Notes
### Ponderhit Example
```
gui -> engine: position p1 [initial position]
gui -> engine: go wtime xxx btime yyy [engine starts searching]
... time passes
gui <- engine: bestmove a2a3 ponder a7a6 [engine stops]
gui -> engine: position p1 moves a2a3 a7a6 [position after ponder move]
gui -> engine: go ponder wtime xxx btime yyy [engine starts searching]
... time passes (engine does not stop searching until 'stop' or 'ponderhit' is received)
gui -> engine: ponderhit [engine may or may not continue searching depending on time management]
... time passes (or not, engine is free to reply instantly)
gui <- engine: bestmove a3a4 ponder a6a5
```

### Pondermiss Example
```
gui -> engine: position p1
gui -> engine: go wtime xxx btime yyy [engine starts searching]
... time passes
gui <- engine: bestmove a2a3 ponder a7a6 [engine stops]
gui -> engine: position p1 moves a2a3 a7a6
gui -> engine: go ponder wtime xxx btime yyy [engine starts searching]
... time passes (engine does not stop until 'stop' or 'ponderhit' is received)
gui -> engine: stop [engine stops searching]
gui <- engine: bestmove m1 ponder m2 [this is discarded by gui -]
gui -> engine: position p1 moves a2a3 b7b6... [- because engine2 played a different move]
gui -> engine: go...
```
