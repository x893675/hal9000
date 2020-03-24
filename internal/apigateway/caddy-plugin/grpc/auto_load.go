package grpc

import "github.com/caddyserver/caddy"

func Setup(c *caddy.Controller) error {
	rule, err := parse(c)

	if err != nil {
		return err
	}
}


func parse(c *caddy.Controller) (*Rule, error) {
	return nil, nil
}