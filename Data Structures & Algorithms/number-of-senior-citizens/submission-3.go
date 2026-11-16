func countSeniors(details []string) int {
    cnt := 0

	for _, d := range details { 
		age, _ := strconv.Atoi(d[11:13])
		if age > 60 { 
			cnt++
		}
	}

	return cnt
}