/*
Реализовать собственную функцию sleep.
*/

package main

import (
	"fmt"
	"time"
)

// mySleep приостанавливает выполнение программы на заданное количество секунд.
func mySleep(sec int) {
	<-time.After(time.Duration(sec) * time.Second)

}

func main() {
	fmt.Println("Начинаем спать...")
	mySleep(2) // "спим" 2 секунды
	fmt.Println("Проснулись!")
}
