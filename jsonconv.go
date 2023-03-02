package main

import (
	"sort"
	"strings"
)

func card_getIntJson(key []int, i int, k int, text string, id string) string {
	if len(id) == 0 {
		return ""
	}
	if len(text) == 0 {
		return ""
	}
	var new string
	if i < (len(key) - 1) {
		a := text[k+len(id) : key[i+1]]
		new = wordtoIntJson(id, a)
		//возможно будет ошибка и надо удалить \n
	}
	if i == (len(key) - 1) {
		a := text[k+len(id):]
		new = wordtoIntJson(id, a)
	}
	return new
}

func wordtoIntJson(id string, text string) string {
	//на вход id, text
	// id = "id:" , "id"; text = "....chtoto"

	var new string

	text = strings.TrimSuffix(text, "\n")

	if id[len(id)-1] == 0x3a {
		new = "{" + "\"" + id[0:(len(id)-1)] + "\"" + ":" + text + " }"
		//		new = strings.Replace(new, "\n", "", 1)

	} else {
		new = "{" + "\"" + id + "\"" + ":" + text + " }"
	}
	//выход jsonword
	return new
}

func card_getStringJson(key []int, i int, k int, text string, id string) string {

	if len(id) == 0 {
		return ""
	}
	if len(text) == 0 {
		return ""
	}

	var new string

	if i < (len(key) - 1) {
		a := text[k+len(id) : key[i+1]]
		new = wordtoStringJson(id, a)
		//возможно будет ошибка и надо удалить \n
	}
	if i == (len(key) - 1) {
		a := text[k+len(id):]
		new = wordtoStringJson(id, a)
	}
	return new
}
func wordtoStringJson(id string, text string) string {

	text = strings.TrimSuffix(text, "\n")

	var new string
	if id[len(id)-1] == 0x3a {

		new = "{" + "\"" + id[0:(len(id)-1)] + "\"" + ":" + "\"" + text + "\"" + "}"
		//	new = strings.Replace(new, "\n", "", 1)

	} else {
		new = "{" + "\"" + id + "\"" + ":" + "\"" + text + "\"" + "}"
	}

	return new
}

func search_word(text string, list_word []string) (map[int]string, []int) {
	//На входе список слов и текст

	var test map[int]string
	var keys []int

	for i := range list_word {
		//proverka := strings.Index(text, slova[i])
		l_slovo := strings.Index(text, list_word[i])
		if l_slovo != -1 {
			test2 := make(map[int]string, len(test)+1)
			for k, v := range test {
				test2[k] = v
				_ = v
			}
			test2[l_slovo] = list_word[i]
			test = test2
			keys = add_massiv(keys, l_slovo)
		}
	}
	sort.Ints(keys) //от сортировали
	return test, keys
	//на выходе должны давать map, keys
}
