package postgreSQL

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"os"
	"strconv"
)

func CreateRecord(id int, address string, price float32) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("postgres://egormelnikov:54236305@localhost:5432/egormelnikov"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	_, err = conn.Exec(context.Background(), "INSERT INTO advert (id,address,price) VALUES ($1,$2,$3)", id, address, price)
	if err != nil {
		log.Fatalln(err)
	}
}
func ReadAllRecords() string {
	conn, err := pgx.Connect(context.Background(), os.Getenv("postgres://egormelnikov:54236305@localhost:5432/egormelnikov"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	rows, err := conn.Query(context.Background(), "SELECT * FROM advert")
	var str string
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatalln(err)
		}
		str += "Advert number: " + strconv.Itoa(int(values[0].(int32))) + ". Is the address advert : " + values[1].(string) + ". Price: " + strconv.Itoa(int(values[2].(float32))) + "\n"
	}
	return str
}
func ReadRecordByNum(id int) string {
	conn, err := pgx.Connect(context.Background(), os.Getenv("postgres://egormelnikov:54236305@localhost:5432/egormelnikov"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	var address string
	var price float32
	err = conn.QueryRow(context.Background(), "SELECT address, price from advert where id=$1", id).Scan(&address, &price)
	if err != nil {
		log.Fatalln(err)
	}
	return "Advert ID: " + strconv.Itoa(id) + ". Is the adress advert: " + address + ". Price: " + strconv.Itoa(int(price)) + "\n"
}
func UpdateRecord(id int, address string, price float32) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("postgres://egormelnikov:54236305@localhost:5432/egormelnikov"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	_, err = conn.Exec(context.Background(), "UPDATE advert SET address =$1,price =$2 WHERE id = $3", address, price, id)
	if err != nil {
		log.Fatalln(err)
	}
}
func DeleteRecord(id int) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("postgres://egormelnikov:54236305@localhost:5432/egormelnikov"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	row, err := conn.Exec(context.Background(), "DELETE FROM advert where id=$1", id)
	if err != nil {
		log.Fatalln(err)
	}
	if row.RowsAffected() != 1 {
		log.Fatalln("nothing to delete")
	}
}
