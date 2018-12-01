package dir_walker
import (
	"fmt"
	"path/filepath"
	"os"
	"encoding/xml"
)

type MyFile struct {
	XMLName xml.Name `xml:"file"`
	Path string `xml:"path,attr"`
	IsDir bool `xml:"isDir,attr"`
}

func Run_DirWalker() {
	lst := make([]MyFile, 20)
	fmt.Println("Hello!!")
	filepath.Walk(`C:\Program Files (x86)\chaiscript 5.7.1\share\chaiscript\samples`, func( pth string,
							 info os.FileInfo,
							 err error) error {
		if(err == nil) {
			f := MyFile{ Path:pth, IsDir:info.IsDir() }
			lst = append(lst, f)
			output, err := xml.Marshal(f)
			if err == nil {
				os.Stdout.Write(output)
				fmt.Println()
			} else {
				fmt.Println("Error", err)
			}
		}
		return err
	})
}
