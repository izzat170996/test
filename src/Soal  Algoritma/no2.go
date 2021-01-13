package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("===========Kelompok 1=========== \n")
	var cetak = 0
	var rataRata = 0
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch  {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e		
				
			}
		}
		return min, max
	}
	kelompok1 := []int {1, 2, 33, 44, 55, 33, 44, 22, 44, 33}
	fmt.Println("Ini adalah data kelompok1 = ", kelompok1)
	sort.Ints(kelompok1)
	fmt.Println("Data kelompok1 yang telah di urutkan = ", kelompok1)
	for i := 0; i < len(kelompok1); i++ {
		cetak += kelompok1[i]
	}
	fmt.Println("Total kelompok1 adalah = ", cetak)

	for i := 0; i < 10; i++ {
		rataRata += kelompok1[i]
	}
	fmt.Println("Nilai rata rata kelompok1 adalah = ", rataRata / 10)

	var min, max = getMinMax(kelompok1)
	fmt.Printf("Nilai Tertinggi kelompok1 adalah = %v\n Nilai Terendah kelompok1 adalah = %v\n", max, min)
	fmt.Println("\n")


	fmt.Println("===========Kelompok 2=========== \n")
	var cetak2 = 0
	var rataRata2 = 0
	var getMinMax2 = func(n []int) (int, int) {
		var min2, max2 int
		for i, d := range n {
			switch  {
			case i == 0:
				max2, min2 = d, d
			case d > max2:
				max2 = d
			case d < min2:
				min2 = d		
				
			}
		}
		return min2, max2
	}
	kelompok2 := []int {2, 77, 66, 1, 2, 3, 4, 5, 6, 7, 89}
	fmt.Println("Ini adalah data kelompok2 = ", kelompok2)
	sort.Ints(kelompok2)
	fmt.Println("Data kelompok2 yang telah di urutkan = ", kelompok2)
	for j := 0; j < len(kelompok2); j++ {
		cetak2 += kelompok2[j]
	}
	fmt.Println("Total kelompok2 adalah = ", cetak2)

	for j := 0; j < 10; j++ {
		rataRata2 += kelompok2[j]
	}
	fmt.Println("Nilai rata rata kelompok2 adalah = ", rataRata2 / 11)

	var min2, max2 = getMinMax2(kelompok2)
	fmt.Printf("Nilai Tertinggi kelompok2 adalah = %v\n Nilai Terendah kelompok2 adalah = %v\n", max2, min2)
	fmt.Println("\n")

	fmt.Println("===========Kelompok 3=========== \n")
	var cetak3 = 0
	var rataRata3 = 0
	var getMinMax3 = func(n []int) (int, int) {
		var min3, max3 int
		for i, f := range n {
			switch  {
			case i == 0:
				max3, min3 = f, f
			case f > max3:
				max3 = f
			case f < min3:
				min3 = f		
				
			}
		}
		return min3, max3
	}
	kelompok3 := []int {3, 3, 8, 9, 75, 4, 3, 2, 2, 1, 3}
	fmt.Println("Ini adalah data kelompok3 = ", kelompok3)
	sort.Ints(kelompok3)
	fmt.Println("Data kelompok3 yang telah di urutkan = ", kelompok3)
	for k := 0; k < len(kelompok3); k++ {
		cetak3 += kelompok3[k]
	}
	fmt.Println("Total kelompok3 adalah = ", cetak3)

	for k := 0; k < 10; k++ {
		rataRata3 += kelompok3[k]
	}
	fmt.Println("Nilai rata rata kelompok3 adalah = ", rataRata3 / 11)

	var min3, max3 = getMinMax3(kelompok3)
	fmt.Printf("Nilai Tertinggi kelompok3 adalah = %v\n Nilai Terendah kelompok3 adalah = %v\n", max3, min3)
	fmt.Println("\n")


	
	
}
  