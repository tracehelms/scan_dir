# Scan Dir
Scan a directory and get an output of files grouped by extension.

## Usage
After cloning the project, in the console in the project directory, build the binary from source:
```
go build scandir.go
```

You can pass in a folder to scan:
```
./scandir -dir=/Users/tracehelms/
```

Or with no arguments it scans the current directory:
```
./scandir
```

## Example Output
```
~/code/golang/src/github.com/tracehelms/scan_dir/ [master]$ ./scandir
Current working directory is: /Users/trace.helms/code/golang/src/github.com/tracehelms/scan_dir
Scanning directory: /Users/trace.helms/code/golang/src/github.com/tracehelms/scan_dir
No Extension:   31
.sample:   9
.md:   1
.go:   1
Total files scanned:   42
```

It's also pretty fast. This is scanning my entire harddrive:
```
~/code/golang/src/github.com/tracehelms/scan_dir/ [master]$ time ./scandir -dir=/
Current working directory is: /Users/trace.helms/code/golang/src/github.com/tracehelms/scan_dir
Scanning directory: /
...
...
...
Total files scanned:   604211

real  1m0.409s
user  0m4.189s
sys 0m41.160s
```

Scanned my entire harddrive in a minute flat! I'm impressed.
