package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/shopspring/decimal"
)

func main() {
	for true {
		var input string
		fmt.Print("Input diskon ganda: ")
		fmt.Scanln(&input)
		ClearTerminal()

		splitInp := strings.Split(input, ",")

		result := CalculateDiscount(splitInp)
		fmt.Printf("%s%%\n", strings.Join(splitInp, "% + "))
		fmt.Printf("Diskon pengganti: %v%% - Copy to clipboard\n", result)
		clipboard.WriteAll(strings.ReplaceAll(result.String(), ".", ","))
	}
}

func CalculateDiscount(splitInp []string) decimal.Decimal {
	// Float land
	floatArr := make([]float64, len(splitInp))
	for idx, child := range splitInp {
		val, err := strconv.ParseFloat(child, 32)
		if err != nil {
			panic(err)
		}
		floatArr[idx] = val
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(floatArr)))

	// Decimal land
	counter := decimal.NewFromInt(100)
	for _, child := range floatArr {
		percent := counter.Mul(decimal.NewFromFloat(child / 100))
		counter = counter.Sub(percent)
	}

	return decimal.NewFromInt(100).Sub(counter)
}

func ClearTerminal() {
	switch runtime.GOOS {
	case "darwin":
		runCmd("clear")
	case "linux":
		runCmd("clear")
	case "windows":
		runCmd("cmd", "/c", "cls")
	default:
		runCmd("clear")
	}
}

func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
