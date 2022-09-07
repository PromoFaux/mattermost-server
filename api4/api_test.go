// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package api4

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mattermost/mattermost-server/v6/model"
)

func TestAPIScopeAssignemt(t *testing.T) {
	th := SetupEnterprise(t)
	defer th.TearDown()

	t.Run("by name", func(t *testing.T) {
		expected := map[string]model.APIScopes{
			"addChannelMember":                     {model.ScopeChannelsJoin},
			"addTeamMember":                        {model.ScopeTeamsJoin},
			"addTeamMembers":                       {model.ScopeTeamsJoin},
			"autocompleteChannelsForTeam":          {model.ScopeChannelsSearch},
			"autocompleteChannelsForTeamForSearch": {model.ScopeChannelsSearch},
			"autocompleteEmojis":                   {model.ScopeEmojisSearch},
			"autocompleteUsers":                    {model.ScopeUsersSearch},
			"connectWebSocket":                     model.ScopeUnrestrictedAPI,
			"createCategoryForTeamForUser":         {model.ScopeTeamsUpdate, model.ScopeUsersUpdate},
			"createChannel":                        {model.ScopeChannelsCreate},
			"createDirectChannel":                  {model.ScopePostsCreateDM},
			"createEmoji":                          {model.ScopeEmojisCreate},
			"createEphemeralPost":                  {model.ScopePostsCreateEphemeral},
			"createPost":                           model.ScopeCheckedByImplementation,
			"createTeam":                           {model.ScopeTeamsCreate},
			"createUpload":                         {model.ScopeFilesCreate},
			"createUser":                           {model.ScopeUsersCreate},
			"deleteCategoryForTeamForUser":         {model.ScopeTeamsUpdate, model.ScopeUsersUpdate},
			"deleteChannel":                        {model.ScopeChannelsDelete},
			"deleteEmoji":                          {model.ScopeEmojisDelete},
			"deletePost":                           {model.ScopePostsDelete},
			"deleteReaction":                       {model.ScopePostsUpdate},
			"deleteTeam":                           {model.ScopeTeamsDelete},
			"deleteUser":                           {model.ScopeUsersDelete},
			"demoteUserToGuest":                    {model.ScopeUsersUpdate},
			"executeCommand":                       {model.ScopeCommandsExecute},
			"getAllChannels":                       {model.ScopeChannelsRead},
			"getAllTeams":                          {model.ScopeTeamsRead},
			"getBrandImage":                        model.ScopeUnrestrictedAPI,
			"getBulkReactions":                     {model.ScopePostsRead},
			"getCategoriesForTeamForUser":          {model.ScopeTeamsRead, model.ScopeUsersRead},
			"getCategoryForTeamForUser":            {model.ScopeTeamsRead, model.ScopeUsersRead},
			"getCategoryOrderForTeamForUser":       {model.ScopeTeamsRead, model.ScopeUsersRead},
			"getChannel":                           {model.ScopeChannelsRead},
			"getChannelByName":                     {model.ScopeChannelsRead},
			"getChannelByNameForTeamName":          {model.ScopeChannelsRead, model.ScopeTeamsRead},
			"getChannelMember":                     {model.ScopeChannelsRead, model.ScopeUsersRead},
			"getChannelMembers":                    {model.ScopeChannelsRead, model.ScopeUsersRead},
			"getChannelMembersByIds":               {model.ScopeChannelsRead, model.ScopeUsersRead},
			"getChannelMembersForTeamForUser":      {model.ScopeChannelsRead, model.ScopeTeamsRead, model.ScopeUsersRead},
			"getChannelMembersForUser":             {model.ScopeChannelsRead, model.ScopeUsersRead},
			"getChannelMembersTimezones":           {model.ScopeChannelsRead, model.ScopeUsersRead},
			"getChannelStats":                      {model.ScopeChannelsRead},
			"getChannelUnread":                     {model.ScopeChannelsRead},
			"getChannelsForTeamForUser":            {model.ScopeChannelsRead, model.ScopeTeamsRead, model.ScopeUsersRead},
			"getChannelsForUser":                   {model.ScopeChannelsRead, model.ScopeUsersRead},
			"getDefaultProfileImage":               {model.ScopeUsersRead},
			"getDeletedChannelsForTeam":            {model.ScopeChannelsRead, model.ScopeTeamsRead},
			"getEmoji":                             {model.ScopeEmojisRead},
			"getEmojiByName":                       {model.ScopeEmojisRead},
			"getEmojiImage":                        {model.ScopeEmojisRead},
			"getEmojiList":                         {model.ScopeEmojisRead},
			"getFile":                              {model.ScopeFilesRead},
			"getFileInfo":                          {model.ScopeFilesRead},
			"getFileInfosForPost":                  {model.ScopeFilesRead, model.ScopePostsRead},
			"getFileLink":                          {model.ScopeFilesRead},
			"getFilePreview":                       {model.ScopeFilesRead},
			"getFileThumbnail":                     {model.ScopeFilesRead},
			"getFlaggedPostsForUser":               {model.ScopePostsRead, model.ScopeUsersRead},
			"getPinnedPosts":                       {model.ScopeChannelsRead, model.ScopePostsRead},
			"getPost":                              {model.ScopePostsRead},
			"getPostThread":                        {model.ScopePostsRead},
			"getPostsByIds":                        {model.ScopePostsRead},
			"getPostsForChannel":                   {model.ScopeChannelsRead, model.ScopePostsRead},
			"getPostsForChannelAroundLastUnread":   {model.ScopeChannelsRead, model.ScopePostsRead},
			"getPrivateChannelsForTeam":            {model.ScopeChannelsRead, model.ScopeTeamsRead},
			"getProfileImage":                      {model.ScopeUsersRead},
			"getPublicChannelsByIdsForTeam":        {model.ScopeChannelsRead, model.ScopeTeamsRead},
			"getPublicChannelsForTeam":             {model.ScopeChannelsRead, model.ScopeTeamsRead},
			"getPublicFile":                        {model.ScopeFilesRead},
			"getReactions":                         {model.ScopePostsRead},
			"getRecentSearches":                    {model.ScopeUsersRead},
			"getSupportedTimezones":                model.ScopeUnrestrictedAPI,
			"getTeam":                              {model.ScopeTeamsRead},
			"getTeamByName":                        {model.ScopeTeamsRead},
			"getTeamIcon":                          {model.ScopeTeamsRead},
			"getTeamMember":                        {model.ScopeTeamsRead, model.ScopeUsersRead},
			"getTeamMembers":                       {model.ScopeTeamsRead, model.ScopeUsersRead},
			"getTeamMembersByIds":                  {model.ScopeTeamsRead, model.ScopeUsersRead},
			"getTeamMembersForUser":                {model.ScopeTeamsRead, model.ScopeUsersRead},
			"getTeamStats":                         {model.ScopeTeamsRead},
			"getTeamUnread":                        {model.ScopeTeamsRead},
			"getTeamsForUser":                      {model.ScopeTeamsRead, model.ScopeUsersRead},
			"getTeamsUnreadForUser":                {model.ScopeTeamsRead, model.ScopeUsersRead},
			"getUpload":                            {model.ScopeFilesRead},
			"getUploadsForUser":                    {model.ScopeFilesRead, model.ScopeUsersRead},
			"getUser":                              {model.ScopeUsersRead},
			"getUserByEmail":                       {model.ScopeUsersRead},
			"getUserByUsername":                    {model.ScopeUsersRead},
			"getUserStatus":                        {model.ScopeUsersRead},
			"getUsers":                             {model.ScopeUsersRead},
			"getUsersByIds":                        {model.ScopeUsersRead},
			"getUsersByNames":                      {model.ScopeUsersRead},
			"login":                                model.ScopeUnrestrictedAPI,
			"loginCWS":                             model.ScopeUnrestrictedAPI,
			"logout":                               model.ScopeUnrestrictedAPI,
			"patchChannel":                         {model.ScopeChannelsUpdate},
			"patchPost":                            {model.ScopePostsUpdate},
			"patchTeam":                            {model.ScopeTeamsUpdate},
			"patchUser":                            {model.ScopeUsersUpdate},
			"pinPost":                              {model.ScopeChannelsUpdate, model.ScopePostsRead},
			"publishUserTyping":                    {model.ScopeUsersUpdate},
			"removeChannelMember":                  {model.ScopeChannelsJoin, model.ScopeUsersUpdate},
			"removeTeamIcon":                       {model.ScopeTeamsUpdate},
			"removeTeamMember":                     {model.ScopeTeamsJoin, model.ScopeUsersUpdate},
			"removeUserCustomStatus":               {model.ScopeUsersUpdate},
			"removeUserRecentCustomStatus":         {model.ScopeUsersUpdate},
			"saveReaction":                         {model.ScopePostsUpdate},
			"searchAllChannels":                    {model.ScopeChannelsSearch},
			"searchArchivedChannelsForTeam":        {model.ScopeChannelsSearch, model.ScopeTeamsRead},
			"searchChannelsForTeam":                {model.ScopeChannelsSearch, model.ScopeTeamsRead},
			"searchEmojis":                         {model.ScopeEmojisSearch},
			"searchFilesForUser":                   {model.ScopeFilesSearch, model.ScopeUsersRead},
			"searchFilesInTeam":                    {model.ScopeFilesSearch, model.ScopeTeamsRead},
			"searchPostsInAllTeams":                {model.ScopePostsSearch},
			"searchPostsInTeam":                    {model.ScopePostsSearch, model.ScopeTeamsRead},
			"searchTeams":                          {model.ScopeTeamsRead, model.ScopeTeamsSearch},
			"searchUsers":                          {model.ScopeUsersRead, model.ScopeUsersSearch},
			"setDefaultProfileImage":               {model.ScopeUsersUpdate},
			"setPostReminder":                      {model.ScopePostsUpdate},
			"setPostUnread":                        {model.ScopePostsUpdate},
			"setProfileImage":                      {model.ScopeUsersUpdate},
			"setTeamIcon":                          {model.ScopeTeamsUpdate},
			"softDeleteTeamsExcept":                {model.ScopeTeamsDelete},
			"teamExists":                           {model.ScopeTeamsRead},
			"unpinPost":                            {model.ScopeChannelsUpdate, model.ScopePostsRead},
			"updateCategoriesForTeamForUser":       {model.ScopeTeamsUpdate, model.ScopeUsersUpdate},
			"updateCategoryForTeamForUser":         {model.ScopeTeamsUpdate, model.ScopeUsersUpdate},
			"updateCategoryOrderForTeamForUser":    {model.ScopeTeamsUpdate, model.ScopeUsersUpdate},
			"updateChannel":                        {model.ScopeChannelsUpdate},
			"updateChannelPrivacy":                 {model.ScopeChannelsUpdate},
			"updatePost":                           {model.ScopePostsUpdate},
			"updateTeam":                           {model.ScopeTeamsUpdate},
			"updateTeamPrivacy":                    {model.ScopeTeamsUpdate},
			"updateUser":                           {model.ScopeUsersUpdate},
			"updateUserActive":                     {model.ScopeUsersUpdate},
			"updateUserCustomStatus":               {model.ScopeUsersUpdate},
			"updateUserStatus":                     {model.ScopeUsersUpdate},
			"uploadFileStream":                     {model.ScopeFilesCreate},
			"getKnownUsers":                        {model.ScopeUsersRead},
			"getUserStatusesByIds":                 {model.ScopeUsersRead},
			"moveChannel":                          {model.ScopeChannelsUpdate},
			"promoteGuestToUser":                   {model.ScopeUsersUpdate},
			"restoreTeam":                          {model.ScopeTeamsUpdate},
			"restoreChannel":                       {model.ScopeChannelsUpdate},
			"uploadData":                           {model.ScopeFilesUpdate},
			"viewChannel":                          {model.ScopeChannelsRead, model.ScopeUsersRead},
		}

		expectedKeys := []string{}
		for k := range expected {
			expectedKeys = append(expectedKeys, k)
		}
		sort.Strings(expectedKeys)
		keys := []string{}
		for k := range th.API.knownAPIsByName {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		require.EqualValues(t, expectedKeys, keys)
		require.EqualValues(t, expected, th.API.knownAPIsByName)
	})

	t.Run("by scope", func(t *testing.T) {
		expected := map[model.Scope][]string{
			"*:*": {
				"login",
				"loginCWS",
				"logout",
				"createPost",
				"getSupportedTimezones",
				"getBrandImage",
				"connectWebSocket",
			},
			"channels:create": {
				"createChannel",
			},
			"channels:delete": {
				"deleteChannel",
			},
			"channels:join": {
				"addChannelMember",
				"removeChannelMember",
			},
			"channels:read": {
				"getChannelMembersForUser",
				"getAllChannels",
				"viewChannel",
				"getPublicChannelsForTeam",
				"getDeletedChannelsForTeam",
				"getPrivateChannelsForTeam",
				"getPublicChannelsByIdsForTeam",
				"getChannelsForTeamForUser",
				"getChannelsForUser",
				"getChannel",
				"getChannelStats",
				"getPinnedPosts",
				"getChannelMembersTimezones",
				"getChannelUnread",
				"getChannelByName",
				"getChannelByNameForTeamName",
				"getChannelMembers",
				"getChannelMembersByIds",
				"getChannelMembersForTeamForUser",
				"getChannelMember",
				"getPostsForChannel",
				"getPostsForChannelAroundLastUnread",
			},
			"channels:search": {
				"searchAllChannels",
				"searchChannelsForTeam",
				"searchArchivedChannelsForTeam",
				"autocompleteChannelsForTeam",
				"autocompleteChannelsForTeamForSearch",
			},
			"channels:update": {
				"updateChannel",
				"patchChannel",
				"updateChannelPrivacy",
				"restoreChannel",
				"moveChannel",
				"pinPost",
				"unpinPost",
			},
			"commands:execute": {
				"executeCommand",
			},
			"emojis:create": {
				"createEmoji",
			},
			"emojis:delete": {
				"deleteEmoji",
			},
			"emojis:read": {
				"getEmojiList",
				"getEmoji",
				"getEmojiByName",
				"getEmojiImage",
			},
			"emojis:search": {
				"searchEmojis",
				"autocompleteEmojis",
			},
			"files:create": {
				"uploadFileStream",
				"createUpload",
			},
			"files:read": {
				"getUploadsForUser",
				"getFileInfosForPost",
				"getFile",
				"getFileThumbnail",
				"getFileLink",
				"getFilePreview",
				"getFileInfo",
				"getPublicFile",
				"getUpload",
			},
			"files:search": {
				"searchFilesForUser",
				"searchFilesInTeam",
			},
			"files:update": {
				"uploadData",
			},
			"internal_api": {
				"addChannelsToPolicy",
				"addLdapPrivateCertificate",
				"addLdapPublicCertificate",
				"addLicense",
				"addSamlIdpCertificate",
				"addSamlPrivateCertificate",
				"addSamlPublicCertificate",
				"addTeamsToPolicy",
				"addUserToTeamFromInvite",
				"appendAncillaryPermissions",
				"assignBot",
				"attachDeviceId",
				"cancelJob",
				"changeSubscription",
				"channelMemberCountsByGroup",
				"channelMembersMinusGroupMembers",
				"clearServerBusy",
				"completeOnboarding",
				"configReload",
				"confirmCustomerPayment",
				"convertBotToUser",
				"convertUserToBot",
				"createBot",
				"createCommand",
				"createComplianceReport",
				"createCustomerPayment",
				"createGroupChannel",
				"createIncomingHook",
				"createJob",
				"createOAuthApp",
				"createOutgoingHook",
				"createPolicy",
				"createScheme",
				"createTermsOfService",
				"createUserAccessToken",
				"databaseRecycle",
				"deleteBrandImage",
				"deleteCommand",
				"deleteExport",
				"deleteIncomingHook",
				"deleteOAuthApp",
				"deleteOutgoingHook",
				"deletePolicy",
				"deletePreferences",
				"deleteScheme",
				"disableBot",
				"disablePlugin",
				"disableUserAccessToken",
				"doPostAction",
				"downloadComplianceReport",
				"downloadExport",
				"downloadJob",
				"enableBot",
				"enablePlugin",
				"enableUserAccessToken",
				"followThreadByUser",
				"func1",
				"generateMfaSecret",
				"generateSupportPacket",
				"getAllRoles",
				"getAnalytics",
				"getAppliedSchemaMigrations",
				"getAudits",
				"getAuthorizedOAuthApps",
				"getBot",
				"getBots",
				"getChannelModerations",
				"getChannelPoliciesForUser",
				"getChannelsForPolicy",
				"getChannelsForScheme",
				"getClientConfig",
				"getClientLicense",
				"getCloudCustomer",
				"getCloudLimits",
				"getCloudProducts",
				"getClusterStatus",
				"getCommand",
				"getComplianceReport",
				"getComplianceReports",
				"getConfig",
				"getEnvironmentConfig",
				"getFilteredUsersStats",
				"getFirstAdminVisitMarketplaceStatus",
				"getGlobalPolicy",
				"getImage",
				"getIncomingHook",
				"getIncomingHooks",
				"getIntegrationsUsage",
				"getInviteInfo",
				"getInvoicesForSubscription",
				"getJob",
				"getJobs",
				"getJobsByType",
				"getLatestTermsOfService",
				"getLatestVersion",
				"getLdapGroups",
				"getLogs",
				"getMarketplacePlugins",
				"getOAuthApp",
				"getOAuthAppInfo",
				"getOAuthApps",
				"getOnboarding",
				"getOpenGraphMetadata",
				"getOutgoingHook",
				"getOutgoingHooks",
				"getPlugins",
				"getPluginStatuses",
				"getPolicies",
				"getPoliciesCount",
				"getPolicy",
				"getPostsUsage",
				"getPreferenceByCategoryAndName",
				"getPreferences",
				"getPreferencesByCategory",
				"getPrevTrialLicense",
				"getProductNotices",
				"getRedirectLocation",
				"getRemoteClusterInfo",
				"getRole",
				"getRoleByName",
				"getRolesByNames",
				"getSamlCertificateStatus",
				"getSamlMetadata",
				"getSamlMetadataFromIdp",
				"getScheme",
				"getSchemes",
				"getServerBusyExpires",
				"getSessions",
				"getSharedChannels",
				"getStorageUsage",
				"getSubscription",
				"getSubscriptionInvoicePDF",
				"getSystemPing",
				"getTeamPoliciesForUser",
				"getTeamsForPolicy",
				"getTeamsForScheme",
				"getTeamsUsage",
				"getThreadForUser",
				"getThreadsForUser",
				"getTotalUsersStats",
				"getUserAccessToken",
				"getUserAccessTokens",
				"getUserAccessTokensForUser",
				"getUserAudits",
				"getUsersByGroupChannelIds",
				"getUsersWithInvalidEmails",
				"getUserTermsOfService",
				"getWarnMetricsStatus",
				"getWebappPlugins",
				"handleCWSWebhook",
				"handleNotifyAdmin",
				"handleTriggerNotifyAdminPosts",
				"importTeam",
				"installMarketplacePlugin",
				"installPluginFromURL",
				"invalidateAllEmailInvites",
				"invalidateCaches",
				"inviteGuestsToChannels",
				"inviteUsersToTeam",
				"linkLdapGroup",
				"listAutocompleteCommands",
				"listCommandAutocompleteSuggestions",
				"listCommands",
				"listExports",
				"listImports",
				"migrateAuthToLDAP",
				"migrateAuthToSaml",
				"migrateIdLdap",
				"moveCommand",
				"openDialog",
				"patchBot",
				"patchChannelModerations",
				"patchConfig",
				"patchPolicy",
				"patchRole",
				"patchScheme",
				"postLog",
				"purgeBleveIndexes",
				"purgeElasticsearchIndexes",
				"pushNotificationAck",
				"regenCommandToken",
				"regenerateOAuthAppSecret",
				"regenerateTeamInviteId",
				"regenOutgoingHookToken",
				"remoteClusterAcceptMessage",
				"remoteClusterConfirmInvite",
				"remoteClusterPing",
				"remoteSetProfileImage",
				"removeChannelsFromPolicy",
				"removeLdapPrivateCertificate",
				"removeLdapPublicCertificate",
				"removeLicense",
				"removePlugin",
				"removeSamlIdpCertificate",
				"removeSamlPrivateCertificate",
				"removeSamlPublicCertificate",
				"removeTeamsFromPolicy",
				"requestCloudTrial",
				"requestRenewalLink",
				"requestTrialLicense",
				"requestTrialLicenseAndAckWarnMetric",
				"resetAuthDataToEmail",
				"resetPassword",
				"restart",
				"revokeAllSessionsAllUsers",
				"revokeAllSessionsForUser",
				"revokeSession",
				"revokeUserAccessToken",
				"saveUserTermsOfService",
				"searchChannelsInPolicy",
				"searchGroupChannels",
				"searchTeamsInPolicy",
				"searchUserAccessTokens",
				"sendPasswordReset",
				"sendVerificationEmail",
				"sendWarnMetricAckEmail",
				"setFirstAdminVisitMarketplaceStatus",
				"setServerBusy",
				"setUnreadThreadByPostId",
				"submitDialog",
				"switchAccountType",
				"syncLdap",
				"teamMembersMinusGroupMembers",
				"testElasticsearch",
				"testEmail",
				"testLdap",
				"testS3",
				"testSiteURL",
				"unfollowThreadByUser",
				"unlinkLdapGroup",
				"updateChannelMemberNotifyProps",
				"updateChannelMemberRoles",
				"updateChannelMemberSchemeRoles",
				"updateChannelScheme",
				"updateCloudCustomer",
				"updateCloudCustomerAddress",
				"updateCommand",
				"updateConfig",
				"updateIncomingHook",
				"updateOAuthApp",
				"updateOutgoingHook",
				"updatePassword",
				"updatePreferences",
				"updateReadStateAllThreadsByUser",
				"updateReadStateThreadByUser",
				"updateTeamMemberRoles",
				"updateTeamMemberSchemeRoles",
				"updateTeamScheme",
				"updateUserAuth",
				"updateUserMfa",
				"updateUserRoles",
				"updateViewedProductNotices",
				"upgradeToEnterprise",
				"upgradeToEnterpriseStatus",
				"uploadBrandImage",
				"uploadPlugin",
				"uploadRemoteData",
				"validateBusinessEmail",
				"validateWorkspaceBusinessEmail",
				"verifyUserEmail",
				"verifyUserEmailWithoutToken",
			},
			"posts:create/dm": {
				"createDirectChannel",
			},
			"posts:create/ephemeral": {
				"createEphemeralPost",
			},
			"posts:delete": {
				"deletePost",
			},
			"posts:read": {
				"getPinnedPosts",
				"getPost",
				"getPostsByIds",
				"getPostThread",
				"getFileInfosForPost",
				"getPostsForChannel",
				"getFlaggedPostsForUser",
				"getPostsForChannelAroundLastUnread",
				"pinPost",
				"unpinPost",
				"getReactions",
				"getBulkReactions",
			},
			"posts:search": {
				"searchPostsInTeam",
				"searchPostsInAllTeams",
			},
			"posts:update": {
				"updatePost",
				"patchPost",
				"setPostUnread",
				"setPostReminder",
				"saveReaction",
				"deleteReaction",
			},
			"teams:create": {
				"createTeam",
			}, "teams:delete": {
				"deleteTeam",
				"softDeleteTeamsExcept",
			},
			"teams:join": {
				"addTeamMember",
				"addTeamMembers",
				"removeTeamMember",
			}, "teams:read": {
				"getAllTeams",
				"searchTeams",
				"getTeamsForUser",
				"getTeamsUnreadForUser",
				"getTeam",
				"getTeamStats",
				"getTeamIcon",
				"getTeamMembers",
				"getTeamMembersByIds",
				"getTeamMembersForUser",
				"getTeamUnread",
				"getTeamByName",
				"getTeamMember",
				"teamExists",
				"getPublicChannelsForTeam",
				"getDeletedChannelsForTeam",
				"getPrivateChannelsForTeam",
				"getPublicChannelsByIdsForTeam",
				"searchChannelsForTeam",
				"searchArchivedChannelsForTeam",
				"getChannelsForTeamForUser",
				"getCategoriesForTeamForUser",
				"getCategoryOrderForTeamForUser",
				"getCategoryForTeamForUser",
				"getChannelByNameForTeamName",
				"getChannelMembersForTeamForUser",
				"searchPostsInTeam",
				"searchFilesInTeam",
			}, "teams:search": {
				"searchTeams",
			},
			"teams:update": {
				"updateTeam",
				"patchTeam",
				"restoreTeam",
				"updateTeamPrivacy",
				"setTeamIcon",
				"removeTeamIcon",
				"createCategoryForTeamForUser",
				"updateCategoriesForTeamForUser",
				"updateCategoryOrderForTeamForUser",
				"updateCategoryForTeamForUser",
				"deleteCategoryForTeamForUser",
			},
			"users:create": {
				"createUser",
			}, "users:delete": {
				"deleteUser",
			},
			"users:read": {
				"getUsers",
				"getUsersByIds",
				"getUsersByNames",
				"getKnownUsers",
				"searchUsers",
				"getUser",
				"getDefaultProfileImage",
				"getProfileImage",
				"getUserByUsername",
				"getUserByEmail",
				"getUploadsForUser",
				"getChannelMembersForUser",
				"getRecentSearches",
				"getTeamsForUser",
				"getTeamsUnreadForUser",
				"getTeamMembers",
				"getTeamMembersByIds",
				"getTeamMembersForUser",
				"getTeamMember",
				"viewChannel",
				"getChannelsForTeamForUser",
				"getChannelsForUser",
				"getCategoriesForTeamForUser",
				"getCategoryOrderForTeamForUser",
				"getCategoryForTeamForUser",
				"getChannelMembersTimezones",
				"getChannelMembers",
				"getChannelMembersByIds",
				"getChannelMembersForTeamForUser",
				"getChannelMember",
				"getFlaggedPostsForUser",
				"searchFilesForUser",
				"getUserStatus",
				"getUserStatusesByIds",
			}, "users:search": {
				"searchUsers",
				"autocompleteUsers",
			},
			"users:update": {
				"setProfileImage",
				"setDefaultProfileImage",
				"updateUser",
				"patchUser",
				"updateUserActive",
				"promoteGuestToUser",
				"demoteUserToGuest",
				"publishUserTyping",
				"removeTeamMember",
				"createCategoryForTeamForUser",
				"updateCategoriesForTeamForUser",
				"updateCategoryOrderForTeamForUser",
				"updateCategoryForTeamForUser",
				"deleteCategoryForTeamForUser",
				"removeChannelMember",
				"updateUserStatus",
				"updateUserCustomStatus",
				"removeUserCustomStatus",
				"removeUserRecentCustomStatus"},
		}

		normalize := func(v map[model.Scope][]string) map[model.Scope][]string {
			for k := range v {
				sort.Slice(v[k], func(i, j int) bool { return v[k][i] < v[k][j] })
			}
			return v
		}

		require.EqualValues(t, normalize(expected), normalize(th.API.knownAPIsByScope))
	})
}
