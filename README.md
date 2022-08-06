# REST API для Управления Списком Книг

### Стэк
- go 1.17
- postgres 
- docker

### Запуск
```go build -o app cmd/main.go && ./app```

Для postgres можно использовать Docker

```docker run -d --name postgres -e POSTGRES_PASSWORD=qwerty123 -p 5432:5432 postgres```