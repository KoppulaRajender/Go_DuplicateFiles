package main

import (
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var directory string = os.Args[1] //Taking directory as an argument for checking duplicate file at root level

var filesMap = make(map[[sha512.Size]byte][]string) //Map to store SHA hash as keys and and duplicate file names as array of strings

//main function
func main() {
	log.SetFlags(log.Lshortfile)
	err := filepath.Walk(directory, scanDuplicates)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, value := range filesMap {
		if len(value) > 1 {
			fmt.Printf(" %v\n", value)
		}
	}
}

//scanDuplicates function
func scanDuplicates(path string, f os.FileInfo, err error) error {
	folder := filepath.Dir(path)
	if err != nil {
		log.Print(err)
		return nil
	}

	if f.IsDir() {
		return nil
	}

	if folder == directory { //TO check dulicate files at root level and skipping all subdirectories
		data, err := ioutil.ReadFile(path)
		if err != nil {
			log.Print(err)
			return nil
		}
		digest := sha512.Sum512(data)

		//fmt.Printf("%s\n", data) --> To displat the data of the files
		//fmt.Printf("%s\n", digest) --> To display the Digest value of the file's content

		//Filling map with duplicate files checking against common Digest
		if v, ok := filesMap[digest]; ok {
			filesMap[digest] = append(v, f.Name())
		} else {
			filesMap[digest] = append(v, f.Name())
		}

	}
	return nil

}
