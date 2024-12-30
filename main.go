package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"reflect"
)

var succPart []byte

func check(bytes []byte) bool {

	// using reflect to check bytes[:6] == succPart
	//fmt.Println(bytes[:6])
	//fmt.Println(succPart)

	if reflect.DeepEqual(bytes[:3], succPart) {
		return true
	}

	return false
}

func brute() (string, []byte) {

	// let's say you want to sign cookie {"user":"admin","msg":"x"} and you use alphabet like [ascii 32-180]{0,4}

	// should work later on multiprocessing / GPU acceleration

	// this program would work relatively much better when hashing files with bigger size (e.g.2MB); then appending chars at the end
	// it isn't remotely close to hashcat performance yet for sure

	part0 := "{\"user\":\"admin\",\"msg\":\""
	// part1 := []byte()
	part2 := "\"}"
	template := sha1.New()
	template.Write([]byte(part0))

	for charNum := 1; charNum <= 4; charNum++ {

		fmt.Println(charNum)

		searching := true

		minChr := uint8(32)
		maxChr := uint8(180)

		part1b := make([]byte, charNum)

		for i := 0; i < charNum; i++ {
			part1b[i] = byte(minChr)
		}

		for searching {

			for chr := minChr; chr <= maxChr; chr++ {

				part1b[charNum-1] = chr

				// fmt.Printf("%s\n", part1b)

				h := copyHash(template)

				h.Write(part1b)

				h.Write([]byte(part2))

				digest := h.Sum([]byte(""))

				// check condition

				if check(digest) {
					return hex.EncodeToString(digest), append(append([]byte(part0), part1b...), part2...)
				}
			}

			part1b[charNum-1] += 1

			// skip last handle below
			for i := charNum - 1; i > 0; i-- {
				if part1b[i] > maxChr {
					part1b[i-1] += 1
					part1b[i] = minChr
				}
			}
			if part1b[0] > maxChr {
				searching = false
			}
		}

	}
	return "Nothing", []byte("")
}

func main() {

	// test()

	succPart, _ = hex.DecodeString("fba810")

	hexdigest, result := brute()

	fmt.Printf("%s %s \n%v\n", hexdigest, result, result)
}

func test() {

	fmt.Println("ok")

	pre := sha1.New()
	pre.Write([]byte("aa"))

	pre2 := copyHash(pre)

	pre.Write([]byte("aa"))

	fmt.Println(pre)
	fmt.Println(pre2)

	fmt.Println(&pre)
	fmt.Println(&pre2)

	fmt.Println(hex.EncodeToString(pre.Sum([]byte(""))))
	fmt.Println(hex.EncodeToString(pre2.Sum([]byte(""))))

	pre3 := sha1.New()
	pre3.Write([]byte("aaaa"))
	digest3 := pre3.Sum([]byte(""))
	fmt.Println(hex.EncodeToString(digest3))

	return
}
