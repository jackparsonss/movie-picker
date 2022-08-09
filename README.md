# movie-picker
## Usage
1. `go install .`
2. `movie`

## Commands
- `add <movie name>`: adds a new movie
- `add-watched <movie name>`: adds a movie to  the watched list
- `list`: lists all movies
- `watched:` lists all watched movies
- `delete <id>`: delete movie based on id
- `pick`: picks a movie, prompts truth(use fun picker), adds to watched, deletes from list, else picks new movie
- `fun`: interactive cli

## Dependencies
- [cobra](https://github.com/spf13/cobra)
- [promptui](https://github.com/manifoldco/promptui)
