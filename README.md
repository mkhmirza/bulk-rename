# Bulk Renamer
Rename multiple files using terminal, based on a pattern.

## Usage
using **go**:
```
cd bulk-renamer
go build bulk-renamer.go
```
After building the binary run the following command for help
```
./bulk-renamer -h
```
For Renaming files use
```
./bulk-renamer -folder <path-to-folder> -pattern <pattern> 
```
Make sure to add the trailing `/` in the folder path for example:<br /> 
```
path/to/some/folder/ # right way
path/to/some/folder  # wrong way
```

### Renaming only folders within a DIR
To rename folders within a directory containing other files use 
```
./bulk-renamer <path-to-folder> -pattern <pattern> -f
```
`-f` flag makes sure that only folders within a directory are renamed.

### Renaming Files With A Specific Extension
To rename all files of a specific extension, run
```
./bulk-renamer -folder <path-to-folder> -pattern <pattern>  -extension <file-extension> 
```

### Starting Point for the Counter
Starting number for the counter can be changed as
```
./bulk-renamer -folder <path-to-folder> -starting-point <starting-number> 
```