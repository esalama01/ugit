package main

import(
	"fmt"
)

func Header_commit(commit *Commit)[]byte{
	s1 := ""
}

func content(commit *Commit)[]byte{
	s1 := "tree" + " " + commit.parent_tree + "\n"
	var s2 string
	if len(commit.parents) >= 1{
		for _, entry := range commit.parents{
			data := "parent" + " " + entry + "\n"
			s2 += data
		}
	}else{
		s2 = ""
	}
	s3 := "author" + " " + commit.Author.Username + " " + "<"+ commit.Author.Useremail + ">"+ " " + commit.Author.EventDate.String() + "\n"
	s4 := "committer" + " " + commit.Author.Username + " " + "<"+ commit.Author.Useremail + ">" + " " + commit.Author.EventDate.String() + "\n"
	s5 := "\n"
	s6 := commit.message + "\n"
	var buffer []byte
	buffer = append(buffer, s1...)
	buffer = append(buffer, s2...)
	buffer = append(buffer, s3...)
	buffer = append(buffer, s4...)
	buffer = append(buffer, s5...)
	buffer = append(buffer, s6...)
	return buffer
}