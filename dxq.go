package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func getcurrentdirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	return dir
}


func main() {
	path := getcurrentdirectory()
	fmt.Println(path)
}
