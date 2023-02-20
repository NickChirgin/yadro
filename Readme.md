
## Запуск программы

#### Склонировать репозиторий

```
git clone https://github.com/NickChirgin/yadro.git
```

#### Cбилдить бинарный файл командой

```
go build
```

#### Установить его

```
go install 
```

#### Найти куда установился файл и запустить его

##### Linux

```
./yadro filepath.csv
```

##### Windows

```
yadro.exe filepath.csv
```

## Замечания к работе

* Работа производится только с целыми числами. Поддерживается только одна операция в одной ячейке.
* Допустимые операции: сложение, вычитание, умножение, деление(при делении двух целых чисел в го берется только целая часть).
* Обработанные ошибки: некорректная ячейка(как строка, так и столбец), деление на 0, ошибки при открытии файлов.
* Возможна ошибка когда будет бесконечный цикл если две ячейки с формулами ссылаются друг на друга. Как такое обработать не придумал.