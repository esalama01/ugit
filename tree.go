package main
import(
	"compress/zlib"
	"os"
	"strings"
	"path/filepath"
	"slices"
	"cmp"
	"crypto/sha1"
	"encoding/hex"
	"strconv"
	"bytes"
)

// Constructor for TreeNode
func NewTreeNode(name string) *TreeNode {
	node := new(TreeNode)
	node.Name = name
	node.Children = make(map[string]*TreeNode)
	return node
}

// Constructor for Tree
func NewTree(root_name string) *Tree {
	return &Tree{root: NewTreeNode(root_name)}
}


func Tree_construction(entry []string, t *Tree){ //a function to implement a word into the Tree
	node := t.root //a pointer to the root of the Tree
	for i := 0; i < len(entry); i++{ //for each file or directory in the path slice
		if _, ok := node.Children[entry[i]]; ok{ //check if entry[i] exists in the current node children
			//if yes, move to the corresponding child node.
			node = node.Children[entry[i]]
		 }else{//If it doesn't exist, create a new node for the entry[i] and link it to the current node.
			node.Children[entry[i]] = NewTreeNode(entry[i])
			node = node.Children[entry[i]]
		}
	}
	node.Is_blob = true //the end or the path is always a blob.
}

type Cin struct{
	sha1_hash string
	name string
	mode string
}

func Sha1_file(name string)string{
	file, err := os.Open(name)
	check(err)
	val,_ := Get_Hash_Blob(file)
	Ugit_hash_object_w(file)
	return val
}

func Header_tree(bita9a *Cin)[]byte{
	var buffer []byte
	s1 := bita9a.mode
	s2 := bita9a.name
	s3 := "\000"
	data, err := hex.DecodeString(bita9a.sha1_hash)
	check(err)
	s4 := data
	buffer = append(buffer, s1...)
	buffer = append(buffer, ' ')
	buffer = append(buffer, s2...)
	buffer = append(buffer, s3...)
	buffer = append(buffer, s4...)
	return buffer
}

func Sha1_tree(my_list []*Cin)string{
	slices.SortFunc(my_list, func(a, b *Cin) int { //sorting the structs of my_list by their names.
			return cmp.Compare(a.name, b.name)
	})
	var conc []byte
	for _, entry := range my_list{
		header := Header_tree(entry)
		conc = append(conc, header...) 
	}
	my_header := "tree" + " " + strconv.Itoa(len(conc)) + "\000"
	data := append([]byte(my_header), conc...)
	hash := sha1.Sum(data)
	str := hex.EncodeToString(hash[:]) //converted hash to string
	folder_name := FirstN(string(str), 2)
	file_name := LastN(string(str), 2)

	objDir := filepath.Join(".ugit", "objects", folder_name)
	objPath := filepath.Join(objDir, file_name)
	os.MkdirAll(objDir, 0755)
	compressed := Compression_tree(data)
	os.WriteFile(objPath, compressed, 0444)
	return string(str) //converted hash to string
}

func Compression_tree(data []byte)[]byte{//compressing the headered data to store as data into the objects directory
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write(data)
	w.Close()
	return b.Bytes()
}


func Post_order_trav(node *TreeNode,prefix string)(string){//a function that takes a Tree node as input and returns it's hash value 
	//base case
	path := filepath.Join(prefix, node.Name)
	if node.Is_blob{
		c := Cin{name : node.Name, mode : "100644", sha1_hash : Sha1_file(path)}
		return c.sha1_hash
	}else{
		var my_list []*Cin
		for _, sub_node := range node.Children{
			my_list = append(my_list, &Cin{
  				name:      sub_node.Name,
    			mode:      "04000",
    			sha1_hash: Post_order_trav(sub_node, path),
			})
		}
		//now i ll begin the logic for building the hash out of my_list
		sha1 := Sha1_tree(my_list)
		return sha1
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

	//--------------constructing an empty tree-------------------------------------------
	
	//-----------------------------------------------------------------------------------


	//--------------constructing an empty Tree-------------------------------------------
	my_Tree := NewTree("") 
	//-----------------------------------------------------------------------------------

	//--------------Loading the staging area along with the paths slice------------------
	area := StagingArea{
        entries: make(map[string]*Index),
    }
	Loadfromtheindex(&area)
	paths := Traversal()
	//-----------------------------------------------------------------------------------

	//---------------slicing each path to fill the Tree----------------------------------
	for _, path := range paths{
		entry := strings.Split(path, "/")
		Tree_construction(entry , my_Tree)
	}
	//-----------------------------------------------------------------------------------

	//----------------Constructing The Tree----------------------------------------------
	root_hash := Post_order_trav(my_Tree.root, "")
	//-----------------------------------------------------------------------------------
}