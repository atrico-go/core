package syncEx

import (
	"context"
	"time"
)

// Context helper
// Allows inline context creation assuming parent will handle cancel
// USAGE
// ctx,cancel := NewParentContext()
// defer cancel()
// AsyncFunc(ctx.WithTimeout(time.Second))
//
type ParentContext struct {
	context.Context
}

func NewParentContext() (ctx ParentContext, cancel context.CancelFunc) {
	return MakeParentContext(context.Background())
}

func MakeParentContext(base context.Context) (ctx ParentContext, cancel context.CancelFunc) {
	parent := ParentContext{}
	parent.Context, cancel = context.WithCancel(base)
	return parent, cancel
}

func (c *ParentContext) WithTimeout(timeout time.Duration) (ctx context.Context) {
	ctx,_ = context.WithTimeout(c, timeout)
	return ctx
}
