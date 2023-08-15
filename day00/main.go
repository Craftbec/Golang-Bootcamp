package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func mode(arr []int) int {
	var num, count int
	if len(arr) == 1 {
		return arr[0]
	}
	m := map[int]int{}
	for key, value := range arr {
		m[value]++
		if key == 0 {
			num = value
			count = 1
		} else {
			if m[value] > count {
				num = value
				count = m[value]
			}
		}
	}
	return num
}

func median(i float64, arr []int) float64 {
	tmp_i := int(i)
	if tmp_i%2 != 0 {
		return float64(arr[tmp_i/2])
	} else {
		return float64((arr[tmp_i/2] + arr[tmp_i/2-1])) / 2.0
	}
}

func sd(i float64, arr []int, mean float64) float64 {
	var sd float64
	if i != 1 {
		for _, value := range arr {
			sd += math.Pow(float64(value)-mean, 2)
		}
		sd = math.Sqrt(sd / (i - 1))
	}
	return sd
}

func main() {
	var arr []int
	var sum, i, mean float64
	var Mean, Median, Mode, SD bool
	flag.BoolVar(&Mean, "mean", false, "Show metric Mean")
	flag.BoolVar(&Median, "median", false, "Show metric Median")
	flag.BoolVar(&Mode, "mode", false, "Show metric Mode")
	flag.BoolVar(&SD, "sd", false, "Show metric SD")
	flag.Parse()
	fmt.Println("To complete the entry, enter go")
	inp := bufio.NewScanner(os.Stdin)
	for inp.Scan() {
		a := inp.Text()
		if a == " " {
			fmt.Println("Empty input")
			os.Exit(3)
		}
		if a == "go" {
			break
		}
		number, err := strconv.Atoi(a)
		if err != nil {
			fmt.Println("Not an integer")
			os.Exit(4)
		}
		if number < -100000 || number > 100000 {
			fmt.Println("Out of bounds")
			os.Exit(5)
		}
		arr = append(arr, number)
		sum += float64(number)
		i++
	}

	if len(arr) == 0 {
		os.Exit(6)
	}
	sort.Ints(arr)
	mean = sum / i
	if Mean {
		fmt.Printf("Mean: %.2f\n", mean)
	}
	if Median {
		fmt.Printf("Median: %.2f\n", median(i, arr))
	}
	if Mode {
		fmt.Printf("Mode: %d\n", mode(arr))
	}
	if SD {
		fmt.Printf("SD: %.2f\n", sd(i, arr, mean))
	}
	if !Mean && !Median && !Mode && !SD {
		fmt.Printf("Mean: %.2f\nMedian: %.2f\nMode: %d\nSD: %.2f\n", mean, median(i, arr), mode(arr), sd(i, arr, mean))
	}
}
