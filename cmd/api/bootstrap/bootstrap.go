package bootstrap

import (
	"database/sql"
	"fmt"
	"github.com/juanegido/hexapi/internal/platform/server"
	"github.com/juanegido/hexapi/internal/platform/storage/mysql"
)

const (
	srvHost = "localhost"
	srvPort = 8080

	dbHost = "localhost"
	dbPort = 3306
	dbUser = "myuser"
	dbPass = "mypass"
	dbName = "mydb"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	courseRepository := mysql.NewCourseRepository(db)

	srv := server.New(srvHost, srvPort, courseRepository)

	return srv.Run()
}
