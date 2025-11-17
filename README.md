# Go Microservice

## Что это за проект

Небольшой микросервис на Go. 
Пока он просто читает настройки из `.env`-файла и выводит их в консоль.

- HTTP-сервер на `gorilla/mux` + `negroni`;
- DI-контейнер на `go.uber.org/dig`;
- пример домена **Book**: создание книги через `POST /books`.

## Структура проекта

```text
.
├── cmd/
│   └── app/
│       └── main.go
├── config/
│   └── local.env
├── internal/
│   ├── app/
│   │   └── app.go
│   ├── config/
│   │   ├── config.go
│   │   ├── config_helper.go
│   │   └── database/
│   ├── entities/
│   │   └── book/
│   │       └── book_entity.go
│   ├── models/
│   │   └── book/
│   │       └── book_create_models.go
│   ├── repositories/
│   │   └── book/
│   │       └── book_repositories.go
│   └── services/
│       └── book/
│           └── book_service.go
├── api/
│   ├── controllers/
│   │   └── book/
│   │       ├── book_controller.go
│   │       └── book_controller_route.go
│   └── router/
│       └── router.go
├── server/
│   └── server.go
├── go.mod
├── go.sum
└── README.md
````


    cmd/app/main.go — точка входа в приложение (отсюда всё запускается).

    configs/local.env — файл с настройками (логин/пароль БД, хост, порт и т.п.).

    internal/config — работа с конфигурацией всего приложения.

    internal/config/database — отдельная структура с настройками для базы данных.
    
    internal/app — сборка DI-контейнера на go.uber.org/dig.

    api/router — описание HTTP-маршрутов (обёртка над gorilla/mux).

    server — HTTP-сервер, который собирает negroni + mux и запускает ListenAndServe.



### Что за папки

* `cmd/app/main.go`
  Точка входа в приложение. Собирает DI-контейнер (`internal/app`), достаёт из него `server.Server` и запускает HTTP-сервер.

* `config/local.env`
  Локальные настройки (порт сервера, параметры БД и т.п.). Читаются в `internal/config`.

---

### internal/

Всё, что не должно использоваться извне (внешними модулями). «Внутренности» микросервиса.

* `internal/config`
  Работа с конфигурацией:

  * чтение `.env`;
  * маппинг в структуру `Config`;
  * отдельный под-пакет `database` под настройки БД.

* `internal/app/app.go`
  Сборка DI-контейнера на `go.uber.org/dig`:

  * регистрируются конфиг, сервер, роутер;
  * отдельно регистрируются все компоненты домена Book (репозиторий, сервис, контроллер, маршрутизация).

* `internal/entities/book/book_entity.go`
  **Entity** — внутренняя сущность книги для БД:

  ```go
  type Entity struct {
      Uuid uuid.UUID
      Name string
  }
  ```

  Это «карточка книги» внутри системы: как она хранится и живёт в логике, независимо от HTTP.

* `internal/models/book/book_create_models.go`
  **CreateModel** — модель входных данных для сервиса:

  ```go
  type CreateModel struct {
      Name string `json:"name" form:"name"`
  }
  ```

  То, что присылает клиент в `POST /books` (JSON/форма).

* `internal/repositories/book/book_repositories.go`
  **Repository** — слой доступа к данным для книг.

  Хранит слайс `[]Entity` (учебная «БД в памяти») и умеет добавлять книги:

  > Repository — отдельный слой в приложении, который отвечает за работу с хранилищем данных.
  > Сервис говорит ему: «Сохрани/найди/удали вот эту `Entity`», а как именно — знает только он.

* `internal/services/book/book_service.go`
  **Service** — бизнес-логика для книг.

  При создании книги:

  * принимает `CreateModel` от контроллера;
  * собирает `Entity` (генерирует `uuid`, копирует `Name`);
  * передаёт `Entity` в репозиторий.

  Сервис — «адвокат по сути дела»: здесь должны жить правила, проверки, транзакции.

---

### api/

Всё, что касается HTTP-слоя.

* `api/controllers/book/book_controller.go`
  **Controller** — обёртка над сервисом для HTTP.

  Делает 4 вещи:

  1. Читает тело HTTP-запроса.
  2. Превращает его в Go-структуру (`CreateModel`).
  3. Передаёт модель в сервис.
  4. Формирует HTTP-ответ.

  Контроллер — это прослойка между HTTP-миром и сервисом.

* `api/controllers/book/book_controller_route.go`
  **ControllerRoute** — описывает, какие URL и методы привязаны к контроллеру:

  * вешает `POST /books` на `Controller.CreateBook`;
  * при необходимости здесь же будут и другие маршруты (`GET /books`, `DELETE /books/{id}` и т.д.).

* `api/router/router.go`
  Общий роутер приложения:

  * создаёт `*mux.Router`;
  * вызывает `BookRoutes.Route(router)`, чтобы зарегистрировать все эндпоинты книг;
  * в будущем сюда можно добавить маршруты других доменов (users, orders и т.п.).

---

### server/

* `server/server.go`
  HTTP-сервер:

  * принимает `*router.Router`, вызывает `InitRoutes()` и получает настроенный `*mux.Router`;
  * оборачивает его в `negroni` (логирование, recovery и др. middleware);
  * запускает `ListenAndServe` на порту из конфига.

---

## Как работает создание книги

Путь запроса `POST /books`:

1. **Клиент** отправляет JSON:

   ```json
   {"name": "Мастер и Маргарита"}
   ```

2. **Router** (`api/router`) направляет запрос на контроллер книг по правилу из `ControllerRoute`.

3. **Controller** (`CreateBook`):

   * декодирует JSON в `CreateModel`;
   * вызывает `service.Create(model)`.

4. **Service**:

   * из `CreateModel` собирает `Entity` (генерирует `uuid.New()`);
   * передаёт `Entity` в репозиторий.

5. **Repository**:

   * добавляет сущность в своё хранилище (слайс в памяти);
   * сейчас просто печатает список книг в консоль.

6. **Controller** возвращает клиенту `HTTP 200 OK` (позже можно отдавать созданную книгу с её `uuid`).

Такой разрез по слоям позволяет легко:

* менять хранилище (репозиторий) без переписывания контроллеров;
* добавлять бизнес-правила в сервис;
* расширять HTTP-API новыми методами, не трогая ядро домена.

