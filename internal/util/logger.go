package util

import (
	"context"
	"fmt"
	"os"
	// "sort"
	"sync"
	"text/tabwriter"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

// ANSI escape codes for colors
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorCyan   = "\033[36m"
)

func LogTracker(ctx context.Context, wg *sync.WaitGroup, done <-chan struct{}, trackers map[common.Address]*Tracker, interval time.Duration) {
	defer wg.Done()

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {

		case <-ctx.Done():
			log.Trace("Logger routine stopped.")
			return

		case <-done:
			log.Trace("Logger routine received done signal. Exiting.")
			return

		case <-ticker.C:
			// Create a tabwriter for pretty-printing
			writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)

			// Print the header with color
			fmt.Fprintf(writer, "%sToken\tTxCount\tAccumulatedAmount\tLastRecordedBlock\tMintTxHashesCount%s\n", ColorBlue, ColorReset)

			// Print the tracker data with alternating row colors
			rowColor := ColorCyan

			for token, tracker := range trackers {
				fmt.Fprintf(writer, "%s%s\t%d\t%s\t%d\t%d%s\n",
					rowColor, token, tracker.TxCount, tracker.AccAmount, tracker.LastRecordedBlock, len(tracker.MintTxHashes), ColorReset)

				// Alternate row colors
				if rowColor == ColorCyan {
					rowColor = ColorGreen
				} else {
					rowColor = ColorCyan
				}
			}

			// Flush the writer to ensure all data is written
			writer.Flush()
		}
	}
}
