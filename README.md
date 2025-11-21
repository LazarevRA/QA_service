1) Запустить docker-compose up -d db
2) проведение миграций goose -dir migrations postgres "host=localhost user=postgres password=password dbname=qa_service port=5432 sslmode=disable" up