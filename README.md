# Anki adder

## Настройка

1. Заполнить config.yaml и положить его по пути `~/.config/anki-adder/config.yaml` (linux/mac; для windows не поддержано)
2. Запустить Anki (иначе не будет работать)
3. Поставить в Anki плагин AnkiConnect (иначе не будет работать)

## Использование

Импортировать заметки из csv файла

```
anki add --file <file_path>
anki add -f <file_path>
```

Импортировать заметки из csv, взяв csv из буфера обмена

```
anki add --clipboard
anki add -c
```

Добавить предложение в файл sentences.txt, чтобы сохранить предложение на будущее
```
anki save <sentence> 
anki save  # предложение будет взято из буфера обмена
```
