package usecases

import (
	"fmt"
	"src/src/ex00/entity"
)

type FinderEx00 struct{}

func (f *FinderEx00) Process(flags interface{}) error {
	finderFlags, ok := flags.(*entity.Flags)
	if !ok {
		return fmt.Errorf("invalid flags type for FinderEx00")
	}

	if finderFlags.ShowFiles && !finderFlags.ShowSymlinks && !finderFlags.ShowDirs {

		FileSearcher(finderFlags.Path, finderFlags.Ext)
	}
	if finderFlags.ShowDirs && !finderFlags.ShowFiles && !finderFlags.ShowSymlinks {
		DirSearcher(finderFlags.Path)
	}
	if !finderFlags.ShowDirs && !finderFlags.ShowFiles && finderFlags.ShowSymlinks {
		ShowLink(finderFlags.Path)
	}

	if finderFlags.ShowDirs && finderFlags.ShowFiles && finderFlags.ShowSymlinks {
		Searcher(finderFlags.Path)
		LinkSearcher(finderFlags.Path)
	}

	return nil
}
