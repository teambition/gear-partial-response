package partial

import (
	"github.com/teambition/gear"
	mask "github.com/teambition/json-mask-go"
)

// Options is the partial response middleware options.
type Options struct {
	// Query specifies the querystring to use. By defaults it is "fields".
	Query string
}

// Sender is to implement gear.Sender interface.
type Sender struct {
	query string
}

// Send is to implement gear.Sender interface.
func (s *Sender) Send(ctx *gear.Context, code int, data interface{}) (err error) {
	if len(ctx.Query(s.query)) == 0 {
		return ctx.JSON(code, data)
	}

	maskedData, err := mask.Mask(data, ctx.Query(s.query))

	if err != nil {
		return ctx.JSON(code, data)
	}

	return ctx.JSON(code, maskedData)
}

// New returns a new partial response middleware for your gear app.
func New(opts Options) *Sender {
	if opts.Query == "" {
		opts.Query = "fields"
	}

	return &Sender{query: opts.Query}
}
