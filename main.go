package main

import (
	"fmt"

	lvl1 "github.com/khunaungpaing/GO_LEARN/level-one"
	lvl2 "github.com/khunaungpaing/GO_LEARN/level-two"
)

/* _______________________
*    GO LEVEL 1 EXERCISES
*  _______________________
 */
func LevelOne() {

	// Variable
	fmt.Println("VARIABLE IN GO\n=================")
	lvl1.GoVariables()

	// Value
	fmt.Println("\n\nVALUE IN GO\n=================")
	lvl1.BasicValueInGo()
	lvl1.CompositeValuesInGo()
	lvl1.MapValuesInGo()
	lvl1.StructValuesInGo()

	// Slice
	fmt.Println("SLICE IN GO\n=================")
	lvl1.SliceInGo()

	// Array
	fmt.Println("\n\nArray IN GO\n=================")
	lvl1.ArrayInGo()

	// rune
	fmt.Println("\n\nRune IN GO\n=================")
	lvl1.RuneInGo()

	// Constant
	fmt.Println("\n\nConstant IN GO\n=================")
	lvl1.ConstInGo()

	// Embedded struct
	fmt.Println("\n\nEmbedd Struct IN GO\n=================")
	lvl1.EmbedStructInGo()

	// Closure
	fmt.Println("\n\nClosure IN GO\n=================")
	lvl1.ClosureInGo()

	// Loop
	fmt.Println("\n\nLoop IN GO\n=================")
	lvl1.LoopInGo()

	// Function
	fmt.Println("\n\nFunctions IN GO\n=================")
	lvl1.FunctionsInGo()

	// Generic
	fmt.Println("\n\nGeneric IN GO\n=================")
	lvl1.GenericInGo()

	// Interface and implementation
	fmt.Println("\n\nGeneric IN GO\n=================")
	lvl1.InterfaceImplInGo()

	// Range
	fmt.Println("\n\nRange IN GO\n=================")
	lvl1.RangeInGo()

	// Switch Case
	fmt.Println("\n\nSwitch case IN GO\n=================")
	lvl1.SwitchInGo()

	// If Else
	fmt.Println("\n\nif else IN GO\n=================")
	lvl1.IfElseInGo(30)

	// Map
	fmt.Println("\n\nMap IN GO\n=================")
	lvl1.MapInGo()

	// Method
	fmt.Println("\n\nMethod IN GO\n=================")
	lvl1.MethodInGo()

	// Multiple Return
	fmt.Println("\n\nMultiple Return IN GO\n=================")
	lvl1.MultipleReturnInGo()

	// Recursive
	fmt.Println("\n\nRecursive IN GO\n=================")
	lvl1.RecursiveInGo()

	// Variadic Function
	fmt.Println("\n\nVariadic Function IN GO\n=================")
	lvl1.VariadicInGo()
}

/* _______________________
*    GO LEVEL 2 EXERCISES
*  _______________________
 */
func LevelTwo() {
	// Atomic
	fmt.Println("\n\nAtomic IN GO\n=================")
	lvl2.AtomicInGo()

	// Channel
	fmt.Println("\n\nChannel IN GO\n=================")
	lvl2.ChannelInGo()
}

func main() {
	// LevelOne()
	// LevelTwo()

	done := make(chan bool)
	defer close(done)

	// Generates a channel sending integers
	// From 0 to 9
	range10 := rangeChannel(done, 10)

	for num := range takeFirstN(done, range10, 2) {
		fmt.Println(num)
	}
}

func rangeChannel(done <-chan bool, n int) <-chan int {
	nums := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			nums <- i
		}
	}()

	return nums
}

func takeFirstN(done <-chan bool, dataSource <-chan int, n int) <-chan int {
	// 1
	takeChannel := make(chan int)

	// 2
	go func() {
		defer close(takeChannel)

		// 3
		for i := 0; i < n; i++ {
			select {
			case val, ok := <-dataSource:
				if !ok {
					return
				}
				takeChannel <- val
			case <-done:
				return
			}
		}
	}()
	return takeChannel
}
