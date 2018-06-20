package main

import (
    "runtime"
    "log"
    "time"

    "github.com/jrallison/go-workers"
    "./db"
)

type Setting struct {
    ID int `gorm:"primary_key" json:"id,omitempty"`
    Title string `json:"title, omitempty"`
    Name string `json:"name, omitempty"`
    Area string `json:"area, omitempty"`
    Description string `json:"description, omitempty"`
    Options string `json:"description, omitempty"`
    StateCurrent string `json:"state_current, omitempty"`
    StateInitial string `json:"state_initial, omitempty"`
    Frontend bool `json:"-"`
    Preferences string `json:"preferences, omitempty"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

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
    settings := []Setting{}
    db.DBCon.Where("area = ?", "Transaction::Backend::Async").Order("name").Find(&settings)

    log.Println("running transactions task", msg)
    log.Println("%+v\n", settings)
}
