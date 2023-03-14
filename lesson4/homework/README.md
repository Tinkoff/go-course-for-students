# Домашняя работа №4

Реализовать программу, которая подсчитывает размер файлов в объектном хранилище.

## Функциональные требования

В [dir_sizer.go](storage%2Fdir_sizer.go) вы видите интерфейс `DirSizer` и структуру `sizer`, которая его реализует. Это и есть ваша программа. 

```go
type DirSizer interface {
    // Size calculate a size of given Dir, receive a ctx and the root Dir instance
    // will return Result or error if happened
    Size(ctx context.Context, d Dir) (Result, error)
}

// sizer implement the DirSizer interface
type sizer struct {
    // maxWorkersCount number of workers for asynchronous run
    maxWorkersCount int
    
    // TODO: add other fields as you wish
}

```

Вам нужно реализовать функцию `Size(ctx context.Context, d Dir) (Result, error)` в [dir_sizer.go](storage%2Fdir_sizer.go)

Объектное хранилище реализовано в файле [storage.go](storage%2Fstorage.go) 

```go
// File represent a file object
type File interface {
	// Name return a fully qualified file name
	Name() string
	// Stat returns a size of file or error
	Stat(ctx context.Context) (int64, error)
}

// Dir represent a dir object
type Dir interface {
	// Name return a fully qualified dir name
	Name() string
	// Ls return a set of Dir and a set of File or error if happened
	Ls(ctx context.Context) ([]Dir, []File, error)
}
```

Для вашего удобства мы реализовали хранилище для работы с реальной файловой системой и файловое хранилище в памяти.

При необходимости в структуру `sizer` вы можете добавлять дополнительные поля.


## Требования к коду

* использование пакета context;
* использование goroutine;
* тесты должны проходить без модификаций;
* допускается использование пакета golang.org/x/sync;
* допускается написание собственных тестов.


## Запуск тестов

1. зайти в терминале в каталог homework
2. вызвать ```go test ./... -v```


## Критерии оценки

* проходят линтеры и тесты, кроме `Test_DirSizerAsync`: до 4-x баллов;
* `Test_DirSizerAsync` тест проходит: до 2-x баллов;
* чистота кода: до 3-x баллов;
* использование `maxWorkersCount` (максимальное количество goroutine обрабатывающих объекты в один момент времени, без учета main): 1 балл;

Максимальное количество баллов, которые можно получить - 10.
