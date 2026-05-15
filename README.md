# visualizationDbDebet

Проект состоит из Go API и встроенного статического frontend.

## Структура

```text
cmd/
  api/
    main.go                 # точка входа backend

internal/
  <feature>/               # debet, contract, customer, contractor, object, ...
    http_handler.go
    service.go
    repository.go
    model.go
    routes.go
  delivery/router/         # сборка роутов
  platform/assets/         # встроенные статические файлы frontend
    embed.go
    web/
      index.html
      src/

go.mod
```

## Правила API и нейминга

- Канонические префиксы:
  - `/debets`
  - `/responses`
  - `/contracts`
  - `/customers`
  - `/contractors`
  - `/objects`
  - `/block-factors`
- Для совместимости оставлены legacy-маршруты (`/debet`, `/response`, `/customer`, `/contractor`, `/contract`, `/blockFactor`).

## Как запускать

1. Создайте `.env` в корне проекта:

```env
DB_HOST=...
DB_PORT=5432
DB_USER=...
DB_PASSWORD=...
DB_NAME=...
DB_SSLMODE=disable
PORT=8080
```

2. Запустите сервер:

```bash
go run ./cmd/api
```

3. Откройте в браузере:

```text
http://localhost:8080
```

Frontend (`internal/platform/assets/web`) раздается встроено через `embed`.
