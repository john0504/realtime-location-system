package service

type Notifier interface {
	Broadcast(playerID string, lat, lng float64, targets []string)
	NotifyLandmark(playerID string, name string)
}
