func countSeniors(details []string) int {
	count := 0
	for _, d := range details { 
		age, _ := strconv.Atoi(d[11:13])
		if age > 60 { 
			count++
		}
	}

	return count
    
}