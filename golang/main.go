package main

import (
    // "encoding/json"
    "database/sql"
    // "fmt"
    // "log"
    // "net/http"
    // "context"
    "testing"
    _ "github.com/go-sql-driver/mysql"
    "github.com/stretchr/testify/assert"
)

func TestGetRegisteredDriver(t *testing.T) {
    assert.Equal(t, []string{"mysql"}, sql.Drivers())
}

func main () {
    t := new(testing.T)
    TestGetRegisteredDriver(t)
}
