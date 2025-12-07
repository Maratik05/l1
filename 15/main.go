package main

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	// Проблема в том что из переменной v используется только последние 100 байтов а оставшуюся память очистить не может сборщик мусора
	//решение такое: скопировать эти последние 100 байтов в слайс и преобразовать в строку
	b := make([]byte, 100)
	copy(b, v[:100])
	justString = string(b)
}

func main() {

}
