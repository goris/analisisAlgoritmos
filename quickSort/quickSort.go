package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math/rand/v2"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"time"
)

func randomizedPartition(arr []int, lo int, hi int) int {
	pivotIndex := lo + rand.IntN(hi-lo+1)
	arr[pivotIndex], arr[hi] = arr[hi], arr[pivotIndex]

	pivot := arr[hi]
	i := lo - 1

	for j := lo; j < hi; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[hi] = arr[hi], arr[i+1]
	return i + 1
}

func randomizedQuickSort(arr []int, lo int, hi int) {
	if lo >= hi {
		return
	}

	p := randomizedPartition(arr, lo, hi)
	randomizedQuickSort(arr, lo, p-1)
	randomizedQuickSort(arr, p+1, hi)
}

// quickSort sorts arr in place, in ascending order.
func quickSort(arr []int) {
	randomizedQuickSort(arr, 0, len(arr)-1)
}

func timeQuickSort(arr []int) time.Duration {
	start := time.Now()
	quickSort(arr)
	return time.Since(start)
}

func average(durations []time.Duration) time.Duration {
	if len(durations) == 0 {
		return 0
	}

	var total time.Duration
	for _, d := range durations {
		total += d
	}
	return total / time.Duration(len(durations))
}

func writeTimesCSV(w io.Writer, sizes []int, times [][]time.Duration) error {
	csvWriter := csv.NewWriter(w)

	if err := csvWriter.Write([]string{"n", "run", "nanoseconds"}); err != nil {
		return err
	}

	for i, size := range sizes {
		for run, d := range times[i] {
			record := []string{
				strconv.Itoa(size),
				strconv.Itoa(run + 1),
				strconv.FormatInt(d.Nanoseconds(), 10),
			}
			if err := csvWriter.Write(record); err != nil {
				return err
			}
		}
	}

	csvWriter.Flush()
	return csvWriter.Error()
}

func main() {
	// Genera arreglos de tamaño
	// n=100, 250, 500, 750, 1000, 1250, 1500, 1750, 2000.
	// De manera que cada uno de estos arreglos contenga
	// los números naturales desde 0 hasta n-1.
	arraysSizes := []int{100, 250, 500, 750, 1000, 1250, 1500, 1750, 2000}
	const runs = 50
	const timesFile = "output/times.csv"
	const summaryFile = "output/summary.txt"
	const errorPrefix = "error:"

	if err := os.MkdirAll(filepath.Dir(timesFile), 0o755); err != nil {
		fmt.Fprintln(os.Stderr, errorPrefix, err)
		os.Exit(1)
	}

	summary, err := os.Create(summaryFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, errorPrefix, err)
		os.Exit(1)
	}

	// Everything written to out goes to the console and to summaryFile.
	out := io.MultiWriter(os.Stdout, summary)

	times := make([][]time.Duration, len(arraysSizes))

	for i, size := range arraysSizes {
		arr := generateArraySorted(size)

		// Para cada uno de los arreglos anteriores realiza el siguiente proceso 50 veces:
		// 1. Permuta de manera aleatoria el arreglo
		// 2. Luego ordénalo con Quicksort, mide el tiempo que tarda este proceso y guárdalo.
		for run := 0; run < runs; run++ {
			shuffleArray(arr)
			times[i] = append(times[i], timeQuickSort(arr))
		}

		// Calcula la media para cada tamaño de n
		fmt.Fprintf(
			out,
			"n=%-5d avg=%-10v min=%-10v max=%v\n",
			size,
			average(times[i]),
			slices.Min(times[i]),
			slices.Max(times[i]),
		)
	}

	f, err := os.Create(timesFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, errorPrefix, err)
		os.Exit(1)
	}
	if err := writeTimesCSV(f, arraysSizes, times); err != nil {
		f.Close()
		fmt.Fprintln(os.Stderr, errorPrefix, err)
		os.Exit(1)
	}
	if err := f.Close(); err != nil {
		fmt.Fprintln(os.Stderr, errorPrefix, err)
		os.Exit(1)
	}

	if err := summary.Close(); err != nil {
		fmt.Fprintln(os.Stderr, errorPrefix, err)
		os.Exit(1)
	}

	fmt.Println("times saved to", timesFile)
	fmt.Println("summary saved to", summaryFile)
}
