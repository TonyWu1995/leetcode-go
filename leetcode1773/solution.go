package leetcode1773

var mp = map[string]int{
	"type":  0,
	"color": 1,
	"name":  2,
}

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	res := 0
	for _, item := range items {
		if item[mp[ruleKey]] == ruleValue {
			res++
		}
	}
	return res
}
