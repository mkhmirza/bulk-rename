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
	"path/filepath"
)

const (
	colorRed = "\033[31m";
	colorReset = "\033[0m";
)

func checkError(err error) {
	if err != nil {
		panic(err)
	} 
}

func getFileExtension(filename string) (bool, string) {
	var extension = filepath.Ext(filename)
	if extension == ""{
		return false, "nil";
	}
	return true, extension;
}

func checkMode(renameFolder bool, colorReset, path, specExt string) string {
	mode := "";
	redPath := fmt.Sprint(string(colorRed), path);
	if renameFolder {
		mode = "Renaming Folders in " + redPath; 
	} else if  !renameFolder {
		mode = "Renaming Files in " + redPath + fmt.Sprint(string(colorReset), "");
		if len(specExt) > 0 {
			ext := fmt.Sprint(string(colorRed), specExt)
			mode +=  " with " + ext + fmt.Sprint(string(colorReset), " extension.");
		}
	}
	return mode;
}

func rename(path *string, newName, fullPath string, counter *int) {
	
	// attaching the full path with the new name
	newPath := *path + newName;
	// renaming the files 
	err := os.Rename(fullPath, newPath);
	checkError(err);
	*counter++;
}

func main() {
	fmt.Println();

	// folder path to files that are going to be renamed
	path := flag.String("path", "~/Desktop", "folder path having files to rename");
	// pattern of naming
	pattern := flag.String("pattern", "", "set a pattern for renaming files");
	// if folders are to be renamed not files 
	renameFolder := flag.Bool("rfolder", false, "rename folders only within a directory");
	// rename files with the given extension
	specExt := flag.String("extension", "", "rename files having passed extension");
	// give a starting value for the counter
	startingPoint := flag.Int("spoint", 0, "starting value for the counter")
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
	}

	temp := *path
	// get the last character
	trailing := string(temp[len(*path) - 1])
	if !(trailing == "/") {
		// add a '/' to not cause a exception
		*path = *path + "/";
	}

	// adding color to path
	pathToPrintRed := *path;
	pathToPrintRed = fmt.Sprint(string(colorRed), pathToPrintRed);

	// reading files and folders of the given path
	files, err := ioutil.ReadDir(*path)
	checkError(err);

	counter := *startingPoint;
	var fileSlice []string;
	
	space := fmt.Sprint(string(colorReset), " ");
	// fmt.Printf("Folder Path: %s %v\n", pathToPrintRed, space);
	mode := checkMode(*renameFolder, colorReset, *path, *specExt);

	fmt.Printf("Mode: %v %v\n", mode, space);
	fmt.Println();

	fmt.Println("File Name\tNew Name");
	fmt.Println("----------------------------------------");

	//loop through the files
	for _, f := range files {
		
		fullPath := *path + f.Name();

		ext := "";
		newName := "";
		isFileExtension := false;

		// if 'f' option is not given it means
		// the current files should be renamed files with their original extension
		if !*renameFolder {

			// ignore directories within the folder
			if f.IsDir(){
				continue;
			}

			// getting the file extension
			isFileExtension, ext = getFileExtension(f.Name());
			// if file extension is false (not given)
			if !(isFileExtension) {
				continue;
			}

			// constructing a new name for a file
			newName = *pattern + strconv.Itoa(counter) + ext;

			if len(*specExt) > 0 {
				// concat a '.' in front of the extension to match
 				currentExtension := "." + *specExt;
			
				// see if the current file extension is same as above file extension
				if ext == currentExtension {
					// rename a file having a specific extension
					rename(path, newName, fullPath, &counter);
					fileSlice = append(fileSlice, f.Name());
				} else {
					// ignore this file, and move on
					continue;
				}

			} else { // rename a all files in the folder
				rename(path, newName, fullPath, &counter);
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
				rename(path, newName, fullPath, &counter);
				fileSlice = append(fileSlice, f.Name());
			}
		} 
		
		oldName := f.Name();
		fmt.Printf("%s\t%s\n", oldName, newName);
	}


	fmt.Println("----------------------------------------");
	fmt.Printf("Total Files Renamed: %d\n", len(fileSlice));
	fmt.Println();	
}
