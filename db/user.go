package db

type User struct {
	ID      int     `db:"id"`
	Name    string  `db:"name"`
	Enabled bool    `db:"enabled"`
	Balance float32 `db:"balance"`
}

func UsersAll() *[]User {
	users := make([]User, 0)
	db.Select(&users, "SELECT u.*, if(SUM(amount) IS NULL, 0.00, SUM(amount)) AS balance FROM users u LEFT JOIN ledger l ON u.id = l.user_id GROUP BY u.id, l.user_id")

	return &users
}

func UsersFind(id int64) (User, error) {
	var user User

	err := db.Get(&user, "SELECT u.*, IF(SUM(amount) IS NULL, 0.00, SUM(amount)) AS balance FROM users u LEFT JOIN ledger l ON u.id = l.user_id WHERE u.id = ? GROUP BY l.user_id", id)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (user *User) FetchBalance() {
	rows := db.QueryRow("SELECT sum(amount) AS balance FROM ledger WHERE user_id = ?", user.ID)
	rows.Scan(&user.Balance)
}
