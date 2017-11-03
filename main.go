package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"./tools"
)

var (
	run = flag.String("run", "", "Two integers delimited by dash representing the set and challenge number to run.")

	dataDir = "data"
)

func readBase64File(filename string) []byte {
	b, err := ioutil.ReadFile(path.Join(dataDir, filename))
	if err != nil {
		log.Fatalf("Error reading file: %v.", err)
	}
	raw, err := base64.StdEncoding.DecodeString(string(b))
	if err != nil {
		log.Fatalf("Error decoding base64 string: %v.", err)
	}
	return raw
}

func readLines(filename string) []string {
	f, err := os.Open(path.Join(dataDir, filename))
	if err != nil {
		log.Fatalf("Error opening file: %v.", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	lines := []string{}
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}
func main() {
	flag.Parse()

	switch *run {
	case "1-3":
		input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
		output, _, err := tools.SingleByteXorDecipher(input)
		if err != nil {
			log.Fatalf("Error decipher: %v.", err)
		}
		fmt.Println(output)
	case "1-4":
		inputFile := "1-4.txt"
		lines := readLines(inputFile)

		bestScore := 0.0
		bestLine := ""
		for _, line := range lines {
			correctLine, score, err := tools.SingleByteXorDecipher(line)
			if err != nil {
				log.Fatalf("Error deciphering: %v.", err)
			}
			if score > bestScore {
				bestScore = score
				bestLine = correctLine
			}
		}
		fmt.Printf("Best line is %q.\n", bestLine)
	case "1-6":
		inputFile := "1-6.txt"
		ct := readBase64File(inputFile)
		answer, err := tools.SolveRepeatXor(ct)
		if err != nil {
			log.Fatalf("Error Solving RepeatXor: %v.", err)
		}
		fmt.Println(answer)
	case "1-7":
		inputFile := "1-7.txt"
		ct := readBase64File(inputFile)
		answer, err := tools.AesEcbDecrypt(ct, []byte("YELLOW SUBMARINE"))
		if err != nil {
			log.Fatalf("Error AES decrypt: %v.", err)
		}
		fmt.Println(string(answer))
	case "1-8":
		inputFile := "1-8.txt"
		lines := readLines(inputFile)

		bestScore := 0
		bestLine := ""
		blockSize := 128
		for _, l := range lines {
			score := 0
			blocks := make(map[string]int)
			blockSizeHex := blockSize / 4
			for i := 0; i+blockSizeHex <= len(l); i += blockSizeHex {
				block := l[i : i+blockSizeHex]
				blocks[block]++
				if blocks[block] > 1 {
					score++
				}
			}
			if score > bestScore {
				bestScore = score
				bestLine = l
			}

		}
		fmt.Printf("The line encrypted with ECB is %q with %d duplicates.\n", bestLine, bestScore)

	}

}
