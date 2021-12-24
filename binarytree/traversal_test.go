package main

import "testing"

/*func TestLevelOrder(root *TreeNode) [][]int {
}*/

func testEq(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func TestAverageOfLevels1(t *testing.T) {
	root := buildTree("3,9,20,15,7")
	display(root)
	actualAvgs := averageOfLevels(root)
	expectedAvgs := []float64{3, 14.5, 11}
	if !testEq(actualAvgs, expectedAvgs) {
		t.Error("TestAverageOfLevels failed: actual: ", actualAvgs, " expected: ", expectedAvgs)
	}
}

func TestAverageOfLevels2(t *testing.T) {
	root := buildTree("3,9,20,null,null,15,7")
	display(root)
	actualAvgs := averageOfLevels(root)
	expectedAvgs := []float64{3, 14.5, 11}
	if !testEq(actualAvgs, expectedAvgs) {
		t.Error("TestAverageOfLevels failed: actual: ", actualAvgs, " expected: ", expectedAvgs)
	}
}
