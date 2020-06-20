# filedivdir

filedivdir divides the files in the base directory by creating a directory for each file timestamp.

## Installation

When you have installed the Go, Please execute following `go get` command:

```sh
go get -u github.com/qt-luigi/filedivdir
cd filedivdir/cmd
go build -o filedivdir
```

## Usage

```sh
$ ./filedivdir
filedivdir divides the files in the base directory by creating a directory for each file timestamp.

Usage:

	filedivdir [argument]

The argument is:

	basedir    the base directory that contains the files to be divide.
```

## Execution samples

```sh
$ ls -l sample/
-rwxrwxrwx  1 user   staff    1024 10 28  2019 a.txt
-rwxrwxrwx  1 user   staff    3072  4  5 14:29 b.txt
-rwxrwxrwx  1 user   staff    4096  6 10 23:45 c.txt

$ filedivdir sample/

$ ls -l sample/
drwxr-xr-x    3 user   staff   128  6 20 21:14 20191028
drwxr-xr-x    3 user   staff   768  6 20 21:14 20200405
drwxr-xr-x    3 user   staff   768  6 20 21:14 20200610

$ ls -l sample/20191028/
-rwxrwxrwx@ 1 user   staff    1024 10 28  2019 a.txt
$ ls -l sample/20200405/
-rwxrwxrwx@ 1 user   staff    2048  4  5 14:29 b.txt
$ ls -l sample/20200610/
-rwxrwxrwx@ 1 user   staff    4096  6 10 23:45 c.txt
```

## License

MIT

## Author

Ryuji Iwata

## Note

This tool is mainly using by myself. :-)
