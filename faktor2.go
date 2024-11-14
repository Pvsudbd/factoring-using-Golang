package main

import "fmt"



func main() {
	var n int
	fmt.Print("Masukkan jumlah siswa: ")
	fmt.Scan(&n)

	
	nilai := make([]int, n)

	
	fmt.Println("Masukkan nilai siswa secara urut:")
	for i := 0; i < n; i++ {
		fmt.Printf("Nilai siswa %d: ", i+1)
		fmt.Scan(&nilai[i])
	}

	
	fmt.Println("\nHasil nilai siswa:")
	for i := 0; i < n; i++ {
		var grade string
		if nilai[i] >= 85 {
			grade = "A"
		} else if nilai[i] >= 70 {
			grade = "B"
		} else if nilai[i] >= 55 {
			grade = "C"
		} else if nilai[i] >= 40 {
			grade = "D"
		} else {
			grade = "E"
		}
		fmt.Printf("siswa %d: %s\n", i+1,grade)
	}
}
