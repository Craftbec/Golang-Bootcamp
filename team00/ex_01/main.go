package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"sync"
)

func main() {
	var k float64
	flag.Float64Var(&k, "k", 0.5, "-k 'float number' for apply coefficient")
	flag.Parse()
	job1(k)
}

func SaveData(num int32, expVal, stdDiv float64) {
	file, err := os.OpenFile("report.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Unable to open file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("Обработано: %d Mean: %f Sd: %f\n", num, expVal, stdDiv))
}

func job1(k float64) {
	var pool = sync.Pool{
		New: func() interface{} { return []float64{} },
	}
	arrf := []float64{}
	sc := bufio.NewScanner(os.Stdin)
	var i int32
	var mean, sd float64
	for sc.Scan() {
		i++
		val, err := strconv.ParseFloat(sc.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}
		arrf = pool.Get().([]float64)
		arrf = append(arrf, val)
		mean = getMean(arrf)
		sd = getSD(arrf)
		if i%10 == 0 {
			SaveData(i, mean, sd)
		}
		fmt.Printf("%.4f %.4f\n", mean, sd)
		pool.Put(arrf)
	}
	SaveData(i, mean, sd)
	var rng float64 = k * sd
	fmt.Printf("Коэффициент: %f\n", rng)
}

func job2() {
	var pool = sync.Pool{
		New: func() interface{} { return []float64{} },
	}

	arr := []float64{
		1.6513, 1.7565, 0.6083, 1.4704, 2.1753, 1.7167, 2.2368, 2.5537,
		0.8521, 2.4381, 2.1716, 2.4907, 0.4413, 1.5685, 1.9973, 0.9721,
		2.1812, 2.7806, 1.9680, 1.6533, 1.4858, 1.8646, 1.3895, 2.5349,
		2.0925, 2.4069, 1.4311, 2.7880, 2.9056, 2.0298, 3.3549, 0.7335,
		1.8123, 1.6946, 2.5190, 2.5859, 2.3343, 2.5242, 1.9325, 1.4543,
		2.3342, 1.8458, 1.9856, 1.6804, 2.3315, 1.5189, 2.3910, 1.7651,
		2.4104, 2.0206, 1.6257, 1.1815, 1.4319, 2.7913, 2.3952, 1.3910,
		1.5083, 1.8050, 2.1738, 1.5019, 1.8716, 1.8310, 1.7531, 2.7549,
		1.6786, 3.1878, 2.2641, 2.3616, 1.5296, 1.5411, 2.6792, 1.8144,
		1.5379, 1.0726, 2.9240, 2.1584, 1.9594, 2.4932, 2.4661, 2.0664,
		1.8146, 2.3391, 1.4690, 2.2529, 1.5426, 2.5530, 1.9883, 0.6890,
		0.6913, 1.3056, 2.3350, 2.3363, 2.4059, 2.8849, 1.9974, 1.8074,
		2.3966, 1.5232, 2.0516,
	}

	arrf := []float64{}
	for _, val := range arr {
		arrf = pool.Get().([]float64)
		arrf = append(arrf, val)
		fmt.Printf("%.4f %.4f\n", getMean(arrf), getSD(arrf))
		pool.Put(arrf)
	}
}

func getMean(arr []float64) (res float64) {
	var out float64 = 0
	for _, v := range arr {
		out += float64(v)
	}
	return out / float64(len(arr))
}

func getSD(arr []float64) float64 {
	if len(arr) == 1 {
		return 0
	}
	var out float64
	avg := getMean(arr)
	for _, v := range arr {
		out += math.Pow((float64(v) - avg), 2)
	}
	out = out / float64(len(arr)-1)
	return math.Sqrt(out)
}
