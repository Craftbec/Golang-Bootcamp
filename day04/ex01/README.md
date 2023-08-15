## Создание сертификатов
-   Устанавливаем центр управления сертификатами `go install github.com/jsha/minica@latest`
-   minica --domains localhost
## Запуск сервера
-   Из папки server `go build` для сборки исполняемого файла `./candy-server`
-   Непосредственно запуск `./candy-server --tls-certificate ../cert/server-cert/cert.pem --tls-key ../cert/server-cert/key.pem --tls-port 3333`

## Запуск клиента
-   Из папки client `go build` для сборки исполняемого файла `./candy-client`
-   Непосредственно запуск `./candy-client -k AA -c 2 -m 50`
-   Ожидаемый ответ: `Thank you! Your change is 20`
