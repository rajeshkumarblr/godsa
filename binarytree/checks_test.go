package main

import (
	"testing"
)

func TestLowestCommonAncestor1(t *testing.T) {
	root := buildTree("6,2,8,0,4,7,9,null,null,3,5")
	expected := 6
	p := root.Left
	q := root.Right
	actual := lowestCommonAncestorIterative(root, p, q).Val
	if actual != expected {
		t.Error("LCA test failed: actual: ", actual, " expected: ", expected)
	}
}

func TestLowestCommonAncestor2(t *testing.T) {
	root := buildTree("6,2,8,0,4,7,9,null,null,3,5")
	expected := 2
	p := root.Left
	q := p.Right
	actual := lowestCommonAncestorIterative(root, p, q).Val
	if actual != expected {
		t.Error("LCA test failed: actual: ", actual, " expected: ", expected)
	}
}

func TestLowestCommonAncestor3(t *testing.T) {
	root := buildTree("6,2,8,0,4,7,9,null,null,3,5")
	expected := 4
	p := root.Left.Right.Left
	q := root.Left.Right.Right
	actualNode := lowestCommonAncestorIterative(root, p, q)
	if actualNode != nil {
		actualVal := actualNode.Val
		if actualVal != expected {
			t.Error("LCA test failed: actual: ", actualVal, " expected: ", expected)
		}
	}
}
