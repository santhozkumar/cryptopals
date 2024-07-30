package main

import (
	// "cryptopals/set1"
	"encoding/hex"
	"fmt"
	"math"
	"strings"
)


type AlphabetScore struct {
    alphabet string
    probs map[rune]float64
}

var chiAlpha AlphabetScore = AlphabetScore{
    alphabet: "abcdefghijklmnopqrstuvwxyz",
    probs: map[rune]float64{
        'a':	0.08167,
        'b':	0.01492,
        'c':	0.02782,
        'd':	0.04253,
        'e':	0.12702,
        'f':	0.2228,
        'g':	0.2015,
        'h':	0.6094,
        'i':	0.6966,
        'j':	0.0253,
        'k':	0.1772,
        'l':	0.4025,
        'm':	0.2406,
        'n':	0.6749,
        'o':	0.7507,
        'p':	0.1929,
        'q':	0.0095,
        'r':	0.5987,
        's':	0.6327,
        't':	0.9056,
        'u':	0.2758,
        'v':	0.0978,
        'w':	0.2360,
        'x':	0.0250,
        'y':	0.1974,
        'z':	0.074, }}



var reverseHexTable = "" +
	"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
	"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
	"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
	"\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\xff\xff\xff\xff\xff\xff" +
	"\xff\x0a\x0b\x0c\x0d\x0e\x0f\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
	"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
	"\xff\x0a\x0b\x0c\x0d\x0e\x0f\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
	"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
	"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
	"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
	"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
	"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
	"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
	"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
	"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
	"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff"


func XORByte(data []byte, k byte) []byte {
    out := make([]byte, len(data))

    for i, _ := range data {
        out[i] = data[i] ^ k
    }
    return out
}


func CalcScore(dec string, chiAlpha AlphabetScore) float64 {
    // (observed - expected ) ^ 2 / expected
    dec = strings.ToLower(dec)
    total := 0
    count := make(map[rune]int)
    for _, s := range dec {
        if strings.Contains(chiAlpha.alphabet, string(s)) {
            total += 1
            if _, ok := count[s]; !ok{
                count[s] = 1
            } else {
                count[s] += 1
            }
        }
    }

    chi2 := 0.0
    fmt.Println(chiAlpha)
    for _, c := range chiAlpha.alphabet {
        expected := float64(total) * float64(chiAlpha.probs[c])
        fmt.Println("Expected ", expected)
        actual := float64(count[c])
        fmt.Println("Actual ", expected)

        val := math.Pow(expected-actual, 2) / expected
        chi2 += val
    }

    fmt.Println("before returning", chi2)
    return chi2
}


func main() {
	// set1.Hex_to_base64()
 //    set1.XOR()


	// fmt.Printf("%x, %s", out, out)
	// fmt.Printf(outString)

    cipherhex := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
    cipher, _ := hex.DecodeString(cipherhex)

    var low float64 = 1000.0
    var key string
    var message string
    for i := 0; i < 256; i++ {
        dec := XORByte(cipher, byte(i))
        fmt.Println(hex.EncodeToString(dec))

        score := CalcScore(string(dec), chiAlpha)
        fmt.Println(score)

        if score < low {
            low = score
            key = string(byte(i))
            message = string(dec)
        }
    }
    fmt.Println(key)
    fmt.Println(message)
	// fmt.Printf("%b, %s", cipher, cipher)
	// var s string = "46447381"
	// fmt.Println([]byte(s))
	// b, _ := hex.DecodeString(s)
	// fmt.Println(b)
	// fmt.Printf("%x, %s", out, out)

}
