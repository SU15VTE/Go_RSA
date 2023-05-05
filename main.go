/*Test RSA*/
package main

import "fmt"

func main() {
	p, q := Getrandom(2048)
	n, e, d := GetKey(p, q)
	str := "Hello,World!"
	fmt.Printf("n:%v\ne:%v\nd:%v\n", n, e, d)
	// plaintext := StrToBigArr(str)
	// fmt.Println(plaintext)
	// ciphertext := Encrypt(plaintext, e, n)
	// fmt.Println(ciphertext)
	// text := Decrypt(ciphertext, d, n)
	// fmt.Println(text)
	// fmt.Println(BigArrToStr(text))
	plaintext := StrToBigInt(str)
	fmt.Println(plaintext)
	ciphertext := EncryptStr(plaintext, e, n)
	fmt.Println(ciphertext)
	text := DecryptStr(ciphertext, d, n)
	fmt.Println(text)
	fmt.Println(BigIntToStr(text))
}
