package main

import (
	_ "github.com/juvoinc/exposure-service/controller"

	_ "github.com/lib/pq"
)



func main() {
	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	//	"password=%s dbname=%s sslmode=disable",
	//	host, port, user, password, dbname)
	//db, err := sql.Open("postgres", psqlInfo)
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()
	//
	//sqlStatement := `
	//	SELECT COUNT(*) FROM users;
	//	`
	//rows, err := db.Query(sqlStatement)
	//if err != nil {
	//	// handle this error better than this
	//	panic(err)
	//}
	//defer rows.Close()
	//for rows.Next() {
	//	var count int
	//	err = rows.Scan(&count)
	//	if err != nil {
	//		// handle this error
	//		panic(err)
	//	}
	//	fmt.Println(count)
	//}
	//// get any error encountered during iteration
	//err = rows.Err()
	//if err != nil {
	//	panic(err)
	//}
}
