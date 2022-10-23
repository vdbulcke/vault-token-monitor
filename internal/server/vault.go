package server

import (
	"time"

	vault "github.com/hashicorp/vault/api"
	"go.uber.org/zap"
)

type ParsedVaultToken struct {
	Token          *vault.Secret
	DisplayName    string
	TTL            time.Duration
	ExpireTimeUnix int64
}

// vaultTokenRenew Renew Vault accessor token, then lookup renewed token
func (v *VaultMonitorServer) vaultTokenRenew(accessor string) (*ParsedVaultToken, error) {

	v.logger.Info("renew vault token", zap.String("accessor", accessor))

	tokenAuth := v.client.Auth().Token()

	tokenInfo, err := tokenAuth.RenewAccessorWithContext(v.ctx, accessor, 0)
	if err != nil {
		return nil, err
	}

	v.logger.Debug("vault renewed token info ", zap.String("accessor", accessor), zap.Any("token", tokenInfo))

	return v.vaultTokenLookup(accessor)
}

// vaultTokenLookup Lookup Vault accessor token,
func (v *VaultMonitorServer) vaultTokenLookup(accessor string) (*ParsedVaultToken, error) {
	v.logger.Debug("Lookup vault token", zap.String("accessor", accessor))

	tokenAuth := v.client.Auth().Token()

	tokenInfo, err := tokenAuth.LookupAccessorWithContext(v.ctx, accessor)
	if err != nil {
		return nil, err
	}

	v.logger.Debug("vault token lookup info ", zap.String("accessor", accessor), zap.Any("token", tokenInfo))
	return v.parseVaultToken(tokenInfo, accessor)
}

// parseVaultToken format Vault token as internal struct ParsedVaultToken
func (v *VaultMonitorServer) parseVaultToken(tokenInfo *vault.Secret, accessor string) (*ParsedVaultToken, error) {

	parsedToken := &ParsedVaultToken{
		Token: tokenInfo,
	}

	tokenDisplayName, ok := tokenInfo.Data["display_name"].(string)
	if !ok {
		v.logger.Info("Could not get display_name from token ", zap.String("accessor", accessor))
	} else {
		parsedToken.DisplayName = tokenDisplayName
	}

	tokenExpireTime, ok := tokenInfo.Data["expire_time"].(string)
	if ok {

		// convert expire Time
		t, err := time.Parse(time.RFC3339, tokenExpireTime)
		if err != nil {
			v.logger.Error("Could not get display_name from token ", zap.String("accessor", accessor), zap.Error(err))
		} else {
			parsedToken.ExpireTimeUnix = t.Unix()
		}
	}

	ttl, err := tokenInfo.TokenTTL()
	if err != nil {
		v.logger.Error("Error parsing Vault token TTL", zap.String("accessor", accessor), zap.Error(err))
		return parsedToken, err
	}

	parsedToken.TTL = ttl

	v.logger.Debug("token info ", zap.String("accessor", accessor), zap.Any("token", parsedToken))

	return parsedToken, nil

}
