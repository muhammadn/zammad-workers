package config

import  "github.com/jinzhu/configor"

var Config = struct {
        DB   struct {
                Name     string `default:"zammad_development"`
                Adapter  string `default:"postgres"`
                User     string
                Password string
        }
}{}

func init() {
        if err := configor.Load(&Config, "config/database.yml"); err != nil {
                panic(err)
        }
}
