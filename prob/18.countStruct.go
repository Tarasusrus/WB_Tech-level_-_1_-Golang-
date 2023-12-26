/*
Реализовать структуру-счетчик,
которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

package main

import (
	"fmt"
	"sync"
)

// Counter - структура счетчика с защищенным мьютексом доступа
type Counter struct {
	num        int // num хранит текущее значение счетчика
	sync.Mutex     // Встроенный мьютекс для защиты доступа к num
}

// Inc инкрементирует счетчик в безопасном для конкуренции режиме
func (c *Counter) Inc() {
	c.Lock()         // Блокируем мьютекс перед изменением счетчика
	defer c.Unlock() // Разблокируем мьютекс

	c.num++ // Увеличиваем счетчик на 1

}

// Value возвращает текущее значение счетчика
func (c *Counter) Value() int {
	return c.num
}
func main() {
	var wg sync.WaitGroup // Используем WaitGroup для ожидания завершения всех горутин
	counter := &Counter{} // Создаем новый счетчик

	// Запускаем 100 горутин для инкрементации счетчика
	for i := 0; i < 100; i++ {
		wg.Add(1) // Уведомляем WaitGroup о новой горутине
		go func() {
			defer wg.Done() // Уведомляем WaitGroup о завершении горутины
			counter.Inc()   // Инкрементируем счетчик

		}()
	}
	wg.Wait() // Ожидаем завершения всех горутин
	fmt.Println("Итоговое значение счетчика:", counter.Value())
}
