# Стек проекта

1. Postgres
2. Go 1.20
3. Gin
4. JWT
5. Swagger
6. Docker
7. Taskfile
8. Air

## Установка зависимостей

1. Taskfile

    ```bash
    go install github.com/go-task/task/v3/cmd/task@latest
    ```

2. Go модули

    - для разработки (с hot reload)

    ```bash
    task dev:install
    ```

    - без dev пакетов

    ```bash
    task
    ```

## Настройка конфигурации  

1.Нужно создать файл `.env` и добавить переменные:

```bash
# DATABASE
PSQL_USER="postgres"
PSQL_PASSWORD="postgres"
PSQL_HOST="localhost"
PSQL_PORT="5432"
PSQL_DATABASE="mydatabse"

# Mail
MAIL_FROM="fromEmail@example.ru"
MAIL_PASSWORD="password"
MAIL_USERNAME="username"
MAIL_HOST="smtp.example.ru"
MAIL_PORT="465"
#MAIL_SSL=bool (optional) #default: false


# (optional for docker)
# PG_ADMIN
PG_ADMIN_EMAIL="admin@admin.ru"
PG_ADMIN_PASSWORD="admin"
```

## Миграции

1. Применить миграции

    ```bash
    task migration -- up
    ```

2. Отменить миграции

    ```bash
    task migration -- down
    ```

## Запуск

1. **Дев режим с (hot reload)**. Если команда не работает, см. ***"Установка зависимостей"*** пункт 2 hot reload

    ```bash
    task dev
    ```

2. **Обычный** (компиляция бинарного файла)

    ```bash
    task run
    ```

**Production with Docker**:

```bash
docker compose up
```

## Прочее

1. Перечень и описание всех скриптов

    ```bash
    task --list-all
    ```

2. Использование другого конфига в папки `configs`

    ***Работает для следующих команд:***
    - task dev
    - task run

    ```bash
    task dev -- -config <path_to_config>
    ```

3. Получить информацию о CLI параметров приложения

    ```bash
    task dev -- -h
    # or
    task run -- -h
    ```

4. Генерация Swagger документации

    ```bash
    task swagger
    ```
