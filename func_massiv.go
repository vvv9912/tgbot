package main

func create_massiv(n int) []int {
	a := make([]int, n)
	return a
}

/*
	example use

a := make([]int, 6)

	a[0] = 1
	a[2] = 3
	a[5] = 5
	//var s map
	//a = delete_massiv(a, 6)
	var cc []int
	a = add_massiv(a, 123)
	fdsfs, sads := find_massiv(a, 123)
*/
func delete_massiv(a []int, izm int) []int {
	c := make([]int, len(a)-1)
	izm = izm - 1
	for i := range a {
		if i < izm {
			c[i] = a[i]
		}
		if i > izm {
			c[i-1] = a[i]
		}
	}
	return c
}
func add_massiv(a []int, cadd int) []int {
	c := make([]int, len(a)+1)
	for i := range a {
		c[i] = a[i]
	}
	c[len(a)] = cadd
	return c
}
func find_massiv(a []int, find int) (int, int) {
	//var n int
	// Возвращает n+1 те элемент каой
	for i := range a {
		if a[i] == find {
			err := 0
			return i + 1, err
		}
	}
	err := 1
	return 0, err
}
