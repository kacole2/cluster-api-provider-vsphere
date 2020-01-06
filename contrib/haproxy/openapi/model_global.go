/*
 * HAProxy Data Plane API
 *
 * API for editing and managing haproxy instances. Provides process information, configuration management, haproxy stats and logs.  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 1.2
 * Contact: support@haproxy.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Global HAProxy global configuration
type Global struct {
	CpuMaps               []GlobalCpuMaps     `json:"cpu_maps,omitempty"`
	Daemon                string              `json:"daemon,omitempty"`
	ExternalCheck         bool                `json:"external_check,omitempty"`
	MasterWorker          bool                `json:"master-worker,omitempty"`
	Maxconn               int32               `json:"maxconn,omitempty"`
	Nbproc                int32               `json:"nbproc,omitempty"`
	Nbthread              int32               `json:"nbthread,omitempty"`
	Pidfile               string              `json:"pidfile,omitempty"`
	RuntimeApis           []GlobalRuntimeApis `json:"runtime_apis,omitempty"`
	SslDefaultBindCiphers string              `json:"ssl_default_bind_ciphers,omitempty"`
	SslDefaultBindOptions string              `json:"ssl_default_bind_options,omitempty"`
	StatsTimeout          *int32              `json:"stats_timeout,omitempty"`
	TuneSslDefaultDhParam int32               `json:"tune_ssl_default_dh_param,omitempty"`
}