package main

import "fmt"

func main() {
	s1 := []string{"a", "b", "c", "d", "e"}
	s2 := []string{"foo", "bar", "c", "d", "e", "f", "g", "h"}
	fmt.Println(getIntersection(s1, s2))
}

func getIntersection(s1, s2 []string) (inter []string) {
	dict := make(map[string]bool)
	for _, v := range s1 {
		dict[v] = true
	}
	for _, v := range s2 {
		if dict[v] {
			inter = append(inter, v)
		}
	}

	return

}
