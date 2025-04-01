package usecases

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func Searcher(path string) {
	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {

		if err != nil {
			return err
		}

		fmt.Println(path)
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}

//func Searcher(path string) {
//	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
//		if err != nil {
//			return err
//		}
//
//		if d.IsDir() {
//			fmt.Println("[DIR] ", path)
//		} else {
//			fmt.Println(" [FILE] ", path)
//		}
//
//		return nil
//	})
//
//	if err != nil {
//		fmt.Println("Ошибка:", err)
//	}
//}
