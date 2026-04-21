---
page_title: "pexip_infinity_global_configuration Resource - terraform-provider-pexip"
subcategory: ""
description: |-
  Manages the Pexip Infinity global system configuration.
---

# pexip_infinity_global_configuration (Resource)

Manages the Pexip Infinity global system configuration. This is a singleton resource that configures system-wide settings including protocol support, security parameters, cloud bursting, port ranges, timeouts, and administrative settings. Only one global configuration exists per Pexip Infinity cluster.

## Example Usage

### Basic Global Configuration

```terraform
resource "pexip_infinity_global_configuration" "basic" {
  enable_webrtc = true
  enable_sip    = true
  enable_h323   = false
  enable_rtmp   = true
  crypto_mode   = "besteffort"
}
```

### Enterprise Security Configuration

```terraform
resource "pexip_infinity_global_configuration" "enterprise" {
  enable_webrtc    = true
  enable_sip       = true
  enable_h323      = true
  enable_rtmp      = true
  crypto_mode      = "required"
  
  # Network configuration
  media_ports_start       = 40000
  media_ports_end         = 49999
  signalling_ports_start  = 5060
  signalling_ports_end    = 5080
  
  # Conference settings
  conference_create_permissions = "user_admin"
  conference_creation_mode      = "per_cluster"
  guests_only_timeout          = 30
  waiting_for_chair_timeout    = 10
  
  # Security and compliance
  enable_analytics      = true
  enable_error_reporting = false
  bandwidth_restrictions = "restricted"
  
  administrator_email = "admin@company.com"
}
```

### Cloud Bursting with AWS

```terraform
resource "pexip_infinity_global_configuration" "aws_bursting" {
  enable_webrtc    = true
  enable_sip       = true
  crypto_mode      = "besteffort"
  
  # Cloud bursting configuration
  bursting_enabled = true
  cloud_provider   = "aws"
  aws_access_key   = var.aws_access_key
  aws_secret_key   = var.aws_secret_key
  
  # Conference settings
  conference_create_permissions = "any_authenticated"
  conference_creation_mode      = "per_cluster"
  
  administrator_email = "pexip-admin@company.com"
}
```

### Azure Cloud Integration

```terraform
resource "pexip_infinity_global_configuration" "azure_integration" {
  enable_webrtc    = true
  enable_sip       = true
  enable_h323      = false
  enable_rtmp      = true
  crypto_mode      = "required"
  
  # Azure cloud bursting
  bursting_enabled = true
  cloud_provider   = "azure"
  azure_client_id  = var.azure_client_id
  azure_secret     = var.azure_secret
  
  # Media configuration
  max_pixels_per_second = "1920x1080x30"
  media_ports_start     = 36000
  media_ports_end       = 59999
  
  # Administrative settings
  enable_analytics       = true
  enable_error_reporting = true
  administrator_email    = "video-admin@company.com"
}
```

### High-Security Configuration

```terraform
resource "pexip_infinity_global_configuration" "high_security" {
  enable_webrtc    = true
  enable_sip       = true
  enable_h323      = false
  enable_rtmp      = false
  crypto_mode      = "required"
  
  # Restricted conference creation
  conference_create_permissions = "admin_only"
  conference_creation_mode      = "per_cluster"
  
  # Short timeouts for security
  guests_only_timeout       = 15
  waiting_for_chair_timeout = 5
  
  # Network restrictions
  bandwidth_restrictions = "restricted"
  media_ports_start      = 50000
  media_ports_end        = 50999
  signalling_ports_start = 5060
  signalling_ports_end   = 5061
  
  # Security monitoring
  enable_analytics       = true
  enable_error_reporting = false
  
  administrator_email = "security-admin@company.com"
  
  # Restrict conference creation to specific groups
  global_conference_create_groups = [
    "CN=Video Admins,OU=Groups,DC=company,DC=com",
    "CN=IT Managers,OU=Groups,DC=company,DC=com"
  ]
}
```

### Development Environment Configuration

```terraform
resource "pexip_infinity_global_configuration" "development" {
  enable_webrtc    = true
  enable_sip       = true
  enable_h323      = true
  enable_rtmp      = true
  crypto_mode      = "besteffort"
  
  # Relaxed conference creation for development
  conference_create_permissions = "any_authenticated"
  conference_creation_mode      = "per_node"
  
  # Extended timeouts for testing
  guests_only_timeout       = 120
  waiting_for_chair_timeout = 30
  
  # Analytics for development insights
  enable_analytics       = true
  enable_error_reporting = true
  
  administrator_email = "dev-team@company.com"
}
```

