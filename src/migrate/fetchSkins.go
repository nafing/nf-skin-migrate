package migrate

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type QBCoreSkin struct {
	Ear struct {
		DefaultTexture int `json:"defaultTexture"`
		DefaultItem    int `json:"defaultItem"`
		Item           int `json:"item"`
		Texture        int `json:"texture"`
	} `json:"ear"`
	Torso2 struct {
		DefaultTexture int `json:"defaultTexture"`
		DefaultItem    int `json:"defaultItem"`
		Item           int `json:"item"`
		Texture        int `json:"texture"`
	} `json:"torso2"`
	Pants struct {
		DefaultTexture int `json:"defaultTexture"`
		DefaultItem    int `json:"defaultItem"`
		Item           int `json:"item"`
		Texture        int `json:"texture"`
	} `json:"pants"`
	Ageing struct {
		DefaultTexture float64 `json:"defaultTexture"`
		DefaultItem    float64 `json:"defaultItem"`
		Item           float64 `json:"item"`
		Texture        float64 `json:"texture"`
	} `json:"ageing"`
	Bag struct {
		DefaultTexture int `json:"defaultTexture"`
		DefaultItem    int `json:"defaultItem"`
		Item           int `json:"item"`
		Texture        int `json:"texture"`
	} `json:"bag"`
	Nose1 struct {
		DefaultTexture float64 `json:"defaultTexture"`
		DefaultItem    float64 `json:"defaultItem"`
		Item           float64 `json:"item"`
		Texture        float64 `json:"texture"`
	} `json:"nose_1"`
	Face2 struct {
		DefaultTexture int `json:"defaultTexture"`
		DefaultItem    int `json:"defaultItem"`
		Item           int `json:"item"`
		Texture        int `json:"texture"`
	} `json:"face2"`
	ChimpHole struct {
		DefaultTexture float64 `json:"defaultTexture"`
		DefaultItem    float64 `json:"defaultItem"`
		Item           float64 `json:"item"`
		Texture        float64 `json:"texture"`
	} `json:"chimp_hole"`
	Nose3 struct {
		DefaultTexture float64 `json:"defaultTexture"`
		DefaultItem    float64 `json:"defaultItem"`
		Item           float64 `json:"item"`
		Texture        float64 `json:"texture"`
	} `json:"nose_3"`
	Hair struct {
		DefaultTexture float64 `json:"defaultTexture"`
		DefaultItem    float64 `json:"defaultItem"`
		Item           float64 `json:"item"`
		Texture        float64 `json:"texture"`
	} `json:"hair"`
	Nose2 struct {
		DefaultTexture float64 `json:"defaultTexture"`
		DefaultItem    float64 `json:"defaultItem"`
		Item           float64 `json:"item"`
		Texture        float64 `json:"texture"`
	} `json:"nose_2"`
	Cheek1 struct {
		DefaultTexture float64 `json:"defaultTexture"`
		DefaultItem    float64 `json:"defaultItem"`
		Item           float64 `json:"item"`
		Texture        float64 `json:"texture"`
	} `json:"cheek_1"`
	Nose5 struct {
		DefaultTexture float64 `json:"defaultTexture"`
		DefaultItem    float64 `json:"defaultItem"`
		Item           float64 `json:"item"`
		Texture        float64 `json:"texture"`
	} `json:"nose_5"`
	Mask struct {
		DefaultTexture int `json:"defaultTexture"`
		DefaultItem    int `json:"defaultItem"`
		Item           int `json:"item"`
		Texture        int `json:"texture"`
	} `json:"mask"`
	Shoes struct {
		DefaultTexture int `json:"defaultTexture"`
		DefaultItem    int `json:"defaultItem"`
		Item           int `json:"item"`
		Texture        int `json:"texture"`
	} `json:"shoes"`
	Bracelet struct {
		DefaultTexture int `json:"defaultTexture"`
		DefaultItem    int `json:"defaultItem"`
		Item           int `json:"item"`
		Texture        int `json:"texture"`
	} `json:"bracelet"`
	JawBoneWidth struct {
		DefaultTexture float64 `json:"defaultTexture"`
		DefaultItem    float64 `json:"defaultItem"`
		Item           float64 `json:"item"`
		Texture        float64 `json:"texture"`
	} `json:"jaw_bone_width"`
	EyebrownForward struct {
		DefaultTexture float64 `json:"defaultTexture"`
		DefaultItem    float64 `json:"defaultItem"`
		Item           float64 `json:"item"`
		Texture        float64 `json:"texture"`
	} `json:"eyebrown_forward"`
	Facemix struct {
		SkinMix         float64 `json:"skinMix"`
		DefaultShapeMix float64 `json:"defaultShapeMix"`
		DefaultSkinMix  float64 `json:"defaultSkinMix"`
		ShapeMix        float64 `json:"shapeMix"`
	} `json:"facemix"`
	Arms struct {
		DefaultTexture int `json:"defaultTexture"`
		DefaultItem    int `json:"defaultItem"`
		Item           int `json:"item"`
		Texture        int `json:"texture"`
	} `json:"arms"`
	NeckThikness struct {
		DefaultTexture float64 `json:"defaultTexture"`
		DefaultItem    float64 `json:"defaultItem"`
		Item           float64 `json:"item"`
		Texture        float64 `json:"texture"`
	} `json:"neck_thikness"`
	Blush struct {
		DefaultTexture float64 `json:"defaultTexture"`
		DefaultItem    float64 `json:"defaultItem"`
		Item           float64 `json:"item"`
		Texture        float64 `json:"texture"`
	} `json:"blush"`
	EyeColor struct {
		DefaultTexture int `json:"defaultTexture"`
		DefaultItem    int `json:"defaultItem"`
		Item           int `json:"item"`
		Texture        int `json:"texture"`
	} `json:"eye_color"`
	Glass struct {
		DefaultTexture int `json:"defaultTexture"`
		DefaultItem    int `json:"defaultItem"`
		Item           int `json:"item"`
		Texture        int `json:"texture"`
	} `json:"glass"`
	Face struct {
		DefaultTexture int `json:"defaultTexture"`
		DefaultItem    int `json:"defaultItem"`
		Item           int `json:"item"`
		Texture        int `json:"texture"`
	} `json:"face"`
	Vest struct {
		DefaultTexture int `json:"defaultTexture"`
		DefaultItem    int `json:"defaultItem"`
		Item           int `json:"item"`
		Texture        int `json:"texture"`
	} `json:"vest"`
	Accessory struct {
		DefaultTexture int `json:"defaultTexture"`
		DefaultItem    int `json:"defaultItem"`
		Item           int `json:"item"`
		Texture        int `json:"texture"`
	} `json:"accessory"`
	Hat struct {
		DefaultTexture int `json:"defaultTexture"`
		DefaultItem    int `json:"defaultItem"`
		Item           int `json:"item"`
		Texture        int `json:"texture"`
	} `json:"hat"`
	LipsThickness struct {
		DefaultTexture float64 `json:"defaultTexture"`
		DefaultItem    float64 `json:"defaultItem"`
		Item           float64 `json:"item"`
		Texture        float64 `json:"texture"`
	} `json:"lips_thickness"`
	EyeOpening struct {
		DefaultTexture float64 `json:"defaultTexture"`
		DefaultItem    float64 `json:"defaultItem"`
		Item           float64 `json:"item"`
		Texture        float64 `json:"texture"`
	} `json:"eye_opening"`
	ChimpBoneWidth struct {
		DefaultTexture float64 `json:"defaultTexture"`
		DefaultItem    float64 `json:"defaultItem"`
		Item           float64 `json:"item"`
		Texture        float64 `json:"texture"`
	} `json:"chimp_bone_width"`
	JawBoneBackLenght struct {
		DefaultTexture float64 `json:"defaultTexture"`
		DefaultItem    float64 `json:"defaultItem"`
		Item           float64 `json:"item"`
		Texture        float64 `json:"texture"`
	} `json:"jaw_bone_back_lenght"`
	Cheek2 struct {
		DefaultTexture float64 `json:"defaultTexture"`
		DefaultItem    float64 `json:"defaultItem"`
		Item           float64 `json:"item"`
		Texture        float64 `json:"texture"`
	} `json:"cheek_2"`
	Moles struct {
		DefaultTexture float64 `json:"defaultTexture"`
		DefaultItem    float64 `json:"defaultItem"`
		Item           float64 `json:"item"`
		Texture        float64 `json:"texture"`
	} `json:"moles"`
	Watch struct {
		DefaultTexture int `json:"defaultTexture"`
		DefaultItem    int `json:"defaultItem"`
		Item           int `json:"item"`
		Texture        int `json:"texture"`
	} `json:"watch"`
	ChimpBoneLowering struct {
		DefaultTexture float64 `json:"defaultTexture"`
		DefaultItem    float64 `json:"defaultItem"`
		Item           float64 `json:"item"`
		Texture        float64 `json:"texture"`
	} `json:"chimp_bone_lowering"`
	TShirt struct {
		DefaultTexture int `json:"defaultTexture"`
		DefaultItem    int `json:"defaultItem"`
		Item           int `json:"item"`
		Texture        int `json:"texture"`
	} `json:"t-shirt"`
	Eyebrows struct {
		DefaultTexture float64 `json:"defaultTexture"`
		DefaultItem    float64 `json:"defaultItem"`
		Item           float64 `json:"item"`
		Texture        float64 `json:"texture"`
	} `json:"eyebrows"`
	Cheek3 struct {
		DefaultTexture float64 `json:"defaultTexture"`
		DefaultItem    float64 `json:"defaultItem"`
		Item           float64 `json:"item"`
		Texture        float64 `json:"texture"`
	} `json:"cheek_3"`
	ChimpBoneLenght struct {
		DefaultTexture float64 `json:"defaultTexture"`
		DefaultItem    float64 `json:"defaultItem"`
		Item           float64 `json:"item"`
		Texture        float64 `json:"texture"`
	} `json:"chimp_bone_lenght"`
	Nose4 struct {
		DefaultTexture float64 `json:"defaultTexture"`
		DefaultItem    float64 `json:"defaultItem"`
		Item           float64 `json:"item"`
		Texture        float64 `json:"texture"`
	} `json:"nose_4"`
	Lipstick struct {
		DefaultTexture float64 `json:"defaultTexture"`
		DefaultItem    float64 `json:"defaultItem"`
		Item           float64 `json:"item"`
		Texture        float64 `json:"texture"`
	} `json:"lipstick"`
	Decals struct {
		DefaultTexture int `json:"defaultTexture"`
		DefaultItem    int `json:"defaultItem"`
		Item           int `json:"item"`
		Texture        int `json:"texture"`
	} `json:"decals"`
	EyebrownHigh struct {
		DefaultTexture float64 `json:"defaultTexture"`
		DefaultItem    float64 `json:"defaultItem"`
		Item           float64 `json:"item"`
		Texture        float64 `json:"texture"`
	} `json:"eyebrown_high"`
	Makeup struct {
		DefaultTexture float64 `json:"defaultTexture"`
		DefaultItem    float64 `json:"defaultItem"`
		Item           float64 `json:"item"`
		Texture        float64 `json:"texture"`
	} `json:"makeup"`
	Nose0 struct {
		DefaultTexture float64 `json:"defaultTexture"`
		DefaultItem    float64 `json:"defaultItem"`
		Item           float64 `json:"item"`
		Texture        float64 `json:"texture"`
	} `json:"nose_0"`
	Beard struct {
		DefaultTexture float64 `json:"defaultTexture"`
		DefaultItem    float64 `json:"defaultItem"`
		Item           float64 `json:"item"`
		Texture        float64 `json:"texture"`
	} `json:"beard"`
}

