package options

import (
	"flag"
	"fmt"
	"os"
	"path"
)

type Options struct {
	Dispersion string
	File string // empty for stdin
	NColors int
	Verbose bool
	Debug bool
	Colorise bool
	Format string
}

const (
	DISPERSION_SPREAD = "spread"
	DISPERSION_VARIANCE = "variance"
	FORMAT_X = "x" // 0090c0
	FORMAT_D = "d" // 55,255,0
	FORMAT_F = "f" // 0.00 0.01 1.0
)

func CliOptions() Options {
	options := Options{ }

	flag.Usage = func() {
		fmt.Printf("\n")
		fmt.Printf("Usage:\n",)
		fmt.Printf("\t%s [-v] [-x] [-c] [-n colors] [-f format] [-d dispersion-measurent] file \n", path.Base(os.Args[0]))
		fmt.Printf("\n")
		flag.PrintDefaults()
	}

	flag.BoolVar(   & options.Verbose,  "v", false, "be verbose.")
	flag.BoolVar(   & options.Debug,    "x", false, "print debug output.  implies -v.")
	flag.BoolVar(   & options.Colorise, "c", false, "colorise output (requires RGB terminal).")
	flag.IntVar(    & options.NColors,  "n", 16,    "number of colors.  defaults to 16.")
	flag.StringVar( & options.Format,
	                "f",
	                FORMAT_X,
	                fmt.Sprintf( "format for color output.  possible values: \"%s\", \"%s\" or \"%s\"",
	                             FORMAT_X, FORMAT_D, FORMAT_F ))

	flag.StringVar( & options.Dispersion,
	                "d",
	                DISPERSION_VARIANCE,
	                fmt.Sprintf( "dispersion measurement for selecting split axis.  possible values: \"%s\", \"%s\"",
	                             DISPERSION_VARIANCE,
	                             DISPERSION_SPREAD ))

	flag.Parse()

	options.File = flag.Arg(0)

	// validate

	if flag.NArg() > 1 {
		fmt.Printf("too many arguments: %v\n", flag.Args())
		flag.Usage()
		os.Exit(1)
	}

	if options.Dispersion != DISPERSION_VARIANCE && options.Dispersion != DISPERSION_SPREAD {
		fmt.Printf("unsupported dispersion measurement: %v\n", options.Dispersion)
		flag.Usage()
		os.Exit(1)
	}

	if options.Format != FORMAT_X && options.Format != FORMAT_D && options.Format != FORMAT_F {
		fmt.Printf("unsupported color format: %v\n", options.Format)
		flag.Usage()
		os.Exit(1)
	}

	// post-process

	if options.Debug { options.Verbose = true }

	return options
}
