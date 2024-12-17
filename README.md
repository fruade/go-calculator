# go-calculator

Проект предоставляет сервис для вычисления арифметических выражений.

## Структура проекта
```
go-calculator/
│
├── cmd/
│   ├── calc_service/
│   │   └── main.go                 // Точка входа приложения
│   ├── tests/
│   │   └── handlers_test.go        // Тесты для обработчиков
│
├── internal/
│   ├── constants/
│   │   └── constants.go            // Константы для ошибок
│   ├── handlers/
│   │   └── handlers.go             // Обработчики
│   ├── middleware/
│   │   └── validation.go           // Middleware для валидации
│   ├── service/
│   │   └── calculator.go           // Основная логика
│
├── go.mod                          // Модуль Go
├── README.md                       // Информация о проекте
└── .gitignore                      // Игнорируемые файлы и директории

```

## Как использовать

### Запуск проекта

1. Склонируйте репозиторий:
   ```bash
   git clone https://github.com/fruade/go-calculator.git
   cd calc_service

2. Запустите сервер
    ```bash
    go run ./cmd/calc_service/main.go

3. Отправьте запросы на `http://localhost:8080/api/v1/calculate`

### Запуск тестов
```bash
go test ./cmd/tests
```

---

## Примеры запросов
Linux/macOS:
```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "2+2*2"
}'
```
Windows:
```bash
curl -Uri 'http://localhost:8080/api/v1/calculate' `
     -Method POST `
     -ContentType 'application/json' `
     -Body '{"expression": "2+2*2"}'
```
PowerShell:
```bash
Invoke-RestMethod `
-Uri http://localhost:8080/api/v1/calculate `
-Method Post `
-ContentType "application/json" `
-Body '{"expression":"2+2*2"}'

```

# Примеры ответа:

## Успех 200
```bash
{
  "result": "6.00"
}
```

### Ошибка 400
```bash
{
  "error": "Invalid request format"
}
or
{
  "error": "Failed to read request body"
}
```

### Ошибка 405
```bash
{
  "error": "Method not allowed"
}
```

### Ошибка 422
```bash
{
  "error": "Expression is not valid"
}
```

### Ошибка 500
```bash
{
  "error": "Internal server error"
}
```
