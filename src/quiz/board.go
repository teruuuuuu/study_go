package quiz

import (
	"fmt"
	"text/scanner"
	"sort"
)

/**
 * rの計算結果を格納
 */
type Point struct {
	r int
	r0 int
	r1 int
	r2 int
	r3 int
	r4 int
}
/**
 * ボードに書き込む領域を格納
 */
type Side struct {
	sqrX int
	sqrY int
	sqrW int
	sqrH int
	brdH int
}

/**
 * 走査線上にスキャンし、一辺の長さを格納
 */
type RasterEvaluation struct {
	start uint16 // 開始位置
	end   uint16 // 終了位置
	point uint16 // 現在の一辺の長さ
}

type ScanSet struct {
	resultMap    map[uint16][]RasterEvaluation
	max          int
}

/*
 * ボードの初期化用
 */
func initArray(sz int) *[][]uint16{
	arr := make([][]uint16, sz)
	for i := 0; i < sz; i++ {
		for j := 0; j < sz; j++ {
			arr[i] = append(arr[i], uint16(1))
		}
	}
	return &arr
}

func genBoadCalc(r int) int {
	return (r % 10009) * 99991
}

/**
 * ボード初期化時の計算用
 */
func nextPoint(current *Point) *Point {
	nextP := new(Point)
	nextP.r0 = genBoadCalc(current.r)
	nextP.r1 = genBoadCalc(nextP.r0)
	nextP.r2 = genBoadCalc(nextP.r1)
	nextP.r3 = genBoadCalc(nextP.r2)
	nextP.r4 = genBoadCalc(nextP.r3)
	nextP.r = nextP.r4
	return nextP
}

/**
 * ボード初期化時の計算用
 */
func nextSide(sz int, h int, point Point) *Side {
	nextSide := new(Side)
	nextSide.sqrX = point.r0 % sz
	nextSide.sqrY = point.r1 % sz
	nextSide.sqrW = point.r2 % (sz - nextSide.sqrX) % 100
	nextSide.sqrH = point.r3 % (sz - nextSide.sqrY) % 100
	nextSide.brdH = (point.r4 % h) + 1
	return nextSide
}

/**
 * ボード初期化時に座標指定での書き込み用
 */
func boxFill(arr [][]uint16, side Side) *[][]uint16{
	for x := side.sqrX; x < side.sqrX + side.sqrW; x++ {
		for y := side.sqrY; y < side.sqrY + side.sqrH; y++ {
			arr[y][x] = uint16(side.brdH)
		}
	}
	return &arr
}
/**
 * ボード初期化の結果確認用
 */
func boxView(arr [][]uint16){
	sz := len(arr)
	for i:= 0; i < sz; i++ {
		//fmt.Println(arr[i])
	}
}

/**
 * ボードを初期化して結果を返す
 */
func genBoard(sz int, h int) [][]uint16 {
	max := sz
	if(max > 100) { max = 100 }
	arr := initArray(sz)
	var r, r0, r1, r2, r3, r4 int = 1, 0, 0, 0, 0, 0
	var currentPoint *Point = &Point{r, r0, r1, r2, r3, r4}
	for i := 0; i < max; i++ {
		currentPoint = nextPoint(currentPoint)
		currentSide := nextSide(sz, h, *currentPoint)
		arr = boxFill(*arr, *currentSide)
	}
	return *arr
}

/**
 * 一行の情報を返す
 */
func rasterEvalute(max int, arr []uint16) ScanSet {
	resultMap := make(map[uint16] []RasterEvaluation)
	currentNum := uint16(arr[0])
	startPosition := uint16(0)
	for i:=0; i < len(arr); i++ {
		if currentNum != uint16(arr[i]) {
			evaluateResult, ok := resultMap[currentNum]
			if !ok { evaluateResult = []RasterEvaluation{} }
			evaluateResult = append(evaluateResult, RasterEvaluation{startPosition,uint16(i-1),1})
			resultMap[currentNum] = evaluateResult
			startPosition = uint16(i)
			currentNum = arr[i]
		}
		if i == len(arr) - 1 {
			evaluateResult, ok := resultMap[currentNum]
			if !ok { evaluateResult = []RasterEvaluation{} }
			evaluateResult = append(evaluateResult, RasterEvaluation{startPosition,uint16(i),1})
			resultMap[currentNum] = evaluateResult
		}
	}
	return ScanSet{resultMap, 1}
}

/**
 * 走査線上にスキャンし一辺の長さが最長になるものを取得する
 */
