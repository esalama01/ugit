package main
import (
	"os"
	"path/filepath"
	"fmt"
	"io/fs"
)
/*
	i will be implementing the git status command. But how?
	I ll have to scan the working directory and the stgaingarea and make a comparison between the files.
	The procedure is as follows:
		1- scan both the working directory (ignore the ugit command) and the staging area.
		2- make a comparison by using the sha1 checksum alg Or by comparing their metadata(for speed reasons).
		3.1- if a file exists on the working directory but not on the staging area it should be labeled as Untracked.
		3.2- If a file exists on the staging area but not on the orking directory then it's been removed. I must label it as deleted.
		3.3- if a file exists on both of them, then kmel 7yatek.
*/


func traversal()([]string){ // a function that stores the path of each file in my directory on a slice. 
	var paths []string //a slice to store the paths in
	root := "." //the current directory
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
            return err
        }
		if d.IsDir() {
    		if d.Name() == ".ugit" {//skipping the ugit directory
            	return filepath.SkipDir
        	}
    	}
        if !d.IsDir() {
            paths = append(paths,path)
        }
        return nil
	})
	if err != nil {
        fmt.Printf("error walking the path %q: %v\n", root, err)
    }
	return paths
}

func compare(area *StagingArea) {
	
}