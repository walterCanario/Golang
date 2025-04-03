package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// ConnectDB establece la conexión con la base de datos según el tipo indicado (mysql o postgres)
func ConnectDB(dbType string) (*sql.DB, error) {
	var dsn string

	switch dbType {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
			os.Getenv("MYSQL_USER"),
			os.Getenv("MYSQL_PASSWORD"),
			os.Getenv("MYSQL_HOST"),
			os.Getenv("MYSQL_DB"),
		)
	case "postgres":
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("PG_HOST"),
			os.Getenv("PG_USER"),
			os.Getenv("PG_PASSWORD"),
			os.Getenv("PG_DB"),
		)
	default:
		log.Fatal("Base de datos no soportada")
	}

	db, err := sql.Open(dbType, dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
