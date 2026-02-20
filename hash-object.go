package main

import(
	"os"
	"fmt"
	"log"
	"compress/zlib"
	"strings"
	"os/exec"
	"crypto/sha1"
	"path/filepath"
	"strconv"
	"bytes"
)
//------------------------------------------------------------------
type Objects interface {
	Ugit_hash_object() string
}
//------------------------------------------------------------------

func firstN(s string, n int) string { //a function that returns the first n characters of a string
	runes := []rune(s)
	 if n >= len(runes) {
        return s
    }
	return string(runes[:n])
}
func lastN(s string, n int) string{ //a function that returns the last n characters of a string
	runes := []rune(s)
	 if n >= len(runes) {
        return s
    }
	return string(runes[n:])
}



func Header(file *os.File) string{//must return a "blob__space_size of the contents of the file in bytes_\0"
	//i'll be using the exec.Command() function
	file_name := file.Name()
	cmd := exec.Command("wc", "-c", file_name)
	out, err := cmd.Output()
	check(err)

	size := string(out)
	var myInt int
	_, err = fmt.Sscanf(size, "%d", &myInt)
	check(err)
	
	
	s1 := "blob"
	s2 := strconv.Itoa(myInt)
	s3 := "\000" //null character
	result := s1 + " " + s2 + s3 //string concatenation
	return result
}

func Sha1(file *os.File) string{
	content, err := os.ReadFile(file.Name())
	check(err)

	header := Header(file)
	data := []byte(header + string(content))
	hash := sha1.Sum(data) //in [32]byte format
	return string(hash[:]) //converted hash to string
}

func Compression(file *os.File) string{
	data, err := os.ReadFile(file.Name())
	check(err)
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write([]byte(data))
	w.Close()
	return string(b.Bytes()[:])
}

func Get_Hash(f *os.File)(string, string){
	id := strings.TrimSpace(Sha1(f))
	return id, Compression(f)
}

func Ugit_hash_object(f *os.File){// when called i should call Get_Hash to generate the hash id  and the compressed content. eq to the git hash-object command
	val1, val2 := Get_Hash(f)
	fmt.Printf("%s%s",val1,val2)
}

func Ugit_hash_object_w(f *os.File){// when called i should call Get_Hash to generate the hash id  and the compressed content. eq to the git hash-object -w command
	val1, val2 := Get_Hash(f)
	b := Blob{
		Blob_ID : val1,
		Content : val2,
	}
	folder_name := firstN(b.Blob_ID, 2)
	file_name := lastN(b.Blob_ID, 38)
	// And then i should create the necessary repos to store the blob in
	//the folders's name is the first two chars of b.Blob_ID
	
	targetDir := ".ugit/objects" //go inside the .ugit/objects directory
	if err := os.Chdir(targetDir); err != nil {
		log.Fatalf("Error : %v\n", err)
	}

	err := os.MkdirAll(folder_name, 0755) //creating the blob's dirtectory
	check(err)
	
	// creating the blob's file
	fullPath := filepath.Join(folder_name, file_name)
	p, err := os.Create(fullPath)
	check(err)
	defer p.Close()
	
	//writing the compressed content to the file
	_, err = p.WriteString(b.Content)
	check(err)
	fmt.Println("Blob created succesfully.")
}