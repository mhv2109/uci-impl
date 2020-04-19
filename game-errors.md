# ~Illegal Opening move for Black~ FIXED
```
2020-04-15 21:08:53.935-->1:ucinewgame
2020-04-15 21:08:53.935-->1:isready
2020-04-15 21:08:53.936<--1:readyok
2020-04-15 21:08:53.942-->1:position startpos moves h2h4
2020-04-15 21:08:53.942-->1:go wtime 300000 btime 300000 winc 0 binc 0
2020-04-15 21:08:53.942<--1:bestmove h2h4 <-- illegal move!
```

# Invalid move
Move already played, no piece in square.
```
2020-04-19 15:02:37.217<--1:-- engine[slotnr].EngineProcess.Running --
2020-04-19 15:02:37.222-->1:uci
2020-04-19 15:02:37.222<--1:id name mhv2109-engine
2020-04-19 15:02:37.222<--1:id author mhv2109
2020-04-19 15:02:37.223<--1:option name hash type spin default 1024 min 1024 max 4096
2020-04-19 15:02:37.223<--1:option name UCI_EngineAboutOption type string default A UCI Chess engine written in Go by mhv2109
2020-04-19 15:02:37.223<--1:uciok
2020-04-19 15:02:37.223<--1:copyprotection checking
2020-04-19 15:02:37.223<--1:copyprotection ok
2020-04-19 15:02:37.227-->1:isready
2020-04-19 15:02:37.263<--1:readyok
2020-04-19 15:02:37.434*1*Start calc, move no: 1
2020-04-19 15:02:37.435-->1:ucinewgame
2020-04-19 15:02:37.435-->1:isready
2020-04-19 15:02:37.435<--1:readyok
2020-04-19 15:02:37.442-->1:position startpos moves d2d4
2020-04-19 15:02:37.442-->1:go wtime 300000 btime 300000 winc 0 binc 0
2020-04-19 15:02:37.443<--1:bestmove b8c6
2020-04-19 15:02:37.443*1*Found move:Nb8-c6
2020-04-19 15:02:44.920*1*Start calc, move no: 3
2020-04-19 15:02:44.920-->1:position startpos moves d2d4 b8c6 e2e4
2020-04-19 15:02:44.920-->1:go wtime 292931 btime 300000 winc 0 binc 0
2020-04-19 15:02:44.921<--1:bestmove b7b6
2020-04-19 15:02:44.921*1*Found move:b7-b6
2020-04-19 15:03:00.205*1*Start calc, move no: 5
2020-04-19 15:03:00.206-->1:position startpos moves d2d4 b8c6 e2e4 b7b6 c2c4
2020-04-19 15:03:00.206-->1:go wtime 278038 btime 300000 winc 0 binc 0
2020-04-19 15:03:00.206<--1:bestmove a7a6
2020-04-19 15:03:00.206*1*Found move:a7-a6
2020-04-19 15:05:15.882*1*Start calc, move no: 7
2020-04-19 15:05:15.882-->1:position startpos moves d2d4 b8c6 e2e4 b7b6 c2c4 a7a6 h2h4
2020-04-19 15:05:15.882-->1:go wtime 142757 btime 300000 winc 0 binc 0
2020-04-19 15:05:15.883<--1:bestmove c6b4
2020-04-19 15:05:15.883*1*Found move:Nc6-b4
2020-04-19 15:05:22.608*1*Start calc, move no: 9
2020-04-19 15:05:22.608-->1:position startpos moves d2d4 b8c6 e2e4 b7b6 c2c4 a7a6 h2h4 c6b4 a2a4
2020-04-19 15:05:22.609-->1:go wtime 136430 btime 300000 winc 0 binc 0
2020-04-19 15:05:22.611<--1:bestmove d7d5
2020-04-19 15:05:22.611*1*Found move:d7-d5
2020-04-19 15:05:25.244*1*Start calc, move no: 11
2020-04-19 15:05:25.245-->1:position startpos moves d2d4 b8c6 e2e4 b7b6 c2c4 a7a6 h2h4 c6b4 a2a4 d7d5 c4d5
2020-04-19 15:05:25.245-->1:go wtime 134253 btime 300000 winc 0 binc 0
2020-04-19 15:05:25.248<--1:bestmove h7h5
2020-04-19 15:05:25.248*1*Found move:h7-h5
2020-04-19 15:05:28.854*1*Start calc, move no: 13
2020-04-19 15:05:28.854-->1:position startpos moves d2d4 b8c6 e2e4 b7b6 c2c4 a7a6 h2h4 c6b4 a2a4 d7d5 c4d5 h7h5 h1h3
2020-04-19 15:05:28.854-->1:go wtime 131068 btime 300000 winc 0 binc 0
2020-04-19 15:05:28.857<--1:bestmove h7h5
2020-04-19 15:05:28.857*1*---------> Arena:Illegal move!: "h7h5" (Feinpruefung)
```
