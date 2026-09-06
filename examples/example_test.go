package algorithms

import (
	"cmp"
	"fmt"
	"slices"
	"sort"
	"strconv"
)

// Во второй части курса модель проще: callback не возвращают error.
func sampleUsers() []User {
	return []User{{Name: "Анна", Age: 17, Active: true}, {Name: "Борис", Age: 25, Active: true}}
}
func adult(user User) bool  { return user.Age >= 18 }
func active(user User) bool { return user.Active }
func name(user User) string { return user.Name }

type UsersByAge []User

func (s UsersByAge) Len() int           { return len(s) }
func (s UsersByAge) Less(i, j int) bool { return s[i].Age < s[j].Age }
func (s UsersByAge) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func Example_concreteAndInterfaces() {
	users := sampleUsers()
	fmt.Println(FindUser(users, adult), AllUsers(users, active), MapUserNames(users, name))
	fmt.Println(FindIndex(len(users), func(i int) bool { return adult(users[i]) }))
	fmt.Println(AllIndex(len(users), func(i int) bool { return active(users[i]) }))
	texts := make([]string, len(users))
	MapInto(len(users), func(i int) { texts[i] = name(users[i]) })
	fmt.Println(texts)
	query := UserQuery{Values: users, Pred: adult}
	fmt.Println(FindInterface(query), AllInterface(query))
	RunMap(UserNameJob{Source: users, Target: texts, Fn: name})
	fmt.Println(texts)
	// Output:
	// 1 true [Анна Борис]
	// 1
	// true
	// [Анна Борис]
	// 1 false
	// [Анна Борис]
}

func Example_emptyInterface() {
	users := sampleUsers()
	boxed := make([]interface{}, len(users))
	for i, user := range users {
		boxed[i] = user
	}
	index := FindAny(boxed, func(x interface{}) bool { return adult(x.(User)) })
	all := AllAny(boxed, func(x interface{}) bool { return active(x.(User)) })
	raw := MapAny(boxed, func(x interface{}) interface{} { return name(x.(User)) })
	names := make([]string, len(raw))
	for i, value := range raw {
		names[i] = value.(string)
	}
	fmt.Println(index, all, names)
	// Output: 1 true [Анна Борис]
}

func Example_reflection() {
	users := sampleUsers()
	index, err := FindReflect(users, adult)
	if err != nil {
		fmt.Println(err)
		return
	}
	all, err := AllReflect(users, active)
	if err != nil {
		fmt.Println(err)
		return
	}
	raw, err := MapReflect(users, name)
	if err != nil {
		fmt.Println(err)
		return
	}
	names, ok := raw.([]string)
	if !ok {
		fmt.Println("unexpected result type")
		return
	}
	fmt.Println(index, all, names)
	_, err = MapReflect(users, func(s string) int { return len(s) })
	fmt.Println("Несовместимый callback:", err != nil)
	// Output:
	// 1 true [Анна Борис]
	// Несовместимый callback: true
}

func Example_generation() {
	users := sampleUsers()
	fmt.Println(FindUsersGenerated(users, adult), AllUsersGenerated(users, active), MapUsersGenerated(users, name))
	// Output: 1 true [Анна Борис]
}

func Example_generics() {
	users := sampleUsers()
	fmt.Println(Find(users, adult), All(users, active), Map(users, name))
	fmt.Println(Map(users, func(u User) int { return u.Age }))
	fmt.Println(Find([]string{"go", "generic"}, func(s string) bool { return len(s) > 2 }))
	// Output:
	// 1 true [Анна Борис]
	// [17 25]
	// 1
}

func Example_standardLibrary() {
	users := sampleUsers()
	sort.Sort(UsersByAge(users))
	sort.Slice(users, func(i, j int) bool { return users[i].Age < users[j].Age })
	slices.SortFunc(users, func(a, b User) int { return cmp.Compare(a.Age, b.Age) })
	numbers := []int{3, 1, 2}
	slices.Sort(numbers)
	fmt.Println(numbers, Map(users, func(u User) int { return u.Age }))
	fmt.Println(slices.IndexFunc(users, adult), !slices.ContainsFunc(users, func(u User) bool { return !active(u) }))
	// Output:
	// [1 2 3] [17 25]
	// 1 true
}

func Example_iterators() {
	users := sampleUsers()
	for user := range slices.Values(users) {
		fmt.Println(user.Name)
		break
	}
	names := MapSeq(slices.Values(users), name)
	for value := range names {
		fmt.Println("Первое имя:", value)
		break
	}
	found, ok := FindSeq(slices.Values(users), adult)
	fmt.Println(found.Name, ok)
	lengths := MapSeq(MapSeq(slices.Values(users), name), func(s string) int { return len(s) })
	fmt.Println(AllSeq(lengths, func(n int) bool { return n >= 4 }))
	fmt.Println(slices.Collect(MapSeq(slices.Values(users), name)))
	// Output:
	// Анна
	// Первое имя: Анна
	// Борис true
	// true
	// [Анна Борис]
}

func Example_seq2AndErrors() {
	for i, user := range slices.All(sampleUsers()) {
		fmt.Println(i, user.Name)
	}
	parse := func(s string) (int, error) { return strconv.Atoi(s) }
	numbers := MapSeqErr(slices.Values([]string{"1", "x", "3"}), parse)
	// Потребитель А: остановиться на первой ошибке.
	for n, err := range numbers {
		if err != nil {
			fmt.Println("стоп:", err)
			break
		}
		fmt.Println("число:", n)
	}
	// Потребитель Б: пропустить ошибку и продолжить.
	for n, err := range numbers {
		if err != nil {
			continue
		}
		fmt.Println("число:", n)
	}
	// Output:
	// 0 Анна
	// 1 Борис
	// число: 1
	// стоп: strconv.Atoi: parsing "x": invalid syntax
	// число: 1
	// число: 3
}

func Example_sourceCleanup() {
	cleanup := func() { fmt.Println("ресурс освобождён") }
	for i := range Count(10, cleanup) {
		if i == 2 {
			break
		}
		fmt.Println("элемент", i)
	}
	fmt.Println("цикл завершён")
	// Output:
	// элемент 0
	// элемент 1
	// ресурс освобождён
	// цикл завершён
}

func Example_yieldAfterFalse() {
	defer func() {
		fmt.Println("panic:", recover() != nil)
	}()
	// Источник нарушает протокол: продолжает после yield == false.
	broken := func(yield func(int) bool) {
		yield(1)
		yield(2)
	}
	for v := range broken {
		fmt.Println("получено", v)
		break
	}
	// Output:
	// получено 1
	// panic: true
}

func Example_interleave() {
	odd := slices.Values([]int{1, 3, 5})
	even := slices.Values([]int{2, 4})
	fmt.Println(slices.Collect(Interleave(odd, even)))
	left := Count(3, func() { fmt.Println("левый источник закрыт") })
	right := Count(3, func() { fmt.Println("правый источник закрыт") })
	for v := range Interleave(left, right) {
		fmt.Println("элемент", v)
		if v == 1 {
			break
		}
	}
	evens := FilterSeq(slices.Values([]int{1, 2, 3, 4}), func(n int) bool { return n%2 == 0 })
	fmt.Println(slices.Collect(evens))
	// Output:
	// [1 2 3 4 5]
	// элемент 0
	// элемент 0
	// элемент 1
	// правый источник закрыт
	// левый источник закрыт
	// [2 4]
}
