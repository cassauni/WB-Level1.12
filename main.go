// Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree"). Создать для неё собственное множество.
// Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.
package main

import "fmt"

func UnicSet(data []string) map[string]struct{} {
	set := make(map[string]struct{}, 0)
	for _, value := range data {
		set[value] = struct{}{}
	}
	return set
}

func main() {

	set := []string{"cat", "cat", "dog", "cat", "tree"}
	unicSet := UnicSet(set)
	for key, _ := range unicSet {
		fmt.Printf("%s ", key)
	}
}
