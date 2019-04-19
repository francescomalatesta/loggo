# Loggo - Learning Project in Go

[![Build Status](https://travis-ci.org/francescomalatesta/loggo.svg?branch=master)](https://travis-ci.org/francescomalatesta/loggo)

This is a simple project I am making as I am learning Go for my everyday job in Hootsuite.

Loggo is a very basic logging service with a TCP server implementation to handle requests.

Such original name, much training.

## Build

Just

```
$ go build
```

## Usage

Once you built it, run it with

```
$ ./loggo 7777
```

where 7777 is the port you want to use for the service. Then, you can open a TCP connection to it with

```
$ nc localhost 7777
```

### Commands

#### HELLO

A very basic "Hello World" command. Put here for learning purposes.

## Testing

Run tests with:

```
$ go test ./...
```
