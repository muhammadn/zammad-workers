package main

import (
    "runtime"
    "log"
    "github.com/jrallison/go-workers"
    "./db"
)

func main() {
    defer db.DBCon.Close()
    runtime.GOMAXPROCS(runtime.NumCPU()) // Use all CPU Cores

    workers.Configure(map[string]string{
        "process": "worker1",
        "server": "localhost:6379",
    })
    workers.Process("transactions", Transactions, 10)
    workers.Run()
}

func Default(msg *workers.Msg) {
    log.Println("running default task", msg)
}

func Transactions(msg *workers.Msg) {
    log.Println("running transactions task", msg)
}
