package main

import (
	"./basics"
	"./pointers"
)

func basic() {
	//basics.HelloWorld()
	//basics.ShortVariableDecolations()
	//fmt.Println(basics.Add(42, 13))
	//fmt.Println(basics.Swap("hello", "world"))
	//fmt.Println(basics.Split(17))
	//basics.Variables()
	//fmt.Print(basics.PHP)
	//basics.VariablesInit()
	//basics.VariableDecolation()
	//basics.BasicTypes()
	//basics.TypeConversion()
	//basics.TypeInteface()
	//basics.IfState()
	//basics.Loop()
	//basics.Switch()
	//basics.Switch2()
	basics.DeferSample()
}

func point() {
	//pointers.GoPointer()
	//pointers.VertexPrint()
	pointers.StructLiteral()
}

func main() {
	//basic()
	point()
}