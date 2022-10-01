package helpers

func StatusMapper(payload map[string]interface{}) map[string]interface{} {
	res := map[string]interface{}{
		"wind_status":  "Aman",
		"water_status": "Aman",
	}

	if _, ok := payload["water"]; ok {
		if payload["water"].(float64) > 8 {
			res["water_status"] = "Bahaya"
		} else if payload["water"].(float64) >= 6 && payload["water"].(float64) <= 8 {
			res["water_status"] = "Siaga"
		}
	}

	if _, ok := payload["wind"]; ok {
		if payload["wind"].(float64) > 15 {
			res["wind_status"] = "Bahaya"
		} else if payload["wind"].(float64) >= 7 && payload["wind"].(float64) <= 15 {
			res["wind_status"] = "Siaga"
		}
	}

	return res
}
