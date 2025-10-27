### Hexlet tests and linter status:
[![Actions Status](https://github.com/AndreyHexMentor/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/AndreyHexMentor/go-project-242/actions)

# Hexlet Path Size

`hexlet-path-size` — утилита для вычисления размера файлов и директорий с различными опциями (рекурсивный подсчёт, вывод в человекочитаемом формате, включая скрытые файлы).

## Установка

### С помощью Go

Если у вас установлен Go, вы можете собрать и установить утилиту, выполнив следующие команды:

```bash
go install github.com/AndreyHexMentor/go-project-242/cmd/hexlet-path-size@latest
```

### Использование

Утилита поддерживает следующие флаги:

- `-r` или `--recursive` — **рекурсивно** вычисляет размер директорий.

- `-H или --human` — выводит размер в **человекочитаемом** **формате** (например, 1.5MB, 500KB).

- `-a или --all` — включает **скрытые** файлы и директории.

### Asciinema

[![asciicast](https://asciinema.org/a/EMMBO9kuVWT9xX2buV49d7fVu.svg)](https://asciinema.org/a/EMMBO9kuVWT9xX2buV49d7fVu)

### Примеры

Получение размера одного файла:

```bash
hexlet-path-size file.txt
```

Получение размера директории (не рекурсивно):

```bash
hexlet-path-size path/to/directory
```

Рекурсивный подсчёт размера директории, включая скрытые файлы и директории:

```bash
hexlet-path-size -r -a path/to/directory
```

Получение размера с человекочитаемым выводом:

```bash
hexlet-path-size -H path/to/directory
```

Получение размера с человекочитаемым выводом для файла:

```bash
hexlet-path-size -H file.txt
```