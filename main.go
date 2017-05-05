package main

import (
	"fmt"

	pdb "piego/db"
)

func main() {
	pdb.InitDB()

	users := pdb.AllUsers()
	for _, user := range users {
		fmt.Printf("%s's balance is %.2f\n", user.Name, user.Balance())

		id, err := pdb.InsertShoplistEntry(user.ID, "Perercevelaat, vers", 1, "2017-04-28")

		if err != nil {
			panic("err")
		}

		fmt.Printf("Added entry %d\n", id)
		entry := pdb.FindShoplistEntry(id)

		fmt.Printf("Added entry %dx %s\n", entry.Qty, entry.Name)
	}

	shoplist := pdb.GetShoplist("2017-04-28")

	user_list := make(map[string][]pdb.ShoplistEntry)

	for _, entry := range shoplist.Entries {
		if user_list[entry.UserID] == nil {
			user_list[entry.UserID] = []pdb.ShoplistEntry{}
		}

		user_list[entry.UserID] = append(user_list[entry.UserID], entry)
	}

	for id, entries := range user_list {
		fmt.Printf("Entries for %s/%s\n", id, entries[0].User().Name)

		for _, entry := range entries {
			fmt.Printf("%dx %s\n", entry.Qty, entry.Name)
		}
	}
}
