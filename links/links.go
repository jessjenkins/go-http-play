package links

import (
	"context"
	"fmt"
)

const (
	ctxHost = "host"
)

func URLBuild(ctx context.Context, url string) string {
	host := ctx.Value(ctxHost)
	return fmt.Sprintf("%s/%s", host, url)
}
