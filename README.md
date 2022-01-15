# l1
l1 task repository

## task 1
Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
## task 2
Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
## task 3
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
## task 4
Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте. <br>
<br>Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
## task 5
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
## task 6
Реализовать все возможные способы остановки выполнения горутины. 
## task 7
Реализовать конкурентную запись данных в map.
## task 8
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
## task 9
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
## task 10
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
## task 11
Реализовать пересечение двух неупорядоченных множеств.
## task 12
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
## task 13
Поменять местами два числа без создания временной переменной.
## task 14
Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
## task 15
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

`var` justString `string`<br>
`func` someFunc() {<br>
  v := createHugeString(1 << 10)<br>
  justString = v[:100]<br>
}<br>
<br>
`func` main() {<br>
  someFunc()<br>
}
## task 16
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
## task 17
Реализовать бинарный поиск встроенными методами языка.
## task 18
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.
## task 19
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.
## task 20
Разработать программу, которая переворачивает слова в строке. 
Пример: «snow dog sun — sun dog snow».
## task 21
Реализовать паттерн «адаптер» на любом примере.
## task 22
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
## task 23
Удалить i-ый элемент из слайса.
## task 24
Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
## task 25
Реализовать собственную функцию sleep.
## task 26
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например: 
abcd — true <br>
abCdefAaf — false <br>
	aabcd — false <br>
# Устные вопросы
### 1.	Какой самый эффективный способ конкатенации строк?

### 2.	Что такое интерфейсы, как они применяются в Go?

### 3.	Чем отличаются `RWMutex` от `Mutex`?

### 4.	Чем отличаются буферизированные и не буферизированные каналы?

### 5.	Какой размер у структуры `struct{}{}`?

### 6.	Есть ли в Go перегрузка методов или операторов?

### 7.	В какой последовательности будут выведены элементы `map[int]int`? <br>
<br>
Пример:		<br>
m[0]=1		<br>
m[1]=124	<br>
m[2]=281	<br>

### 8.	В чем разница `make` и `new`?

### 9.	Сколько существует способов задать переменную типа `slice` или `map`?

### 10.	Что выведет данная программа и почему? <br>
						<br>
func update(p *int) {	<br>
  b := 2		<br>
  p = &b		<br>
}			<br>
			<br>
func main() {		<br>
  var (			<br>
     a = 1		<br>
     p = &a		<br>
  )			<br>
  fmt.Println(*p)	<br>
  update(p)		<br>
  fmt.Println(*p)	<br>
}			<br>

### 11.	Что выведет данная программа и почему?

func main() {					<br>
  wg := sync.WaitGroup{}			<br>
  for i := 0; i < 5; i++ {			<br>
     wg.Add(1)					<br>
     go func(wg sync.WaitGroup, i int) {	<br>
        fmt.Println(i)				<br>
        wg.Done()				<br>
     }(wg, i)					<br>
  }						<br>
  wg.Wait()					<br>
  fmt.Println("exit")				<br>
}						<br>

### 12.	Что выведет данная программа и почему?

func main() {		<br>
  n := 0		<br>
  if true {		<br>
     n := 1		<br>
     n++		<br>
  }			<br>
  fmt.Println(n)	<br>
}			<br>

### 13.	Что выведет данная программа и почему?

func someAction(v []int8, b int8) {	<br>
  v[0] = 100				<br>
  v = append(v, b)			<br>
}					<br>
					<br>
func main() {				<br>
  var a = []int8{1, 2, 3, 4, 5}		<br>
  someAction(a, 6)			<br>
  fmt.Println(a)			<br>
}					<br>

### 14.	Что выведет данная программа и почему?

func main() {				<br>
  slice := []string{"a", "a"}		<br>
					<br>
  func(slice []string) {		<br>
     slice = append(slice, "a")		<br>
     slice[0] = "b"			<br>
     slice[1] = "b"			<br>
     fmt.Print(slice)			<br>
  }(slice)				<br>
  fmt.Print(slice)			<br>
}					<br>
