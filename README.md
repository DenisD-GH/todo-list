REST API для управления задачами (TODO-лист).  

Реализовано на Go с использованием fiber и PostgreSQL (pgx)  


Запуск программы.  

1. База данных подключается через переменную окружения  
DATABASE_URL=postgres://username:password@localhost:5432/postgres  
Где:  
-   username — имя пользователя PostgreSQL.  
-   password — пароль.  
-   localhost:5432 — хост и порт PostgreSQL.  
-   postgres — имя базы данных.  

На linux и macOS в терминале:  
export DATABASE_URL="postgres://username:password@localhost:5432/postgres"  

На windows в командной строке:  
set DATABASE_URL=postgres://username:password@localhost:5432/postgres  

2. Установка зависимостей  
    go get github.com/jackc/pgx/v5  
    go get github.com/gofiber/fiber/v2

3. Запуск программы с помощью команды
    go run main.go



Пример использования API через postman

GET метод
![GET](screenshots/get-method.png)

POST метод
![POST](screenshots/post-method.png)
![POST](screenshots/result_of_post-method.png)

PUT метод
![PUT](screenshots/put-method.png)
![PUT](screenshots/result_of_put-method.png)

DELETE метод
![POST](screenshots/delete-method.png)

База данных
![DB](screenshots/db.png)