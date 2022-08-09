# movie-picker
## Usage
1. `go install .`
2. `movie <command>`

## Commands
- `movie`: help
- `movie add <movie name>`: adds a new movie
- `movie add-watched <movie name>`: adds a movie to  the watched list
- `movie list`: lists all movies
- `movie watched:` lists all watched movies
- `movie delete <id>`: delete movie based on id
- `movie pick`: picks a movie, prompts truth(use fun picker), adds to watched, deletes from list, else picks new movie
- `movie fun`: interactive cli

## Dependencies
- [cobra](https://github.com/spf13/cobra)
- [promptui](https://github.com/manifoldco/promptui)
- [boltdb](https://github.com/boltdb/bolt)
- [color](https://github.com/fatih/color)
