## Создание сревера из swagger файла
-   Поместим файл `swagger.yml` в папку `src`
-   Создаем сервер `swagger generate server -A server`
-   Подключаем пакеты `go mod tidy`

## Запуск сервера
-   Команда `go build` для сборки исполняемого файла `./server`
-   Непосредственно запуск `./server --port 3333`

## Изменения и дополнения в файлах
-   Функции для обработки входящих данных и получения ответа ( 3 штуки) `server_api.go`
-   Удаление из файла buy_candy.go для поля `change` записи `,omitempty`

## Тест строка клиента
-   curl -XPOST -H "Content-Type: application/json" -d '{"money": 20, "candyType": "AA", "candyCount": 1}' http://127.0.0.1:3333/buy_candy
-   Ожидаемый ответ: {"change":5,"thanks":"Thank you!"}

-   curl -XPOST -H "Content-Type: application/json" -d '{"money": 46, "candyType": "YR", "candyCount": 2}' http://127.0.0.1:3333/buy_candy

-   Ожидаемый ответ: {"change":0,"thanks":"Thank you!"}