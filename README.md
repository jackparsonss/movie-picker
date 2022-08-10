# movie-picker
## Usage
1. `go install .`
2. `movie <command>`

## Commands
- `movie`: help
- `movie add <movie name>`: adds a new movie
- `movie add <movie name> --watched`: adds a movie to  the watched list
- `movie watch`: moves movie from movie list to watched list
- `movie list`: lists all movies
- `movie list --watched:` lists all watched movies
- `movie delete <id>`: delete movie based on id
- `movie delete <id> --watched`: delete movie from watched list based on id
- `movie pick`: picks a random movie
- `movie fun`: interactive cli

## Dependencies
- [cobra](https://github.com/spf13/cobra)
- [promptui](https://github.com/manifoldco/promptui)
- [boltdb](https://github.com/boltdb/bolt)
- [color](https://github.com/fatih/color)
