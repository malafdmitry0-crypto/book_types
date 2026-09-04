# 15. Соседние механизмы, которые тоже называют приведением

> Цель главы: Распознавать библиотечные преобразования и адаптеры.

### 15.1. `errors.As`: извлечение ошибки из дерева

```go
// imports: "errors", "io/fs"
func pathFromError(err error) (string, bool) {
    var target *fs.PathError
    if errors.As(err, &target) {
        return target.Path, true
    }
    return "", false
}
```

`err.(*fs.PathError)` проверяет только непосредственный динамический тип. `errors.As` проходит по обёрнутым ошибкам и учитывает пользовательский метод `As`. Это механизм поиска, а не `T(x)`. Он появился в Go 1.13; пример использует пакет `io/fs` из Go 1.16. Источник: [errors](https://pkg.go.dev/errors#As).

### 15.2. JSON и другие форматы

`json.Unmarshal` декодирует данные по правилам JSON. Для традиционного API `encoding/json` число при декодировании в `any` обычно становится `float64`; `Decoder.UseNumber` позволяет получить `json.Number`. Преобразование через JSON может терять сведения о типах и точность чисел, поэтому это отдельный контракт сериализации. Источник: [encoding/json](https://pkg.go.dev/encoding/json#Unmarshal).

### 15.3. Адаптер функции к интерфейсу

```go
// import "net/http"
func serve(w http.ResponseWriter, r *http.Request) {}

var handler http.Handler = http.HandlerFunc(serve)
```

Здесь сначала функция конвертируется в определённый функциональный тип `http.HandlerFunc`, у которого есть метод `ServeHTTP`, затем присваивается интерфейсу. Это сочетание двух обычных механизмов. Источник: [Effective Go: interfaces](https://go.dev/doc/effective_go#interfaces_and_types).

### 15.4. Ручной маппинг, генерация, cgo

Для структур с разными полями пишут `ToDomain`, `ToDTO`, конструктор или генерируют такой код. Генерация не вводит новых правил типизации. На границе с C действуют также размеры C-типов, правила владения и передачи указателей; функции вроде `C.GoString` выполняют специальное преобразование данных. Это отдельная тема межъязыкового взаимодействия, не универсальный cast. См. [cgo](https://pkg.go.dev/cmd/cgo).

## Самопроверка

Почему errors.As находит ошибку, которую прямой assertion может не найти?

---

[← `unsafe`: интерпретация памяти](14-unsafe.md) · [Оглавление](../README.md) · [Таблица выбора →](16-reference.md)
