package db

import (
	"errors"
)

type ShoplistEntry struct {
	ID     uint   `db:"id"`
	UserID uint   `db:"user_id"`
	ShopID uint   `db:"shop_id"`
	Name   string `db:"name"`
	Qty    uint   `db:"qty"`
	Date   string `db:"date"`

	user *User
}

func ShoplistForDate(date string) (*[]ShoplistEntry, error) {
	entries := make([]ShoplistEntry, 0)
	err := db.Select(&entries, "select * from shoplist where date = ?", date)

	return &entries, err
}

func ShoplistEntryFind(id int64) (*ShoplistEntry, error) {
	var entry ShoplistEntry
	err := db.Get(&entry, "select * from shoplist where id = ?", id)

	return &entry, err
}

func ShoplistEntryCreate(e ShoplistEntry) (int64, error) {
	result, err := db.Exec("INSERT INTO shoplist (user_id, shop_id, `name`, qty, `date`) VALUES(?,?,?,?,?)",
		e.UserID, e.ShopID, e.Name, e.Qty, e.Date)
	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()
	return id, nil
}

func ShoplistEntryDelete(id int64) error {
	result, err := db.Exec("DELETE FROM shoplist WHERE id = ?", id)
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

func ShoplistEntryUpdate(id int64, e ShoplistEntry) error {
	_, err := db.Exec("UPDATE shoplist SET user_id = ?, shop_id = ?, name = ?, qty = ?, date = ? WHERE id = ?",
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
