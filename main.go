package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/mgutz/ansi"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"regexp"
)

var CONFIG_FILE string

type config struct {
	Crayons []struct {
		Expr   string `yaml:"expr"`
		Crayon string `yaml:"crayon"`
	} `yaml:"crayons"`
}

type colorRule struct {
	expr   *regexp.Regexp
	crayon func(string) string
}

func (c colorRule) Transform(line string) string {
	return c.expr.ReplaceAllStringFunc(line, c.crayon)
}

func newColorRule(expr *regexp.Regexp, color string) colorRule {
	return colorRule{expr, ansi.ColorFunc(color)}
}

func readRules(cfgFileName string) ([]colorRule, error) {
	if file, err := os.Open(cfgFileName); err != nil {
		return nil, err
	} else if data, err := ioutil.ReadAll(file); err != nil {
		return nil, err
	} else {
		cfg := config{}
		if err := yaml.Unmarshal(data, &cfg); err != nil {
			return nil, err
		}

		rules := make([]colorRule, len(cfg.Crayons))
		for idx, crayon := range cfg.Crayons {
			rules[idx] = newColorRule(regexp.MustCompile(crayon.Expr), crayon.Crayon)
		}
		return rules, nil

	}
}

func init() {
	config_dir := "./"
	current_user, err := user.Current()
	if err == nil {
		config_dir = current_user.HomeDir
	}

	config_file := path.Join(config_dir, ".config.yml")

	flag.StringVar(&CONFIG_FILE, "c", config_file, "config file location")
}

func main() {
	flag.Parse()
	rdr := os.Stdin
	scanner := bufio.NewScanner(rdr)

	if rules, err := readRules(CONFIG_FILE); err != nil {
		fmt.Println("Error while reading rules:", err)
		fmt.Println("Aborting.")
	} else {

		for scanner.Scan() {
			line := scanner.Text()
			for _, rule := range rules {
				line = rule.Transform(line)
			}
			fmt.Println(line)
		}
	}

}
