# palette

A command line utility that extracts a color palette from an image.

License: [0BSD](https://spdx.org/licenses/0BSD.html).

## install

```
go get github.com/giorgiga/palette
go build github.com/giorgiga/palette
```

I think :)

## usage

```
Usage:
	palette [-v] [-x] [-c] [-n colors] [-f format] [-d dispersion-measurent] file

  -c	colorise output (requires RGB terminal).
  -d string
    	dispersion measurement for selecting split axis.  possible values: "variance", "spread" (default "variance")
  -f string
    	format for color output.  possible values: "x", "d" or "f" (default "x")
  -n int
    	number of colors.  defaults to 16. (default 16)
  -v	be verbose.
  -x	print debug output.  implies -v.
```

## contributing

I made this utility mainly to practice go and not only it's my first go program, it's also my first "desktop" (ie: non-server) program.

All help is greatly appreciated.
