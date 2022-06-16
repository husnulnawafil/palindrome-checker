package main

import "fmt"

/*
================================
 -------------------------------
 TAKE HOME SESSION
 -------------------------------
================================
*/
/*	Agar proses recursive dapat dikontrol pengulangannya
	dibutuhkan sebuah variabel pembantu. Dalam kasus ini
	saya menggunakan variabel i yang bernilai 0<= i < len(w0rd)/2 */
func recursiveChecker(word string, i int) bool {
	if i < len(word)/2 {
		if word[i] == word[len(word)-i-1] {
			return recursiveChecker(word, i+1)
		} else {
			return false
		}
	} else {
		return true
	}
}
func main() {
	/*	Dalam pemanggilan fungsi recursiveChecker nilai parameter i idealnya bernilai 0
		Jika tidak maka hasil bisa saja tidak seperti yang diharapkan*/
	fmt.Println(recursiveChecker("katak", 0))
	fmt.Println(recursiveChecker("basi", 0))
	fmt.Println(recursiveChecker("isi", 0))
	fmt.Println(recursiveChecker("kasur", 0))
	fmt.Println(recursiveChecker("akakakakakakaka", 0))
	fmt.Println(recursiveChecker("akakakaakakakaka", 0))
}

/*
=============================
 ----------------------------
 ONLINE INTERVIEW SESSION
 ----------------------------
=============================
*/

// func wordChecker(word string) bool {
// 	for i := 0; i < len(word)/2; i++ {
// 		if word[i] == word[len(word)-i-1] {
// 			return true
// 		}
// 	}
// 	return false
// }

// func wordChecker2(word string) bool {
// 	for i, j := 0, len(word)-1; i < len(word) && j >= 0; i, j = i+1, j-1 {
// 		if word[i] == word[j] {
// 			return true
// 		}
// 	}
// 	return false

// }

// func main() {

// 	// // fmt.Println(wordChecker2("katak"))
// 	// // fmt.Println(wordChecker2("basi"))
// 	// // fmt.Println(wordChecker2("isi"))
// 	// fmt.Println(wordChecker("katak"))
// 	// fmt.Println(wordChecker("basi"))
// 	// fmt.Println(wordChecker("isi"))

// 	// // sumDiagonal([4][4]int{{11, 2, 4, 1}, {4, 5, 6, 1}, {10, 8, 12, 1}, {1, 1, 1, 1}})
// }

// func sumDiagonal(array [3][3]int) {
// 	res := 0
// 	for i := 0; i < len(array); i++ {
// 		if i%2 == 0 {
// 			res += array[i][0]
// 			res += array[i][2]
// 		} else {
// 			res += array[i][i]
// 			res += array[i][i]
// 		}
// 	}
// 	// result := array[0][0] + array[0][2] + (2 * array[1][1]) + array[2][0] + array[2][2]
// 	fmt.Println(res)
// }

// func sumDiagonal(array [4][4]int) {
// 	res := 0
// 	for i := 0; i < len(array); i++ {
// 		res += array[i][i] + array[i][len(array)-1-i]
// 	}
// 	fmt.Println(res)
// }

// 11 2 4 1
// 4 5 6 1
// 10 8 12 1
// 1 1 1 1
