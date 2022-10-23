package server

import (
	"context"

	vault "github.com/hashicorp/vault/api"
	"github.com/vdbulcke/vault-token-monitor/internal"
	"go.uber.org/zap"
)

type VaultMonitorServer struct {
	logger *zap.Logger
	config *internal.Config
	ctx    context.Context
	client *vault.Client
}

// NewVaultMonitorServer Creates new NewVaultMonitorServer with background context
func NewVaultMonitorServer(l *zap.Logger, c *internal.Config) (*VaultMonitorServer, error) {

	config := vault.DefaultConfig()
	config.Address = c.Address

	// tls
	tlsConfig := &vault.TLSConfig{
		CACert:   c.CACertPEMPath,
		Insecure: c.SkipTLSValidation,
	}

	err := config.ConfigureTLS(tlsConfig)
	if err != nil {
		return nil, err
	}

	// create new client
	client, err := vault.NewClient(config)
	if err != nil {
		l.Error("unable to initialize Vault client:", zap.Error(err))
		return nil, err
	}

	// Authentication Token
	client.SetToken(c.Token)

	return &VaultMonitorServer{
		logger: l,
		config: c,
		ctx:    context.Background(),
		client: client,
	}, nil
}

// StartServer Starts the monitoring server as a blocking
// function call
func (v *VaultMonitorServer) StartServer() error {

	// create scheduler
	err := v.StartScheduler()
	if err != nil {
		v.logger.Error("error starting scheduler", zap.Error(err))
		return err
	}

	// start prometheus server
	// This is a block function call, until
	// OS Signal Interrupt are sent
	err = v.startPrometheusServer()
	if err != nil {
		v.logger.Error("error starting scheduler", zap.Error(err))
		return err
	}

	return nil

}
