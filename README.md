# Simple program to generate strong, random passwords.

This program uses the `github.com/sethvargo/go-password/password` library to generate strong passwords.

## Prequisites

[Golang](https://golang.org) - tested with Go version 1.11.

[github.com/sethvargo/go-password/password](https://godoc.org/github.com/sethvargo/go-password/password) package.
```
go get -u github.com/sethvargo/go-password/password
```

## Building

You may build this program by running:

```
go build
```

## Usage

```
$ password -h
Usage of password:
  -allow_repeat
    	Allow repeating characters (default: false)
  -digits int
    	number of digits (default 6)
  -length int
    	length of password (default 32)
  -no_upper
    	No uppercase characters (default: false)
  -symbols int
    	number of symbols (default 8)
```

### Examples:

```
$ password -length=30
ruPNDA81K^*eT!&t72j}b3(hiO6f?/
```

```
$ password -length=52 -symbols=7 -digits=9
HCl"D8VZ7^F9v<6rgtP1XoY|]sbGKEIwmhL2d)3S0JWxRz[Oyq5u
```

```
$ password -length=100 -symbols=7 -digits=25 -allow_repeat=true
hEjBSy6szx4r1`htfaxt*ay6F18ldGDvhRqXS2U4-m/I{3lA0PdzIC~4g9CoK17v4lcsvg3SY73L2CjgvHi8pYoo1L4lY32I=7LD
```

Note that the limitations you empose for the password must make sense. Otherwise, you may get an error such as the ones below.

```
$ password -length=100 -symbols=7
2019/10/02 01:11:12 number of letters exceeds available letters and repeats are not allowed
```

```
$ password -length=52 -symbols=7 -digits=17
2019/10/02 01:11:45 number of digits exceeds available digits and repeats are not allowed
```


