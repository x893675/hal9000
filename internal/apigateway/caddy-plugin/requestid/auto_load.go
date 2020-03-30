package requestid

import (
	"fmt"
	"github.com/caddyserver/caddy"
	"github.com/caddyserver/caddy/caddyhttp/httpserver"
)

func Setup(c *caddy.Controller) error{
	rule, err := parse(c)

	if err != nil {
		return err
	}

	c.OnStartup(func() error {
		fmt.Println("request id middleware is start...")
		return nil
	})

	c.OnShutdown(func() error {
		fmt.Println("request id middleware shutdown...")
		return nil
	})

	httpserver.GetConfig(c).AddMiddleware(func(next httpserver.Handler) httpserver.Handler {
		return &Rule{Next: next, key: rule.key}
	})

	return nil
}


func parse(c *caddy.Controller) (*Rule, error) {
	rule := &Rule{}
	if c.Next() {
		args := c.RemainingArgs()
		switch len(args) {
		case 0:
			for c.NextBlock() {
				switch c.Val() {
				case "key":
					if !c.NextArg() {
						return nil, c.ArgErr()
					}

					rule.key = c.Val()

					if c.NextArg() {
						return nil, c.ArgErr()
					}
				}
			}
		default:
			return nil, c.ArgErr()
		}
	}

	if c.Next() {
		return nil, c.ArgErr()
	}

	return rule, nil
}