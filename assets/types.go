package assets

type Request struct {
	DisplayFieldName string `json:"displayFieldName"`
	FieldAliases     struct {
		OidPermit                 string `json:"oid_permit"`
		RequestNum                string `json:"request_num"`
		PermissionDate            string `json:"permission_date"`
		PermissionNum             string `json:"permission_num"`
		ExpiryDate                string `json:"expiry_date"`
		OpenRequest               string `json:"open_request"`
		URLHadmaya                string `json:"url_hadmaya"`
		BuildingNum               string `json:"building_num"`
		SwTama38                  string `json:"sw_tama_38"`
		SugBakasha                string `json:"sug_bakasha"`
		TochenBakasha             string `json:"tochen_bakasha"`
		Finished                  string `json:"finished"`
		Occupation                string `json:"occupation"`
		TrHathalatBniya           string `json:"tr_hathalat_bniya"`
		Protected                 string `json:"protected"`
		BuildingStage             string `json:"building_stage"`
		Koteret                   string `json:"koteret"`
		Progress                  string `json:"progress"`
		MaslulRishuy              string `json:"maslul_rishuy"`
		DateImport                string `json:"date_import"`
		YechidotDiyur             string `json:"yechidot_diyur"`
		ShapeLength               string `json:"shape_Length"`
		ShapeArea                 string `json:"shape_Area"`
		MsKlali                   string `json:"ms_klali"`
		HearaTeudatGmar           string `json:"heara_teudat_gmar"`
		KSivugMakor               string `json:"k_sivug_makor"`
		SivugMakor                string `json:"sivug_makor"`
		RequestStage              string `json:"request_stage"`
		MsTikBinyan               string `json:"ms_tik_binyan"`
		Addresses                 string `json:"addresses"`
		HakalaTosefetAchuzShetach string `json:"hakala_tosefet_achuz_shetach"`
		HakalaYdHagdalaAchuz      string `json:"hakala_yd_hagdala_achuz"`
		HakalaYdMevukash          string `json:"hakala_yd_mevukash"`
		HakalaYdMutar             string `json:"hakala_yd_mutar"`
		HakalaMelel               string `json:"hakala_melel"`
		HakalaNimuk               string `json:"hakala_nimuk"`
		SwTama38Chadash           string `json:"sw_tama_38_chadash"`
		SwTama38Tosefet           string `json:"sw_tama_38_tosefet"`
		TikTipul1                 string `json:"tik_tipul_1"`
		TikTipul2                 string `json:"tik_tipul_2"`
		TikTipul3                 string `json:"tik_tipul_3"`
		TikTipul4                 string `json:"tik_tipul_4"`
		TikTipul5                 string `json:"tik_tipul_5"`
	} `json:"fieldAliases"`
	GeometryType     string `json:"geometryType"`
	SpatialReference struct {
		Wkid       int `json:"wkid"`
		LatestWkid int `json:"latestWkid"`
	} `json:"spatialReference"`
	Fields []struct {
		Name   string `json:"name"`
		Type   string `json:"type"`
		Alias  string `json:"alias"`
		Length int    `json:"length,omitempty"`
	} `json:"fields"`
	Features []struct {
		Attributes struct {
			OidPermit                 int     `json:"oid_permit"`
			RequestNum                int     `json:"request_num"`
			PermissionDate            int64   `json:"permission_date"`
			PermissionNum             int     `json:"permission_num"`
			ExpiryDate                int64   `json:"expiry_date"`
			OpenRequest               int64   `json:"open_request"`
			URLHadmaya                string  `json:"url_hadmaya"`
			BuildingNum               int     `json:"building_num"`
			SwTama38                  string  `json:"sw_tama_38"`
			SugBakasha                string  `json:"sug_bakasha"`
			TochenBakasha             string  `json:"tochen_bakasha"`
			Finished                  string  `json:"finished"`
			Occupation                string  `json:"occupation"`
			TrHathalatBniya           int64   `json:"tr_hathalat_bniya"`
			Protected                 int     `json:"protected"`
			BuildingStage             string  `json:"building_stage"`
			Koteret                   string  `json:"koteret"`
			Progress                  int     `json:"progress"`
			MaslulRishuy              string  `json:"maslul_rishuy"`
			DateImport                string  `json:"date_import"`
			YechidotDiyur             int     `json:"yechidot_diyur"`
			ShapeLength               float64 `json:"shape_Length"`
			ShapeArea                 float64 `json:"shape_Area"`
			MsKlali                   any     `json:"ms_klali"`
			HearaTeudatGmar           any     `json:"heara_teudat_gmar"`
			KSivugMakor               int     `json:"k_sivug_makor"`
			SivugMakor                string  `json:"sivug_makor"`
			RequestStage              string  `json:"request_stage"`
			MsTikBinyan               int     `json:"ms_tik_binyan"`
			Addresses                 string  `json:"addresses"`
			HakalaTosefetAchuzShetach string  `json:"hakala_tosefet_achuz_shetach"`
			HakalaYdHagdalaAchuz      any     `json:"hakala_yd_hagdala_achuz"`
			HakalaYdMevukash          string  `json:"hakala_yd_mevukash"`
			HakalaYdMutar             string  `json:"hakala_yd_mutar"`
			HakalaMelel               string  `json:"hakala_melel"`
			HakalaNimuk               string  `json:"hakala_nimuk"`
			SwTama38Chadash           string  `json:"sw_tama_38_chadash"`
			SwTama38Tosefet           string  `json:"sw_tama_38_tosefet"`
			TikTipul1                 string  `json:"tik_tipul_1"`
			TikTipul2                 string  `json:"tik_tipul_2"`
			TikTipul3                 string  `json:"tik_tipul_3"`
			TikTipul4                 string  `json:"tik_tipul_4"`
			TikTipul5                 string  `json:"tik_tipul_5"`
		} `json:"attributes"`
		Geometry struct {
			Rings [][][]float64 `json:"rings"`
		} `json:"geometry"`
	} `json:"features"`
}

// This method return a list of addresses for all rquests
func (req *Request) GetAddresses() []string {

	var addresses []string
	for _, request := range req.Features {
		addresses = append(addresses, request.Attributes.Addresses)
		addresses = append(addresses, "\n")
	}
	return addresses
}

/*
func (req *Request) GetCentroid() string {

	rings := req.Features[0].Geometry.Rings

	coords := []geom.Coord{}
	for x,_ := range rings{
		for y, _ := range rings {
			coords = append(coords,rings[x][y])
		}
	}

	polygon := geom.NewPolygon(geom.XY)
	polygon.Centr
}
*/
