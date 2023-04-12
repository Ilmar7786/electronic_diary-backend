## Установка пакетов

    $ go mod download

## Настройка конфигурации  
Нужно создать файл .env и добавить переменные:

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

## Запуск
**Локально**: `make run`


**Docker**: `make docker-up`
