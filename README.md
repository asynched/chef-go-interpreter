# CHEFGI - CHEF GO INTERPRETER

A chef interpreter written in Golang

## About the project

### How does it work?

The main application is an interpreter for the CHEF language, you pass a file for it to parse and interpret. Example:

-   Source file:

```
Gibberish:

Main:
	104 ml of nothing
	105 g of void
```

-   Command:

```
$ ./main.out $SOURCE_FILE
```

> This will print "hi" to the terminal

## Build

You can use the cmake program to build the application, to build it just run:

```
$ make build
```

Running the command above will create a file named `main` in the root of the project

## Author

| ![Eder Lima](https://github.com/asynched.png?size=100) |
| ------------------------------------------------------ |
| [Eder Lima](https://github.com/asynched)               |
