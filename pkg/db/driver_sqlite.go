package db

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/mattn/go-sqlite3"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const DriverSqlite = "sqlite3"

func init() {
	connectionFactories[DriverSqlite] = NewSqliteDriverFactory()
}

func NewSqliteDriverFactory() DriverFactory {
	return &sqliteDriverFactory{}
}

type sqliteDriverFactory struct{}

func (m *sqliteDriverFactory) GetDSN(settings Settings) string {
	var databaseFilePath string
	if (strings.Contains(settings.Uri.Host, "~")) {
		dirname, _ := os.UserHomeDir()
		databaseFilePath = fmt.Sprintf("%s", strings.ReplaceAll(settings.Uri.Host, "~", dirname))
	} else {
		ex, _ := os.Executable()
		databaseFilePath = filepath.Join(filepath.Dir(ex), settings.Uri.Host)
	}

	dsn := url.URL{
		User: url.UserPassword(settings.Uri.User, settings.Uri.Password),
		Host: fmt.Sprintf("file:%s", databaseFilePath),
		Path: settings.Uri.Database,
	}

	qry := dsn.Query()
	dsn.RawQuery = qry.Encode()

	uri := dsn.String()

	return uri[4:]
}

func (m *sqliteDriverFactory) GetMigrationDriver(db *sql.DB, database string, migrationsTable string) (database.Driver, error) {
	return sqlite3.WithInstance(db, &sqlite3.Config{
		DatabaseName:    database,
		MigrationsTable: migrationsTable,
	})
}
