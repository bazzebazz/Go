package main

import (
	"flag"
	"fmt"
)

func main() {
	// filter pattern
	flagPattern := flag.String("p", "", "filter by pattern")
	flagAll := flag.Bool("a", false, "all files including hide files")
	flagNumberRecords := flag.Int("n", 0, "number of records")

	// order flags
	hasOrderTime := flag.Bool("t", false, "sort by time, oldest to first")
	hasOrderSize := flag.Bool("s", false, "sort by file size, smallest to bigger")
	hasOrderReverse := flag.Bool("r", false, "reverse order sorting")

	flag.Parse()
	fmt.Println("Flag Pattern:", *flagPattern)
	fmt.Println("Flag All:", *flagAll)
	fmt.Println("Flag Number Records:", *flagNumberRecords)
	fmt.Println("Flag Order Time:", *hasOrderTime)
	fmt.Println("Flag Order Size:", *hasOrderSize)
	fmt.Println("Flag Order Reverse:", *hasOrderReverse)
}
