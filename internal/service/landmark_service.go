package service

import (
	"log"
	"math"
	"realtime-location/pkg/db"
)

type Landmark struct {
	ID        int
	Name      string
	Latitude  float64
	Longitude float64
	Radius    int
}

// 取得所有地標
func GetLandmarks() ([]Landmark, error) {
	rows, err := db.DB.Query("SELECT id, name, latitude, longitude, radius FROM landmarks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Landmark

	for rows.Next() {
		var l Landmark
		rows.Scan(&l.ID, &l.Name, &l.Latitude, &l.Longitude, &l.Radius)
		list = append(list, l)
	}

	return list, nil
}

func calculateDistance(lat1, lng1, lat2, lng2 float64) float64 {
	const R = 6371000 // 地球半徑（公尺）

	dLat := (lat2 - lat1) * math.Pi / 180
	dLng := (lng2 - lng1) * math.Pi / 180

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*math.Pi/180)*math.Cos(lat2*math.Pi/180)*
			math.Sin(dLng/2)*math.Sin(dLng/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}

func CheckLandmarks(playerID string, lat, lng float64) {

	landmarks, err := GetLandmarks()
	if err != nil {
		log.Println(err)
		return
	}

	for _, l := range landmarks {

		dist := calculateDistance(lat, lng, l.Latitude, l.Longitude)

		if dist <= float64(l.Radius) {
			// 觸發事件
			NotifyPlayerEnterLandmark(playerID, l)
		}
	}
}

func NotifyPlayerEnterLandmark(playerID string, l Landmark) {
	notifier.NotifyLandmark(playerID, l.Name)
}
