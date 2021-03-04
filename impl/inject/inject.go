package inject

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/kuritka/golic/utils/log"

	"github.com/denormal/go-gitignore"
)

type Inject struct {
	opts   Options
	ctx    context.Context
	ignore gitignore.GitIgnore
}

var logger = log.Log

func New(ctx context.Context, options Options) *Inject {
	return &Inject{
		ctx:  ctx,
		opts: options,
	}
}

func (i *Inject) Run() (err error) {
	logger.Info().Msgf("reading %s", i.opts.License)
	i.ignore, err = gitignore.NewFromFile(i.opts.License)
	if err != nil {
		return err
	}
	i.traverse()
	return nil
}

func (i *Inject) String() string {
	return "inject"
}

func (i *Inject) traverse() {
	p := func(path string, i gitignore.GitIgnore) {
		if i.Ignore(path) {
			fmt.Println(" - " + path)
		} else {
			fmt.Println(" + " + path)
		}
	}

	items, _ := ioutil.ReadDir(".")
	for _, item := range items {
		if item.IsDir() {
			subitems, _ := ioutil.ReadDir(item.Name())
			for _, subitem := range subitems {
				if !subitem.IsDir() {
					// handle file there
					p(item.Name()+"/"+subitem.Name(), i.ignore)
				}
			}
		} else {
			// handle file there
			p(item.Name(), i.ignore)
		}
	}
}
