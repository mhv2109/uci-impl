# Chess Engine(s) written in Go
## About
A chess engine writen in Go that implements the
[Universal Chess Interface](./docs/engine-interface.txt).

Currently builds a single engine that picks a valid
move at random.

## Installing
### Prerequisites
* Go v1.14
* Make

### Building
Download dependencies using the [vend](https://github.com/nomad-software/vend)
tool to get CGo dependencies.

Run `make all` to build & run tests. The output binary
can be found in `bin/` as `mhv2109-uci-random`.

### Configuring GUIs
More configuration will be required, depending on your
chess GUI.  To configure
[GNOME Chess](https://wiki.gnome.org/Apps/Chess), place
the binary on your path (for example, in `/usr/bin`)
and use [this example config](./conf/mhv2109-uci-random.conf).
Other chess GUIs (like [Arena](http://www.playwitharena.de/))
may allow you to just select the binary and protocol (UCI).
