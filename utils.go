package main

func removeEmptyStrings(strs []string) []string {
	res := make([]string, 0)

	for _, str := range strs {
		if str != "" {
			res = append(res, str)
		}
	}

	return res
}
