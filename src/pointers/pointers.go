package pointers

import "fmt"

func GoPointer() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}


type Vertex struct {
	X int
	Y int
}

func VertexPrint() {
	fmt.Println(Vertex{1, 2})
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func StructLiteral() {
	fmt.Println(v1, p, v2, v3)
	fmt.Printf("v1 is of type %T\n", v1)
	fmt.Printf("p is of type %T\n", p)
}
