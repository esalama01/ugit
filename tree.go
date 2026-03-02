package main
import(
	"os"
	"strings"
	"path/filepath"
)
//i will construct the tree first ny assigning names to it and adding subtrees and blobs relatiopns by using the trie ds, and after that i will make the tree_ID by using a traversal algorithm.(Merkel)

type TrieNode struct {//i'll define a trie data structure.
	Name	string
	Children	map[string]*TrieNode //mapping between 
	Is_blob	bool
}

type Trie struct {
	root	*TrieNode
}

// Constructor for TrieNode
func NewTrieNode(name string) *TrieNode {
	node := new(TrieNode)
	node.Name = name
	return node
}

// Constructor for Trie
func NewTrie(root_name string) *Trie {
	return &Trie{root: NewTrieNode(root_name)}
}


func Tree_construction(entry []string, t *Trie){ //a function to implement a word into the trie
	node := t.root //a pointer to the root of the trie
	for i := 0; i < len(entry); i++{ //for each file or directory in the path slice
		if _, ok := node.Children[entry[i]]; ok{ //check if entry[i] exists in the current node children
			//if yes, move to the corresponding child node.
			node = node.Children[entry[i]]
		 }else{//If it doesn't exist, create a new node for the entry[i] and link it to the current node.
			node = NewTrieNode(entry[i])
		}
		node.Is_blob = true //the end or the path is always a blob.
	}
}

func directory_name()(string){
	dir, err := os.Getwd()
	check(err)

	// Extract just the name of the directory from the full path
	dirName := filepath.Base(dir)
	return dirName

}

//i'll be implementing the git write-tree command.
func Ugit_write_tree(){//the function takes as input the index file and reads it.

	//--------------Initializing an empty tree-------------------------------------------
	my_tree := new(Tree)
	my_trie := NewTrie(directory_name())
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
		entry := strings.Split(path, "/")
		Tree_construction(entry , my_trie)
	}
	//-----------------------------------------------------------------------------------
}