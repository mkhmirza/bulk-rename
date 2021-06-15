package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// folder path to files that are going to be rename
	folderPath := flag.String("folder", "~/Desktop", "folder path having files to rename");
	// pattern of naming
	pattern := flag.String("pattern", "pic", "set a pattern for renaming files",)
	flag.Parse();
	
	// reading files and folders of the given path
	files, err := ioutil.ReadDir(*folderPath)
	checkError(err);

	counter := 0;
	var fileSlice []string;
	
	fmt.Printf("Folder Path: %s\n", *folderPath);
	fmt.Println("File Name\tNew Name");
	//loop through the files
	for _, f := range files {
		fullPath := *folderPath + f.Name();
		// getting the file extension
		ext :=  "." + strings.Split(f.Name(), ".")[1];
		// constructing a new name
		newName := *pattern + strconv.Itoa(counter) + ext;
		// attaching the full path with the new name
		newPath := *folderPath + newName;
		// renaming the files 
		err = os.Rename(fullPath, newPath);
		checkError(err);
		counter++;

		fmt.Printf("%s\t%s\n", f.Name(), newName)
		fileSlice = append(fileSlice, f.Name());
	}

	fmt.Println("----------------------------------------");
	fmt.Printf("Total Files Renamed: %d\n", len(fileSlice));	
}
