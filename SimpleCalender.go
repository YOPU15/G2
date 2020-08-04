package main

import (
	"fmt"
)

func main() {

	var bulan = [12]string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
	var hari = [7]string{"senin", "selasa", "rabu", "kamis", "jum'at", "sabtu", "Minggu"}

	for a := 0; a < len(bulan); a++ {
		fmt.Println(" ")
		fmt.Println("\n=============")
		fmt.Println(bulan[a])
		fmt.Println("=============")
		if a == 1 || a == 3 || a == 5 || a == 7 || a == 8 || a == 10 || a == 12 {
			for i := 0; i < len(hari); i++ {
				fmt.Print(hari[i], "\t")
			}
			fmt.Println("\n======================================================")
			for j := 1; j <= 34; j++ {
				// if j%7 == 0 {
				// 	fmt.Println(j, "\n")
				// } else {
				// 	fmt.Print(j, "\t")
				// }
				if j == 1 || j == 2 || j == 3 {
					fmt.Print("*\t")
				} else if j == 7 || j == 14 || j == 28 || j == 21 {
					fmt.Print(j-3, "\n")
				} else {
					fmt.Print(j-3, "\t")
				}

			}
		} else {
			for i := 0; i < len(hari); i++ {
				fmt.Print(hari[i], "\t")
			}
			fmt.Println("\n======================================================")
			for j := 1; j <= 30; j++ {
				if j == 7 || j == 14 || j == 21 || j == 28 {
					fmt.Print(j, "\n")
				} else {
					fmt.Print(j, "\t")
				}
			}
		}
	}

}
