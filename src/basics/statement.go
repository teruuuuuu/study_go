package basics

import (
	"math"
	"fmt"
	"runtime"
	"time"
)

func Loop() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}


func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func IfState() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

func Switch() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}


func Switch2() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func DeferSample() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		// deferは呼び出し元関数終了後にfiloの順番で処理される
		defer fmt.Println(i)
	}

	fmt.Println("done")
}