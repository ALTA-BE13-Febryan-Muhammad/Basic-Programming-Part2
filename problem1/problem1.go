package main

import "fmt"

func KonversiNilai(nilai int) string {
	// your code here
	var NilaiUjian string
	if nilai >= 90 && nilai <= 100 {
		NilaiUjian = "A+"
	} else if nilai >= 80 && nilai < 90 {
		NilaiUjian = "A"
	} else if nilai >= 70 && nilai < 80 {
		NilaiUjian = "B+"
	} else if nilai >= 0 && nilai < 70 {
		NilaiUjian = "C"
	}
	return NilaiUjian
}

func main() {
	var nilai int = 80
	var nilai2 int = 34
	var nilai3 int = 69
	var nilai4 int = 70
	var nilai5 int = 99

	fmt.Println(KonversiNilai(nilai))
	fmt.Println(KonversiNilai(nilai2))
	fmt.Println(KonversiNilai(nilai3))
	fmt.Println(KonversiNilai(nilai4))
	fmt.Println(KonversiNilai(nilai5))
}
