package awesomeapi

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/silastgoes/client-server-api/models"
)

type USDBRL struct {
	Bid *models.Bid `json:"USDBRL"`
}

func GetBidUSDBRL() (*models.Bid, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2000*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	bid := &USDBRL{}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(bid)
	if err != nil {
		return nil, err
	}
	return bid.Bid, nil
}
