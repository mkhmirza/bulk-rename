package main

/* 
	Bulk Renamer for renaming multiple files/folder using terminal 
*/

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

func checkMode(renameFolder bool, path, specExt string) string {
	mode := "";
	if renameFolder {
		mode = "Renaming Folders in " + path; 
	} else if  !renameFolder {
		mode = "Renaming Files in " + path;
		if len(specExt) > 0 {
			mode += " with " + specExt + " extension";
		}
	}
	return mode;
}

func rename(oldName string,path *string, newName, fullPath string, counter *int) {
	// attaching the full path with the new name
	newPath := *path + newName;
	// renaming the files 
	err := os.Rename(fullPath, newPath);
	checkError(err);
	*counter++;
	fmt.Printf("%s\t%s\n", oldName, newName);
}

func main() {

	// folder path to files that are going to be renamed
	path := flag.String("path", "~/Desktop", "folder path having files to rename");
	// pattern of naming
	pattern := flag.String("pattern", "", "set a pattern for renaming files");
	// if folders are to be renamed not files 
	renameFolder := flag.Bool("rename-folder", false, "rename folders only within a directory");
	// rename files with the given extension
	specExt := flag.String("extension", "", "rename files having passed extension");
	// give a starting value for the counter
	startingPoint := flag.Int("starting-point", 0, "starting value for the counter")
	// dry run option
	dryRun := flag.Bool("dry-run", false, "dry run on test/ folder");
	flag.Parse();
	

	// if dry run option is given then set the folder to test/
	if *dryRun {
		// set test/ as the dir
		*path = "./test/";
	}

	// if path is given '.', it means current directory
	if *path == "." {
		currentDir, err := os.Getwd();
		checkError(err);
		*path = currentDir + "/";
		//fmt.Printf("Current Directory: %v\n", currentDir);
	}

	// reading files and folders of the given path
	files, err := ioutil.ReadDir(*path)
	checkError(err);

	counter := *startingPoint;
	var fileSlice []string;
	
	fmt.Printf("Folder Path: %s\n", *path);
	mode := checkMode(*renameFolder, *path, *specExt);
	fmt.Printf("Mode: %v\n", mode);
	fmt.Println("File Name\tNew Name");
	fmt.Println("----------------------------------------");

	//loop through the files
	for _, f := range files {
		
		fullPath := *path + f.Name();

		ext := "";
		newName := "";

		// if 'f' option is not given it means
		// the current files should be renamed files with their original extension
		if !*renameFolder {

			// ignore directories within the folder
			if f.IsDir(){
				continue;
			}

			// getting the file extension
			ext =  "." + strings.Split(f.Name(), ".")[1];
			// constructing a new name for a file
			newName = *pattern + strconv.Itoa(counter) + ext;


			if len(*specExt) > 0 {
				// concat a '.' in front of the extension to match
 				currentExtension := "." + *specExt;
			
				// see if the current file extension is same as above file extension
				if ext == currentExtension {
					// rename a file having a specific extension
					rename(f.Name(), path, newName, fullPath, &counter);
					fileSlice = append(fileSlice, f.Name());
				}

			} else { // rename a all files in the folder
				rename(f.Name(), path, newName, fullPath, &counter);
				fileSlice = append(fileSlice, f.Name());
			}

		}  else if *renameFolder {
			// only rename the folders in the directory
			if f.IsDir() {

				// ignore normal files within the directory
				if !f.IsDir() {
					continue;
				}

				// constructing a new name for a folder
				newName = *pattern + strconv.Itoa(counter);
				// rename 
				rename(f.Name(), path, newName, fullPath, &counter);
				fileSlice = append(fileSlice, f.Name());
			}
		} 

	}


	fmt.Println("----------------------------------------");
	fmt.Printf("Total Files Renamed: %d\n", len(fileSlice));	
}
