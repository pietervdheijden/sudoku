package util

func Equals(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for _, e := range s2 {
		if !Contains(s1, e) {
			return false
		}
	}
	return true
}

func Subset(s1 []int, s2 []int) bool {
	if len(s2) > len(s1) || len(s2) == 0 {
		return false
	}
	for _, e := range s2 {
		if !Contains(s1, e) {
			return false
		}
	}
	return true
}

func Merge(s1 []int, s2 []int) []int {
	s3 := []int{}
	for _, v := range s1 {
		if !Contains(s3, v) {
			s3 = append(s3, v)
		}
	}
	for _, v := range s2 {
		if !Contains(s3, v) {
			s3 = append(s3, v)
		}
	}
	return s3
}

// source: https://stackoverflow.com/a/10485970/3737152
func Contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

// returns elements in both s1 and s2
func Intersect(s1 []int, s2 []int) []int {
	s3 := []int{}
	for _, v := range s1 {
		if Contains(s2, v) {
			s3 = append(s3, v)
		}
	}
	return s3
}

func IntersectSlices(slices ...[]int) []int {
	s := []int{}
	for _, slice := range slices {
		intersect := true
		for _, v := range slice {
			for _, otherSlice := range slices {
				if !Contains(otherSlice, v) {
					intersect = false
				}
			}
			if intersect && !Contains(s, v) {
				s = append(s, v)
			}
		}
	}
	return s
}

// returns element in s1, but not in s2
func Except(s1 []int, s2 []int) []int {
	s3 := []int{}
	for _, v := range s1 {
		if !Contains(s2, v) {
			s3 = append(s3, v)
		}
	}
	return s3
}