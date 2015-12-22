package main
import (
	"fmt"
	"path/filepath"
	"os"
	"encoding/xml"
)

type FInfo struct {
	path string
	isDir bool
}

func main() {
	lst := make([]FInfo, 20)
	fmt.Println("Hello!!")
	filepath.Walk("/home/varun/Bin/idea/bin/", func( path string,
							 info os.FileInfo,
							 err error) error {
		if(err == nil) {
			lst = append(lst, FInfo{path, info.IsDir()})
		}
		return err
	})
	for info := range lst {
		output, err := xml.MarshalIndent(info, "	", "	")
		if err == nil {
			os.Stdout.Write(output)
		} else {
			fmt.Println("Error", err)
		}
	}
}
