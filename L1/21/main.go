package main

import "fmt"

/*
Допустим, у нас есть старый тип логгера, который пишет лог через метод WriteLog(string),
а наш новый код ожидает интерфейс Logger с методом Log(string).
*/

// OldLogger Старый логгер
type OldLogger struct{}

// Logger Новый интерфейс
type Logger interface {
	Log(message string)
}

/*
Адаптер реализует интерфейс Logger,
но внутри содержит ссылку на OldLogger и делегирует вызов:
*/

// LoggerAdapter Адаптер для OldLogger
type LoggerAdapter struct {
	oldLogger *OldLogger
}

// Клиент работает с любым Logger
func useLogger(l Logger) {
	l.Log("Привет, адаптер!")
}

func (o *OldLogger) WriteLog(message string) {
	fmt.Println("[OLD LOG]:", message)
}

func (a *LoggerAdapter) Log(message string) {
	// Перенаправляем вызов
	a.oldLogger.WriteLog(message)
}

func main() {
	// Используем старый логгер через адаптер
	old := &OldLogger{}
	adapter := &LoggerAdapter{oldLogger: old}

	useLogger(adapter)
	// Вывод [OLD LOG]: Привет, адаптер!
}

/*
Можно использовать Adapter, когда:
У нас уже есть существующий код, менять который нельзя
Нужно сделать его совместимым с новым интерфейсом
Необходимо использовать чужую библиотеку с неудобным интерфейсом
Все это позволяет переиспользовать старый код, устраняя дублирование
делая систему более гибкой и разсиряемой, но при этом добавяет еще
один уровень абстракции немного усложняя код, иногда приводит к избыточным
врапперам, если адаптеров слишком много.
Пример реального использования:
В Go стандартная библиотека io часто использует адаптеры, например:
bytes.Buffer реализует io.Writer, io.Reader, io.ReaderFrom и т.д.;
bufio.NewReader оборачивает io.Reader, добавляя буферизацию (адаптирует поведение).
*/
