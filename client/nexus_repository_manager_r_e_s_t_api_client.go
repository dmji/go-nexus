// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"nexus/client/assets"
	"nexus/client/azure_blob_store"
	"nexus/client/blob_store"
	"nexus/client/community_edition_eula"
	"nexus/client/components"
	"nexus/client/content_selectors"
	"nexus/client/email"
	formatsops "nexus/client/formats"
	"nexus/client/lifecycle"
	"nexus/client/malicious_risk_on_disk"
	"nexus/client/malicious_risk_visualizer"
	"nexus/client/manage_sonatype_repository_firewall_configuration"
	"nexus/client/product_licensing"
	"nexus/client/read_only"
	"nexus/client/repository_management"
	"nexus/client/routing_rules"
	"nexus/client/script"
	"nexus/client/search"
	"nexus/client/security_certificates"
	"nexus/client/security_management"
	"nexus/client/security_management_anonymous_access"
	"nexus/client/security_management_l_d_a_p"
	"nexus/client/security_management_privileges"
	"nexus/client/security_management_realms"
	"nexus/client/security_management_roles"
	"nexus/client/security_management_secrets_encryption"
	"nexus/client/security_management_users"
	"nexus/client/status"
	"nexus/client/support"
	"nexus/client/system_nodes"
	"nexus/client/tasks"
)

// Default nexus repository manager r e s t API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/service/rest/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new nexus repository manager r e s t API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *NexusRepositoryManagerRESTAPI {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new nexus repository manager r e s t API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *NexusRepositoryManagerRESTAPI {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new nexus repository manager r e s t API client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *NexusRepositoryManagerRESTAPI {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(NexusRepositoryManagerRESTAPI)
	cli.Transport = transport
	cli.Assets = assets.New(transport, formats)
	cli.AzureBlobStore = azure_blob_store.New(transport, formats)
	cli.BlobStore = blob_store.New(transport, formats)
	cli.CommunityEditionEula = community_edition_eula.New(transport, formats)
	cli.Components = components.New(transport, formats)
	cli.ContentSelectors = content_selectors.New(transport, formats)
	cli.Email = email.New(transport, formats)
	cli.Formats = formatsops.New(transport, formats)
	cli.Lifecycle = lifecycle.New(transport, formats)
	cli.MaliciousRiskOnDisk = malicious_risk_on_disk.New(transport, formats)
	cli.MaliciousRiskVisualizer = malicious_risk_visualizer.New(transport, formats)
	cli.ManageSonatypeRepositoryFirewallConfiguration = manage_sonatype_repository_firewall_configuration.New(transport, formats)
	cli.ProductLicensing = product_licensing.New(transport, formats)
	cli.ReadOnly = read_only.New(transport, formats)
	cli.RepositoryManagement = repository_management.New(transport, formats)
	cli.RoutingRules = routing_rules.New(transport, formats)
	cli.Script = script.New(transport, formats)
	cli.Search = search.New(transport, formats)
	cli.SecurityCertificates = security_certificates.New(transport, formats)
	cli.SecurityManagement = security_management.New(transport, formats)
	cli.SecurityManagementAnonymousAccess = security_management_anonymous_access.New(transport, formats)
	cli.SecurityManagementldap = security_management_l_d_a_p.New(transport, formats)
	cli.SecurityManagementPrivileges = security_management_privileges.New(transport, formats)
	cli.SecurityManagementRealms = security_management_realms.New(transport, formats)
	cli.SecurityManagementRoles = security_management_roles.New(transport, formats)
	cli.SecurityManagementSecretsEncryption = security_management_secrets_encryption.New(transport, formats)
	cli.SecurityManagementUsers = security_management_users.New(transport, formats)
	cli.Status = status.New(transport, formats)
	cli.Support = support.New(transport, formats)
	cli.SystemNodes = system_nodes.New(transport, formats)
	cli.Tasks = tasks.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// NexusRepositoryManagerRESTAPI is a client for nexus repository manager r e s t API
type NexusRepositoryManagerRESTAPI struct {
	Assets assets.ClientService

	AzureBlobStore azure_blob_store.ClientService

	BlobStore blob_store.ClientService

	CommunityEditionEula community_edition_eula.ClientService

	Components components.ClientService

	ContentSelectors content_selectors.ClientService

	Email email.ClientService

	Formats formatsops.ClientService

	Lifecycle lifecycle.ClientService

	MaliciousRiskOnDisk malicious_risk_on_disk.ClientService

	MaliciousRiskVisualizer malicious_risk_visualizer.ClientService

	ManageSonatypeRepositoryFirewallConfiguration manage_sonatype_repository_firewall_configuration.ClientService

	ProductLicensing product_licensing.ClientService

	ReadOnly read_only.ClientService

	RepositoryManagement repository_management.ClientService

	RoutingRules routing_rules.ClientService

	Script script.ClientService

	Search search.ClientService

	SecurityCertificates security_certificates.ClientService

	SecurityManagement security_management.ClientService

	SecurityManagementAnonymousAccess security_management_anonymous_access.ClientService

	SecurityManagementldap security_management_l_d_a_p.ClientService

	SecurityManagementPrivileges security_management_privileges.ClientService

	SecurityManagementRealms security_management_realms.ClientService

	SecurityManagementRoles security_management_roles.ClientService

	SecurityManagementSecretsEncryption security_management_secrets_encryption.ClientService

	SecurityManagementUsers security_management_users.ClientService

	Status status.ClientService

	Support support.ClientService

	SystemNodes system_nodes.ClientService

	Tasks tasks.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *NexusRepositoryManagerRESTAPI) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Assets.SetTransport(transport)
	c.AzureBlobStore.SetTransport(transport)
	c.BlobStore.SetTransport(transport)
	c.CommunityEditionEula.SetTransport(transport)
	c.Components.SetTransport(transport)
	c.ContentSelectors.SetTransport(transport)
	c.Email.SetTransport(transport)
	c.Formats.SetTransport(transport)
	c.Lifecycle.SetTransport(transport)
	c.MaliciousRiskOnDisk.SetTransport(transport)
	c.MaliciousRiskVisualizer.SetTransport(transport)
	c.ManageSonatypeRepositoryFirewallConfiguration.SetTransport(transport)
	c.ProductLicensing.SetTransport(transport)
	c.ReadOnly.SetTransport(transport)
	c.RepositoryManagement.SetTransport(transport)
	c.RoutingRules.SetTransport(transport)
	c.Script.SetTransport(transport)
	c.Search.SetTransport(transport)
	c.SecurityCertificates.SetTransport(transport)
	c.SecurityManagement.SetTransport(transport)
	c.SecurityManagementAnonymousAccess.SetTransport(transport)
	c.SecurityManagementldap.SetTransport(transport)
	c.SecurityManagementPrivileges.SetTransport(transport)
	c.SecurityManagementRealms.SetTransport(transport)
	c.SecurityManagementRoles.SetTransport(transport)
	c.SecurityManagementSecretsEncryption.SetTransport(transport)
	c.SecurityManagementUsers.SetTransport(transport)
	c.Status.SetTransport(transport)
	c.Support.SetTransport(transport)
	c.SystemNodes.SetTransport(transport)
	c.Tasks.SetTransport(transport)
}
