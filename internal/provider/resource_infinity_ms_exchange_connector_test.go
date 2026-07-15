/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package provider

import (
	"os"
	"testing"

	"github.com/pexip/go-infinity-sdk/v41/config"
	"github.com/pexip/go-infinity-sdk/v41/types"
	"github.com/stretchr/testify/mock"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/pexip/go-infinity-sdk/v41"

	"github.com/pexip/terraform-provider-infinity/internal/test"
)

func TestInfinityMsExchangeConnector(t *testing.T) {
	t.Parallel()
	_ = os.Setenv("TF_ACC", "1")

	client := infinity.NewClientMock()

	createResponse := &types.PostResponse{
		Body:        []byte(""),
		ResourceURI: "/api/admin/configuration/v1/ms_exchange_connector/1/",
	}
	client.On("PostWithResponse", mock.Anything, "configuration/v1/ms_exchange_connector/", mock.Anything, mock.Anything).Return(createResponse, nil)

	stringPtr := func(s string) *string { return &s }

	mockState := &config.MsExchangeConnector{
		ID:                             1,
		ResourceURI:                    "/api/admin/configuration/v1/ms_exchange_connector/1/",
		Name:                           "tf-test-ms-exchange-connector",
		Description:                    "tf-test MS Exchange Connector",
		MeetingBufferBefore:            60,
		MeetingBufferAfter:             90,
		ScheduledAliasSuffixLength:     8,
		RoomMailboxEmailAddress:        stringPtr("test@example.com"),
		RoomMailboxName:                "tf-test-ms-exchange-connector",
		URL:                            "https://example.com",
		Username:                       "tf-test-ms-exchange-connector",
		Password:                       "test-value",
		AuthenticationMethod:           "OAUTH",
		AuthProvider:                   "AZURE",
		UUID:                           "test-uuid",
		ScheduledAliasPrefix:           stringPtr("test-value"),
		ScheduledAliasDomain:           "example.com",
		EnableDynamicVmrs:              true,
		EnablePersonalVmrs:             true,
		AllowNewUsers:                  false,
		DisableProxy:                   true,
		UseCustomAddInSources:          true,
		EnableAddinDebugLogs:           true,
		OauthClientID:                  stringPtr("test-value"),
		OauthClientSecret:              "test-value",
		OauthAuthEndpoint:              "test-value",
		OauthTokenEndpoint:             "test-value",
		OauthRedirectURI:               "test-value",
		OauthRefreshToken:              "",
		OauthState:                     nil,
		KerberosRealm:                  "test-value",
		KerberosKdc:                    "test-value",
		KerberosKdcHttpsProxy:          "test-value",
		KerberosExchangeSpn:            "test-value",
		KerberosEnableTls:              true,
		KerberosAuthEveryRequest:       true,
		KerberosVerifyTlsUsingCustomCa: true,
		AddinServerDomain:              "test-value",
		AddinDisplayName:               "tf-test-ms-exchange-connector",
		AddinDescription:               "tf-test MS Exchange Connector",
		AddinProviderName:              "tf-test-ms-exchange-connector",
		AddinButtonLabel:               "test-value",
		AddinGroupLabel:                "test-value",
		AddinSupertipTitle:             "test-value",
		AddinSupertipDescription:       "tf-test MS Exchange Connector",
		AddinApplicationID:             stringPtr("test-value"),
		AddinAuthorityURL:              "https://example.com",
		AddinOidcMetadataURL:           "https://example.com",
		AddinAuthenticationMethod:      "EXCHANGE_USER_ID_TOKEN",
		AddinNaaWebApiApplicationID:    stringPtr("test-value"),
		PersonalVmrOauthClientID:       stringPtr("test-value"),
		PersonalVmrOauthClientSecret:   "test-value",
		PersonalVmrOauthAuthEndpoint:   "test-value",
		PersonalVmrOauthTokenEndpoint:  "test-value",
		PersonalVmrAdfsRelyingPartyTrustIdentifier: "test-value",
		OfficeJsURL:                  "https://example.com",
		MicrosoftFabricURL:           "https://example.com",
		MicrosoftFabricComponentsURL: "https://example.com",
		AdditionalAddInScriptSources: "test-value",
		HostIdentityProviderGroup:    stringPtr("test-server.example.com"),
		IvrTheme:                     stringPtr("test-value"),
		NonIdpParticipants:           "disallow_all",
		// Template fields with API defaults
		AcceptEditedOccurrenceTemplate:      "<div style=\"font-size:11.0pt; color:#000000; font-family:Calibri,Arial,Helvetica,sans-serif;\">\r\nThis meeting occurrence in a recurring series has been successfully rescheduled using the aliases: {{alias}} and {{numeric_alias}}.<br>\r\n</div>",
		AcceptEditedRecurringSeriesTemplate: "<div style=\"font-size:11.0pt; color:#000000; font-family:Calibri,Arial,Helvetica,sans-serif;\">\r\nThis recurring meeting series has been successfully rescheduled.<br>\r\nAll meetings in this series will use the aliases: {{alias}} and {{numeric_alias}}.<br>\r\n</div>",
		AcceptEditedSingleMeetingTemplate:   "<div style=\"font-size:11.0pt; color:#000000; font-family:Calibri,Arial,Helvetica,sans-serif;\">\r\nThis meeting has been successfully rescheduled using the aliases: {{alias}} and {{numeric_alias}}.<br>\r\n</div>",
		AcceptNewRecurringSeriesTemplate:    "<div style=\"font-size:11.0pt; color:#000000; font-family:Calibri,Arial,Helvetica,sans-serif;\">\r\nThis recurring meeting series has been successfully scheduled.<br>\r\nAll meetings in this series will use the aliases: {{alias}} and {{numeric_alias}}.<br>\r\n</div>",
		AcceptNewSingleMeetingTemplate:      "<div style=\"font-size:11.0pt; color:#000000; font-family:Calibri,Arial,Helvetica,sans-serif;\">\r\nThis meeting has been successfully scheduled using the aliases: {{alias}} and {{numeric_alias}}.<br>\r\n</div>",
		ConferenceDescriptionTemplate:       "Scheduled Conference booked by {{organizer_email}}",
		ConferenceNameTemplate:              "{{subject}} ({{organizer_name}})",
		ConferenceSubjectTemplate:           "{{subject}}",
		MeetingInstructionsTemplate:         "<br>\r\n<div style=\"font-size:11.0pt; color:#000000; font-family:Calibri,Arial,Helvetica,sans-serif;\">\r\n<b>Please join my Pexip Virtual Meeting Room in one of the following ways:</b><br>\r\n<br>\r\nFrom a VC endpoint or a Skype/Lync client:<br>\r\n{{alias}}<br>\r\n<br>\r\nFrom a web browser:<br>\r\n<a href=\"https://{{addin_server_domain}}/webapp/#/?conference={{alias}}\">https://{{addin_server_domain}}/webapp/#/?conference={{alias}}</a><br>\r\n<br>\r\nFrom a Pexip Infinity Connect client:<br>\r\npexip://{{alias}}<br>\r\n<br>\r\nFrom a telephone:<br>\r\n[Your number], then {{numeric_alias}} #<br>\r\n<br>\r\n{{alias_uuid}}<br>\r\n</div>",
		PersonalVmrDescriptionTemplate:      "{{description}}",
		PersonalVmrInstructionsTemplate:     "{% if domain_aliases %}\r\n    {% set alias = domain_aliases[0] %}\r\n{% elif other_aliases %}\r\n    {% set alias = other_aliases[0] %}\r\n{% else %}\r\n    {% set alias = numeric_aliases[0] %}\r\n{% endif %}\r\n{% if (not allow_guests) and pin %}\r\n    {% set meeting_pin = pin %}\r\n{% elif allow_guests and guest_pin %}\r\n    {% set meeting_pin = guest_pin %}\r\n{% else %}\r\n    {% set meeting_pin = \"\" %}\r\n{% endif %}\r\n<br>\r\n<div style=\"font-size:11.0pt; color:#000000; font-family:Calibri,Arial,Helvetica,sans-serif;\">\r\n<b>Please join my Pexip Virtual Meeting Room in one of the following ways:</b><br>\r\n<br>\r\nFrom a VC endpoint or a Skype/Lync client:<br>\r\n{{alias}}<br>\r\n<br>\r\nFrom a web browser:<br>\r\n<a href=\"https://{{addin_server_domain}}/webapp/#/?conference={{alias}}\">https://{{addin_server_domain}}/webapp/#/?conference={{alias}}</a><br>\r\n<br>\r\nFrom a Pexip Infinity Connect client:<br>\r\npexip://{{alias}}<br>\r\n<br>\r\n{% if numeric_aliases %}\r\nFrom a telephone:<br>\r\n[Your number], then {{numeric_aliases[0]}} #<br>\r\n<br>\r\n{% endif %}\r\n{% if meeting_pin %}\r\nPlease join using the PIN <b>{{meeting_pin}}</b><br>\r\n<br>\r\n{% endif %}\r\n</div>",
		PersonalVmrLocationTemplate:         "{% if domain_aliases %}\r\n    {% set alias = domain_aliases[0] %}\r\n{% elif other_aliases %}\r\n    {% set alias = other_aliases[0] %}\r\n{% else %}\r\n    {% set alias = numeric_aliases[0] %}\r\n{% endif %}\r\nhttps://{{addin_server_domain}}/webapp/#/?conference={{alias}}",
		PersonalVmrNameTemplate:             "{{name}}",
		PlaceholderInstructionsTemplate:     "<div style=\"font-size:11.0pt; color:#000000; font-family:Calibri,Arial,Helvetica,sans-serif;\">\r\nThis meeting will be hosted in a Virtual Meeting Room. Joining instructions will be<br>\r\nsent to you soon in a separate email.<br>\r\n</div>",
		RejectAliasConflictTemplate:         "<div style=\"font-size:11.0pt; color:#000000; font-family:Calibri,Arial,Helvetica,sans-serif;\">\r\nWe are unable to schedule this meeting because the alias: {{alias}} is already <br>\r\nin use by another Pexip Virtual Meeting Room. Please try creating a new meeting.<br>\r\n</div>",
		RejectAliasDeletedTemplate:          "<div style=\"font-size:11.0pt; color:#000000; font-family:Calibri,Arial,Helvetica,sans-serif;\">\r\nWe are unable to schedule this meeting because its alias has been deleted.<br>\r\nPlease try creating a new meeting.<br>\r\n</div>",
		RejectGeneralErrorTemplate:          "<div style=\"font-size:11.0pt; color:#000000; font-family:Calibri,Arial,Helvetica,sans-serif;\">\r\nWe are unable to schedule this meeting. Please try creating a new meeting.<br>\r\nIf this issue continues, please forward this message to your system administrator, including the following ID:<br>\r\nCorrelationID=\"{{correlation_id}}\".<br>\r\n</div>",
		RejectInvalidAliasIDTemplate:        "<div style=\"font-size:11.0pt; color:#000000; font-family:Calibri,Arial,Helvetica,sans-serif;\">\r\nThis meeting request does not contain currently valid scheduling data, and therefore cannot be processed.<br>\r\nPlease use the add-in to create a new meeting request, without editing any of the content that is inserted by the add-in.<br>\r\nIf this issue continues, please contact your system administrator.<br>\r\n</div>",
		RejectRecurringSeriesPastTemplate:   "<div style=\"font-size:11.0pt; color:#000000; font-family:Calibri,Arial,Helvetica,sans-serif;\">\r\nThis recurring series cannot be scheduled because all<br>\r\noccurrences happen in the past.<br>\r\n</div>",
		RejectSingleMeetingPast:             "<div style=\"font-size:11.0pt; color:#000000; font-family:Calibri,Arial,Helvetica,sans-serif;\">\r\nThis meeting cannot be scheduled because it occurs in the past.<br>\r\n</div>",
		ScheduledAliasDescriptionTemplate:   "Scheduled Conference booked by {{organizer_email}}",
		// Add-in pane fields with defaults
		AddinPaneTitle:                                   "Add a VMR",
		AddinPaneDescription:                             "This assigns a Virtual Meeting Room for your meeting",
		AddinPaneButtonTitle:                             "Add a Single-use VMR",
		AddinPaneSuccessHeading:                          "Success",
		AddinPaneSuccessMessage:                          "This meeting is now set up to be hosted as an audio or video conference in a Virtual Meeting Room. Please note this conference is not scheduled until you select Send.",
		AddinPaneAlreadyVideoMeetingHeading:              "VMR already assigned",
		AddinPaneAlreadyVideoMeetingMessage:              "It looks like this meeting has already been set up to be hosted in a Virtual Meeting Room. If this is a new meeting, select Send to schedule the conference.",
		AddinPaneGeneralErrorHeading:                     "Error",
		AddinPaneGeneralErrorMessage:                     "There was a problem adding the joining instructions. Please try again.",
		AddinPaneManagementNodeDownHeading:               "Cannot assign a VMR right now",
		AddinPaneManagementNodeDownMessage:               "Sorry, we are unable to assign a Virtual Meeting Room at this time. Select Send to schedule the meeting, and all attendees will be sent joining instructions later.",
		AddinPanePersonalVmrAddButton:                    "Add a Personal VMR",
		AddinPanePersonalVmrSignInButton:                 "Sign In",
		AddinPanePersonalVmrSelectMessage:                "Select the VMR you want to add to the meeting",
		AddinPanePersonalVmrNoneMessage:                  "You do not have any personal VMRs",
		AddinPanePersonalVmrErrorGettingMessage:          "There was a problem getting your personal VMRs. Please try again.",
		AddinPanePersonalVmrErrorSigningInMessage:        "There was a problem signing you in. Please try again.",
		AddinPanePersonalVmrErrorInsertingMeetingMessage: "There was a problem adding the joining instructions. Please try again.",
	}

	client.On("GetJSON", mock.Anything, "configuration/v1/ms_exchange_connector/1/", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		dest := args.Get(3).(*config.MsExchangeConnector)
		*dest = *mockState
	}).Maybe()

	client.On("PutJSON", mock.Anything, "configuration/v1/ms_exchange_connector/1/", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		req := args.Get(2).(*config.MsExchangeConnectorUpdateRequest)
		dest := args.Get(3).(*config.MsExchangeConnector)

		// Non-pointer string/bool fields (always set from plan)
		mockState.Name = req.Name
		mockState.Description = req.Description
		mockState.RoomMailboxName = req.RoomMailboxName
		mockState.URL = req.URL
		mockState.Username = req.Username
		mockState.Password = req.Password
		mockState.AuthenticationMethod = req.AuthenticationMethod
		mockState.AuthProvider = req.AuthProvider
		mockState.ScheduledAliasDomain = req.ScheduledAliasDomain
		mockState.OauthClientSecret = req.OauthClientSecret
		mockState.OauthAuthEndpoint = req.OauthAuthEndpoint
		mockState.OauthTokenEndpoint = req.OauthTokenEndpoint
		mockState.OauthRedirectURI = req.OauthRedirectURI
		mockState.KerberosRealm = req.KerberosRealm
		mockState.KerberosKdc = req.KerberosKdc
		mockState.KerberosKdcHttpsProxy = req.KerberosKdcHttpsProxy
		mockState.KerberosExchangeSpn = req.KerberosExchangeSpn
		mockState.AddinServerDomain = req.AddinServerDomain
		mockState.AddinDisplayName = req.AddinDisplayName
		mockState.AddinDescription = req.AddinDescription
		mockState.AddinProviderName = req.AddinProviderName
		mockState.AddinButtonLabel = req.AddinButtonLabel
		mockState.AddinGroupLabel = req.AddinGroupLabel
		mockState.AddinSupertipTitle = req.AddinSupertipTitle
		mockState.AddinSupertipDescription = req.AddinSupertipDescription
		mockState.AddinAuthorityURL = req.AddinAuthorityURL
		mockState.AddinOidcMetadataURL = req.AddinOidcMetadataURL
		mockState.AddinAuthenticationMethod = req.AddinAuthenticationMethod
		mockState.PersonalVmrOauthClientSecret = req.PersonalVmrOauthClientSecret
		mockState.PersonalVmrOauthAuthEndpoint = req.PersonalVmrOauthAuthEndpoint
		mockState.PersonalVmrOauthTokenEndpoint = req.PersonalVmrOauthTokenEndpoint
		mockState.PersonalVmrAdfsRelyingPartyTrustIdentifier = req.PersonalVmrAdfsRelyingPartyTrustIdentifier
		mockState.OfficeJsURL = req.OfficeJsURL
		mockState.MicrosoftFabricURL = req.MicrosoftFabricURL
		mockState.MicrosoftFabricComponentsURL = req.MicrosoftFabricComponentsURL
		mockState.AdditionalAddInScriptSources = req.AdditionalAddInScriptSources
		mockState.NonIdpParticipants = req.NonIdpParticipants

		// Pointer fields - unconditionally assign so they are cleared when nil
		mockState.RoomMailboxEmailAddress = req.RoomMailboxEmailAddress
		mockState.ScheduledAliasPrefix = req.ScheduledAliasPrefix
		mockState.OauthClientID = req.OauthClientID
		mockState.AddinApplicationID = req.AddinApplicationID
		mockState.AddinNaaWebApiApplicationID = req.AddinNaaWebApiApplicationID
		mockState.PersonalVmrOauthClientID = req.PersonalVmrOauthClientID
		mockState.HostIdentityProviderGroup = req.HostIdentityProviderGroup
		mockState.IvrTheme = req.IvrTheme

		// Int/bool pointer fields (always set via schema defaults)
		if req.ScheduledAliasSuffixLength != nil {
			mockState.ScheduledAliasSuffixLength = *req.ScheduledAliasSuffixLength
		}
		if req.MeetingBufferBefore != nil {
			mockState.MeetingBufferBefore = *req.MeetingBufferBefore
		}
		if req.MeetingBufferAfter != nil {
			mockState.MeetingBufferAfter = *req.MeetingBufferAfter
		}
		if req.EnableDynamicVmrs != nil {
			mockState.EnableDynamicVmrs = *req.EnableDynamicVmrs
		}
		if req.EnablePersonalVmrs != nil {
			mockState.EnablePersonalVmrs = *req.EnablePersonalVmrs
		}
		if req.AllowNewUsers != nil {
			mockState.AllowNewUsers = *req.AllowNewUsers
		}
		if req.DisableProxy != nil {
			mockState.DisableProxy = *req.DisableProxy
		}
		if req.UseCustomAddInSources != nil {
			mockState.UseCustomAddInSources = *req.UseCustomAddInSources
		}
		if req.EnableAddinDebugLogs != nil {
			mockState.EnableAddinDebugLogs = *req.EnableAddinDebugLogs
		}
		if req.KerberosEnableTls != nil {
			mockState.KerberosEnableTls = *req.KerberosEnableTls
		}
		if req.KerberosAuthEveryRequest != nil {
			mockState.KerberosAuthEveryRequest = *req.KerberosAuthEveryRequest
		}
		if req.KerberosVerifyTlsUsingCustomCa != nil {
			mockState.KerberosVerifyTlsUsingCustomCa = *req.KerberosVerifyTlsUsingCustomCa
		}

		*dest = *mockState
	}).Maybe()

	client.On("DeleteJSON", mock.Anything, "configuration/v1/ms_exchange_connector/1/", mock.Anything).Return(nil)

	testInfinityMsExchangeConnector(t, client)
}

