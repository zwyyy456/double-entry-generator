package config

import (
	"github.com/deb-sig/double-entry-generator/pkg/provider/alipay"
	"github.com/deb-sig/double-entry-generator/pkg/provider/bmo"
	"github.com/deb-sig/double-entry-generator/pkg/provider/citic"
	"github.com/deb-sig/double-entry-generator/pkg/provider/hsbchk"
	"github.com/deb-sig/double-entry-generator/pkg/provider/htsec"
	"github.com/deb-sig/double-entry-generator/pkg/provider/huobi"
	"github.com/deb-sig/double-entry-generator/pkg/provider/icbc"
	"github.com/deb-sig/double-entry-generator/pkg/provider/jd"
	"github.com/deb-sig/double-entry-generator/pkg/provider/mt"
	"github.com/deb-sig/double-entry-generator/pkg/provider/td"
	"github.com/deb-sig/double-entry-generator/pkg/provider/wechat"
)

// Config is the global configuration.
type Config struct {
	Title                    string         `yaml:"title,omitempty"`
	DefaultMinusAccount      string         `yaml:"defaultMinusAccount,omitempty"`
	DefaultPlusAccount       string         `yaml:"defaultPlusAccount,omitempty"`
	DefaultCashAccount       string         `yaml:"defaultCashAccount,omitempty"`
	DefaultPositionAccount   string         `yaml:"defaultPositionAccount,omitempty"`
	DefaultCommissionAccount string         `yaml:"defaultCommissionAccount,omitempty"`
	DefaultPnlAccount        string         `yaml:"defaultPnlAccount,omitempty"`
	DefaultCurrency          string         `yaml:"defaultCurrency,omitempty"`
	Alipay                   *alipay.Config `yaml:"alipay,omitempty"`
	Wechat                   *wechat.Config `yaml:"wechat,omitempty"`
	Huobi                    *huobi.Config  `yaml:"huobi,omitempty"`
	Htsec                    *htsec.Config  `yaml:"htsec,omitempty"`
	Icbc                     *icbc.Config   `yaml:"icbc,omitempty"`
	Td                       *td.Config     `yaml:"td,omitempty"`
	Bmo                      *bmo.Config    `yaml:"bmo,omitempty"`
	JD                       *jd.Config     `yaml:"jd,omitempty"`
	Citic                    *citic.Config  `yaml:"citic,omitempty"`
	HsbcHK                   *hsbchk.Config `yaml:"hsbchk,omitempty"`
	MT                       *mt.Config     `yaml:"mt,omitempty"`
}
