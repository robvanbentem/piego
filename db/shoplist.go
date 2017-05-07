package db

import (
	"errors"
)

type ShoplistEntry struct {
	ID     int    `db:"id"`
	UserID int    `db:"user_id"`
	ShopID int    `db:"shop_id"`
	Name   string `db:"name"`
	Qty    int    `db:"qty"`
	Date   string `db:"date"`

	user *User
}

func ShoplistForDate(date string) ([]ShoplistEntry, error) {
	entries := make([]ShoplistEntry, 0)
	err := db.Select(&entries, "select * from shoplist where date = ?", date)

	return entries, err
}

func ShoplistEntryFind(id int64) (ShoplistEntry, error) {
	var entry ShoplistEntry
	err := db.Get(&entry, "select * from shoplist where id = ?", id)

	return entry, err
}

func ShoplistEntryCreate(user int, shop int, name string, qty int, date string) (int64, error) {
	result, err := db.Exec("insert into shoplist (user_id, shop_id, `name`, qty, `date`) values(?,?,?,?,?)", user, shop, name, qty, date)

	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()
	return id, nil
}

func ShoplistEntryDelete(id int64) error {
	result, err := db.Exec("delete from shoplist where id = ?", id)

	if err != nil {
		return err
	}

	if aff, err := result.RowsAffected(); err != nil {
		return err
	} else if aff != 1 {
		return errors.New("could not delete or already deleted")
	}

	return nil
}

func ShoplistEntryUpdate(id int64, e *ShoplistEntry) error {
	_, err := db.Exec("update shoplist set user_id = ?, shop_id = ?, name = ?, qty = ?, date = ? where id = ?",
		e.UserID, e.ShopID, e.Name, e.Qty, e.Date, id)

	if err != nil {
		return err
	}

	return nil
}

func (entry ShoplistEntry) User() *User {
	if entry.user == nil {
		var user User
		db.Get(&user, "select * from users where id = ?", entry.UserID)

		entry.user = &user
	}

	return entry.user
}
