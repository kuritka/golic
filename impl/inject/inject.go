package inject

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

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
	i.opts.template,err =  read(i.opts.Template)
	if err == nil {
		i.traverse()
	}
	return
}

func (i *Inject) String() string {
	return "inject"
}

func read(f string) (s string, err error) {
	content, err := ioutil.ReadFile(f)
	if err != nil {
		return
	}
	// Convert []byte to string and print to screen
	return string(content), nil
}

func (i *Inject) traverse() {
	p := func(path string, i gitignore.GitIgnore, o Options) (err error) {
		if !i.Ignore(path) {
			fmt.Println(" + " + path)
			err = inject(path,o)
		}
		return
	}

	err := filepath.Walk("./",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				return p(path, i.ignore, i.opts)
			}
			return nil
		})
	if err != nil {
		logger.Err(err).Msg("")
	}
}


func inject(path string, o Options) (err error) {
	c,err := read(path)
	if err != nil {
		return err
	}
	l := fmt.Sprintf("/*\n%s\n*/",o.template)
	if strings.HasPrefix(c, l) {
		fmt.Printf(" -> skip")
		return
	}
	if !o.Dry {
		data := []byte(fmt.Sprintf("/*\n%s\n*/\n%s",o.template,c))
		err = ioutil.WriteFile(path,data, os.ModeExclusive)
	}
	return
}


