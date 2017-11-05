package basics

// importはファイル毎で行う
import (
	"fmt"
	"math/cmplx"
	"math"
)

func Add(x, y int) int {
	return x + y
}

func Swap(x, y string) (string, string) {
	return y, x
}

func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java, PHP bool
var i, j int = 1, 2

func Variables() {
	var i int
	fmt.Println(i, c, python, java)
}

func VariablesInit() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

func VariableDecolation() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func BasicTypes() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}

func TypeConversion() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

func TypeInteface() {
	v := "42" // change me!
	fmt.Printf("v is of type %T\n", v)

	w := 42 // change me!
	fmt.Printf("w is of type %T\n", w)
}