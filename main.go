package main

import (
    "log"
    "github.com/jrallison/go-workers"
)

func main() {
    workers.Configure(map[string]string{
        "process": "worker1",
        "server": "localhost:6379",
    })
    workers.Process("communicate_email", CommunicateEmail, 10)
    workers.Process("default", Default, 10)
    workers.Process("transactions", Transactions, 10)
    workers.Run()
}

func CommunicateEmail(msg *workers.Msg) {
    log.Println("running communicate email task", msg)
}

func Default(msg *workers.Msg) {
    log.Println("running default task", msg)
}

func Transactions(msg *workers.Msg) {
    log.Println("running transactions task", msg)
}
