type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    res := ""
    for _, str := range strs {
        res += fmt.Sprintf("%d#%s", len(str), str)
    }
    return res
}

func (s *Solution) Decode(str string) []string {
    res := []string{}
    i := 0

    for i < len(str) {
        j := i
        for str[j] != '#' {
            j++
        }
        var length int
        fmt.Sscanf(str[i:j], "%d", &length)
        res = append(res, str[j+1:j+1+length])
        i = j + 1 + length
    }

    return res
}