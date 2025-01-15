# Grep CLI Application

## Overview

This is a simple Go program that replicates the functionality of the `grep` command, allowing you to search for lines in a file that match a specified regular expression pattern.

## Features

- Reads files line by line for efficient memory usage.
- Matches lines against a user-provided regular expression.
- Outputs matching lines along with their line numbers.
- Handles errors gracefully, such as invalid file paths or patterns.

## Setup and Usage

### Build the Application

To build the application, use the provided `Makefile`. Run:

```sh
make build
```

The binary will be generated in the `bin/` directory with the name `grep`.

### Run the Application

To execute the application, use:

```sh
./bin/grep <file> <pattern>
```

### Clean Up

To remove the built binary, run:

```sh
make clean
```

## Example Output

Given a file `example.txt` with the following content:

```
1: This is a test file.
2: An error occurred here.
3: Testing again.
```

Running the command:

```sh
./bin/grep example.txt error
```

Will output:

```
2: An error occurred here.
```
