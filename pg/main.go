package pg

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	//_ "github.com/go-sql-driver/mysql"
)

type myD struct {
}

func (a myD) Open(name string) (driver.Conn, error) {
	panic("implement me")
}

func failErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	sql.Register("pgmock", &myD{})

	fmt.Println(sql.Drivers())

	db, err := sql.Open("pgmock", "")
	failErr(err)

	fmt.Printf("%+v", db)
}
