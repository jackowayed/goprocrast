Toggles /etc/hosts entries to block distracting sites.

A port of [noprocrast](https://github.com/rfwatson/noprocrast) in Golang.

Uses setuid so that you can run it without sudoing each time.

## Installation

Download the code and then to test or run within that directory

```
$ make
```

If you pass environment variable `DEV=1`, it will use `./hosts.dev`
as your hosts file instead of `/etc/hosts`, for testing.

To install (requries sudo at install time, but then not when you run it).

```
$ make install
```