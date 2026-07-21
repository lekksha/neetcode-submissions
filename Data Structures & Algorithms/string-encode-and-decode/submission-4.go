type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    var builder strings.Builder
    for _, str := range strs {
        builder.WriteString(strconv.Itoa(len(str)))
        builder.WriteByte('#')
        builder.WriteString(str)
    }
    return builder.String()
}

func (s *Solution) Decode(encoded string) []string {
	const delimiter string = "#"
	res := []string{}

	prefix := "" 
	curStr := ""
	counter := 0
	c := ""
	i := 0
	for i < len(encoded) {
		c = string(encoded[i])
		if c == delimiter {
			counter, _ = strconv.Atoi(prefix)
			prefix = ""
			curStr = encoded[i+1:i+counter+1]
			res = append(res, curStr)
			i = i + counter
			counter = 0
		} else {
			prefix = prefix + c
		}
		i++
	}
	return res
}
