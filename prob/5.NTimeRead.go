/*
5. Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/

package main

import (
	"fmt"
	"time"
)

var n = 5

func main() {
	// Проверка, что N > 0
	if n <= 0 {
		fmt.Println("Введите n > 0")
		return
	}
	// Создаем канал в который будем отправлять сообщения
	ch := make(chan int)

	//запускае функцию для отправки сообщения в канал
	sendMesInChanForNSec(ch, n)

	//читаем сообщения из канала
	for msg := range ch {
		fmt.Println(msg)
	}
}

// Реализуем функцию для отправки сообщений в канал. Принимает канал и n - секунд работы
func sendMesInChanForNSec(ch chan<- int, n int) {
	// создаем таймер
	timer := time.NewTimer(time.Duration(n) * time.Second)

	go func() {
		defer close(ch)
		for i := 0; ; i++ {
			select {
			case <-timer.C:
				return
			case ch <- i:
			}
		}
	}()
}
