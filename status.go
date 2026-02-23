package main
import (
	"path/filepath"
	"fmt"
	"io/fs"
	"os"
)
func difference(a, b []string) []string {// a functionn to calculate the difference between two slices of strings.
    mb := make(map[string]struct{}, len(b))
    for _, x := range b {
        mb[x] = struct{}{}
    }
    var diff []string
    for _, x := range a {
        if _, found := mb[x]; !found {
            diff = append(diff, x)
        }
    }
    return diff
}

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

func compare(area *StagingArea, paths []string) (int, string){//a function to compare whatever s in the index file with my current working directory.
	
	var exists []string //i ll need it to ppoint to removed files
	//i ll traverse the paths slice  and check if each path exists in my index
	for _, path := range paths {
		f, err := os.Open(path) //open the file to pass it as a parameter in the Get_Hash function.
		check(err)
		defer f.Close()
		value, ok := area.entries[path]
		if ok { //if it exists
			val1, _ := Get_Hash(f)
			leftovers = append(leftovers, path)
			if val1 == value.Id{ //if they re the same
				return 2, path
			}else{
				return 3,path
			}
		}else{//if it doesn t eexist
			return 1, path
		}
	}
	//now i ll look for files that should be removed
	keys := make([]int, len(area.entries)) //a slice containig the keys from my map.
	i := 0
	for k := range area.entries {
    keys[i] = k
    i++
	}
	//i ll substract 
}