### Multi-Protocol Configuration

```terraform
resource "pexip_infinity_global_configuration" "multi_protocol" {
  # Enable all protocols
  enable_webrtc = true
  enable_sip    = true
  enable_h323   = true
  enable_rtmp   = true
  
  # Security settings
  crypto_mode = "besteffort"
  
  # Optimized port ranges for different protocols
  media_ports_start      = 32768
  media_ports_end        = 65535
  signalling_ports_start = 1024
  signalling_ports_end   = 65535
  
  # Video quality settings
  max_pixels_per_second = "3840x2160x30"  # 4K support
  
  # Conference management
  conference_create_permissions = "user_admin"
  conference_creation_mode      = "per_cluster"
  
  # Operational settings
  enable_analytics       = true
  enable_error_reporting = true
  bandwidth_restrictions = "none"
  
  administrator_email = "operations@company.com"
}
```

### Variable-Driven Configuration

```terraform
variable "global_config" {
  type = object({
    enable_webrtc                = bool
    enable_sip                   = bool
    enable_h323                  = bool
    enable_rtmp                  = bool
    crypto_mode                  = string
    conference_create_permissions = string
    conference_creation_mode     = string
    enable_analytics             = bool
    administrator_email          = string
  })
  default = {
    enable_webrtc                = true
    enable_sip                   = true
    enable_h323                  = false
    enable_rtmp                  = true
    crypto_mode                  = "besteffort"
    conference_create_permissions = "user_admin"
    conference_creation_mode     = "per_cluster"
    enable_analytics             = true
    administrator_email          = "admin@company.com"
  }
}

resource "pexip_infinity_global_configuration" "from_variable" {
  enable_webrtc                = var.global_config.enable_webrtc
  enable_sip                   = var.global_config.enable_sip
  enable_h323                  = var.global_config.enable_h323
  enable_rtmp                  = var.global_config.enable_rtmp
  crypto_mode                  = var.global_config.crypto_mode
  conference_create_permissions = var.global_config.conference_create_permissions
  conference_creation_mode     = var.global_config.conference_creation_mode
  enable_analytics             = var.global_config.enable_analytics
  administrator_email          = var.global_config.administrator_email
}
```


