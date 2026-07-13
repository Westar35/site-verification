# dns-check-service

Сервис подтверждения владения доменом через DNS TXT-запись (по аналогии с Google/Yandex Search Console). Пользователь передаёт домен, сервис выдаёт TXT-запись для DNS-зоны, после добавления записи проверяет её и сохраняет факт подтверждения за аккаунтом.

Стек: Go, Protocol Buffers + [ConnectRPC](https://connectrpc.com/) (генерация через [buf](https://buf.build/)), MongoDB.

## Как это работает

1. `GetDomainVerificationRecord` — по домену возвращает TXT-запись (`record_name`/`record_value`), которую нужно добавить в DNS-зону. В базу не пишет.
2. Пользователь добавляет TXT-запись в DNS-зону своего домена.
3. `VerifyDomain` — резолвит DNS, ищет ожидаемую запись. Если нашёл — сохраняет подтверждение за текущим `Account ID`; если нет — возвращает `NOT_VERIFIED`, ничего не меняя в базе.
4. `GetDomainStatus` / `ListVerifiedDomains` — чтение статуса/списка подтверждённых доменов текущего аккаунта.
5. `DeleteDomainVerification` — идемпотентное удаление подтверждения текущего аккаунта.

`Account ID` передаётся в заголовке `Account-Id` на каждый запрос (кроме `/healthz`).

## Конфигурация

Все параметры читаются из переменных окружения. При локальном запуске (`go run`)
они подхватываются из файла `.env` (через `godotenv`). Шаблон — `.env.example`:

```
cp .env.example .env
```

| Переменная | Дефолт | Назначение |
|---|---|---|
| `HTTP_ADDR` | `:8080` | адрес, на котором слушает HTTP-сервер |
| `MONGO_URI` | `mongodb://localhost:27017` | строка подключения к MongoDB |
| `MONGO_DATABASE` | `dns-check-service` | имя базы |
| `MONGO_COLLECTION` | `domains` | имя коллекции |
| `DNS_SERVERS` | `8.8.8.8:53` | DNS-сервер для проверки TXT-записей |

## Запуск

### Вариант 1 — Docker

```
docker compose up --build
```
Поднимает MongoDB и сам сервис; сервис доступен на `http://localhost:8080`.
`MONGO_URI` внутри compose переопределяется на `mongodb://mongo:27017`.

### Вариант 2 — Mongo в Docker, сервис локально

```
make mongo-up   # поднимает MongoDB в контейнере на localhost:27017
make run        # go run ./cmd  (читает .env)
```

## Проверка через curl

# health-check
curl.exe http://localhost:8080/healthz

# 1. получить TXT-запись
curl.exe -X POST http://localhost:8080/site_verification.v1.SiteVerificationService/GetDomainVerificationRecord -H "Content-Type: application/json" -H "Account-Id: acc-1" -d '{"domain":"HTTPS://Example.COM:8080/path?x=1"}'

# 2. проверка домена
curl.exe -X POST http://localhost:8080/site_verification.v1.SiteVerificationService/VerifyDomain -H "Content-Type: application/json" -H "Account-Id: acc-1" -d '{"domain":"example.com"}'

# 3. статус
curl.exe -X POST http://localhost:8080/site_verification.v1.SiteVerificationService/GetDomainStatus -H "Content-Type: application/json" -H "Account-Id: acc-1" -d '{"domain":"example.com"}'

# 4. список подтверждённых
curl.exe -X POST http://localhost:8080/site_verification.v1.SiteVerificationService/ListVerifiedDomains -H "Content-Type: application/json" -H "Account-Id: acc-1" -d '{}'

# 5. удаление
curl.exe -X POST http://localhost:8080/site_verification.v1.SiteVerificationService/DeleteDomainVerification -H "Content-Type: application/json" -H "Account-Id: acc-1" -d '{"domain":"example.com"}'

# 6. без заголовка Account-Id -> unauthenticated
curl.exe -X POST http://localhost:8080/site_verification.v1.SiteVerificationService/GetDomainStatus -H "Content-Type: application/json" -d '{"domain":"example.com"}'

## Docker

```
make docker-build
docker run --rm -p 8080:8080 --env MONGO_URI=mongodb://host.docker.internal:27017 dns-check-service
```
