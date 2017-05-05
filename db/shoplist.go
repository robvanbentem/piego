package db

type ShoplistEntry struct {
	ID     int    `db:"id"`
	UserID string `db:"user_id"`
	Name   string `db:"name"`
	Qty    int    `db:"qty"`
	Date   string `db:"date"`

	user *User
}

type Shoplist struct {
	Entries []ShoplistEntry
	Date    string
}

func GetShoplist(date string) Shoplist {
	var entries []ShoplistEntry
	db.Select(&entries, "select * from shoplist where date = ?", date)

	return Shoplist{entries, date}
}

func FindShoplistEntry(id int64) ShoplistEntry {
	var entry ShoplistEntry
	db.Get(&entry, "SELECT * FROM shoplist WHERE id = ?", id)

	return entry
}

func InsertShoplistEntry(user int, name string, qty int, date string) (int64, error) {
	result, err := db.Exec("INSERT INTO shoplist (user_id, name, qty, date) VALUES(?,?,?,?)", user, name, qty, date)

	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()
	return id, nil
}

func (entry ShoplistEntry) User() *User {
	if entry.user == nil {
		var user User
		db.Get(&user, "select * from users where id = ?", entry.UserID)

		entry.user = &user
	}

	return entry.user
}