func rasterScan(arr [][]uint16) int {
	currentColSetTemp := ScanSet{make(map[uint16] []RasterEvaluation) ,1}
	upset := &currentColSetTemp
	//one := 1
	max := 1
	for i:=0; i < len(arr); i++ {
		rasterMapTemp := rasterEvalute(max, arr[i])
		downSet := &rasterMapTemp
		colEvalutateMerge(upset, downSet)
		//fmt.Println("i:", i, "; max:", downSet.max)
		//rasterView(upset.resultMap)
		upset = downSet
		max = downSet.max
	}
	return max
}

/**
 * 一行の情報と上位の行をマージする
 */
func colEvalutateMerge(upSet, downSet *ScanSet) {
	downSet.max = upSet.max
	keys := []int{}
	for key := range downSet.resultMap { keys = append(keys,int(key)) }
	sort.Ints(keys)
	for _,key := range keys {
		//go merge( key, upSet, downSet)
		merge( key, upSet, downSet)
	}
	//return max, currentCol
}

func merge( key int, upSet, downSet *ScanSet) {
	mergeResult := []RasterEvaluation{}
	downValue := downSet.resultMap[uint16(key)]
	upValue, upOk := upSet.resultMap[uint16(key)]
	if !upOk {
		upValue = []RasterEvaluation{}
	}
	for i:=0; i < len(downValue); i++ {
		coverd := false
		for j:=0; j<len(upValue); j++ {
			startPosition := downValue[i].start
			if startPosition < upValue[j].start { startPosition = upValue[j].start}
			endPosition := downValue[i].end
			if endPosition > upValue[j].end { endPosition = upValue[j].end}
			if endPosition > startPosition && (endPosition - startPosition) >= uint16(downSet.max) {
				mergeResult = append(mergeResult, RasterEvaluation{startPosition,endPosition,upValue[j].point + 1})
				if(int(upValue[j].point + 1) > downSet.max){
					downSet.max = int(upValue[j].point + 1)
				}
			}

			if(downValue[i].start >= upValue[j].start && downValue[i].end <= upValue[j].end){
				coverd = true
			}
			if(downValue[i].end < upValue[j].start){continue}
		}
		if !coverd && (downValue[i].end - downValue[i].start) > uint16(downSet.max) {
			mergeResult = append(mergeResult, downValue[i])
		}
	}
	downSet.resultMap[uint16(key)] = mergeResult
}

/**
 * 行情報の確認用
 */
func rasterView(rasterMap map[uint16] []RasterEvaluation) {
	keys := []int{}
	for key := range rasterMap { keys = append(keys,int(key)) }
	sort.Ints(keys)


	for _,key := range keys {
		for i:=0; i < len(rasterMap[uint16(key)]); i++ {
			//fmt.Println("key=", key, "value:", rasterMap[uint16(key)][i])
		}
	}
}

func test(){
	arr := genBoard(10,4)
	boxView(arr)
	max := rasterScan(arr)
	fmt.Println("resut:", max)
	if max != 4 {
		fmt.Println("fail")
	}else{
		fmt.Println("success")
	}

	arr = genBoard(13,5)
	boxView(arr)
	max = rasterScan(arr)
	fmt.Println("resut:", max)
	if max != 5 {
		fmt.Println("fail")
	}else{
		fmt.Println("success")
	}
}



func test2(){
	answer(40,6)

	answer(100,5)

	answer(500,4)

	answer(1000,4)

	answer(1000,3)

	answer(2000,4)

	answer(2000,3)

	answer(3000,3)
}

var sin scanner.Scanner
func scan() string{
	tok:=sin.Scan()
	if tok==scanner.EOF {return ""}
	return sin.TokenText()
}

func answer(sz, h int){
	//startTime := time.Now()
	arr := genBoard(sz, h)
	//genBoadTime := time.Now()
	//fmt.Println("genBoad:", (genBoadTime.Sub(startTime)))
	max := rasterScan(arr)
	//scanEndTime := time.Now()
	fmt.Println(max * max)
	//fmt.Println("ScanEnd:", (scanEndTime.Sub(genBoadTime)))
}
func main() {
	//test()
	test2()


	/*
	sin.Init(os.Stdin)
	var s string
	sz := -1
	h := -1
	for {
		s=scan()
		if s=="" {break}
		if s=="," {continue}
		if sz == -1 {
			sz,_ = strconv.Atoi(s)
		}else{
			h,_ = strconv.Atoi(s)
			answer(sz, h)
			sz = -1
			h = -1
		}
	}
	*/
}