package main

import (
	"errors"
	"fmt"

	"github.com/golang-ninja-courses/error-handling-mastery/tasks/05-errors-best-practices/empty-struct-problem/pkga"
	"github.com/golang-ninja-courses/error-handling-mastery/tasks/05-errors-best-practices/empty-struct-problem/pkgb"
)

var (
	aVal  error = pkga.EOF{}
	aPtr  error = new(pkga.EOF)
	aPtr2 error = new(pkga.EOF)

	bVal error = pkgb.EOF{}
	bPtr error = new(pkgb.EOF)
)

func main() {
	fmt.Println("\nValues errors from diff packages:")
	fmt.Printf("%p %p\n", &aVal, &bVal) // 0x102a08f70 0x102a08f70 - разные адреса, разные пакеты.
	fmt.Println(aVal == bVal)           // Очевидно - false
	fmt.Println(errors.Is(aVal, bVal))  // false

	fmt.Println("\nPointers errors from diff packages:")
	fmt.Printf("%p %p\n", aPtr, bPtr) // 0x102a3edb0 0x102a3edb0 - одинаковые адреса при разных пакетах!
	// Неочевидно - тоже false, несмотря на один и тот же адрес:
	fmt.Println(aPtr == bPtr)
	fmt.Println(errors.Is(aPtr, bPtr)) // false
	// Обратимся к спецификации:
	//   Interface values are comparable. Two interface values are equal if they have identical dynamic types
	//   and equal dynamic values or if both have value nil.
	//
	// Несмотря на равенство указателей (dynamic values), типы этих указателей (dynamic types) различны,
	// соответственно две ошибки (два интерфейса error) тоже различны:
	fmt.Printf("%T(%p) %T(%p)\n", aPtr, aPtr, bPtr, bPtr) // *pkga.EOF(0x102a3edb0) *pkgb.EOF(0x102a3edb0)

	fmt.Println("\nCompare values & pointers:")
	// Сравнение указателя и неуказателя - очевидно false:
	fmt.Println(aPtr == aVal)          // false
	fmt.Println(errors.Is(aPtr, aVal)) // false
	fmt.Println(aPtr == bVal)          // false
	fmt.Println(errors.Is(aPtr, bVal)) // false

	fmt.Println("\nCompare errors from one package:")
	// Получаем возможное равенство всё-таки только в пределах одного пакета:
	fmt.Println(aPtr == aPtr2)      // Возможно true
	fmt.Println(aVal == pkga.EOF{}) // true
}
