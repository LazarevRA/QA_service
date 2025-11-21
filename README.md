Описание проекта:

Проект представляет собой микро сервис вопросов и ответов, где можно публиковать вопросы,
добавлять на них ответы, и удалять вопросы. 

Предварительные требования:

-Docker, Docker compose, Make (опционально)

Используемые технологии

    Backend: Go 1.25

    База данных: PostgreSQL 15

    Миграции: Goose

    Контейнеризация: Docker + Docker Compose

    Автоматизация: Make

Запуск приложения:

    1) скопируйте репозиторий
        git clone https://github.com/LazarevRA/hitalent_test.git

    2) Для быстрого запуска используется команда make service-start

Make команды:

    Основные команды:

        make service-start  //Полный запуск, включает в себя сборку, поднятие БД, Миграции, запуск приложения
        make down           // Остановка контейнеров 

    Команды для docker compose:
        
        make db-up   //Запуск базы данных   
        make app-up  //Запуск сервиса
        make build   //Пересборка Docker образов

    Команды для миграций и управление БД:

        make migrate-up     //Применить все миграции
        make migrate-down   //Откатить последнюю миграцию
        make migrate-status //Посмотреть статус миграций

Доступ:

    Приложение работает на порту 8080: http://localhost:8080
    База данных PostgreSQL: localhost:5432
        
        Для доступа к БД:

            Пользователь: postgres
            Пароль: password
            База данных: qa_service
    
    Конфиг находится в /internal/config

При отсутствии make можно запустить сервис следующим способом:
   
    1) docker-compose up -d db
    2) goose -dir migrations postgres "host=localhost user=postgres password=password dbname=qa_service port=5432 sslmode=disable" up
    3) docker compose up -d app

    Остановка всех сервисов:
        docker compose down
    
    Откат миграций 
        docker compose run --rm migrations goose -dir /migrations postgres "host=db user=postgres password=password dbname=qa_service port=5432 sslmode=disable" down