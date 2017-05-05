package db

type User struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	Enabled int    `db:"enabled"`
}

func AllUsers() []User {
	users := []User{}
	db.Select(&users, "select * from users")

	return users
}

func (user User) Balance() float32 {
	var balance float32

	rows := db.QueryRow("SELECT sum(amount) AS balance FROM ledger WHERE user_id = ?", user.ID)
	rows.Scan(&balance)

	return balance
}
