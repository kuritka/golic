package inject

import (
	"context"
	"fmt"
	"io/ioutil"
)

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
	items, _ := ioutil.ReadDir(".")
	for _, item := range items {
		if item.IsDir() {
			subitems, _ := ioutil.ReadDir(item.Name())
			for _, subitem := range subitems {
				if !subitem.IsDir() {
					// handle file there
					fmt.Println(item.Name() + "/" + subitem.Name())
				}
			}
		} else {
			// handle file there
			fmt.Println(item.Name())
		}
	}
}
