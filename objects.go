import "encoding/json"

type Blob struct {
	Blob_ID string //sha1 hash key of this blob
	value string
}

type Tree struct {
	Tree_ID string //sha1 hash key of this tree
	blob map[string]*string //map between blob names and pointers to their ids inside this tree 
	subtree map[string]*string //maps between subtrees names and pointer of their id's inside this tree
}

type User struct {
	user_name string `json:"username"`
	user_email string `json:"email address"`
	EventDate time.Time `json:"event date"`
}

type Commit struct {
	Commit_ID string //sha1 hash key of this commit
	Author *User
	parent_tree *string //points to parent tree id
	parents  []*string //slice of pointer to parents commit id's
	message string
}