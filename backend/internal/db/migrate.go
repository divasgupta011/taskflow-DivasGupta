package db

import (
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(databaseURL string) {
	var m *migrate.Migrate
	var err error

	// retry logic
	for i := 0; i < 5; i++ {
		// m, err = migrate.New("file://migrations", databaseURL)
		m, err = migrate.New("file:///app/migrations", databaseURL)
		if err == nil {
			break
		}
		log.Println("waiting for migration setup...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("migration init error:", err)
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal("migration failed:", err)
	}

	log.Println("migrations applied successfully")
}
