package main

import "testing"

func TestMemFsAddContentToFile(t *testing.T) {
	fs := NewMemFs()
	fs.addContentToFile("/root/dir1/file1.txt", []byte{'b', 'y', 't', 'e'})
	fs.addContentToFile("/root/dir1/file2.txt", []byte{'b', 'y', 't', 'e'})
}
