package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("./ignr/pairs.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	fmt.Println("type MyTradeHistorys struct {")

	r := bufio.NewReader(f)
	for {
		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		// BTCZEC  MyTradeHistory `json:"BTC_ZEC"`

		// pair := string(l)

		fmt.Print("\t")
		// l = "BTC/ZEC"

		// strings.Replace(pair, "/")
		coins := strings.Split(string(l), "/")

		fmt.Print(coins[1] + coins[0])

		// fmt.Print("BTCZEC")

		fmt.Print(" MyTradeHistory `json:\"")

		fmt.Print(coins[1] + "_" + coins[0])
		// fmt.Print("BTC_ZEC")

		fmt.Println("\"`")
	}

	fmt.Println("}")

}
