package piscine

func ConcatParams(args []string) string {
	var str string
	for i := 0; i < len(args); i++ {
		if i == len(args)-1 {
			str = str + string(args[i])
		}
		str = str + string(args[i]) + "\n"
	}
	return str
}
