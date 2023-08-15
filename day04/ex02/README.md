## Запуск сервера
-   Из папки server `go build` для сборки исполняемого файла `./server`
-   Непосредственно запуск `./server --tls-certificate ../cert/server/cert.pem --tls-key ../cert/server/key.pem --tls-port 3333`

## Тест строка клиента
-   curl -s --key cert/client/key.pem --cert cert/client/cert.pem --cacert cert/minica.pem -XPOST -H curl -s --key ../cert/client/key.pem --cert ../cert/client/cert.pem --cacert ../cert/minica.pem -XPOST -H "Content-Type: application/json" -d '{"candyType": "NT", "candyCount": 2, "money": 34}' "https://localhost:3333/buy_candy"
-   Ожидаемый ответ: {"change":0,"thanks":" ____________\n< Thank you! >\n ------------\n        \\   ^__^\n         \\  (oo)\\_______\n            (__)\\       )\\/\\\n                ||----w |\n                ||     ||\n"}

## Запуск клиента
-   Из папки client `go build` для сборки исполняемого файла `./client`
-   Непосредственно запуск  `./client -k NT -c 2 -m 34`
-   Ожидаемый ответ:
_____________
< Thank you ! >
 -------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
 Your change is 0