package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func Average(data []float64) float64 {
	myData := data
	n := len(data)
	if n == 0 {
		return 0
	}
	var sum float64
	for _, num := range myData {
		sum += num
	}
	average := sum / float64(n)
	return average
}

func Median(data []float64) float64 {
	numericData := data
	sort.Float64s(numericData)

	n := len(numericData)
	if n == 0 {
		return 0
	}
	if n%2 == 0 {
		mid1 := numericData[n/2-1]
		mid2 := numericData[n/2]
		return (mid1 + mid2) / 2
	}
	return numericData[n/2]
}

func Variance(data []float64) float64 {
	numericData := data

	n := len(numericData)
	if n == 0 {
		return 0
	}
	var sum float64
	for _, num := range numericData {
		sum += num
	}
	mean := sum / float64(len(numericData))
	var varianceSum float64
	for _, v := range numericData {
		varianceSum += math.Pow(v-mean, 2)
	}
	return math.Round(varianceSum / float64(n))
}

func StdDev(num float64) float64 {
	mySqrt := math.Sqrt(num)
	return float64(math.Round(mySqrt))
}

func outRange(num float64) bool {
	if num < 0 || num > 1000 {
		return true
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var data []float64
	isFirst := true

	for scanner.Scan() {
		num, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			continue
		}
		if outRange(num) {
			num = Median(data)

		}
		data = append(data, num)

		average := Average(data)
		if isFirst {
			lowLim := num - 40
			highLim := num + 40
			fmt.Printf("%d %d\n", int(lowLim), int(highLim))
			isFirst = false
		} else {
			lowLim := average - StdDev(Variance(data))
			highLim := average + StdDev(Variance(data))
			fmt.Printf("%d %d\n", int(lowLim), int(highLim))

		}
	}
}
