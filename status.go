package main
import (
	"path/filepath"
	"fmt"
	"io/fs"
	"os"
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
/*
	finished the first implementation of my compare function but i ll need to heavily modify it. Don t forget to add the os.Stat() func to check if a file exists instead of all the slices shit.
*/

func Compare(area *StagingArea, paths []string) (map[string][]string){//a function to compare whatever s in the index file with my current working directory.
	m := make(map[string][]string)
	var exists []string //i ll need it to ppoint to removed files
	//i ll traverse the paths slice  and check if each path exists in my index
	for _, path := range paths { 
		f, err := os.Open(path) //open the file to pass it as a parameter in the Get_Hash function.
		check(err)
		value, ok := area.entries[path]
		if ok { //if it exists
			val1, _ := Get_Hash(f)
			f.Close()
			exists = append(exists, path)
			if val1 == value.Id{ //if they re the same
				m["2"] = append(m["2"],path)
			}else{
				m["3"] = append(m["3"],path)
			}
		}else{//if it doesn t exist
			m["1"] = append(m["1"],path)
		}
		
	}
	//now look for file that should be deleted
	for path := range area.entries{
		_, err := os.Stat(path) //look for files that should be removed
    	if err != nil {
        	m["4"] = append(m["4"], path) 
    	}
	}
	return m
}

func Ugit_status() {//crating the basic git status command
	area := StagingArea{
        entries: make(map[string]*Index),
    }
	Loadfromtheindex(&area)
	paths := traversal()
	status_map := Compare(&area, paths)
	fmt.Printf("On branch main\n") //to be modified later when implementing branching
	for key, paths := range status_map{
		switch key{
		case "1":
			fmt.Printf("Untracked files:\n	(use 'git add <file>...' to include in what will be committed)")
			for _, path := range paths{
				fmt.Printf("%s\n",path)
			}
		case "3","4":
			fmt.Printf("(use 'git restore <file>...' to discard changes in working directory)")
			if key == "3"{
				fmt.Printf("modified:\n")
				for _,path := range paths{
				fmt.Printf("%s\n",path)
				}
			}else{
				fmt.Printf("deleted:\n")
				for _,path := range paths{
				fmt.Printf("%s\n",path)
				}
			}
		default:

		}
	} 
}