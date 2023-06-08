package database

import (
	"fmt"
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/verma-kartik/marketplace-api/internal/models"
)

func (d *Database) MigrateDb() {
	fmt.Println("migrating the db")

	err := d.gClient.AutoMigrate(&models.Product{})
	if err != nil {
		fmt.Println(err)
	}
	//driver, err := postgres.WithInstance(d.Client.DB, &postgres.Config{})
	//if err != nil {
	//	return fmt.Errorf("could not create the postgres driver: %w", err)
	//}
	//
	//m, err := migrate.NewWithDatabaseInstance(
	//	"file:///migrations",
	//	"postgres",
	//	driver,
	//)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return err
	//}
	//
	//if err := m.Up(); err != nil {
	//	return fmt.Errorf("could not run up migration %w", err)
	//}
	//
	//fmt.Println("successfully migrated the database")

	//return nil
}
