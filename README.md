# Keygoller

> Simple keylogger implementation in Go

## Installation

```sh
go get -u github.com/rotsix/keygoller
```

## Usage

```sh
Usage of keygoller:
  -channel string
    	channel to send strokes to (IRC) (default "#keylolgger")
  -host string
    	host to send strokes to (default "chat.freenode.net")
  -pass string
    	password associated to the user (default "hunter2")
  -port int
    	port to send strokes to (default 6667)
  -protocol string
    	protocol to use to send strokes (default "debug")
  -request string
    	type of request to send (HTTP) (default "post")
  -size int
    	number of strokes to get before sending (default 512)
  -url string
    	url to send strokes to (HTTP) (default "/data")
  -user string
    	user to connect with (default "Guest74629")
```

It only works on Linux for now, and must be run as root.

The logger supports the following protocols:

- debug: print to `stdout`
- IRC: pass the `-host`, `-port`, `-channel` and `-user` options
- HTTP: not yet
