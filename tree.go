package main
import(
	"os"
	"strings"
	"path/filepath"
)
//i will construct the tree first ny assigning names to it and adding subtrees and blobs relatiopns by using the trie ds, and after that i will make the tree_ID by using a traversal algorithm.(Merkel)

type TrieNode struct {//i'll define a trie data structure.
	Name	string
	Children	map[string]*TrieNode //mapping between direcotiries' and files' names and their node struct.
	Is_blob	bool
}

type Trie struct {
	root	*TrieNode
}

// Constructor for TrieNode
func NewTrieNode(name string) *TrieNode {
	node := new(TrieNode)
	node.Name = name
	node.Children = make(map[string]*TrieNode)
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
			node.Children[entry[i]] = NewTrieNode(entry[i])
			node = node.Children[entry[i]]
		}
	}
	node.Is_blob = true //the end or the path is always a blob.
}

func Post_order_trav(node *TrieNode, t *Tree)(*Tree){//a function that takes a trie node as input and returns it's hash value 
	//base case
	if node.Is_blob{
		t.Blob[node.Name] = //a function that computes the hash value for the blob
	}else{
		new_tree := new(Tree)
		for _, val := range node.Children{
			new_tree[val.Name] = //a function that computes the hash value for the tree
		}
	}
	return t
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

	//--------------constructing an empty tree-------------------------------------------
	my_tree := new(Tree)
	my_tree.Blob = make(map[string]string)
	my_tree.Subtree = make(map[string]string)
	//-----------------------------------------------------------------------------------

	//--------------constructing an empty trie-------------------------------------------
	my_trie := NewTrie(directory_name())
	//-----------------------------------------------------------------------------------

	//--------------Loading the staging area along with the paths slice------------------
	area := StagingArea{
        entries: make(map[string]*Index),
    }
	Loadfromtheindex(&area)
	paths := Traversal()
	//-----------------------------------------------------------------------------------

	//---------------slicing each path to fill the trie----------------------------------
	for _, path := range paths{
		entry := strings.Split(path, "/")
		Tree_construction(entry , my_trie)
	}
	//-----------------------------------------------------------------------------------

	//----------------Constructing The Tree----------------------------------------------

	//-----------------------------------------------------------------------------------
}