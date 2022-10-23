package server

import (
	"time"

	"github.com/carlescere/scheduler"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

// StartScheduler starts the scheduler with schedulerJob()
// as scheduled job
func (v *VaultMonitorServer) StartScheduler() error {
	// scheduler
	frequencyDuration, err := time.ParseDuration(v.config.SchedulerPeriodDuration)
	if err != nil {
		v.logger.Error("Error parsing  SchedulerPeriodDuration", zap.Error(err))
		return err
	}

	// start scheduler with job
	_, err = scheduler.Every(int(frequencyDuration.Seconds())).Seconds().Run(v.schedulerJob)
	if err != nil {
		v.logger.Error("Error starting scheduler", zap.Error(err))
		return err
	}

	return nil
}

// schedulerJob the actual steps perform by scheduler
// For each token in config, lookup Token TTL
// if token as a auto renew threshold, then renew the token
// Expose the token TTL as a prometheus metric
func (v *VaultMonitorServer) schedulerJob() {
	v.logger.Debug("Executing scheduler job")

	// for each token in config
	for _, token := range v.config.VaultAccessorMonitors {

		// lookup token
		vToken, err := v.vaultTokenLookup(token.Accessor)
		if err != nil {
			v.handleError(err, token.Accessor)
			continue
		}

		isAutoRenew := "false"
		if token.AutoRenewThresholdDuration != "" {
			isAutoRenew = "true"

			renewLimitDuration, err := time.ParseDuration(token.AutoRenewThresholdDuration)
			if err != nil {
				v.handleError(err, token.Accessor)
				continue
			}

			// if TTL is lower than auto renew threshold
			// then renew token
			if vToken.TTL < renewLimitDuration {
				vToken, err = v.vaultTokenRenew(token.Accessor)
				if err != nil {
					v.handleError(err, token.Accessor)
					continue
				}

				// count number of token renew per token
				promMetricTokenRenewCount.With(prometheus.Labels{
					"accessor":     token.Accessor,
					"display_name": vToken.DisplayName,
				}).Inc()

			}

		}

		// set prometheus metrics for  token
		promMetricTokenExpirationSeconds.With(prometheus.Labels{
			"display_name": vToken.DisplayName,
			"accessor":     token.Accessor,
			"auto_renew":   isAutoRenew,
		}).Set(float64(vToken.ExpireTimeUnix))

	}

}

// handleError record error looking up or renewing token as prometheus metric
func (v *VaultMonitorServer) handleError(err error, accessor string) {

	// increment error counter
	promMetricTokenErrorCount.With(prometheus.Labels{
		"accessor": accessor,
	}).Inc()

	v.logger.Error("error for token ", zap.String("accessor", accessor), zap.Error(err))
}
