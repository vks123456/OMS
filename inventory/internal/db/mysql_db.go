package db

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
)

type MysqlDB struct {
	MysqlDB *sql.DB
}

func (d *MysqlDB) InitializeMySQL() {
	dBConnection, err := sql.Open("mysql", "root:root@(localhost:3306)/order_management")
	if err != nil {
		log.Fatal().Err(err).Msg("Connection Failed!!")
	}
	err = dBConnection.Ping()
	if err != nil {
		log.Fatal().Err(err).Msg("Ping Failed!!")
	}
	dBConnection.SetMaxOpenConns(10)
	dBConnection.SetMaxIdleConns(5)
	dBConnection.SetConnMaxLifetime(time.Second * 10)
	d.MysqlDB = dBConnection
}

func (d *MysqlDB) GetProductQuantity(id string) (int, error) {
	sqlQuery := "SELECT quantity FROM item_info WHERE id = ?"
	stmt, err := d.MysqlDB.Prepare(sqlQuery)
	if err != nil {
		return -1, err
	}
	res, err := stmt.Query(id)
	if err != nil {
		return -1, err
	}

	for res.Next() {
		q := 0
		err := res.Scan(&q)
		if err != nil {
			return -1, err
		}
		return q, nil
	}
	return -1, errors.New("NO DATA FOUND")
}

func (d *MysqlDB) GetAndBlockProductQuantity(id string, blockQuantity int) (int, error) {

	// Get Quantity
	sqlQuery := "SELECT quantity FROM item_info WHERE id = ?"
	stmt, err := d.MysqlDB.Prepare(sqlQuery)
	if err != nil {
		return -1, err
	}
	res, err := stmt.Query(id)
	if err != nil {
		return -1, err
	}
	var q int
	for res.Next() {
		err := res.Scan(&q)
		if err != nil {
			return -1, err
		}
		break
	}

	//Block Quantity
	if q < blockQuantity {
		return -1, errors.New("ITEM IS OUT OF STOCK")
	}
	sqlQuery = "UPDATE item_info SET quantity = ? WHERE id = ?"
	stmt, err = d.MysqlDB.Prepare(sqlQuery)
	if err != nil {
		return -1, err
	}
	res, err = stmt.Query(q-blockQuantity, id)
	if err != nil {
		return -1, err
	}

	return q, nil
}

func (d *MysqlDB) UpdateProductQuantity(id string, addQuantity int) error {

	// Get Quantity
	sqlQuery := "SELECT quantity FROM item_info WHERE id = ?"
	stmt, err := d.MysqlDB.Prepare(sqlQuery)
	if err != nil {
		return err
	}
	res, err := stmt.Query(id)
	if err != nil {
		return err
	}
	var q int
	for res.Next() {
		err := res.Scan(&q)
		if err != nil {
			return err
		}
		break
	}

	//Update Quantity
	sqlQuery = "UPDATE item_info SET quantity = ? WHERE id = ?"
	stmt, err = d.MysqlDB.Prepare(sqlQuery)
	if err != nil {
		return err
	}
	res, err = stmt.Query(q+addQuantity, id)
	if err != nil {
		return err
	}

	return nil
}

func (d *MysqlDB) AddNegativeProductQuantity(id string, addQuantity int) error {

	// Get Quantity
	sqlQuery := "SELECT quantity FROM item_info WHERE id = ?"
	stmt, err := d.MysqlDB.Prepare(sqlQuery)
	if err != nil {
		return err
	}
	res, err := stmt.Query(id)
	if err != nil {
		return err
	}
	var q int
	for res.Next() {
		err := res.Scan(&q)
		if err != nil {
			return err
		}
		break
	}

	//  Inventory is not negative
	if q >= 0 {
		return nil
	}

	//Update Quantity
	sqlQuery = "UPDATE item_info SET quantity = ? WHERE id = ?"
	stmt, err = d.MysqlDB.Prepare(sqlQuery)
	if err != nil {
		return err
	}
	res, err = stmt.Query(q+addQuantity, id)
	if err != nil {
		return err
	}

	return errors.New("INVENTORY IN NEGATIVE")
}
