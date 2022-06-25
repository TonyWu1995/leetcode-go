package leetcode1773

import "testing"

func Test_countMatches(t *testing.T) {
	items := [][]string{
		{"phone", "blue", "pixel"},
		{"computer", "silver", "phone"},
		{"phone", "gold", "iphone"},
	}
	ruleKey := "type"
	ruleValue := "phone"
	res := countMatches(items, ruleKey, ruleValue)

	if res != 2 {
		t.Errorf("no match")
	}

}
