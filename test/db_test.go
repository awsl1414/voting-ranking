package test

import (
	"fmt"
	"testing"
	"voting-ranking/pkg/db"
)

func TestDb(t *testing.T) {
	_ = db.SetupDBLink()
	ok := db.Db.Migrator().HasTable("users")
	fmt.Println("ok:", ok)
}
