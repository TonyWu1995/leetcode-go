package leetcode811

import (
	"fmt"
	"testing"
)

func Test_subdomainVisits_case1(t *testing.T) {
	cpdomains := []string{"9001 discuss.leetcode.com"}

	res := subdomainVisits(cpdomains)

	fmt.Println(res)

}

func Test_subdomainVisits_case2(t *testing.T) {
	cpdomains := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}

	res := subdomainVisits(cpdomains)

	fmt.Println(res)

}
