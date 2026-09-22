func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens { 
		val, err := strconv.Atoi(token)
    if err != nil {
        // Jika error, berarti token bukan angka (misalnya operator '+', '*', dll)
        // Tangani sesuai kebutuhan evaluasi stack kamu di sini
        continue 
    }
	
		switch token { 
			case "+":
				res := int(stack[len(stack) - 1]) + int(stack[len(stack) - 2])
				stack = stack[:len(stack)-2]
				stack = append(stack, res)
			case "-":
			res := int(stack[len(stack) - 1]) - int(stack[len(stack) - 2])
				stack = stack[:len(stack)-2]
				stack = append(stack, res)
			case "*":
			res := int(stack[len(stack) - 1]) * int(stack[len(stack) - 2])
				stack = stack[:len(stack)-2]
				stack = append(stack, res)
			case "/":
			res := int(stack[len(stack) - 1]) / int(stack[len(stack) - 2])
				stack = stack[:len(stack)-2]
				stack = append(stack, res)
			default:
				stack = append(stack, strconv.Atoi(token))
		}
	}

	return stack[0]

}
