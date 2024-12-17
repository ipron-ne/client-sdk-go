package config

import (
	"strings"
	"time"
)

type Config struct {
	BaseURLOption
	TimeoutOption
	DebugOption
	AppTokenOption
	TenantIDOption
}

var DefaultConfig = Config{
	BaseURLOption: BaseURLOption{BaseURL: ""},
	TimeoutOption: TimeoutOption{Timeout: 30 * time.Second},
	DebugOption:   DebugOption{IsDebug: false},
	AppTokenOption: AppTokenOption{AppToken: ""},
}

// Base URL Option

type BaseURLOption struct {
	BaseURL string
}

func (o *BaseURLOption) ApplyOption(opt *Config) {
	opt.BaseURL = o.BaseURL
}

func WithBaseURL(baseURL string) *BaseURLOption {
	baseURL = strings.TrimRight(baseURL, "/")
	return &BaseURLOption{BaseURL: baseURL}
}

// Timeout Option

type TimeoutOption struct {
	Timeout time.Duration
}

func (o *TimeoutOption) ApplyOption(opt *Config) {
	opt.Timeout = o.Timeout
}

func WithTimeout(timeout time.Duration) *TimeoutOption {
    if timeout == 0 {
    	timeout = 10000 * time.Millisecond
    }
	return &TimeoutOption{Timeout: timeout}
}

// Debug Option

type DebugOption struct {
	IsDebug bool
}

func (o *DebugOption) ApplyOption(opt *Config) {
	opt.IsDebug = o.IsDebug
}

func WithDebug(isDebug bool) *DebugOption {
	return &DebugOption{IsDebug: isDebug}
}

// App Token Option

type AppTokenOption struct {
	AppToken string
}

func (o *AppTokenOption) ApplyOption(opt *Config) {
	opt.AppToken = o.AppToken
}

func WithAppToken(appToken string) *AppTokenOption {
	return &AppTokenOption{AppToken: appToken}
}

// NewConfig creates a new Config object with the default values.

func NewConfig(options ...interface{}) Config {
	config := DefaultConfig
	for _, opt := range options {
		if o, ok := opt.(Option); ok {
			o.ApplyOption(&config)
		}
	}
	return config
}


type TenantIDOption struct {
	TenantID string
}

func (o *TenantIDOption) ApplyOption(opt *Config) {
	opt.TenantID = o.TenantID
}

func WithTenantID(tenantID string) *TenantIDOption {
	return &TenantIDOption{TenantID: tenantID}
}

// Option is an interface for the options of the Config object.

type Option interface {
	ApplyOption(*Config)
}


// GetConfig returns the Config object.

func GetConfig() Config {
	return DefaultConfig
}

// SetConfig sets the Config object.

func SetConfig(config Config) {
	DefaultConfig = config
}
