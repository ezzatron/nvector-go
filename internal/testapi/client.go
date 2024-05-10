package testapi

import (
	"context"
	"errors"

	"github.com/gorilla/websocket"
)

// Client is a websocket client for calling the n-vector test API.
type Client struct {
	conn *websocket.Conn
}

// NewClient creates a new websocket client for calling the n-vector test API.
func NewClient() (*Client, error) {
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:17357/", nil)
	if err != nil {
		return nil, err
	}

	return &Client{c}, nil
}

// Close closes the websocket connection.
func (c *Client) Close() error {
	err := c.conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
	)

	return errors.Join(err, c.conn.Close())
}

func call[J, G any](
	_ context.Context,
	client *Client,
	unmarshal func(J) G,
	fn string,
	args map[string]any,
) (result G, _ error) {
	err := client.conn.WriteJSON(map[string]any{
		"id":   1,
		"fn":   fn,
		"args": args,
	})
	if err != nil {
		return result, err
	}

	var res struct {
		Error  string `json:"error"`
		Result J      `json:"result"`
	}
	err = client.conn.ReadJSON(&res)
	if err != nil {
		return result, err
	}
	if res.Error != "" {
		return result, errors.New(res.Error)
	}

	return unmarshal(res.Result), nil
}
