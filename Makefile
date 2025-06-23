PROJECT_NAME = $(notdir $(CURDIR))

# Определение операционной системы
ifeq ($(OS),Windows_NT)
	#Windows
	RM = del /F /Q
	BINARY_NAME = $(PROJECT_NAME).exe
	MKDIR = mkdir
	RMDIR = rmdir /S /Q
else
	#Linux/Mac
	RM = rm -f
	BINARY_NAME = $(PROJECT_NAME)
	MKDIR = mkdir -p
	RMDIR = rm -rf
endif

.PHONY: build-raw run test clean-raw deps fmt lint docs coverage bin build clean migrate clean-logs logs

migrate:
	@echo "Migrating..."
	@go run ./migrations/auto.go

#Команда для сборки приложения
build-raw:
	@echo "Building..."
	@go build -o $(BINARY_NAME) ./cmd/main.go

#Команда для запуска приложения
run:
	@echo "Running..."
	@go run ./cmd/main.go

#Команда для запуска тестов
test:
	@echo "Running tests..."
	@go test -v ./...

# Команда для очистки скомпилированных файлов
clean-raw:
	@echo "Cleaning..."
	@go clean
	@$(RM) $(BINARY_NAME)
	@$(RM) coverage.out

#Команда для установки зависимостей
deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy

# Команда для проверки форматирования кода
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Команда для проверки линтером
lint:
	@echo "Running linter..."
	@go vet ./...

# Команда для генерации документации
docs:
	@echo "Generating documentation..."
	@go doc -all ./...

# Команда для проверки покрытия тестами
coverage:
	@echo "Running tests with coverage..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out

# Команда для сборки в директорию bin
build:
	@echo "Building to bin directory..."
	@go build -o bin/$(BINARY_NAME) ./cmd/main.go

# Команда для очистки директории bin
clean:
	@echo "Cleaning bin directory..."
	@$(RMDIR) bin

clean-logs:
	@echo "Cleaning logs directory..."
	@$(RMDIR) logs
	@$(MKDIR) logs

logs:
	@echo "Showing logs..."
	@cat logs/app.log