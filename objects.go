package main
import (
	//"encoding/json"
	"time"
)

type Blob struct {
	Blob_ID string //sha1 hash key of this blob
	Content string
}

type User struct {
	Username string `json:"user_name"`
	Useremail string `json:"email_address"`
	EventDate time.Time `json:"event_date"`
}

type Commit struct {
	Commit_ID string //sha1 hash key of this commit
	Author *User
	parent_tree *string //points to parent tree's id
	parents  []*string //slice of pointers to parents commit's ids
	message string
}
