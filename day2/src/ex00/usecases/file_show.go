package usecases

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"src/src/ex00/utils"
)

func FileSearcher(path string, ext string) {

	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if !info.IsDir() {
			if ext == "" {
				fmt.Println(path)
			} else {
				walkerExt := filepath.Ext(path)
				walkerExt = utils.FinderPoint(walkerExt)
				ext = utils.FinderPoint(ext)

				if ext == walkerExt {
					fmt.Println(path)
				}
			}

		}
		return err
	})

	if err != nil {
		fmt.Println(err)
	}
}
