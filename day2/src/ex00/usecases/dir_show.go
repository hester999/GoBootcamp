package usecases

import (
	"fmt"
	"os"
	"strings"
)

func DirSearcher(path string) {
	root := func(path string) string {
		tmp := strings.SplitN(path, "/", len(path)-1)
		str := "./"
		str += tmp[len(tmp)-1]
		return str
	}(path)

	dirSearcher(path, root)

}

func dirSearcher(path, root string) {

	dir, err := os.Open(path)
	if err != nil {
		fmt.Println("Dir open error:", err)
		return
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Dir read error:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {

			root += "/" + file.Name()
			fmt.Println(root)
			dirSearcher(path+"/"+file.Name(), root)
		}
	}
}
