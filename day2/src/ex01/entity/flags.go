package entity

import (
	"errors"
	"flag"
)

type Flags struct {
	LineCount bool
	CharCount bool
	WordCount bool
	Path      []string
}

//
//func (f *Flags) GetAbsPath() string {
//	abs, _ := filepath.Abs(f.Path)
//	return abs
//}

func (f *Flags) Parse() error {
	flag.BoolVar(&f.LineCount, "l", false, "Line count")
	flag.BoolVar(&f.CharCount, "m", false, "Char count")
	flag.BoolVar(&f.WordCount, "w", false, "Word count")

	flag.Parse()

	if flag.NArg() > 0 {
		f.Path = flag.Args()
	} else {
		return errors.New("no args")

	}

	err := f.checkTrue()

	if err != nil {
		return err
	}
	return nil
}

func (f *Flags) checkTrue() error {
	tCount := 0

	if f.CharCount {
		tCount++
	}
	if f.WordCount {
		tCount++
	}

	if f.LineCount {
		tCount++
	}

	if tCount != 1 {
		return errors.New("only one flag is allowed")
	}
	return nil
}
