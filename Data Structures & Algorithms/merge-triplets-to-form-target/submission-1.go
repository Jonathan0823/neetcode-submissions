func mergeTriplets(triplets [][]int, target []int) bool {
    x, y, z := false, false, false

	for _, t := range triplets { 
		if t[0] > target[0] || t[1] > target[1] || t[2] > target[2] { 
			continue
		} 

		x = x || t[0] == target[0]
		y = y || t[1] == target[1]
		z = z || t[2] == target[2]
	} 

	return x == true && y == true && z == true
}
