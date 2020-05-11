package helpers

import (
	"bytes"
	"net/http"
	"os"
	"time"
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

func HttpRequest(methodType string, Url string,body []byte) (*http.Response, error){
	req, err := http.NewRequest(methodType, Url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	var client http.Client
	client.Timeout = 15 * time.Second
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

