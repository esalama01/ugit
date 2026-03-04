package main

import(
	"fmt"
	"strconv"
	"crypto/sha1"
	"encoding/hex"
	"compress/zlib"
	"bytes"
)

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

func Header_commit(commit *Commit)[]byte{
	s1 := "commit"
	s2 := " "
	s3 := strconv.Itoa(len(content(commit)))
	s4 := "\000"
	var buffer []byte
	buffer = append(buffer, s1...)
	buffer = append(buffer, s2...)
	buffer = append(buffer, s3...)
	buffer = append(buffer, s4...)
	return buffer
}
func get_data(commit *Commit)[]byte{
	var data []byte
	data = append(data, Header_commit(commit)...)
	data = append(data, content(commit)...)
	return data
}
func Sha1_commit(commit *Commit)string{
	data := get_data(commit)
	hash := sha1.Sum(data)
	str := hex.EncodeToString(hash[:])
	return str
}

func Compression_commit(data []byte)[]byte{//compressing the headered data to store as data into the objects directory
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write(data)
	w.Close()
	return b.Bytes()
}

func ugit_commit(message string){

}