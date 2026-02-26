package main
import(
	"os"
	"strings"
)


func Tree_construction(entry []string, tree *Tree)(*Tree){ // a function that constructs a tree object
	//an entry example : [testing/mehdi/test.go]
	for _, name := range entry {
		fileinfo, err := os.Stat(name)
		check(err)
		if (!fileinfo.IsDir()){
			f, err := os.Open(name)
			check(err)
			defer f.Close()
			val1, val2 := Get_Hash_Blob(f)
			b := Blob{
			Blob_ID : val1,
			Content : val2,
			}
			tree.blob[name] = &b.Blob_ID
			return tree
		}
		else{
			//my recursion logic
		}
	}
}


//i'll be implementing the git write-tree command.
func Ugit_write_tree(){//the function takes as input the index file and reads it.

	//--------------Initializing an empty tree-------------------------------------------
	tree := new(Tree)
	//-----------------------------------------------------------------------------------

	//--------------Loading the staging area along with the paths slice------------------
	area := StagingArea{
        entries: make(map[string]*Index),
    }
	Loadfromtheindex(&area)
	paths := Traversal()
	//-----------------------------------------------------------------------------------

	//---------------slicing each path to construct a tree-------------------------------
	for _, path := range paths{
		entries := strings.Split(path, "/")

	}
	//-----------------------------------------------------------------------------------
}