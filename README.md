# Go-1-ATourOfGo

Exercises from official A [Tour of Go website](https://go.dev/tour).

## Description

This repository holds my solutions to the exercises, they might not be the best solutions but they may serve as a reference for someone else.

## Usage

To run the exercises, you first need to have Go installed in your computer, then you can run the following command:

```bash
go run exercises/<exercise-name>.go
```

There's also a devcontainer configuration for VSCode, so you don't need to bother about installing the project dependencies.

If you're using Linux or the devcontainer you may find the Makefile commands useful:

```bash
make list # Lists all the exercises
make run number=<exercise-number-prefix> # Runs the specified exercise
```
