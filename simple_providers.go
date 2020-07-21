package katyusha

import (
	"context"
	"io/ioutil"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type PlainTextProvider struct {
	BaseProvider `mapstructure:",squash"`
	File         string
}

func NewPlainTextProvider(name string) HostProvider {
	return &PlainTextProvider{BaseProvider: BaseProvider{Name: name}}
}

func (p *PlainTextProvider) Equals(o HostProvider) bool {
	if c, ok := o.(*Cache); ok {
		o = c.Source
	}
	op, ok := o.(*PlainTextProvider)
	return ok &&
		p.BaseProvider.Equals(&op.BaseProvider) &&
		p.File == op.File
}

func (p *PlainTextProvider) ParseViper(v *viper.Viper) error {
	return v.Unmarshal(p)
}

func (p *PlainTextProvider) Load(ctx context.Context, mc chan CacheMessage) (Hosts, error) {
	hosts := make(Hosts, 0)
	data, err := ioutil.ReadFile(p.File)
	if err != nil {
		logrus.Errorf("Could not load %s data in %s: %s", p.Name, p.File, err)
		return hosts, err
	}
	for _, line := range strings.Split(string(data), "\n") {
		line := strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		host := NewHost(line, HostAttributes{})
		hosts = append(hosts, host)
	}
	return hosts, nil
}
