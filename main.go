package main

import "fmt"

func main() {
	//hello world
	fmt.Println("hello, world!")

	//var const dan string
	var nama string = "joni"
	nama = "jono"
	fmt.Println(nama)

	const jeneng string = "deni"
	fmt.Println(jeneng)

	//conversion
	var kata = "hallo"
	var k = kata[0]
	var kKata = string(k)
	fmt.Println(kata)
	fmt.Println(kKata)

	//perbandingan

	var murid1 = 100
	var murid2 = 90
	var perbandingan = murid1 == murid2
	fmt.Println(perbandingan)


	//operasi matematika
	var i = 10
	i++
	i += 9
	fmt.Println(i)

	//type data number


	var nilai8 int8 = 127
	var nilai16 int16 = 32767
	var nilai32 int32 = 4020202
	var nilai64 int64 = 32324242424
	fmt.Println(nilai8)
	fmt.Println(nilai16)
	fmt.Println(nilai32)
	fmt.Println(nilai64)

	//boolean, operasi boolean, array
	var pelajar bool = true
	fmt.Println("apakah kamu pelajar", pelajar)

	var codingMobile = 80
	var codingWeb = 75

	var hasilCodingMobile = codingMobile >= 75
	var hasilCOdingWeb = codingWeb >= 75

	var diterima = hasilCOdingWeb || hasilCodingMobile
	fmt.Println(diterima)

	var skill [3]string
	skill[0] = "javascript"
	skill[1] = "golang"
	skill[2] = "desainer"
	fmt.Println(skill)

	var nilaii = [03]int8{
		80,
		80,
		90,
	}
	fmt.Println(nilaii)
	//type declaration

	type str string
	type num int8

	var produk str = "susu"
	var harga num = 10
	fmt.Println(produk)
	fmt.Println(harga)


}