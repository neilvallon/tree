package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	path := "./"

	tree, err := NewTree(path)
	if err != nil {
		log.Fatalln(err)
	}

	tree.Print("")
}

type tree struct {
	name  string
	files []os.FileInfo
	dirs  []tree
}

func NewTree(root string) (tree, error) {
	item, err := os.Stat(root)
	if err != nil {
		return tree{}, err
	}

	return newTree("", item)
}

func newTree(path string, dir os.FileInfo) (tree, error) {
	prefix := path + dir.Name() + "/"

	list, err := ioutil.ReadDir(prefix)
	if err != nil {
		return tree{}, err
	}

	parent := tree{name: dir.Name()}
	for _, child := range list {
		if child.IsDir() {
			t, err := newTree(prefix, child)
			if err != nil {
				return tree{}, err
			}
			parent.dirs = append(parent.dirs, t)
		} else {
			parent.files = append(parent.files, child)
		}
	}

	return parent, nil
}

func (t tree) Print(prefix string) {
	for _, child := range t.dirs {
		fmt.Println(prefix+"|", child.name)
		child.Print(prefix + "  ")
	}
	for _, child := range t.files {
		fmt.Println(prefix+"-", child.Name())
	}
}
