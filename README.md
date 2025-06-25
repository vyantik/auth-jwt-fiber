# JWT Authentication Fiber

Современный микросервис для аутентификации и авторизации с использованием JWT токенов, построенный на Go с использованием Fiber framework.

## 🚀 Возможности

- **Регистрация пользователей** с валидацией данных
- **Аутентификация** с использованием email и пароля
- **JWT токены** (Access и Refresh токены)
- **Обновление токенов** через refresh token
- **Защищенные маршруты** с middleware аутентификации
- **Валидация данных** с кастомными правилами
- **Логирование** с использованием zerolog
- **PostgreSQL** для хранения данных пользователей
- **Redis** для кэширования и управления сессиями
- **Docker Compose** для простого развертывания

## 🛠 Технологии

- **Go 1.24.3** - основной язык программирования
- **Fiber v2** - веб-фреймворк
- **GORM** - ORM для работы с базой данных
- **PostgreSQL** - основная база данных
- **Redis** - кэш и управление сессиями
- **JWT** - токены аутентификации
- **Zerolog** - структурированное логирование
- **Validator** - валидация данных

## 📋 Требования

- Go 1.24.3 или выше
- Docker и Docker Compose
- PostgreSQL 15.2
- Redis 6.0

## 🚀 Быстрый старт

### 1. Клонирование репозитория

```bash
git clone <repository-url>
cd auth-jwt-service
```

### 2. Настройка переменных окружения

Создайте файл `.env` в корневой директории:

```env
# Application Configuration
APPLICATION_PORT=3000

# Database
POSTGRES_USER=postgres
POSTGRES_PASSWORD=your_password
POSTGRES_DB=auth_service
POSTGRES_PORT=5432
DATABASE_URL=postgres://postgres:your_password@localhost:5432/auth_service?sslmode=disable

# Redis
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=your_redis_password
REDIS_DB=0

# JWT
JWT_ACCESS_SECRET=your_access_secret_key
JWT_REFRESH_SECRET=your_refresh_secret_key

# Logging
LOG_LEVEL=0 # Все логи
LOG_FORMAT=json
```

### 3. Запуск с Docker Compose

```bash
# Запуск базы данных и Redis
docker-compose up -d

# Установка зависимостей
make deps

# Выполнение миграций
make migrate

# Запуск приложения
make run
```

### 4. Альтернативный запуск без Docker

```bash
# Установка зависимостей
make deps

# Сборка приложения
make build

# Запуск
./bin/auth-jwt-service
```

## 📚 API Endpoints

### Публичные маршруты (без аутентификации)

#### Регистрация пользователя

```http
POST /api/register
Content-Type: application/x-www-form-urlencoded

email=user@example.com&username=username&password=password123
```

**Ответ:**

```json
{
  "message": "User registered successfully"
}
```

#### Вход в систему

```http
POST /api/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

**Ответ:**

```json
{
  "message": "Login successful",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

#### Обновление токенов

```http
POST /api/refresh
Content-Type: application/json

{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Ответ:**

```json
{
  "message": "Tokens refreshed successfully",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### Защищенные маршруты (требуют аутентификации)

#### Получение профиля пользователя

```http
GET /api/profile
Authorization: Bearer <access_token>
```

**Ответ:**

```json
{
  "user_id": 1,
  "email": "user@example.com"
}
```

## 🔧 Команды Makefile

```bash
# Установка зависимостей
make deps

# Сборка приложения
make build

# Запуск приложения
make run

# Запуск тестов
make test

# Проверка покрытия тестами
make coverage

# Форматирование кода
make fmt

# Проверка линтером
make lint

# Генерация документации
make docs

# Очистка скомпилированных файлов
make clean

# Выполнение миграций
make migrate

# Очистка логов
make clean-logs

# Отображение логов в консоль
make logs
```

## 🏗 Архитектура проекта

```
auth-jwt-service/
├── cmd/
│   └── main.go              # Точка входа приложения
├── config/
│   ├── env.go               # Загрузка переменных окружения
│   └── structs.go           # Структуры конфигурации
├── internal/
│   ├── auth/                # Модуль аутентификации
│   │   ├── handler.go       # HTTP обработчики
│   │   ├── model.go         # Модели запросов/ответов
│   │   └── service.go       # Бизнес-логика
│   ├── jwt/                 # JWT сервис
│   │   ├── model.go         # JWT модели
│   │   └── service.go       # JWT операции
│   ├── middleware/          # Middleware
│   │   ├── auth.go          # Аутентификация middleware
│   │   └── structs.go       # Структуры middleware
│   └── user/                # Модуль пользователей
│       ├── model.go         # Модель пользователя
│       ├── repository.go    # Репозиторий пользователей
│       └── service.go       # Сервис пользователей
├── migrations/
│   └── auto.go              # Автоматические миграции
├── pkg/
│   ├── db/                  # Подключение к БД
│   │   ├── db.go            # PostgreSQL подключение
│   │   └── redis.go         # Redis подключение
│   ├── logger/              # Логирование
│   │   └── logger.go        # Конфигурация логгера
│   └── validators/          # Валидация
│       ├── password.go      # Валидация паролей
│       └── validator.go     # Основной валидатор
├── docker-compose.yaml      # Docker Compose конфигурация
├── go.mod                   # Go модули
├── go.sum                   # Go зависимости
├── Makefile                 # Команды сборки
└── README.md                # Документация
```

## 🔐 Безопасность

- **Хеширование паролей** с использованием bcrypt
- **JWT токены** с разделением на access и refresh
- **Валидация данных** на всех уровнях
- **Защищенные маршруты** с middleware аутентификации
- **Переменные окружения** для конфиденциальных данных

## 📝 Валидация данных

### Регистрация

- **Email**: обязательное поле, валидный email формат
- **Username**: 3-15 символов, обязательное поле
- **Password**: 8-20 символов, обязательное поле, сильный пароль

### Вход

- **Email**: обязательное поле, валидный email формат
- **Password**: 8-20 символов, обязательное поле

## 🐳 Docker

### Запуск с Docker Compose

```bash
# Запуск всех сервисов
docker-compose up -d

# Просмотр логов
docker-compose logs -f

# Остановка сервисов
docker-compose down
```

## 🧪 Тестирование

```bash
# Запуск всех тестов
make test

# Запуск тестов с покрытием
make coverage

# Запуск конкретного теста
go test ./internal/auth -v
```

## 📊 Мониторинг и логирование

Сервис использует структурированное логирование с zerolog:

- **Форматы**: json, console
- **Контекстная информация**: request ID, user ID, timestamps
