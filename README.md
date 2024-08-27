# Directory Management CLI

This project is a CLI tool for managing directories using the tree data structure. 
It simulates directory operations without interacting with the filesystem.

## Project Structure

```text
|__ endpoint-interview/     
  |__ cmd/ 
    |__ commands.go              # Handles processing commands from a file
    |__ create.go                # Implements the 'create' command
    |__ list.go                  # Implements the 'list' command
    |__ root.go                  # The root command for the CLI
    |__ move.go                  # Implements the 'move' command
  |__ directory/
    |__ directory.go             # Contains the directory data structure and operations
    |__ directory_test.go        # Unit tests for the directory operations
  |__ main.go                    # Entry point for the CLI
  |__ resources/    
    |__ input.txt                # Example input file with commands
```

## Requirements

- Go 1.19 

## Setup

1. Clone the repo and `cd` into it.

```bash
git clone git@github.com:llama-del-rey/endpoint-interview.git
cd endpoint-interview
```

2. From the root of the repo run the `FROMFILE` command to directory operations from an input file:

```bash
./dirmanager FROMFILE resources/input.txt
```

The input.txt file should contain one command per line. Commands are case-sensitive.

```text
CREATE fruits
CREATE vegetables
CREATE fruits/apples
LIST
MOVE fruits/apples vegetables
LIST
DELETE vegetables/apples
LIST
```

### Available commands

1. `CREATE [path]` creates a directory at the specified path.

```bash
- ./dirmanager create fruits
```

2. `MOVE [source] [destination]` moves a directory from the source path to the destination path. 

```bash
`./dirmanager move fruits/apples vegetables`
```

3. `DELETE [path]` deletes the directory at the specified path.

```bash
./dirmanager delete vegetables/apples
```

4. `LIST` lists all directories starting from the root.

```bash
./dirmanager list
```

5. `FROMFILE` runs commands from the provided file

```bash
./dirmanager FROMFILE some-file.txt
```

### Build the binary

The executable is bundled with the project. If you need to build the executable again:

```bash
rm dirmanager 
go build -o dirmanager .
chmod +x dirmanager
```

### Testing

To run unit tests:

```bash
go test ./...
```

