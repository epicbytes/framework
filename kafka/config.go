package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
	"go.uber.org/config"
)

type Config struct {
	Addrs              []string `yaml:"addrs"`
	Version            string   `yaml:"version"`
	Username           string   `yaml:"username"`
	Password           string   `yaml:"password"`
	CertificateCrtPath string   `yaml:"certificate_path"`
	CertificateKeyPath string   `yaml:"certificate_key_path"`
	CertificateCAPath  string   `yaml:"certificate_ca_path"`
}

func (c *Config) ParsedVersion() (sarama.KafkaVersion, error) {
	return sarama.ParseKafkaVersion(c.Version)
}

func (c *Config) IsAuthRequired() bool {
	return c.Username != "" && c.Password != ""
}

func (c *Config) IsCertificateAuth() bool {
	return len(c.CertificateCrtPath) > 0 && len(c.CertificateKeyPath) > 0 && len(c.CertificateCAPath) > 0
}

func newConfig(provider config.Provider) (*Config, error) {
	var cfg Config
	if err := provider.Get("kafka").Populate(&cfg); err != nil {
		return nil, fmt.Errorf("kafka config: %w", err)
	}
	return &cfg, nil
}
