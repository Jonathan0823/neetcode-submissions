func isNStraightHand(hand []int, groupSize int) bool {
 if len(hand)%groupSize != 0 { 
	return false
 }   

 count := map[int]int{}
 for _, n := range hand { 
	count[n]++
 }

 sort.Ints(hand)
 for _, n := range hand { 
	if count[n] > 0 { 
		for i := n; i < n + groupSize; i++ { 
			if count[i] == 0 { 
				return false
			}
			count[i]--
		}
	}
 }
 return true
}
