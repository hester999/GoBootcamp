package entity

import (
	"errors"
	"flag"
	"fmt"
	"path/filepath"
)

type Flags struct {
	ShowSymlinks bool
	ShowDirs     bool
	ShowFiles    bool
	Ext          string
	Path         string
}

func (f *Flags) Parse() {
	flag.BoolVar(&f.ShowSymlinks, "sl", false, "Show only symbolic links")
	flag.BoolVar(&f.ShowDirs, "d", false, "Show only directories")
	flag.BoolVar(&f.ShowFiles, "f", false, "Show only files")
	flag.StringVar(&f.Ext, "ext", "", "File extension filter (requires -f)")
	flag.Parse()

	if flag.NArg() > 0 {
		f.Path = flag.Arg(flag.NArg() - 1)
	}
	err := f.checkFileFlag()
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	if (!f.ShowSymlinks && f.Ext == "") && !f.ShowDirs && !f.ShowFiles {
		f.ShowSymlinks = true
		f.ShowDirs = true
		f.ShowFiles = true
	}
}

func (f *Flags) GetAbsPath() string {
	abs, _ := filepath.Abs(f.Path)
	return abs
}

func (f *Flags) checkFileFlag() error {
	if f.Ext != "" {
		if !f.ShowFiles {
			f.Ext = ""
			return errors.New("flag -f is required")
		}
	}
	return nil
}
