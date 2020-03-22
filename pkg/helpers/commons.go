package helpers

import (
	"os"
)

const (
	MONGOCONNECTIONSTRING = "MongoConnectionStr"
)

func ReadEnv(envType string) string {
	switch envType {
	case MONGOCONNECTIONSTRING:
		return os.Getenv("MONGO_CONN")
	default:
		return ""
	}
}

func SetEnv() {
	os.Setenv("MONGO_CONN", "localhost:27017")
}
