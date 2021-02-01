package microcli

import (
	"context"

	"github.com/kimwsj/cli"
	"github.com/kimwsj/go-config/source"
)

type contextKey struct{}

// Context sets the microcli context
func Context(c *cli.Context) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, contextKey{}, c)
	}
}
