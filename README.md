# Grabble

Grabble is a implementation of Scrabble with Go. Right now it has no support for normal players, this project focused on adding "AI" to Scrabble. However, package "grabble" provides everything that is needed to implement normal game. Feel free to use this package and implement support with e.g. websockets.

This project utilized:

* Interesting data structure: GADDAG
* Complicated backtracking algorithms
* TDD with use of golden files and fixtures

All available commands are under `make help` command.

## Running example game

You can run this application with commands:

> NOTE: You have to provide dictionary to use. By default program is expecting dict in `fixtures/english_dict.txt`

```bash
make game-run # Run game once, additionally there will be screenshot after each turn.
make NUM=1000 game-run-X # Run game 1000 times, no screenshots.
```

also there are some basic statistics:

```bash
make game-get-average          # Get average points from log file
make game-get-average-player   # Get average player winner
```

## Tests

Call `make unit` to run all unit tests.

Call `make bench` to run all benchmark tests.

## Golden files

To update golden files call `make golden-update`

