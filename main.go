package main

import "fmt"

func main() {
    //init 
	i := 21
	k := 123.456
	j := true


    fmt.Printf("%v\n", i) // Menampilkan nilai i (21)
	fmt.Printf("%T\n", i) // Menampilkan tipe data i (int)
	fmt.Printf("%%\n") // Menampilkan % 

	
	
    fmt.Println(j)//Menampilkan Bolean

	fmt.Printf("%b\n", i)//menampilkan 10101
    fmt.Printf("%d\n", i)//menampilkan base 10
    fmt.Printf("%o\n", i)//Menampilkan base 8
    fmt.Printf("%x\n", 15)//Menampilkan f kecil dengan base 16
    fmt.Printf("%X\n", 15)//Menampilkan F besar dengan base 16
	fmt.Printf("U+%04X\n", 'Я')//Menampilkan Я  

	//VAR untuk angka di belakang koma
    fmt.Printf("%.6f\n", k)
	fmt.Printf("%.6E\n", k)


}