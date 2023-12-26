/*
Реализовать все возможные способы остановки выполнения горутины.
*/

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	go func() {
		fmt.Println("Первая горутина закрылась")
	}()

	// Использование канала для сигналов о завершении
	ch := make(chan bool)
	go func() {
		for {
			select {
			case <-ch:
				fmt.Println("Вторая горутина закрылась")
				return
			default:
				// В реальном приложении здесь была бы полезная работа.
				time.Sleep(100 * time.Millisecond) // Чтобы избежать бесконечного цикла
			}
		}
	}()
	ch <- true

	// Использование контекста для управления несколькими горутинами
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Третья горутина закрылась")
				return
			default:
				// В реальном приложении здесь была бы полезная работа.
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
	cancel()

	// Использование таймаутов и дедлайнов
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Четвертая горутина закрылась")
				return
			default:
				// В реальном приложении здесь была бы полезная работа.
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// Использование sync.WaitGroup для ожидания завершения горутин
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Пятая горутина закрылась")
	}()
	wg.Wait()

	// Добавляем задержку перед завершением main для завершения всех горутин
	time.Sleep(5 * time.Second)
	fmt.Println("Главная функция завершена")
}
