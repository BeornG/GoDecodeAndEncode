package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("--------------------ᶘ ᵒᴥᵒᶅ--------------------")

	// decode commands
	decodeCMD := flag.NewFlagSet("dec", flag.ExitOnError)
	decodeRot13 := decodeCMD.Bool("rot13", false, "Decode rot13, must be one string")
	decodeBase2 := decodeCMD.Bool("base2", false, "Decode base2(binary)")
	decodeBase64 := decodeCMD.Bool("base64", false, "Decode base64")
	decodeHex := decodeCMD.Bool("hex", false, "Decode hex")

	// encode commands
	encodeCMD := flag.NewFlagSet("enc", flag.ExitOnError)
	encodeRot13 := encodeCMD.Bool("rot13", false, "Encode rot13, must be one string")
	encodeMD5 := encodeCMD.Bool("md5", false, "Encode md5, must be one string")

	fmt.Println("dec")
	decodeCMD.PrintDefaults()
	fmt.Println("enc")
	encodeCMD.PrintDefaults()
	fmt.Println("--------------------ᶘ ᵒᴥᵒᶅ--------------------")

	if len(os.Args) < 3 {
		os.Exit(1)
	}

	switch os.Args[1] {
	case "dec":
		handleDecode(decodeCMD, decodeRot13, decodeBase2, decodeBase64, decodeHex)
	case "enc":
		handleEncode(encodeCMD, encodeRot13, encodeMD5)
	default:
		fmt.Println("Not a command")
	}

}

func handleDecode(decodeCMD *flag.FlagSet, rot13 *bool, base2 *bool, base64 *bool, hex *bool) {
	decodeCMD.Parse(os.Args[2:])
	input := os.Args[3:]

	switch {
	case !*rot13 && !*base2 && !*base64 && !*hex:
		decodeCMD.PrintDefaults()
		os.Exit(1)
	case *rot13:
		rot13printer()
	case *base2:
		base2decoder(input)
	case *base64:
		base64decoder(input)
	case *hex:
		hexdecoder(input)
	default:
		return
	}
}

func handleEncode(encodeCMD *flag.FlagSet, rot13 *bool, md5 *bool) {
	encodeCMD.Parse(os.Args[2:])

	switch {
	case !*rot13 && !*md5:
		encodeCMD.PrintDefaults()
		os.Exit(1)
	case *rot13:
		rot13printer()
	case *md5:
		GetMD5Hash(os.Args[3])
	default:
		return
	}
}

func GetMD5Hash(input string) {
	hasher := md5.New()
	hasher.Write([]byte(input))
	output := hex.EncodeToString(hasher.Sum(nil))
	fmt.Println(output)
}

func base2decoder(input []string) {
	fmt.Println("Base2 decoded:")
	for i := range input {
		output, err := strconv.ParseInt(input[i], 2, 64) // input, base2, bitsize
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%c", output)

	}
	fmt.Println()
}

func base64decoder(input []string) {
	fmt.Println("Base64 decoded:")
	for i := range input {
		output, err := base64.StdEncoding.DecodeString(input[i])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s", output)
	}
	fmt.Println()
}

func hexdecoder(input []string) {
	fmt.Println("Hex decoded:")
	for i := range input {
		output, err := hex.DecodeString(input[i])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s", output)
	}
	fmt.Println()
}

func rot13printer() {
	input := os.Args[3]
	mapped := strings.Map(rot13, input)
	fmt.Println("Input:")
	fmt.Println(input)
	fmt.Println("Output:")
	fmt.Println(mapped)
}

func rot13(r rune) rune {
	lowercase := r >= 'a' && r <= 'z'
	uppercase := r >= 'A' && r <= 'Z'

	switch {
	case lowercase: // Rotate lowercase letters 13 places.
		if r >= 'm' {
			return r - 13
		} else {
			return r + 13
		}
	case uppercase: // Rotate uppercase letters 13 places.
		if r >= 'M' {
			return r - 13
		} else {
			return r + 13
		}
	}
	return r
}
