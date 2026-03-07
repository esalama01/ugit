package main

import(
	"os"
	"strconv"
	"crypto/sha1"
	"encoding/hex"
	"compress/zlib"
	"bytes"
	"path/filepath"
	"fmt"
	"time"
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

func Sha1_commit(commit *Commit)string{//generating the sha-1 commit id.
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

func insert_into_db(my_hash string, data []byte){//to insert the commit int the objects db
	folder_name := FirstN(string(my_hash), 2)
	file_name := LastN(string(my_hash), 2)

	objDir := filepath.Join(".ugit", "objects", folder_name)
	objPath := filepath.Join(objDir, file_name)
	os.MkdirAll(objDir, 0755)
	compressed := Compression_commit(data)
	os.WriteFile(objPath, compressed, 0444)
}

func get_info()*User{//to retrieve the user infos
	var username string
	var usermail string
	fmt.Print("Enter your UserName: ")
	fmt.Scanln(&username)
	fmt.Print("Enter your email address: ")
	fmt.Scanln(&usermail)
	new_user := User{
		Username : username,
		Useremail : usermail,
		EventDate : time.Now(),
	}
	return &new_user
}

func get_parent()(string,int){// a function to return the sha-1 id of the parent commit
	data, err := os.ReadFile(".ugit/refs/heads/main") //read the refs/heads/main 
	if err != nil {
		return "", 0
	}
	return string(data), 1
}

func ugit_commit(input_message string){
	var parents_list []string
	parent_tree_id := Ugit_write_tree()
	user_info := get_info()
	parent, exists := get_parent()
	if exists == 1{
		parents_list = append(parents_list, parent)
	}else{
		//i must first create the main file
		cont := []byte("")
		objDir := filepath.Join(".ugit", "refs", "heads")
		objPath := filepath.Join(objDir, "main")
		os.WriteFile(objPath, cont, 0644)
	}
	new_commit := Commit{
		Author : user_info,
		parent_tree : parent_tree_id,
		parents : parents_list,
		message : input_message,
	}
	id := Sha1_commit(&new_commit)
	new_commit.Commit_ID = id
	data := get_data(&new_commit)
	insert_into_db(id, data)
//-------commit object created succesfully.-----------
	//writing to the refs/heads/main
	objDir := filepath.Join(".ugit", "refs", "heads")
	objPath := filepath.Join(objDir, "main")
	err := os.Truncate(objPath, 0) //deleting it's content first.
	check(err)
	content := []byte(id)
	os.WriteFile(objPath, content, 0644)
}