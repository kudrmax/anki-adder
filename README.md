# Anki adder (WIP)

## Настройка

1. Заполнить config.yaml и положить его по пути `~/.config/anki-adder/config.yaml` (linux/mac; для windows не поддержано)
2. Запустить Anki (иначе не будет работать)
3. Поставить в Anki плагин AnkiConnect (иначе не будет работать)

## Использование через GUI

<img width="1728" height="1080" alt="image" src="https://github.com/user-attachments/assets/2d9c1c65-22a9-4ffb-ae53-04ef4230185c" />

<img width="1728" height="1080" alt="image" src="https://github.com/user-attachments/assets/92fc6e49-7c1f-4f0c-b3c2-f60995b15e75" />

<img width="1728" height="1080" alt="image" src="https://github.com/user-attachments/assets/f2cebd1b-2781-4292-a6ad-e4348258ce1c" />

## Использование командами

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

Добавить предложение в файл .txt, чтобы сохранить предложение на будущее
```
anki save <sentence> 
anki save  # предложение будет взято из буфера обмена
```
