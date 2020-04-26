### Бот для тинькофф инвестиций

Данный пакет позволяет слушать определенную акцию по FIGI (env BOT_FIGI) и отсылать ее изменения на ваш Back-End (env BOT_API), для того, чтобы работать с ботом вам необходимо иметь аккаунт на тинькоф инвестициях, [ЗАРЕГИСТРИРОВАТЬСЯ](https://www.tinkoff.ru/sl/I1I5JfDYTL) 

После того как вы [зарегистрировались](https://www.tinkoff.ru/sl/I1I5JfDYTL) вы можете получить Token (env STOCK_TOKEN) для доступа к Open API, ищите на настройках на сайте [Tinkoff инвестиций](https://www.tinkoff.ru/sl/I1I5JfDYTL) 

Данный репозиторий содержит файл example.json
```json
{
    "message": {
        "event": "candle",
        "payload": {
            "o": 127.15,
            "c": 126.71,
            "h": 127.17,
            "l": 125.88,
            "v": 15912,
            "time": "2020-04-03T15:40:00Z",
            "interval": "5min",
            "figi": "BBG000BCSST7"
        },
        "time": "2020-04-16T19:13:36.650279639Z"
    }
}
```

Это пример body POST запроса который бот будет слать на ваш Back-End (env BOT_API)

Докер образ `amashukov/tinkoff-bot:latest` представляет из себя готовый image для запуска данного бота, используйте файл `docker-compose.yml` как основу для вашего проекта.

Чтобы запустить бота достаточно выполнить команду
```bash
docker-compose up -d
```

Вывод бота вы можете посмотреть с помощью команды 
```bash
docker logs -f bot_writer
```
пример вывода:
```bash
2020/04/23 22:28:01 Candle: {"payload":{"o":137,"c":135.89,"h":137,"l":135.7,"v":8722,"time":"2020-04-23T22:25:00Z","interval":"5min","figi":"BBG000BCSST7"},"event":"candle","time":"2020-04-23T22:28:01.052929006Z"}
response Status: 200 OK
response Body: <your backend response body here>
```

Переменные окружения:

* BOT_API - URL вашего Back-End сервера
* STOCK_API - URL Tinkoff Open API (поддерживается только websocket)
* STOCK_TOKEN - токен для доступа на биржу
* BOT_FIGI - FIGI акции которой собираетесь торговать или слушать изменения
