package db

type LedgerEntry struct {
	ID     uint    `db:"id"`
	UserID uint    `db:"user_id"`
	Name   string  `db:"name"`
	Amount float32 `db:"amount"`
	Date   string  `db:"date"`
}

func LedgerAll() (*[]LedgerEntry, error) {
	entries := make([]LedgerEntry, 0)

	err := db.Select(&entries, "SELECT * FROM ledger")
	return &entries, err
}

func LedgerForDate(date string) (*[]LedgerEntry, error) {
	entries := make([]LedgerEntry, 0)

	err := db.Select(&entries, "SELECT * FROM ledger WHERE `date` = ?", date)
	return &entries, err
}

func LedgerEntryCreate(e LedgerEntry) (int64, error) {
	result, err := db.Exec("INSERT INTO ledger (user_id, name, amount, date) VALUES(?,?,?,?)", e.UserID, e.Name, e.Amount, e.Date)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func LedgerEntryUpdate(id int64, e LedgerEntry) error {
	_, err := db.Exec("UPDATE ledger SET user_id = ?, name = ?, amount = ?, date = ? WHERE id = ?", e.UserID, e.Name, e.Amount, e.Date, id)
	return err
}
