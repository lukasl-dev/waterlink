package waterlink

import (
	"errors"
	"strings"
)

var (
	ErrShardsNegative = errors.New("total shards cannot be negative")
)

type Option func(client *Client) error

func formatHost(host, normal, secured string) string {
	var builder strings.Builder

	if !strings.HasPrefix(host, normal+"://") && !strings.HasPrefix(host, secured+"://") {
		builder.WriteString(normal)
		builder.WriteString("://")
	}

	if strings.HasPrefix(host, ":") {
		builder.WriteString("127.0.0.1")
	}

	builder.WriteString(host)

	return builder.String()
}

// HTTP defines the server's http url address.
func HTTP(host string) Option {
	return func(client *Client) error {
		client.httpHost = formatHost(host, "http", "https")
		return nil
	}
}

// WS defines the server's websocket url address.
func WS(host string) Option {
	return func(client *Client) error {
		client.wsHost = formatHost(host, "ws", "wss")
		return nil
	}
}

// Password defines the server's password.
func Password(password string) Option {
	return func(client *Client) error {
		client.password = password
		return nil
	}
}

// TotalShards defines the total number of shards your bot is operating on.
func TotalShards(totalShards int) Option {
	return func(client *Client) error {
		if totalShards < 0 {
			return ErrShardsNegative
		}
		client.totalShards = totalShards
		return nil
	}
}

// UserID defines the bot user's id.
func UserID(userID string) Option {
	return func(client *Client) error {
		client.userID = userID
		return nil
	}
}
