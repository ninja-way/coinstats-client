package coinstats

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Client struct {
	client *http.Client
}

// NewClient create http.Client with custom timeout and logging transport
func NewClient(timeout time.Duration) (*Client, error) {
	if timeout <= 0 {
		return nil, errors.New("invalid timeout")
	}

	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil
}

// GetCoins does api request and return specified coins range
func (c Client) GetCoins(limit int) ([]Coin, error) {
	url := fmt.Sprintf("https://api.coinstats.app/public/v1/coins?limit=%d", limit)

	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var coinsResp coinsResponse
	err = json.Unmarshal(body, &coinsResp)
	if err != nil {
		return nil, err
	}

	return coinsResp.Coins, nil
}
