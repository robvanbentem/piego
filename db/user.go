package db

type User struct {
	ID      int     `db:"id"`
	Name    string  `db:"name"`
	Enabled bool    `db:"enabled"`
	Balance float32 `db:"balance"`
}

func UsersAll() []User {
	users := []User{}
	db.Select(&users, "select u.*, if(sum(amount) IS NULL, 0.00, sum(amount)) as balance from users u left join ledger l on u.id = l.user_id group by u.id, l.user_id")

	return users
}

func UsersFind(id int64) (User, error) {
	var user User

	err := db.Get(&user, "select u.*, sum(amount) as balance from users u join ledger l on u.id = l.user_id where u.id = ? group by l.user_id", id)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (user *User) FetchBalance() {
	rows := db.QueryRow("select sum(amount) as balance from ledger where user_id = ?", user.ID)
	rows.Scan(&user.Balance)
}
