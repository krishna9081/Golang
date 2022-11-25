package main

import "fmt"

func groupStrings(strings []string) map[string][]string {
	dict := map[string][]string{}

	for _, str := range strings {
		key := ""
		for i := 0; i < len(str)-1; i++ {
			diff := (26 + int(str[i]-'a') - int(str[i+1]-'a')) % 26

			key += fmt.Sprintf("%d", diff)
		}

		if val, ok := dict[key]; ok {
			if len(val[0]) != len(str) {
				key = key + " "
			}
		}
		dict[key] = append(dict[key], str)
	}
	res := [][]string{}
	for _, v := range dict {
		res = append(res, v)
	}

	return dict
}

func main() {
	strings := []string{"cba", "la"}
	fmt.Println(groupStrings(strings))

}
