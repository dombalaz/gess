package lichess

import (
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	li http.Client
}

// TODO return custom struct
func (cl *Client) DailyPuzzle() (string, error) {
	resp, err := cl.li.Get("https://lichess.org/api/puzzle/daily")
	if err != nil {
		return "", err
	}
	c, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println("status code:", resp.StatusCode)
	fmt.Println("resp size:", resp.ContentLength)
	return string(c), nil
}
