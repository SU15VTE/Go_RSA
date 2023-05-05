package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

func Getrandom(n int) (p, q *big.Int) {
	// Generate the first large prime number:p,with n bits
	p, err := rand.Prime(rand.Reader, n)
	if err != nil {
		panic(err)
	}
	fmt.Printf("p value：%v\n", p)
	//Generate the next large prime number:q, with n bits
	q, err = rand.Prime(rand.Reader, n)
	if err != nil {
		panic(err)
	}
	fmt.Printf("q value：%v\n", q)
	return p, q
}

func GetKey(p, q *big.Int) (*big.Int, *big.Int, *big.Int) {
	// n=p*q
	n := new(big.Int).Mul(p, q)
	fmt.Println(n)
	//φ(n) = (p-1)*(q-1)
	phi := new(big.Int).Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))
	fmt.Println(phi)
	//1 < e <φ(n)
gete:
	e, _ := rand.Int(rand.Reader, new(big.Int).Sub(phi, big.NewInt(2)))
	d := new(big.Int).ModInverse(e, phi)
	if d == nil {
		fmt.Println("Unable to calculate inverse element, regenerate e...")
		goto gete
	}
	fmt.Println(d)
	return n, e, d
}

// encrypt (planinText ** e )mod n
func Encrypt(plainTextArr []*big.Int, e, n *big.Int) []*big.Int {
	cipherTextArr := make([]*big.Int, len(plainTextArr))
	for i, text := range plainTextArr {
		cipherText := new(big.Int).Exp(text, e, n)
		cipherTextArr[i] = cipherText

	}
	return cipherTextArr
}

// decrypt (cipherText ** d )mod n
func Decrypt(cipherTextArr []*big.Int, d, n *big.Int) []*big.Int {
	plainTextArr := make([]*big.Int, len(cipherTextArr))
	for i, text := range cipherTextArr {
		plainText := new(big.Int).Exp(text, d, n)
		plainTextArr[i] = plainText
	}
	return plainTextArr
}

//Convert a string to an array of type big.Int
func StrToBigArr(str string) []*big.Int {

	result := make([]*big.Int, len(str))
	for i, s := range str {
		result[i] = big.NewInt(int64(s))
	}
	return result
}

//Convert an array of type big.Int to a string
func BigArrToStr(BigIntArr []*big.Int) string {
	var strArr []string
	for _, s := range BigIntArr {
		strArr = append(strArr, string(s.Bytes()))
	}
	return strings.Join(strArr, "")
}

func StrToBigInt(s string) *big.Int {
	return new(big.Int).SetBytes([]byte(s))
}

func BigIntToStr(i *big.Int) string {
	return string(i.Bytes())
}

func EncryptStr(plainText, e, n *big.Int) *big.Int {
	cipherText := new(big.Int).Exp(plainText, e, n)
	return cipherText
}

func DecryptStr(cipherText, d, n *big.Int) *big.Int {
	plainText := new(big.Int).Exp(cipherText, d, n)
	return plainText
}