<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `aws_access_key` (String) The Amazon Web Services access key ID for the AWS user that the Pexip Infinity Management Node will use to log in to AWS and start and stop the node instances. Maximum length: 40 characters.
- `aws_secret_key` (String, Sensitive) The Amazon Web Services secret access key that is associated with the AWS access key ID.
- `azure_client_id` (String, Sensitive) The ID used to identify the client (sometimes referred to as Application ID).
- `azure_secret` (String, Sensitive) The Azure secret key that is associated with the Azure client ID.
- `azure_subscription_id` (String) The ID of an Azure subscription.
- `azure_tenant` (String) The Azure tenant ID that is associated with the Azure client ID.
- `bdpm_max_pin_failures_per_window` (Number) Sets the maximum number of PIN failures per service (e.g. VMR) in any sliding 10 minute windowed period, that are allowed from participants at unknown source addresses, before protective action is taken for that service. Range: 5 to 200.
- `bdpm_max_scan_attempts_per_window` (Number) Sets the maximum number of incorrect alias dial attempts in any sliding 10-minute windowed period, that are allowed from an unknown source address, before protective action is taken against that address. Range: 5 to 200.
- `bdpm_pin_checks_enabled` (Boolean) Select this option to instruct Pexip Infinity's Break-in Defense Policy Manager to temporarily block all access to a VMR that receives a significant number of incorrect PIN entry attempts (and thus may perhaps be under attack from a malicious actor). By default, this will block ALL new access attempts to a VMR for up to 10 minutes if more than 20 incorrect PIN entry attempts are made against that VMR in a 10 minute window.
- `bdpm_scan_quarantine_enabled` (Boolean) Select this option to instruct Pexip Infinity's Break-in Defense Policy Manager to temporarily block service access attempts from any source IP address that dials a significant number of incorrect aliases in a short period (and thus may perhaps be attempting to scan your deployment to discover valid aliases).
- `bursting_enabled` (Boolean) Select this option to instruct Pexip Infinity to monitor the system locations and start up / shutdown overflow Conferencing Nodes hosted in either Amazon Web Services (AWS) or Microsoft Azure when in need of extra capacity. For more information, see the Admin Guide section 'Dynamic bursting to a cloud service'.
- `bursting_min_lifetime` (Number) The minimum number of minutes that a cloud bursting node is kept powered on. Note that newly started cloud Conferencing Nodes can take up to 5 minutes to fully startup. Minimum: 5.
- `bursting_threshold` (Number) The bursting threshold controls when your overflow Conferencing Nodes in the cloud are automatically started up so that they can provide additional conferencing capacity. Minimum: 1.
- `cloud_provider` (String) Choose the cloud service provider to use for bursting. Valid values: `AWS`, `AZURE`, `GCP`.
- `contact_email_address` (String) An email address to be added to incident reports to allow Pexip to contact the system administrator for further information. Maximum length: 100 characters.
- `content_security_policy_header` (String) HTTP Content-Security-Policy header contents for Conferencing Nodes. Maximum length: 4096 characters.
- `content_security_policy_state` (Boolean) Enable HTTP Content-Security-Policy for Conferencing Nodes.
- `crypto_mode` (String) Controls the media encryption requirements for participants connecting to Pexip Infinity services. `on`: All participants must use media encryption. `besteffort`: Each participant will use media encryption if their device supports it. `off`: All H.323, SIP and MS-SIP participants must use unencrypted media. You can override this global setting for each individual service.
- `default_theme` (String) The theme to use for services that have no specific theme selected.
- `default_webapp_alias` (String) The web app path to use by default on conferencing nodes.
- `disabled_codecs` (Set of String) Choose codecs to disable.
- `eject_last_participant_backstop_timeout` (Number) The length of time (in seconds) for which a conference will continue with only one participant remaining (independent of Host/Guest role). Must be 0 (never eject) or between 60 and 86400. Default: 0.
- `enable_analytics` (Boolean) Select this option to allow submission of deployment and usage statistics to Pexip. This will help us improve the product.
- `enable_application_api` (Boolean) Enable or disable support for Pexip Infinity Client API. This is required for integration with Infinity Connect browser-based and desktop clients, the Pexip Mobile App for iOS and Android, and any other third-party applications that use the client API, as well as for integration with Microsoft Teams.
- `enable_breakout_rooms` (Boolean) Enable the Breakout Rooms feature on VMRs.
- `enable_chat` (Boolean) Enables relay of chat messages between conference participants using Skype for Business and Infinity Connect clients. You can also configure this setting on individual Virtual Meeting Rooms and Virtual Auditoriums.
- `enable_clock` (Boolean) Enables support for displaying an in-conference timer or countdown clock.
- `enable_denoise` (Boolean) Enable server side denoising for speech from noisy participants (see documentation for ways to enable it for a VMR).
- `enable_dialout` (Boolean) Enables calls via the Distributed Gateway, and allows users of Pexip Infinity Connect, the Pexip Mobile Apps for iOS and Android, and the Pexip management web interface to add participants to a conference.
- `enable_directory` (Boolean) When disabled, Infinity Connect clients will display aliases from their own call history only. When enabled, registered Infinity Connect clients will additionally display the aliases of VMRs, Virtual Auditoriums, Virtual Receptions, and devices registered to the Pexip Infinity deployment.
- `enable_edge_non_mesh` (Boolean) Enable the restricted IPsec network routing requirements of Proxying Edge Nodes. When enabled, if a location only contains Proxying Edge Nodes, then those nodes only require IPsec connectivity with other nodes in that location, the transcoding location, the primary and secondary overflow locations, and with the Management Node.
- `enable_fecc` (Boolean) Enables Connect apps and SIP/H.323 endpoints to send Far-End Camera Control (FECC) signals to supporting endpoints, in order to pan, tilt and zoom the device's camera.
- `enable_h323` (Boolean) Enable the H323 protocol on all Conferencing Nodes.
- `enable_legacy_dialout_api` (Boolean) Enables outbound calls from a VMR using the legacy dialout API. When disabled, outbound calls are only permitted by following Call Routing Rules.
- `enable_lync_auto_escalate` (Boolean) Determines whether a Skype for Business audio call is automatically escalated so that it receives video from a conference.
- `enable_lync_vbss` (Boolean) Determines whether Video-based Screen Sharing (VbSS) is enabled for Skype for Business calls.
- `enable_mlvad` (Boolean) Enable Voice Focus for advanced voice activity detection.
- `enable_rtmp` (Boolean) Enables RTMP calls on all Conferencing Nodes. This allows Infinity Connect clients that use RTMP to access Pexip Infinity services, and allows conference content to be output to streaming and recording services.
- `enable_sip` (Boolean) Enable the SIP protocol over TCP and TLS on all Conferencing Nodes.
- `enable_sip_udp` (Boolean) Enable incoming calls using the SIP protocol over UDP on all Conferencing Nodes. If changing from enabled to disabled, all Conferencing Nodes must be rebooted.
- `enable_softmute` (Boolean) Enable Softmute for advance speech-aware audio gating (see documentation for ways to enable it for a VMR). Note that this does not remove any noise from the audio.
- `enable_ssh` (Boolean) Allows an administrator to log in to the Management and Conferencing Nodes over SSH. This setting can be overridden on individual nodes.
- `enable_turn_443` (Boolean) Enable media relay on TCP port 443 for WebRTC clients as a fallback.
- `enable_webrtc` (Boolean) Enables WebRTC calls on all Conferencing Nodes. This allows access to Pexip Infinity services from Infinity Connect clients that use WebRTC, including Google Chrome, Microsoft Edge, Firefox, Opera and Safari (version 11 onwards) browsers, and the Infinity Connect desktop client.
- `error_reporting_enabled` (Boolean) Select this option to permit submission of incident reports.
- `error_reporting_url` (String) The URL to which incident reports will be sent. Maximum length: 255 characters.
- `es_connection_timeout` (Number) Maximum number of seconds allowed to connect, send, and wait for a response.
- `es_initial_retry_backoff` (Number) Initial time, in seconds, for the first retry attempt when an event cannot be delivered.
- `es_maximum_retry_backoff` (Number) Maximum number of seconds allowed for the retry backoff before raising an alarm and stopping the event publisher.
- `es_media_streams_wait` (Number) Maximum time, in seconds, to wait for an end-of-call media stream message.
- `es_metrics_update_interval` (Number) Time between metrics updates. To disable eventsink metrics, enter 0.
- `es_short_term_memory_expiration` (Number) Internal cache expiration time in seconds. Used to briefly store 'participant_disconnected' events in order to gather end-of-call media statistics.
- `external_participant_avatar_lookup` (Boolean) Determines whether or not avatars for external participants will be retrieved using the method appropriate for the external meeting type.
- `gcp_client_email` (String) The GCP service account ID.
- `gcp_private_key` (String) The private key for the Google Cloud Platform service account user that the Pexip Infinity Management Node will use to log in to GCP and start and stop the node instances. Maximum length: 12288 characters.
- `gcp_project_id` (String) The ID of the GCP project containing bursting nodes.
- `guests_only_timeout` (Number) The length of time (in seconds) for which a conference will continue with only Guest participants, after all Host participants have left. Range: 0 to 86400. Default: 60.
- `legacy_api_password` (String) The password presented to Pexip Infinity by external systems attempting to authenticate with it. Maximum length: 100 characters.
- `legacy_api_username` (String) The username presented to Pexip Infinity by external systems attempting to authenticate with it. Maximum length: 100 characters.
- `live_captions_vmr_default` (Boolean) This option controls whether live captions are enabled by default on all VMRs, Virtual Auditoriums and Call Routing Rules. You can override this setting on each service individually.
- `liveview_show_conferences` (Boolean) Whether to show conferences and backplanes in Live View.
- `local_mssip_domain` (String) The name of the SIP domain that is routed from Skype for Business to Pexip Infinity, either as a static route or via federation. It is also used as the default domain in the From address for outgoing SIP gateway calls and outbound SIP calls from conferences without a valid SIP URI as an alias. Maximum length: 255 characters.
- `logon_banner` (String) Text of the message to display on the login page of the Pexip Infinity administrator web interface. Maximum length: 4096 characters.
- `logs_max_age` (Number) The maximum number of days of logs and call history to retain on Pexip nodes. 0 to disable. Range: 0 to 3650 days. Default: 0 (disabled).
- `management_qos` (Number) The DSCP value for management traffic sent from the Management Node and Conferencing Nodes. Range: 0 to 63.
- `management_session_timeout` (Number) The number of minutes a browser session may remain idle before the user is logged out of the Management Node administration interface. Range: 5 to 1440. Default: 30 minutes.
- `management_start_page` (String) The first page you see after logging into the Management Web.
- `max_callrate_in` (Number) This optional field allows you to limit the bandwidth of media being received by Pexip Infinity from individual participants, for calls where bandwidth limits have not otherwise been specified. Range: 128 to 8192.
- `max_callrate_out` (Number) This optional field allows you to limit the bandwidth of media being sent by Pexip Infinity to individual participants, for calls where bandwidth limits have not otherwise been specified. Range: 128 to 8192. Default: 4128.
- `max_pixels_per_second` (String) Sets the maximum call quality for participants connecting to Pexip Infinity services (VMRs, gateway calls etc.). You can also override this setting on individual services and call routing rules. Valid values: `sd`, `hd`, `fullhd`.
- `max_presentation_bandwidth_ratio` (Number) The maximum percentage of call bandwidth to be allocated to sending presentation. Range: 25 to 75. Default: 75.
- `media_ports_end` (Number) The end value for the range of ports (UDP and TCP) that all Conferencing Nodes will use to send media (for all call protocols). The media port range must contain at least 100 ports. Range: 10000 to 49999. Default: 49999.
- `media_ports_start` (Number) The start value for the range of ports (UDP and TCP) that all Conferencing Nodes will use to send media (for all call protocols). The media port range must contain at least 100 ports. Range: 10000 to 49999. Default: 40000.
- `ocsp_responder_url` (String) The URL to which OCSP requests will be sent either if the OCSP state is set to Override, or if the OCSP state is set to On but there is no URL specified in the TLS certificate. Maximum length: 255 characters.
- `ocsp_state` (String) Whether to use OCSP when checking the validity of TLS certificates. `ON`: An OCSP request will be sent to the URL specified in the TLS certificate. `OVERRIDE`: An OCSP request will be sent to the URL specified in the OCSP responder URL field. Valid values: `OFF`, `ON`, `OVERRIDE`.
- `pin_entry_timeout` (Number) The length of time (in seconds) for which a participant will be permitted to remain at the PIN entry screen before being disconnected. Range: 30 to 86400. Default: 120.
- `session_timeout_enabled` (Boolean) Determines whether inactive users are automatically logged out of the Management Node administration interface after a period of time. If disabled, users of the administrator interface are never timed out.
- `signalling_ports_end` (Number) The end value for the range of ports (UDP and TCP) that all Conferencing Nodes will use to send signaling (for H.323, H.245 and SIP). Range: 10000 to 49999. Default: 39999.
- `signalling_ports_start` (Number) The start value for the range of ports (UDP and TCP) that all Conferencing Nodes will use to send signaling (for H.323, H.245 and SIP). Range: 10000 to 49999. Default: 33000.
- `sip_tls_cert_verify_mode` (String) Determines whether to verify the peer certificate for connections over SIP TLS. `OFF`: the peer certificate will not be verified; all connections will be allowed. `ON`: the peer certificate will be verified, and the peer's remote identities will be compared against the Application Unique String (AUS). Valid values: `OFF`, `ON`.
- `site_banner` (String) Text of the banner to display on the top of every page of this Pexip Infinity administrator web interface. Maximum length: 255 characters.
- `site_banner_bg` (String) The background color for the site banner.
- `site_banner_fg` (String) The text color for the site banner.
- `teams_enable_powerpoint_render` (Boolean) This setting is intended for future use to enable PowerPoint Live content in Microsoft Teams calls. Check the online documentation for the latest status for this feature.
- `waiting_for_chair_timeout` (Number) The length of time (in seconds) for which a Guest participant will remain at the waiting screen if a Host does not join, before being disconnected. Range: 0 to 86400. Default: 900.

