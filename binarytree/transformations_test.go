package main

import "testing"

func TestInvertTree1(t *testing.T) {
	root := buildTree("4,2,7,1,3,6,9")
	expectedTree := buildTree("4,7,2,9,6,3,1")
	expectedTreeStr := getTreeString(expectedTree)
	actualNode := invertTree(root)
	actualTreeStr := getTreeString(actualNode)
	if actualTreeStr != expectedTreeStr {
		t.Error("LCA test failed: actual: ", actualTreeStr, " expected: ", expectedTreeStr)
	}
}
