# go-news-apigateway

go-news-apigateway это агрегатор новостей на микросервисной архитектуре.

## Описание:
go-news-apigateway содержит 5 компонентов:
1. rss-agg
2. api-gateway
3. news
4. comments
5. censor


![api-gateway](https://github.com/al-melnikov/go-news-apigateway/assets/5961306/874ec6dc-ee14-458f-88bc-87a076116872)


## rss-agg
***rss-agg*** занимается скраппингом rss лент, указанных в ***config.json***, и их записью в базу данных

## api-gateway
***api-gateway*** это входная точка для поьзователя. 
Он запускается на localhost на порту 8080. 

Он работает с 4 методами:
### 1. GET /news/{news_id} 
В ответе содержит полную информацию о новости и список всех комменариев к ней.
формат ID UUID.

### 2. GET /news/tree/{news_id} 
В ответе содержит полную информацию о новости и дерево комменариев к ней.

### 3. GET /news
В параметрах запроса можно передать ***reg_exp*** - регулярное выражение для поиска постов.
И page - номер страницы, которую следует вернуть.
На 1 страницу приходится 10 постов.

### 4. POST /comments
В теле запроса необходимо указать ***news_id*** - id новости,
***parent_id*** - id верхнего комментария(можно указать null).
***content*** - текст комментария.

## news
Микросервис news возвращает новости по запросу.
У него 2 метода:

### 1. GET /news/id
Требует ID овости в теле запроса, возвращает полную информацию об этой новости.
### 2. GET /news/reg_exp
Требует текст регулярного выражения в теле запроса и номер желаемой сраницы. 
Возвращает массив постов и информацию о пагинации.

## comments
Микросервис ***comments*** работает с комментариями. Он содержит 3 метода:
### 1. POST /comments 
Требует ***news_id***, ***parent_id***, ***content*** в теле запроса.
Добавляет комментарий в базу данных
### 2. GET /comments
Требует ***news_id*** в теле запроса. Возвращает массив всех комментариев к этой новости.
### 3. GET /comments/tree 
Требует ***news_id*** в теле запроса. Возвращает дерево всех комментариев к этой новости.

## censor
Микросервис ***censor*** проверяет комментарий на удовлетворение правилам 
перед добавлением в базу данных

1 метод:
### PUT /censor
Прнимает ***content*** - текст комментария в теле запроса.
Возвращает ***is_censored*** - true или false.


# Запуск
***rss-agg***,  ***news***, ***comments*** требуют запущенную ***postgres***.
База данных для них всех одна общая.
Строка подключения для них всех выглядит так:
```
"postgres://postgres:password@localhost:5432/apigateway?sslmode=disable"
```
- user: ***postgres***
- пароль: ***password***
- название базы данных: ***apigateway***

Для запуска их надо заменить на актульные для вас. 

Все микросервисы запускаются на localhost на разных портах:
- api-gateway: ***8080***
- comments: ***8081***
- news: ***8082***
- censor: ***8083***


# docker-compose

Все микросервисы можно запустить с помощью docker-compose. 
Поскольку база данных здесь общая, ее все так же необходимо запустить локально.
Строка подключения к базе данных берется из окружения.

```
sudo docker-compose up
```