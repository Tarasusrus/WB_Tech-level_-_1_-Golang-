## WB Tech: level # 1 (Golang)
### Как делать задания

- В заданиях никаких устных решений — только код. Одно решение — один файл с хорошо откомментированным кодом. Каждое решение или невозможность решения надо объяснить.

- Разрешается и приветствуется использование любых справочных ресурсов, привлечение сторонних экспертов и т.д. и т.п.


- Основной критерий оценки — четкое понимание «как это работает». Некоторые задачи можно решить несколькими способами, в этом случае требуется привести максимально возможное количество вариантов.

- Можно задавать вопросы, как по условию задач, так и об их решении. Идеальный вариант — продемонстрировать свои решения и получить максимальный фидбэк от опытных разработчиков Wildberries.
## Задания

1. Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

2. Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.


3. Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.


4. Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.

5. Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.



6. Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.


7. Реализовать все возможные способы остановки выполнения горутины.


8. Реализовать конкурентную запись данных в map.


9. Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.


10. Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.


11. Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.


12. Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.


13. Реализовать пересечение двух неупорядоченных множеств.


14. Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.


15. Поменять местами два числа без создания временной переменной.


16. Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.


17. К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.


var justString string
func someFunc() {
v := createHugeString(1 << 10)
justString = v[:100]
}

func main() {
someFunc()
}


18. Реализовать быструю сортировку массива (quicksort) встроенными методами языка.


19. Реализовать бинарный поиск встроенными методами языка.


20. Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.


21. Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.


22. Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».


23. Реализовать паттерн «адаптер» на любом примере.


24. Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.


25. Удалить i-ый элемент из слайса.


26. Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.


27. Реализовать собственную функцию sleep.


28. Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false

## Устные вопросы

1. Какой самый эффективный способ конкатенации строк?


2. Что такое интерфейсы, как они применяются в Go?


3. Чем отличаются RWMutex от Mutex?


4. Чем отличаются буферизированные и не буферизированные каналы?


5. Какой размер у структуры struct{}{}?


6. Есть ли в Go перегрузка методов или операторов?


7. В какой последовательности будут выведены элементы map[int]int?

Пример:
m[0]=1
m[1]=124
m[2]=281


8. В чем разница make и new?


9. Сколько существует способов задать переменную типа slice или map?


10. Что выведет данная программа и почему?


func update(p *int) {
b := 2
p = &b
}

func main() {
var (
a = 1
p = &a
)
fmt.Println(*p)
update(p)
fmt.Println(*p)
}

11. Что выведет данная программа и почему?


func main() {
wg := sync.WaitGroup{}
for i := 0; i < 5; i++ {
wg.Add(1)
go func(wg sync.WaitGroup, i int) {
fmt.Println(i)
wg.Done()
}(wg, i)
}
wg.Wait()
fmt.Println("exit")
}

12. Что выведет данная программа и почему?


func main() {
n := 0
if true {
n := 1
n++
}
fmt.Println(n)
}


13. Что выведет данная программа и почему?


func someAction(v []int8, b int8) {
v[0] = 100
v = append(v, b)
}

func main() {
var a = []int8{1, 2, 3, 4, 5}
someAction(a, 6)
fmt.Println(a)
}


14. Что выведет данная программа и почему?


func main() {
slice := []string{"a", "a"}

func(slice []string) {
slice = append(slice, "a")
slice[0] = "b"
slice[1] = "b"
fmt.Print(slice)
}(slice)
fmt.Print(slice)
}