### Read-Only

- `deployment_uuid` (String) The ID of the deployment.
- `id` (String) Resource URI for the global configuration in Infinity
- `resource_uri` (String) The URI that identifies this resource.


## Import

Import is supported using any value as the import ID (since this is a singleton resource):

```shell
terraform import pexip_infinity_global_configuration.example global
```

## Usage Notes

### Singleton Resource

This is a singleton resource, meaning only one global configuration exists per Pexip Infinity cluster. Creating multiple instances of this resource in your Terraform configuration will result in conflicts.

### Protocol Support

- **WebRTC**: Required for browser-based video conferences and web applications
- **SIP**: Required for SIP-based endpoints and telephony integration
- **H.323**: Required for legacy H.323 video conferencing systems
- **RTMP**: Required for streaming and recording functionality

### Crypto Modes

- **disabled**: No encryption requirements (not recommended for production)
- **besteffort**: Encryption preferred but not required
- **required**: Encryption mandatory for all connections

### Port Configuration

- **Media Ports**: Used for RTP/RTCP media traffic (audio/video)
- **Signalling Ports**: Used for SIP and H.323 signalling
- **Range Planning**: Ensure adequate port ranges for expected concurrent calls
- **Firewall Coordination**: Coordinate port ranges with firewall and security teams

