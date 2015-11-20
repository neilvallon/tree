package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	path := "./"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalln(err)
	}

	printFiles(path, "", files)
}

func printFiles(path, prefix string, files []os.FileInfo) {
	for _, f := range files {
		// path
		if f.IsDir() {
			p := path + "/" + f.Name()
			files, err := ioutil.ReadDir(p)

			if err != nil {
				log.Fatalln(err)
			}

			printFiles(p, prefix+"-", files)
			continue
		}

		// file
		fmt.Println(prefix, f.Name())
	}
}
