package set1

import (
	"encoding/base64"
	"encoding/hex"
    "bytes"
	// "log"
)

func Hex_to_base64() []byte {

	in, _ := hex.DecodeString("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	var out []byte
	buf := bytes.NewBuffer(out)
	enc := base64.NewEncoder(base64.RawStdEncoding, buf)
	enc.Write(in)
	return buf.Bytes()
	// base64.RawStdEncoding.Encode(out, in)
	// return out
}

func XOR() string {

    a := "1c0111001f010100061a024b53535009181c"
    b := "686974207468652062756c6c277320657965"

    aByte, _ := hex.DecodeString(a)
    bByte, _ := hex.DecodeString(b)

    if len(aByte) != len(bByte) {
        panic("lenght incorrect")
    }

    out := make([]byte, len(aByte))
    n := 0
    for n < len(aByte) {
        out[n] = aByte[n] ^ bByte[n]
        n++
    }
    outString := hex.EncodeToString(out)
    return outString
}
