package bot

func formatPrint(arr []string) string {
	s := ""
	for _, v := range arr {
		s += (v + " ")
	}
	return s
}
