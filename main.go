package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/cipepser/plot/myutil"
)

type Executions []struct {
	GlobalTradeID int    `json:"globalTradeID"`
	TradeID       int    `json:"tradeID"`
	Date          string `json:"date"`
	Type          string `json:"type"`
	Rate          string `json:"rate"`
	Amount        string `json:"amount"`
	Total         string `json:"total"`
}

func main() {
	f, err := os.Open("./polo.json")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	es := &Executions{}

	dec := json.NewDecoder(f)
	err = dec.Decode(es)
	if err != nil {
		panic(err)
	}

	// fes := *es
	fes := make([]float64, len(*es))
	for i, e := range *es {
		fes[i], _ = strconv.ParseFloat(e.Rate, 64)
	}

	// fmt.Println(fes)
	var max, min float64
	min = math.Inf(1)
	max = math.Inf(-1)
	for _, v := range fes {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	fmt.Println("max: ", max)
	fmt.Println("min: ", min)

	myutil.MySinglePlot(fes)

	// var low, max float64
	// low = 99999999999999
	//
	// for _, e := range *es {
	//   if strconve.Amount
	// }

	// for _, s := range r.Document.Sentences.Sentence {
	// 	for _, t := range s.Tokens.Token {
	// 		if t.NER.Text == "PERSON" {
	// 			fmt.Println(t.Word.Text)
	// 		}
	// 	}
	// if i == 0 {
	// 	break
	// }
	// }

}
