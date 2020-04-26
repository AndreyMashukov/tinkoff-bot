package main

import (
    "flag"
    "log"
    "net/http"
    "fmt"
    "bytes"
    "io/ioutil"
    "os"

    "github.com/joho/godotenv"
    "github.com/gorilla/websocket"
)

// init is invoked before main()
func init() {
    // loads values from .env into the system
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}

func main() {
    var botApi, e1 = os.LookupEnv("BOT_API")

    if (!e1) {
        log.Println("BOT_API variable is not set.")

	    os.Exit(1);
    }

    var addr, e2 = os.LookupEnv("STOCK_API")

    if (!e2) {
        log.Println("STOCK_API variable is not set.")

	    os.Exit(1);
    }

    var token, e3 = os.LookupEnv("STOCK_TOKEN")

    if (!e3) {
        log.Println("STOCK_TOKEN variable is not set.")

        os.Exit(1);
    }

    var figi, e4 = os.LookupEnv("BOT_FIGI")

    if (!e4) {
        log.Println("BOT_FIGI variable is not set.")

        os.Exit(1);
    }

    flag.Parse()

    c, _, err := websocket.DefaultDialer.Dial(addr, http.Header{"Authorization": {"Bearer " + token}})
    if err != nil {
        log.Fatal("dial:", err)
    }
    defer c.Close()

    go func() {
        for {
            _, message, err := c.ReadMessage()
            if err != nil {
                log.Println("read: ", err)

                os.Exit(1);
            }

            log.Printf("Candle: %s\n", message)

            result := fmt.Sprintf(`{"message": %s}`, message)

            var jsonStr = []byte(result)
            req, err := http.NewRequest("POST", botApi, bytes.NewBuffer(jsonStr))

            client := &http.Client{}
            resp, err := client.Do(req)

            if err != nil {
                panic(err)
            }

            defer resp.Body.Close()

            fmt.Println("response Status:", resp.Status)
            body, _ := ioutil.ReadAll(resp.Body)
            fmt.Println("response Body:", string(body))
        }
    }()

    sub := fmt.Sprintf(`{"event": "candle:subscribe", "figi": "%s", "interval": "5min"}`, figi)
    err = c.WriteMessage(websocket.TextMessage, []byte(sub))

    if err != nil {
        log.Println("write: ", err)

        os.Exit(1);
    }

    select {}
}
