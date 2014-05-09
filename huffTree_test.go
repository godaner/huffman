// Ben Eggers
// GNU GPL'd

package huffman

// Tests the huffTree

import (
	"testing"
	"io/ioutil"
	"os"
)

////////////////////////////////////////////////////////////////////////////////
// makeTreeFromText tests
////////////////////////////////////////////////////////////////////////////////

func TestMakeTreeFromTextEmpty(t *testing.T) {
	b := make([]byte, 0)
	err := ioutil.WriteFile(".test", b, 0644)
	if err != nil { t.Error(err) }
	defer os.Remove(".test")
	tree, err := makeTreeFromText(".test")
	if err != nil {
		t.Error("Got non-nil error from makeTreeFromText: ", err)
	}
	if tree != nil {
		t.Error("Tree should be nil! Got: ", tree)
	}
}

func TestMakeTreeFromTextSingleChar(t *testing.T) {
	b := []byte{255}
	err := ioutil.WriteFile(".test", b, 0644)
	if err != nil { t.Error(err) }
	defer os.Remove(".test")
	tree, err := makeTreeFromText(".test")
	if err != nil {
		t.Error("Got non-nil error from makeTreeFromText: ", err)
	}
	if tree == nil {
		t.Error("Tree should be nil! Got: ", tree)
	}
	if tree.count != 1 || tree.char != 255 {
		t.Error("Tree was built improperly! Expected: { char: 255, count: 1 },",
			"got { char:", tree.char, ", count:", tree.count, "}")
	}
}

////////////////////////////////////////////////////////////////////////////////
// makeTreeFromNodeSlice tests
////////////////////////////////////////////////////////////////////////////////

func TestMakeTreeFromNodesEmpty(t *testing.T) {
	nodes := []*huffNode{}
	tree := makeTreeFromNodeSlice(nodes)
	if tree != nil {
		t.Error("Tree wasn't nil! tree: ", tree, ".")
	}
}

func TestMakeTreeFromNodesOneNode(t *testing.T) {
	node := &huffNode{char: 120, count: 10}
	nodes := []*huffNode{node}
	tree := makeTreeFromNodeSlice(nodes)
	if tree.char != 120 || tree.count != 10 {
		t.Error("Unexpected! Got node with count:", tree.count, "and char:",
			tree.char, "instead of count:", node.count, "and char:", node.char)
	}
}

func TestMakeTreeFromNodesBasicTree(t *testing.T) {
	nodes := []*huffNode{{char: 120, count: 2}, {char: 120, count: 2}}
	tree := makeTreeFromNodeSlice(nodes)
	if tree.count != 4 {
		t.Error("Tree root count should have been 4, was: ", tree.count, ".")
	}
	if tree.left.count != 2 || tree.left.char != 120 {
		t.Error("Tree's left node was wrong! Expected { 120, 2 }, got {",
			tree.left.char, ",", tree.left.count, "}")
	}
	if tree.right.count != 2 || tree.right.char != 120 {
		t.Error("Tree's right node was wrong! Expected { 120, 2 }, got {",
			tree.right.char, ",", tree.right.count, "}")
	}
}

// This test is fairly tied to the implementation, but tests of something
// internal (like this) often have to be :(
func TestMakeTreeFromNodesMultiLevelTree(t *testing.T) {
	nodes := []*huffNode{{char: 120, count: 2},
		{char: 120, count: 2},
		{char: 121, count: 3}}
	tree := makeTreeFromNodeSlice(nodes)
	if tree.count != 7 {
		t.Error("Tree root count should have been 7, was:", tree.count)
	}
	if tree.left.count != 3 || tree.left.char != 121 {
		t.Error("Tree's left node was wrong! Expected { 121, 3 }, got {",
			tree.left.char, ", ", tree.left.count, "}")
	}
	if tree.right.count != 4 {
		t.Error("Right subtree count should have been 4, was:", tree.right.count)
	}
	if tree.right.right.count != 2 || tree.right.right.char != 120 {
		t.Error("Tree's right node's right node was wrong! Expected { 120, 2 }, got {",
			tree.right.right.char, ",", tree.right.right.count)
	}
	if tree.right.left.count != 2 || tree.right.left.char != 120 {
		t.Error("Tree's right node's left node left node was wrong!",
			"Expected { 120, 2 }, got { ", tree.right.left.char,
			",", tree.right.left.count, "}")
	}
}