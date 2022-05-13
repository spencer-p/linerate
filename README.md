# linerate: Print the rate of lines read from stdin

`linerate` is a command-line utility to print the rate of lines being printed.

For example, if you want to measure the QPS of your server:

```
$ tail /var/log/my-server.log | grep GET | linerate
9 lines/1s
3 lines/1s
16 lines/1s
15 lines/1s
12 lines/1s
...
```

##### Installation

Install with `go install github.com/spencer-p/linerate@latest`.

##### Usage:

```
Usage of ./linerate:
  -d string
    	Duration of time between printing line rates, i.e. inverted frequency (default "1s")
  -w int
    	Number of durations to window together (default 1)
```
