package migrate

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

func Outfits(db *sql.DB) {
	fmt.Println("[Outfits] Fetching Outfits...")

	rows, err := db.Query("SELECT citizenid, outfitname, model, skin FROM player_outfits")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	fmt.Println("[Outfits] Migrating Outfits...")

	for rows.Next() {
		var citizenid string
		var outfitname string
		var model string
		var skin json.RawMessage

		err = rows.Scan(&citizenid, &outfitname, &model, &skin)
		if err != nil {
			fmt.Println(err)
		}

		var activeSkin QBCoreSkin
		err = json.Unmarshal(skin, &activeSkin)
		if err != nil {
			fmt.Println(err)
		}

		hbM, ffM, ecM, hoM, pcM, ppM := Migrate(activeSkin)

		_, err = db.Exec("INSERT INTO `nf-outfits` (citizenid, outfit_name, model, head_blend, face_feature, eye_color, head_overlay, ped_component, ped_prop, tattoos) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", citizenid, outfitname, model, hbM, ffM, ecM, hoM, pcM, ppM, "{}")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("[Outfits] Migrated Outfit: " + outfitname)
	}

	fmt.Println("[Outfits] Outfits Migrated!")
}
