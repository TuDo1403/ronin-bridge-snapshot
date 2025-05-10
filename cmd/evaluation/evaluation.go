// go run evaluate.go -csv ./res.csv -bank 1000 -price 1.00 -kelly 0.17
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
)

/* -------------------------------------------------- */
/* ---------------- CONFIG FLAGS  ------------------- */
var (
	csvFile   = flag.String("csv", "", "path to CSV with flipId,resultSide (HEAD/TAIL)")
	startBank = flag.Float64("bank", 1000, "initial bankroll in tokens")
	// priceUSD  = flag.Float64("price", 1.0, "token price in USD (for $0.01 VRF fee)")
	kellyF = flag.Float64("kelly", 0.17, "Kelly fraction for our bot")
)

/* ---------------- CONSTANTS  ---------------------- */
var (
// vrfUSD = 0.01 // VRF fee in USD
// theirIdx = 2 // index of their bot in CSV
)

var chips = [...]float64{1, 5, 10, 25, 50, 100}

/* -------------- STRATEGY A: OUR BOT --------------- */
type OurBot struct {
	bank  float64
	fee   float64 // tokens per flip
	f     float64 // Kelly fraction
	round int
}

func (k *OurBot) play(win bool) {
	// Skip if we can't even cover the smallest chip + fee
	minCost := chips[0] + k.fee
	if k.bank < minCost {
		fmt.Println("[OurBot] insufficient roll", "round", k.round)
		return
	}
	target := k.f * k.bank // raw Kelly stake

	// choose the largest chip ≤ target (else smallest chip)
	stake := chips[0]
	for i := len(chips) - 1; i >= 0; i-- {
		if chips[i] <= target {
			stake = chips[i]
			break
		}
	}
	cost := stake + k.fee
	if cost > k.bank {
		fmt.Println("[OurBot] insufficient roll", "round", k.round)
		return
	}

	k.round++
	k.bank -= cost
	if win {
		k.bank += stake * 2
	}

	fmt.Printf("true,%d,%t,%d\n", int(stake), win, int(k.bank))
}

/* -------------- STRATEGY B: THEIR BOT ------------- */

type LadderBot struct {
	bank  float64
	index int
	fee   float64
	round int
}

func (l *LadderBot) play(win bool) {
	if l.bank <= 0 {
		fmt.Println("[LadderBot] insufficient roll", "round", l.round)
		return
	}

	stake := chips[l.index]
	cost := stake + l.fee
	if cost > l.bank {
		fmt.Println("[LadderBot] insufficient roll", "round", l.round)
		return
	}

	l.round++

	l.bank -= cost
	if win {
		l.bank += stake * 2

		// fmt.Printf("true,%d,%t,%d\n", int(chips[l.index]), win, int(l.bank))

		l.index = 0 // reset to first chip

		return
	}

	// fmt.Printf("true,%d,%t,%d\n", int(chips[l.index]), win, int(l.bank))

	if l.index == len(chips)-1 && !win {
		return
	}

	// advance pointer regardless of result
	l.index = (l.index + 1) % len(chips)
}

/* -------------------------------------------------- */
func main() {
	flag.Parse()
	if *csvFile == "" {
		log.Fatal("Give -csv path")
	}
	// feeToken := vrfUSD / *priceUSD

	// startBank := 100
	// endBank := 10000
	// steps := 100
	b := *startBank

	// for i := startBank; i <= endBank; i += steps {
	// b := float64(i)

	our := OurBot{bank: b, fee: 0, f: *kellyF}
	them := LadderBot{bank: b, fee: 0}

	file, err := os.Open(*csvFile)
	if err != nil {
		log.Fatalf("open csv: %v", err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalf("read csv: %v", err)
	}

	total := len(records) // number of flips
	fmt.Println("Total flips:", total)
	for i, rec := range records {
		// skip header
		if i == 0 {
			fmt.Println("choice,bet_amount,won,bankroll")
			continue
		}

		result := rec[2]
		won := result == "true"

		// both bots always pick HEADS
		// fmt.Println("Logging for our bot")
		our.play(won)
		// fmt.Println("Logging for their bot")
		them.play(won)
	}

	fmt.Println("Iteration with bank", b)
	printStats("OUR   ", b, our.bank, total)
	printStats("THEIR ", b, them.bank, total)
	// }
}

func printStats(tag string, start, end float64, flips int) {
	profit := end - start
	cagr := math.Pow(end/start, 1/float64(flips)) - 1
	fmt.Printf("%s  final=%.2f  profit=%.2f  flips=%d  CAGR=%.4f%%\n",
		tag, end, profit, flips, cagr*100)
}
