package leetcode567

type multiSet [26]int

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	need, have := buildMultiSet(s1), buildMultiSet(s2[0:len(s1)])
	start, end := 0, len(s1)-1
	for need != have && end < len(s2)-1 {
		toRemove := rune(s2[start])
		have.remove(toRemove)
		start++
		end++
		toAdd := rune(s2[end])
		have.add(toAdd)
	}
	return need == have
}

func buildMultiSet(s string) multiSet {
	var set multiSet
	for _, r := range s {
		set.add(r)
	}
	return set
}

func (ms *multiSet) add(r rune) {
	ms[runePosition(r)]++
}

func (ms *multiSet) remove(r rune) {
	ms[runePosition(r)]--
}

func runePosition(r rune) int {
	return int(r - 'a')
}
