package inject

import (
	"context"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net/http"
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
	cfg *Config
}

var logger = log.Log

func New(ctx context.Context, options Options) *Inject {
	return &Inject{
		ctx:  ctx,
		opts: options,
	}
}

func (i *Inject) Run() (err error) {
	logger.Info().Msgf("reading %s", i.opts.LicIgnore)
	i.ignore, err = gitignore.NewFromFile(i.opts.LicIgnore)
	if err != nil {
		return err
	}
	logger.Info().Msgf("reading %s", i.opts.ConfigURL)
	if i.cfg, err = i.readConfig(); err != nil {
		return
	}
	i.traverse()
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
	p := func(path string, i gitignore.GitIgnore, o Options, config *Config) (err error) {
		if !i.Ignore(path) {
			fmt.Printf(" + " + path)
			var skip bool
			if err,skip  = inject(path,o, config); skip {
				fmt.Printf(" -> skip")
			}
			fmt.Println()
		}
		return
	}

	err := filepath.Walk("./",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				return p(path, i.ignore, i.opts, i.cfg)
			}
			return nil
		})
	if err != nil {
		logger.Err(err).Msg("")
	}
}

func inject(path string, o Options, config *Config) (err error, skip bool) {
	source,err := read(path)
	if err != nil {
		return err,false
	}
	rule := getRule(path)
	license,err := getCommentedLicense(config, o, rule)
	if err != nil {
		return err, false
	}
	if config.LicenseStartsAfterHeader(rule) {
		// gets first line of source and the rest of source code
		l1, lx := splitSource(source)
		// file is not empty and file contains header
		if  headerContains(l1, config.Golic.Rules[rule].Under) {
			license = fmt.Sprintf("%s\n%s", l1, license)
			if strings.HasPrefix(source, license) {
				return nil, true
			}
			source = lx
		}
	}
	if strings.HasPrefix(source, license) {
		return nil, true
	}
	if !o.Dry {
		data := []byte(fmt.Sprintf("%s%s", license, source))
		err = ioutil.WriteFile(path,data, os.ModeExclusive)
	}
	return
}

func headerContains(header string, values []string) bool{
	for _,v := range values {
		if strings.Contains(header, v) {
			return true
		}
	}
	return false
}

func getCommentedLicense(config *Config, o Options, rule string) (string, error) {
	var ok bool
	var template string
	if template, ok = config.Golic.Licenses[o.Template]; !ok {
		return "",fmt.Errorf("no license found for %s, check configuration (%s)",o.Template,o.ConfigURL)
	}
	if _, ok = config.Golic.Rules[rule]; !ok {
		return "",fmt.Errorf("no rule found for %s, check configuration (%s)", rule,o.ConfigURL)
	}
	template = strings.ReplaceAll(template,"{{copyright}}", o.Copyright)
	if config.IsWrapped(rule) {
		return fmt.Sprintf("%s\n%s%s\n",
			config.Golic.Rules[rule].Prefix,
			template,
			config.Golic.Rules[rule].Suffix),
			nil
	}
	// `\r\n` -> `\r\n #`, `\n` -> `\n #`
	content := strings.ReplaceAll(template,"\n",fmt.Sprintf("\n%s ", config.Golic.Rules[rule].Prefix))
	content = strings.TrimSuffix(content, config.Golic.Rules[rule].Prefix+" ")
	return config.Golic.Rules[rule].Prefix + " " + content,nil
}

func splitSource(source string) (firstLine, rest string){
	lines := strings.Split(source,"\n")
	if len(lines) > 0 {
		firstLine = lines[0]
		rest = strings.Join(lines[1:],"\n")
		return
	}
	return "",source
}

func getRule(path string) (rule string) {
	rule = filepath.Ext(path)
	if rule == "" {
		rule = filepath.Base(path)
	}
	return
}

func (i *Inject) readConfig() (c *Config, err error) {
	var client http.Client
	var resp *http.Response
	var b []byte
	c = new(Config)
	resp, err = client.Get(i.opts.ConfigURL)
	if err != nil {
		return
	}
	if resp.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("%s: %s returns %d", http.MethodGet, i.opts.ConfigURL, resp.StatusCode)
	}
	defer resp.Body.Close()
	if b, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	err = yaml.Unmarshal(b, c)
	return
}