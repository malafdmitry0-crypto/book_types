package basics

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode/utf8"
	"unsafe"

	"gotypes/internal/checkedint"
)

func Example_numbers() {
	var n int = 42
	wide := int64(n)
	fraction := float64(n)

	var f float64 = -3.9
	whole := int(f) // -3: дробная часть отбрасывается к нулю

	var big uint16 = 300
	small := uint8(big) // 44: остаются младшие 8 бит

	var negative int = -1
	unsigned := uint8(negative) // 255

	var z complex128 = 3 + 4i
	compact := complex64(z)
	realPart := real(z)
	imagPart := imag(z)
	constructed := complex(float64(n), 0)

	// var bad int64 = n // разные типы: нужна int64(n)
	// _ = int(true)    // bool → число не предусмотрено
	// _ = bool(1)      // число → bool не предусмотрено
	// _ = float64(z)   // сначала явно выбираем real(z) или imag(z)
	fmt.Println(wide, fraction, whole, small, unsigned, compact, realPart, imagPart, constructed)
	// Output: 42 42 -3 44 255 (3+4i) 3 4 (42+0i)
}
func Example_wraparound() {
	// import "math"
	// import "gotypes/internal/checkedint"
	var top int64 = math.MaxInt64
	wrapped := top + 1 // переполнение определено языком: ошибки нет
	checked, err := checkedint.Add(top, 1)
	fmt.Println(wrapped, checked, err)
	// Output: -9223372036854775808 0 integer overflow
}
func Example_constants() {
	const answer = 42
	var a int8 = answer
	var b float64 = answer
	var c complex128 = answer

	inferred := 42 // int: тип по умолчанию
	const typed int = 42
	// var d int64 = typed // typed уже имеет тип int

	// _ = uint8(300) // ошибка компиляции: константа не помещается
	// _ = int(3.9)   // ошибка компиляции: константа не представима как int
	_ = int(3.0) // допустимо: значение константы целое
	fmt.Println(a, b, c, inferred, typed)
	// Output: 42 42 (42+0i) 42 42
}
func Example_assignment() {
	type Numbers []int
	var named Numbers = []int{1, 2}
	var plain []int = named // одинаковые underlying types, один тип неименованный

	type OtherNumbers []int
	// var other OtherNumbers = named // оба типа определённые
	other := OtherNumbers(named) // явная конверсия допустима
	fmt.Println(named, plain, other)
	// Output: [1 2] [1 2] [1 2]
}
func Example_nil() {
	var p *int = nil
	typedNil := (*int)(nil)
	var s []int = nil
	var m map[string]int = nil
	var fn func() = nil
	var ch chan int = nil
	var x any = nil

	n := 42
	ptr := &n
	value := *ptr
	// _ = int(nil) // недопустимо
	fmt.Println(p == nil, typedNil == nil, s == nil, m == nil, fn == nil, ch == nil, x == nil, *ptr, value)
	// Output: true true true true true true true 42 42
}
func Example_strings() {
	text := "Привет"
	bytes := []byte(text) // UTF-8 байты
	runes := []rune(text) // кодовые точки Unicode
	fromBytes := string(bytes)
	fromRunes := string(runes)

	letter := string(rune(65))      // "A", не "65"
	replacement := string(rune(-1)) // "�": недопустимая кодовая точка

	type Label string
	label := Label(text)
	original := string(label)
	fmt.Println(len(bytes), len(runes), fromBytes, fromRunes, letter, replacement, original)
	// Output: 12 6 Привет Привет A � Привет
}
func Example_runes() {
	// import "unicode/utf8"
	text := "Go\xffя" // 'G', 'o', недопустимый байт, 'я' (два байта)
	var parts []string
	for offset, r := range text { // offset — байтовый индекс, r — руна
		parts = append(parts, fmt.Sprintf("%d:%U", offset, r))
	}
	fmt.Println(strings.Join(parts, " "))
	fmt.Println(len(text), utf8.RuneCountInString(text), utf8.ValidString(text))
	// Output:
	// 0:U+0047 1:U+006F 2:U+FFFD 3:U+044F
	// 5 4 false
}
func Example_builder() {
	// import "strings"
	var b strings.Builder
	for i := 0; i < 3; i++ {
		b.WriteString("go")
	}
	built := b.String()                // без копирования накопленного буфера
	cloned := strings.Clone(built[:2]) // отдельная аллокация для короткой подстроки
	fmt.Println(built, cloned)
	// Output: gogogo go
}
func Example_parse() {
	// import "strconv"
	n, err := strconv.Atoi("42")
	if err != nil {
		// Обработать ошибку формата или диапазона.
	}
	decimal := strconv.Itoa(42)
	hex := strconv.FormatInt(255, 16) // "ff"
	parsed, parseErr := strconv.ParseInt("ff", 16, 64)
	flag, boolErr := strconv.ParseBool("true")
	number, floatErr := strconv.ParseFloat("3.14", 64)
	fmt.Println(n, err, decimal, hex, parsed, parseErr, flag, boolErr, number, floatErr)
	// Output: 42 <nil> 42 ff 255 <nil> true <nil> 3.14 <nil>
}
func Example_containers() {
	src := []int{1, 2, 3}
	// dst := []int64(src) // запрещено
	dst := make([]int64, len(src))
	for i, n := range src {
		dst[i] = int64(n)
	}

	// Не работают также []string → []any и map[string]int → map[string]int64.
	fmt.Println(src, dst)
	// Output: [1 2 3] [1 2 3]
}
func Example_channels() {
	ch := make(chan int)
	var recv <-chan int = ch
	var send chan<- int = ch
	explicit := (<-chan int)(ch)
	// both := (chan int)(recv) // восстановить оба направления нельзя
	fmt.Println(recv == explicit, send == ch)
	// Output: true true
}
func Example_arrayPointer() {
	s := []int{10, 20, 30}
	p := (*[2]int)(s)
	p[0] = 99 // s[0] тоже стал 99

	// _ = (*[4]int)(s) // panic: len(s) < 4
	fmt.Println(s, p)
	// Output: [99 20 30] &[99 20]
}
func Example_arrayPointerAppend() {
	s := make([]int, 2, 4)
	p := (*[2]int)(s)
	s = append(s, 3) // cap достаточно: массив общий
	s[0] = 7
	fmt.Println(p[0])
	s = append(s, 4, 5) // cap не хватает: append выделяет новый массив
	s[0] = 8
	fmt.Println(p[0], s[0]) // p по-прежнему смотрит на старый массив
	// Output:
	// 7
	// 7 8
}
func Example_arrayCopy() {
	s := []int{10, 20, 30}
	a := [2]int(s)
	a[0] = 99 // s[0] остаётся 10

	// _ = [4]int(s) // panic: len(s) < 4
	fmt.Println(s, a)
	// Output: [10 20 30] [99 20]
}
func Example_arrayKey() {
	// import "crypto/sha256"
	digest := sha256.New()
	digest.Write([]byte("go"))
	sum := digest.Sum(nil) // []byte: срез не сравним и не годится в ключ map
	key := [32]byte(sum)   // копия фиксированного размера: сравнима
	seen := map[[32]byte]bool{key: true}
	fmt.Println(len(sum), seen[sha256.Sum256([]byte("go"))])
	// Output: 32 true
}
func Example_arrayShallow() {
	one, two := 1, 2
	s := []*int{&one, &two}
	a := [2]*int(s)
	*a[0] = 10 // объект общий: s[0] тоже указывает на 10
	a[1] = nil // элемент массива свой: s[1] не меняется
	fmt.Println(*s[0], s[1] != nil)
	// Output: 10 true
}
func Example_unsafe() {
	// import "unsafe"
	bits := uint32(0x3f800000)
	number := *(*float32)(unsafe.Pointer(&bits)) // 1 на соответствующем представлении
	fmt.Println(number)
	// Output: 1
}
func Example_layout() {
	// import "unsafe"
	type header struct {
		flag byte
		size uint32
	}
	var h header
	fmt.Println(unsafe.Sizeof(h), unsafe.Alignof(h.size), unsafe.Offsetof(h.size))
	// Output: 8 4 4
}
func Example_stdlibInsteadOfUnsafe() {
	// import "encoding/binary"
	// import "math"
	number := math.Float32frombits(0x3f800000)
	raw := []byte{0, 0, 0, 1}
	value := binary.BigEndian.Uint32(raw)
	buf := []byte("view")
	view := unsafe.String(unsafe.SliceData(buf), len(buf)) // buf нельзя менять, пока жив view
	fmt.Println(number, value, view)
	// Output: 1 1 view
}
