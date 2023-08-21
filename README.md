# Note Speller - Test task

# Содержание
1. [Задача](#Задача)
2. [Скачивание приложения](#Installation)
3. [Команды для запуска](#Command)
4. [Структура приложения](#Structure)
5. [Описание HTTP API](#HTTP)

## Задача

(кратко) создать REST сервис добавления и просмотра заметок

## Installation
```bash
# Download this project
git clone https://github.com/Snegniy/notespeller_testtask.git && cd notespeller_testtask

# HTTP API Endpoint : http://127.0.0.1:8000/
```

## Command
```bash
# Запустить приложение в контейнере
make
```

## Structure
```
├── api // Swagger
│   ├── meta.go
│   ├── reqres.go
│   ├── swagger.go
├── cmd
│   ├── main.go          // запуск приложения
├── internal
│   ├── config
│   │   ├── config.go   // инициализация конфигурации приложения 
│   ├── handlers
│   │   ├── handlers.go // HTTP обработчики
│   │   ├── jwt.go // Генерация jwt токенов
│   │   ├── response.go // отправка ответа в формате JSON
│   ├── model
│   │   ├── model.go // модель данных
│   ├── service
│   │   ├── password.go // хэширование и сравнение паролей']
│   │   ├── service.go // бизнес-логика
│   │   ├── speller.go // валидация заметки в yandex.speller
│   ├── storage
│   │   ├── postgres
│   │   │  ├── postgres.go // postgreSQL хранилище
├── migrations
│   ├── init.sql        // начальные настройки БД
├── pkg
│   ├── logger
│   │   ├── logger.go // инициализация логгера
│   ├── server
│   │   ├── server.go  // запуск graceful HTTP сервера
├── .dockerignore
├── .env  // конфигурационные установки по умолчанию
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── Makefile
```

## HTTP
При дефолтных настройках сервер доступен по localhost:8000
#### /register
* `POST` /register 
* `{"username":"(string)", "password":""(string)" }`   - Регистрация пользователя

#### /login
* `POST` /login
* `{"username":"(string)", "password":""(string)" }`   - Логин пользователя

#### /logout
* `GET` /logout - Выход из профиля

#### /
* `POST` /
* `{"note":"(string)" }`   - Добавление заметки

#### /
* `GET` / - Вывод заметок пользователя

