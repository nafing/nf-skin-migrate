package main

import (
	"fmt"
	"nf-skin-migrate/src/database"
	"nf-skin-migrate/src/migrate"
)

func main() {
	fmt.Println("[Main] Starting Migration...")

	var db = database.CreateConnection()
	migrate.Skins(db)
	migrate.Outfits(db)

	fmt.Println("[Main] Migration Complete!")

	defer db.Close()
}
