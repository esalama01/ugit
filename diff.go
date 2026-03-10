package main

import(
	"github.com/neticdk/go-stdlib/diff/myers"
	"os"
	"bytes"
	"compress/zlib"
	"strings"
	"io"
)

func Ugit_cat_file(path string) string{
	content, err := os.ReadFile(path)
	check(err)
	compressed := []byte(content)
	b := bytes.NewReader(compressed)
	r, err := zlib.NewReader(b)
	check(err)
	defer r.Close()
	var sb strings.Builder
	_, err = io.Copy(&sb, r)
	check(err)
	result := sb.String()
	return result
}