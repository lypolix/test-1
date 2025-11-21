package singleton

import (
    "sync"
    "time"
)

type Database struct {
    connection string
}

var (
    instance *Database
    once     sync.Once
)

func GetInstance() *Database {
    once.Do(func() {
        time.Sleep(2 * time.Second)
        instance = &Database{connection: "postgres://localhost:5432/mydb"}
        println("Database instance created")
    })
    return instance
}

func (db *Database) Query(sql string) {
    println("Executing query:", sql, "on connection:", db.connection)
}