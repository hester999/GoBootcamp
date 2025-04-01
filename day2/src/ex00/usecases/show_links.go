package usecases

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func ShowLink(path string) {

	link, err := os.Readlink(path)

	if err != nil {
		fmt.Printf("%s -> [broken]\n", path)
		return
	}
	fmt.Printf("%s -> %s\n", path, link)
}

//func ShowLinksRecursive(path string) {
//	info, _ := os.Lstat(name)
//
//	fmt.Println(info.Mode()&os.ModeSymlink != 0)
//
//}

func LinkSearcher(path string) {

	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if !info.IsDir() {
			if info.Mode()&os.ModeSymlink == os.ModeSymlink {
				ShowLink(path)
			}
		}
		return err
	})

	if err != nil {
		fmt.Println(err)
	}
}
