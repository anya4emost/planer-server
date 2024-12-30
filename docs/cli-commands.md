# Docker

## Подключиться к контейнеру к базе
docker exec -it planer_db psql -U postgres -d planer_db

# Postgres

## Вывести описание таблицы
\d ${table_name}
Пример: \d users