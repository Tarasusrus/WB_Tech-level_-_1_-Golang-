/*
Реализовать конкурентную запись данных в map.
data race - карта не безопасный тип данных несколько горутин могу начать менять данные в результате возникнет паника
используя мьютекс можно обеспечить безопасный доступ
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	//Реализуем через sync.map
	var myMap sync.Map // специально разработан для использования в конкурентных сценариях

	var wg sync.WaitGroup

	//запускаем несколько горутин которые будут записывать в карту через цикл
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			myMap.Store(i, i*2)
		}(i)

	}
	wg.Wait() // ждем завершения работы горутин

	// //Использование Range гарантирует, что итерация происходит безопасно и не приведет к неопределенному поведению.
	myMap.Range(func(key, value interface{}) bool {
		fmt.Printf("Ключ: %v, Значение: %v\n", key, value)
		return true
	})

	//Использование sync.Mutex для защиты стандартной карты
	var mutex sync.Mutex

	myMap2 := make(map[int]int)

	// Запускаем несколько горутин, которые будут записывать в карту
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mutex.Lock() // Блокируем доступ к карте
			myMap2[i] = i * 2
			mutex.Unlock() // // Освобождаем доступ
		}(i)
	}
	wg.Wait() // ждем завершения работы горутин

	fmt.Println("Данные в карте:", myMap2)
}