### Cloud Bursting

- **AWS Integration**: Requires AWS access key and secret key
- **Azure Integration**: Requires Azure client ID and secret
- **Google Cloud**: Supported for cloud bursting (credentials vary)
- **Scalability**: Enables automatic scaling to cloud resources during peak usage

### Conference Creation Permissions

- **none**: Conferences cannot be created dynamically
- **admin_only**: Only administrators can create conferences
- **user_admin**: Administrators and designated users can create conferences
- **any_authenticated**: Any authenticated user can create conferences

### Timeout Configuration

- **Guests Only**: Automatically ends conferences with only guest participants
- **Waiting for Chair**: Ends waiting rooms when no chair joins
- **Zero Values**: Setting to 0 disables the timeout

### Security Considerations

- **Encryption**: Use "required" crypto mode for sensitive environments
- **Access Control**: Restrict conference creation based on organizational needs
- **Monitoring**: Enable analytics for security monitoring and compliance
- **Credentials**: Store cloud credentials securely using Terraform variables

## Troubleshooting

### Common Issues

**Global Configuration Update Fails**
- Verify all enum values are from the valid options list
- Check that port ranges are valid and start <= end
- Ensure timeout values are within the valid range (0-1440)
- Verify cloud credentials are correctly formatted

**Protocol Support Issues**
- Ensure required protocols are enabled for your use case
- Check licensing requirements for certain protocols
- Verify client compatibility with enabled protocols
- Monitor protocol-specific logs for connection issues

