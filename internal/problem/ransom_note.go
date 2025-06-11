package problem

// CanConstructOne uses 2 maps to identify if we have reached enough occurrances of a character.
// It loops through both strings to crearte a map of each. tc: O(n + m)
// We store both those loops in 2 sepearte maps. sc: O(n + m)
// Then we loop though one of the maps again so the. tc: O(n + m + m)

func CanConstructOne(ransomNote string, magazine string) bool {
	rMap := make(map[rune]int)

	for _, v := range ransomNote {
		if _, ok := rMap[v]; !ok {
			rMap[v] = 1
		} else {
			rMap[v] += 1
		}
	}

	mMap := make(map[rune]int)
	for _, v := range magazine {
		if _, ok := mMap[v]; !ok {
			mMap[v] = 1
		} else {
			mMap[v] += 1
		}
	}

	for k, v := range rMap {
		if mMap[k] < v {
			return false
		}
	}

	return true
}
