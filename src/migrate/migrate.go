package migrate

import (
	"encoding/json"
	"fmt"
)

func Migrate(activeSkin QBCoreSkin) ([]byte, []byte, EyeColor, []byte, []byte, []byte) {
	fmt.Println(activeSkin.Face.Item, activeSkin.Face2.Item, activeSkin.Face.Texture, activeSkin.Face2.Texture)

	headBlend := HeadBlend{
		ShapeFirstID:  activeSkin.Face.Item,
		ShapeSecondID: activeSkin.Face2.Item,
		ShapeThirdID:  0,
		SkinFirstID:   activeSkin.Face.Texture,
		SkinSecondID:  activeSkin.Face2.Texture,
		SkinThirdID:   0,
		ShapeMix:      activeSkin.Facemix.ShapeMix,
		SkinMix:       activeSkin.Facemix.SkinMix,
		ThirdMix:      0.5,
	}

	hbM, err := json.Marshal(headBlend)
	if err != nil {
		fmt.Println(err)
	}

	FaceFeature := FaceFeature{
		Eyebrow:            []float64{activeSkin.EyebrownForward.Item, activeSkin.EyebrownHigh.Item},
		CheekWidthHeight:   []float64{activeSkin.Cheek3.Item, activeSkin.Cheek1.Item},
		ChinLengthHeight:   []float64{activeSkin.ChimpBoneLenght.Item, activeSkin.ChimpBoneLowering.Item},
		ChinShapeHole:      []float64{activeSkin.ChimpBoneWidth.Item, activeSkin.ChimpHole.Item},
		NoseTwistCurveness: []float64{activeSkin.Nose5.Item, activeSkin.Nose3.Item},
		LipThickness:       activeSkin.LipsThickness.Item,
		NeckThickness:      activeSkin.NeckThikness.Item,
		NoseLengthTip:      []float64{activeSkin.Nose2.Item, activeSkin.Nose4.Item},
		NoseWidthPeak:      []float64{activeSkin.Nose0.Item, activeSkin.Nose1.Item},
		CheekSize:          activeSkin.Cheek2.Item,
		JawWidthShape:      []float64{activeSkin.JawBoneWidth.Item, activeSkin.JawBoneBackLenght.Item},
		EyeOpening:         activeSkin.EyeOpening.Item,
	}

	ffM, err := json.Marshal(FaceFeature)
	if err != nil {
		fmt.Println(err)
	}

	ecM := EyeColor(activeSkin.EyeColor.Item)

	HeadOverlay := HeadOverlay{
		Blemishes:      []float64{-1, 1.0},
		FacialHair:     []float64{activeSkin.Beard.Item, 1.0, activeSkin.Beard.Texture},
		Eyebrows:       []float64{activeSkin.Eyebrows.Item, 1.0, activeSkin.Eyebrows.Texture},
		Ageing:         []float64{activeSkin.Ageing.Item, 1.0, activeSkin.Ageing.Texture},
		Makeup:         []float64{activeSkin.Makeup.Item, 1.0, activeSkin.Makeup.Texture},
		Blush:          []float64{activeSkin.Blush.Item, 1.0, activeSkin.Blush.Texture},
		Complexion:     []float64{-1, 1.0, 0},
		SunDamage:      []float64{-1, 1.0, 0},
		Lipstick:       []float64{activeSkin.Lipstick.Item, 1.0, activeSkin.Lipstick.Texture},
		MolesFreckles:  []float64{activeSkin.Moles.Item, 1.0, activeSkin.Moles.Texture},
		ChestHair:      []float64{-1, 1.0, 0},
		BodyBlemishes:  []float64{-1, 1.0, 0},
		AddBodyBlemish: []float64{-1, 1.0, 0},
	}

	hoM, err := json.Marshal(HeadOverlay)
	if err != nil {
		fmt.Println(err)
	}

	PedComponent := PedComponent{
		Mask:          []int{activeSkin.Mask.Item, activeSkin.Mask.Texture},
		Hair:          []float64{activeSkin.Hair.Item, activeSkin.Hair.Texture, activeSkin.Hair.Texture},
		Arms:          []int{activeSkin.Arms.Item, activeSkin.Arms.Texture},
		Legs:          []int{activeSkin.Pants.Item, activeSkin.Pants.Texture},
		ParachuteBags: []int{activeSkin.Bag.Item, activeSkin.Bag.Texture},
		Shoes:         []int{activeSkin.Shoes.Item, activeSkin.Shoes.Texture},
		Accessory:     []int{activeSkin.Accessory.Item, activeSkin.Accessory.Texture},
		Undershirt:    []int{activeSkin.TShirt.Item, activeSkin.TShirt.Texture},
		Kevlar:        []int{activeSkin.Vest.Item, activeSkin.Vest.Texture},
		Badge:         []int{activeSkin.Decals.Item, activeSkin.Decals.Texture},
		Torso:         []int{activeSkin.Torso2.Item, activeSkin.Torso2.Texture},
	}

	pcM, err := json.Marshal(PedComponent)
	if err != nil {
		fmt.Println(err)
	}

	PedProp := PedProp{
		Hat:      []int{activeSkin.Hat.Item, activeSkin.Hat.Texture},
		Glasses:  []int{activeSkin.Glass.Item, activeSkin.Glass.Texture},
		Ears:     []int{activeSkin.Ear.Item, activeSkin.Ear.Texture},
		Watch:    []int{activeSkin.Watch.Item, activeSkin.Watch.Texture},
		Bracelet: []int{activeSkin.Bracelet.Item, activeSkin.Bracelet.Texture},
	}

	ppM, err := json.Marshal(PedProp)
	if err != nil {
		fmt.Println(err)
	}

	return hbM, ffM, ecM, hoM, pcM, ppM
}
