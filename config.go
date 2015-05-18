package horizon

import (
	"github.com/PuerkitoBio/throttled"
	"github.com/Sirupsen/logrus"
)

// Config is the configuration for go-horizon.  It get's populated by the
// app's main function and is provided to NewApp.
type Config struct {
	DatabaseUrl            string
	StellarCoreDatabaseUrl string
	RubyHorizonUrl         string
	Port                   int
	Autopump               bool
	RateLimit              throttled.Quota
	RedisUrl               string
	LogLevel               logrus.Level
}
