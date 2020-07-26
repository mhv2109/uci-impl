# A couple of chess engines written in Go
## About
A pair of chess engines writen in Go that implements the
[Universal Chess Interface](./docs/engine-interface.txt).

`mhv2109-uci-random` engine picks a valid move at random.
`mhv2109-uci-minimax` engine uses the Minimax algorithm to select the best move
based on the relative advantage of the computer player in terms of summed value
of pieces on the board in centi-pawns.

## Installing
### Prerequisites
* Go v1.14
* Make

### Building
Run `make all` to build & run tests. The output binaries can be found in `bin/`
as `mhv2109-uci-random` and `mhv2109-uci-minimax`.

### Configuring GUIs
More configuration will be required, depending on your chess GUI.  To configure
[GNOME Chess](https://wiki.gnome.org/Apps/Chess), place the binary on your path
(for example, in `/usr/bin`) and use 
[this example config for mhv2109-uci-randon](./conf/mhv2109-uci-random.conf) or
[this example config for mhv2109-uci-minimax](./conf/mhv2109-uci-minimax.conf).
Other chess GUIs (like [Arena](http://www.playwitharena.de/))
may allow you to just select the binary and protocol (UCI).
