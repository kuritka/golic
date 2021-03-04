package inject

import "context"

type Inject struct {
	opts Options
	ctx  context.Context
}

func New(ctx context.Context, options Options) *Inject {
	return &Inject{
		ctx:  ctx,
		opts: options,
	}
}

func (i *Inject) Run() error {
	return nil
}

func (i *Inject) String() string {
	return "inject"
}

func (i *Inject) traverse() {

}
