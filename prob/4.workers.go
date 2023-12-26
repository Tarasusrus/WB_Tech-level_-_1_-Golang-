/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	// проверка ввода
	if len(os.Args) != 2 {
		fmt.Println("Введите количество воркеров: go run 4.workers.go <numOfWorkers>")
		return
	}
	// Получаем данные от пользователя
	numOfWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Ошибка при преобразовании числа воркеров:", err)
		return
	}

	//создание канала для передачи данных
	dataChannel := make(chan int)

	// обработка Ctrl+C
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	//Запуск главного потока для записи данных в канал
	go func() {
		defer close(dataChannel)
		for i := 0; ; i++ {
			dataChannel <- i
		}
	}()

	//Запуск воркеров
	for i := 0; i < numOfWorkers; i++ {
		go worker(i, dataChannel)
	}

	//Ожидание сигнала для завершения
	<-stopChan
	fmt.Println("Программа завершена")

}

// Функция воркера, читающего данные и выводящего их на экран.
func worker(i int, channel <-chan int) {
	for data := range channel {
		fmt.Printf("Воркер %d получил данные: %d\n", i, data)
		time.Sleep(1 * time.Second)
	}

}