func testInfinityMsExchangeConnector(t *testing.T, client InfinityClient) {
	fullConfig := test.LoadTestFolder(t, "resource_infinity_ms_exchange_connector_full")
	minConfig := test.LoadTestFolder(t, "resource_infinity_ms_exchange_connector_min")

	fullChecks := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttrSet("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "id"),
		resource.TestCheckResourceAttrSet("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "resource_id"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "name", "tf-test-ms-exchange-connector"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "description", "tf-test MS Exchange Connector"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "meeting_buffer_before", "60"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "meeting_buffer_after", "90"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "scheduled_alias_suffix_length", "8"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "authentication_method", "OAUTH"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "auth_provider", "AZURE"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "enable_dynamic_vmrs", "true"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "enable_personal_vmrs", "true"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "allow_new_users", "false"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "disable_proxy", "true"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "use_custom_add_in_sources", "true"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "enable_addin_debug_logs", "true"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "kerberos_enable_tls", "true"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "kerberos_auth_every_request", "true"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "kerberos_verify_tls_using_custom_ca", "true"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "non_idp_participants", "disallow_all"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "addin_authentication_method", "EXCHANGE_USER_ID_TOKEN"),
	)

	minChecks := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttrSet("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "id"),
		resource.TestCheckResourceAttrSet("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "resource_id"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "name", "tf-test-ms-exchange-connector"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "description", ""),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "meeting_buffer_before", "30"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "meeting_buffer_after", "60"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "scheduled_alias_suffix_length", "6"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "authentication_method", "BASIC"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "auth_provider", "ADFS"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "enable_dynamic_vmrs", "true"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "enable_personal_vmrs", "false"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "allow_new_users", "true"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "disable_proxy", "false"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "use_custom_add_in_sources", "false"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "enable_addin_debug_logs", "false"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "kerberos_enable_tls", "true"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "kerberos_auth_every_request", "false"),
		resource.TestCheckResourceAttr("infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "kerberos_verify_tls_using_custom_ca", "false"),
	)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: getTestProtoV6ProviderFactories(client),
		Steps: []resource.TestStep{
			// Step 1: Create with full config
			{
				Config: fullConfig,
				Check:  fullChecks,
			},
			// Step 2: Update to min config
			{
				Config: minConfig,
				Check:  minChecks,
			},
			// Step 3: Destroy
			{
				Config:  minConfig,
				Destroy: true,
			},
			// Step 4: Create with min config
			{
				Config: minConfig,
				Check:  minChecks,
			},
			// Step 5: Update to full config
			{
				Config: fullConfig,
				Check:  fullChecks,
			},
		},
	})
}
