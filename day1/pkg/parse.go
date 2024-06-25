package pkg

import (
	"flag"
	"fmt"
	"os"
)

type Flag struct {
	ReadFlag bool
}

func (f *Flag) parseFlag(fs *flag.FlagSet) error {

	fs.BoolVar(&f.ReadFlag, "f", false, "Read DataBase file")

	err := fs.Parse(os.Args[1:])

	if err != nil {
		fmt.Errorf("flag parsing error: %s", err)
		return err
	}
	return nil

}

func (f *Flag) ParseCommandLine() string {
	var path string
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	err := f.parseFlag(fs)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", "Invalid command line flag")
		os.Exit(1)
	}

	if f.ReadFlag {
		if fs.NArg() > 0 {
			path = fs.Arg(0)
		} else {
			fmt.Fprintf(os.Stderr, "Empty file name")
		}
	} else {
		fmt.Println("Read flag not set.")
	}
	return path
}