type HeadBlend struct {
	ShapeFirstID  int     `json:"shapeFirstID"`
	ShapeSecondID int     `json:"shapeSecondID"`
	ShapeThirdID  int     `json:"shapeThirdID"`
	SkinFirstID   int     `json:"skinFirstID"`
	SkinSecondID  int     `json:"skinSecondID"`
	SkinThirdID   int     `json:"skinThirdID"`
	ShapeMix      float64 `json:"shapeMix"`
	SkinMix       float64 `json:"skinMix"`
	ThirdMix      float64 `json:"thirdMix"`
}
type FaceFeature struct {
	Eyebrow            []float64 `json:"eyebrow"`
	CheekWidthHeight   []float64 `json:"cheek_width_height"`
	ChinLengthHeight   []float64 `json:"chin_length_height"`
	ChinShapeHole      []float64 `json:"chin_shape_hole"`
	NoseTwistCurveness []float64 `json:"nose_twist_curveness"`
	LipThickness       float64   `json:"lip_thickness"`
	NeckThickness      float64   `json:"neck_thickness"`
	NoseLengthTip      []float64 `json:"nose_length_tip"`
	NoseWidthPeak      []float64 `json:"nose_width_peak"`
	CheekSize          float64   `json:"cheek_size"`
	JawWidthShape      []float64 `json:"jaw_width_shape"`
	EyeOpening         float64   `json:"eye_opening"`
}
type EyeColor int
type HeadOverlay struct {
	Blemishes      []float64 `json:"blemishes"`
	FacialHair     []float64 `json:"facial_hair"`
	Eyebrows       []float64 `json:"eyebrows"`
	Ageing         []float64 `json:"ageing"`
	Makeup         []float64 `json:"makeup"`
	Blush          []float64 `json:"blush"`
	Complexion     []float64 `json:"complexion"`
	SunDamage      []float64 `json:"sun_damage"`
	Lipstick       []float64 `json:"lipstick"`
	MolesFreckles  []float64 `json:"moles_freckles"`
	ChestHair      []float64 `json:"chest_hair"`
	BodyBlemishes  []float64 `json:"body_blemishes"`
	AddBodyBlemish []float64 `json:"add_body_blemishes"`
}
type PedComponent struct {
	Mask          []int     `json:"mask"`
	Hair          []float64 `json:"hair"`
	Arms          []int     `json:"arms"`
	Legs          []int     `json:"legs"`
	ParachuteBags []int     `json:"parachute_bags"`
	Shoes         []int     `json:"shoes"`
	Accessory     []int     `json:"accessory"`
	Undershirt    []int     `json:"undershirt"`
	Kevlar        []int     `json:"kevlar"`
	Badge         []int     `json:"badge"`
	Torso         []int     `json:"torso"`
}
type PedProp struct {
	Hat      []int `json:"hat"`
	Glasses  []int `json:"glasses"`
	Ears     []int `json:"ears"`
	Watch    []int `json:"watch"`
	Bracelet []int `json:"bracelet"`
}

func Skins(db *sql.DB) {
	fmt.Println("[Skins] Fetching Skins...")

	rows, err := db.Query("SELECT citizenid, model, skin FROM playerskins")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	fmt.Println("[Skins] Migrating Skins...")

	for rows.Next() {
		var citizenid string
		var model int
		var skin json.RawMessage

		err := rows.Scan(&citizenid, &model, &skin)
		if err != nil {
			fmt.Println(err)
		}

		var activeSkin QBCoreSkin
		err = json.Unmarshal(skin, &activeSkin)
		if err != nil {
			fmt.Println(err)
		}

		hbM, ffM, ecM, hoM, pcM, ppM := Migrate(activeSkin)

		_, err = db.Exec("INSERT INTO `nf-skins` (citizenid, model, head_blend, face_feature, eye_color, head_overlay, ped_component, ped_prop, tattoos) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", citizenid, model, hbM, ffM, ecM, hoM, pcM, ppM, "{}")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("[Skins] Migrated Skin for CitizenID: " + citizenid)
	}

	fmt.Println("[Skins] Migration Complete!")
}
