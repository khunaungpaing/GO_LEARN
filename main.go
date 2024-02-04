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
	// // Atomic
	// fmt.Println("\n\nAtomic IN GO\n=================")
	// lvl2.AtomicInGo()

	// // Channel
	// fmt.Println("\n\nChannel IN GO\n=================")
	// lvl2.ChannelInGo()

	// fmt.Println("\n\nChannel IN GO\n=================")
	// lvl2.ChannelBuffer()

	// fmt.Println("\n\nErrors IN GO\n=================")
	// lvl2.ErrorInGo()

	// fmt.Println("\n\nGoroutines IN GO\n=================")
	// lvl2.GoRoutines()

	// fmt.Println("\n\nChannel Synchronization IN GO\n=================")
	// lvl2.ChannelSynchronizationInGo()

	// fmt.Println("\n\nChannel Direction IN GO\n==============================")
	// lvl2.ChannelDirectionInGo()

	// fmt.Println("\n\nSelect IN GO\n==============================")
	// lvl2.SelectInGo()

	// fmt.Println("\n\nTimeouts IN GO\n==============================")
	// lvl2.TimeoutInGo()

	// fmt.Println("\n\nNon-blocking Operation IN GO\n==============================")
	// lvl2.NonBlockOpera()

	// fmt.Println("\n\nClosed Channel IN GO\n==============================")
	// lvl2.CloseChannel()

	// fmt.Println("\n\nRange Over Channel IN GO\n==============================")
	// lvl2.RangeOverChannel()

	// fmt.Println("\n\nTimer IN GO\n==============================")
	// lvl2.TimerInGo()

	// fmt.Println("\n\nTicker IN GO\n==============================")
	// lvl2.TickerInGo()

	// fmt.Println("\n\nWorker Pool IN GO\n==============================")
	// lvl2.WorkerPool()

	// fmt.Println("\n\nWait Group IN GO\n=======================")
	// lvl2.WaitGroupInGo()

	// fmt.Println("\n\nRate Limiter IN GO\n===================")
	// lvl2.RateLimiterInGo()

	// fmt.Println("\n\nMutex IN GO\n=================")
	// lvl2.MutexInGo()

	// fmt.Println("\n\nStateful Goroutines IN GO\n=================")
	// lvl2.StatefulGoroutine()

	// fmt.Println("\n\nSorting IN GO\n=================")
	// lvl2.SortingInGo()

	fmt.Println("\n\nCustom Sorting IN GO\n=================")
	lvl2.CustomSorting()
}

func main() {
	// LevelOne()
	LevelTwo()
}
