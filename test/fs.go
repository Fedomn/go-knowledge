package main

import "strings"

// 内存文件系统实现：
// addContentToFile
// addContentToFile(path, content)
// ls(path)
// makedir(path, name)
func main() {
	// /root/fs1/1.txt
	// /root/fs2/1.txt
}

type MemFs struct {
	root *MemFsNode
}

func NewMemFs() *MemFs {
	return &MemFs{root: &MemFsNode{
		prefix:  "/",
		isDir:   true,
		content: nil,
		nodes:   make([]*MemFsNode, 0),
	}}
}

type MemFsNode struct {
	prefix  string
	isDir   bool
	content []byte
	nodes   []*MemFsNode
}

func (m *MemFsNode) ensurePath(path string) *MemFsNode {
	if m.prefix == path {
		return m
	}
	hasPath := false
	for idx, node := range m.nodes {
		if node.prefix == path {
			hasPath = true
			return m.nodes[idx]
		}
	}
	if !hasPath {
		node := MemFsNode{
			prefix:  path,
			isDir:   true,
			content: nil,
			nodes:   make([]*MemFsNode, 0),
		}
		m.nodes = append(m.nodes, &node)
		return &node
	}
	return nil
}

func (f *MemFs) addContentToFile(path string, content []byte) {
	splitPaths := strings.Split(path, "/")
	cursor := f.root
	for _, path := range splitPaths {
		if path == "" {
			continue
		}
		isFile := strings.Contains(path, ".")
		if isFile {
			fileNode := MemFsNode{
				prefix:  path,
				isDir:   false,
				content: content,
				nodes:   nil,
			}
			cursor.nodes = append(cursor.nodes, &fileNode)
			return
		}
		cursor = cursor.ensurePath(path)
	}
}