**Port Range Conflicts**
- Ensure media and signalling port ranges don't overlap inappropriately
- Verify port ranges don't conflict with other system services
- Check firewall rules allow traffic on configured port ranges
- Monitor port utilization to ensure adequate range size

**Cloud Bursting Authentication Fails**
- Verify cloud provider credentials are correct and current
- Ensure cloud provider account has necessary permissions
- Check network connectivity to cloud provider APIs
- Monitor cloud provider service status for outages

**Conference Creation Issues**
- Verify conference_create_permissions setting allows intended users
- Check that global_conference_create_groups contains correct group DNs
- Ensure users have appropriate authentication and authorization
- Monitor conference creation logs for detailed error information

**Timeout Configuration Problems**
- Verify timeout values are appropriate for your use case
- Check that timeout settings don't conflict with user expectations
- Monitor conference duration to optimize timeout values
- Ensure timeout notifications are properly configured

**Analytics and Reporting Issues**
- Verify analytics collection is enabled if reporting is required
- Check that error reporting settings meet compliance requirements
- Ensure adequate storage for analytics data
- Monitor analytics pipeline for data collection issues

**Performance Degradation**
- Monitor system performance after configuration changes
- Check media quality settings against available bandwidth
- Verify port range sizes are adequate for concurrent users
- Monitor cloud bursting utilization and costs

**Import Issues**
- Use any string value as the import ID for this singleton resource
- Verify provider credentials have access to global configuration
- Check that the Pexip Infinity cluster is accessible
- Note that sensitive cloud credentials may not be imported

**Encryption Configuration Problems**
- Verify crypto_mode setting is compatible with all endpoints
- Check that "required" mode doesn't block legitimate connections
- Monitor encryption negotiation logs for compatibility issues
- Ensure proper certificate configuration for encrypted connections

**Email Notification Failures**
- Verify administrator_email address is valid and accessible
- Check SMTP configuration for email delivery
- Ensure email notifications are not blocked by spam filters
- Monitor email delivery logs for notification issues

**Bandwidth Restriction Issues**
- Verify bandwidth_restrictions setting matches network capacity
- Check that "restricted" mode doesn't overly limit call quality
- Monitor bandwidth utilization and quality metrics
- Adjust restrictions based on network performance data
