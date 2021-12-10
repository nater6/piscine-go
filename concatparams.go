package piscine

func ConcatParams(args []string) string {
	var str string
	for i := 0; i < len(args); i++ {
		str = str + string(args[i]) + "\n"
	}
	return str
}
