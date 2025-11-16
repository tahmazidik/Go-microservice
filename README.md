# Go Microservice

## Что это за проект

Небольшой микросервис на Go.  
Пока он просто читает настройки из `.env`-файла и выводит их в консоль.  
Дальше к нему добавим HTTP-сервер, БД.

## Структура проекта

```text
.
├── README.md
└── src
    ├── go.mod
    ├── go.sum
    ├── config/
    │   └── local.env
    ├── cmd/
    │   └── app/
    │       └── main.go
    ├── internal/
    │   ├── app/
    │   │   └── app.go
    │   └── config/
    │       ├── config.go
    │       ├── config_helper.go
    │       └── database/
    ├── api/
    │   └── router/
    │       └── router.go
    └── server/
        └── server.go


    cmd/app/main.go — точка входа в приложение (отсюда всё запускается).

    configs/local.env — файл с настройками (логин/пароль БД, хост, порт и т.п.).

    internal/config — работа с конфигурацией всего приложения.

    internal/config/database — отдельная структура с настройками для базы данных.
    
    internal/app — сборка DI-контейнера на go.uber.org/dig.

    api/router — описание HTTP-маршрутов (обёртка над gorilla/mux).

    server — HTTP-сервер, который собирает negroni + mux и запускает ListenAndServe.
