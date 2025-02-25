// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package client

import (
	"io"
	"net/http"

	"github.com/mattermost/mattermost-server/v6/model"
)

type Client interface {
	GetTeamsForUser(userId, etag string) ([]*model.Team, *model.Response, error)
	CreateChannel(channel *model.Channel) (*model.Channel, *model.Response, error)
	RemoveUserFromChannel(channelId, userId string) (*model.Response, error)
	GetChannelMembers(channelId string, page, perPage int, etag string) (model.ChannelMembers, *model.Response, error)
	AddChannelMember(channelId, userId string) (*model.ChannelMember, *model.Response, error)
	DeleteChannel(channelId string) (*model.Response, error)
	PermanentDeleteChannel(channelId string) (*model.Response, error)
	MoveChannel(channelId, teamId string, force bool) (*model.Channel, *model.Response, error)
	GetPublicChannelsForTeam(teamId string, page int, perPage int, etag string) ([]*model.Channel, *model.Response, error)
	GetDeletedChannelsForTeam(teamId string, page int, perPage int, etag string) ([]*model.Channel, *model.Response, error)
	GetPrivateChannelsForTeam(teamId string, page int, perPage int, etag string) ([]*model.Channel, *model.Response, error)
	GetChannelsForTeamForUser(teamId, userId string, includeDeleted bool, etag string) ([]*model.Channel, *model.Response, error)
	RestoreChannel(channelId string) (*model.Channel, *model.Response, error)
	PatchChannel(channelId string, patch *model.ChannelPatch) (*model.Channel, *model.Response, error)
	GetChannelByName(channelName, teamId string, etag string) (*model.Channel, *model.Response, error)
	GetChannelByNameIncludeDeleted(channelName, teamId string, etag string) (*model.Channel, *model.Response, error)
	GetChannel(channelId, etag string) (*model.Channel, *model.Response, error)
	GetTeam(teamId, etag string) (*model.Team, *model.Response, error)
	GetTeamByName(name, etag string) (*model.Team, *model.Response, error)
	GetAllTeams(etag string, page int, perPage int) ([]*model.Team, *model.Response, error)
	CreateTeam(team *model.Team) (*model.Team, *model.Response, error)
	PatchTeam(teamId string, patch *model.TeamPatch) (*model.Team, *model.Response, error)
	AddTeamMember(teamId, userId string) (*model.TeamMember, *model.Response, error)
	RemoveTeamMember(teamId, userId string) (*model.Response, error)
	SoftDeleteTeam(teamId string) (*model.Response, error)
	PermanentDeleteTeam(teamId string) (*model.Response, error)
	RestoreTeam(teamId string) (*model.Team, *model.Response, error)
	UpdateTeamPrivacy(teamId string, privacy string) (*model.Team, *model.Response, error)
	SearchTeams(search *model.TeamSearch) ([]*model.Team, *model.Response, error)
	GetPost(postId string, etag string) (*model.Post, *model.Response, error)
	CreatePost(post *model.Post) (*model.Post, *model.Response, error)
	GetPostsForChannel(channelId string, page, perPage int, etag string, collapsedThreads bool) (*model.PostList, *model.Response, error)
	DoAPIPost(url string, data string) (*http.Response, error)
	GetLdapGroups() ([]*model.Group, *model.Response, error)
	GetGroupsByChannel(channelId string, groupOpts model.GroupSearchOpts) ([]*model.GroupWithSchemeAdmin, int, *model.Response, error)
	GetGroupsByTeam(teamId string, groupOpts model.GroupSearchOpts) ([]*model.GroupWithSchemeAdmin, int, *model.Response, error)
	UploadLicenseFile(data []byte) (*model.Response, error)
	RemoveLicenseFile() (*model.Response, error)
	GetLogs(page, perPage int) ([]string, *model.Response, error)
	GetRoleByName(name string) (*model.Role, *model.Response, error)
	PatchRole(roleId string, patch *model.RolePatch) (*model.Role, *model.Response, error)
	UploadPlugin(file io.Reader) (*model.Manifest, *model.Response, error)
	UploadPluginForced(file io.Reader) (*model.Manifest, *model.Response, error)
	RemovePlugin(id string) (*model.Response, error)
	EnablePlugin(id string) (*model.Response, error)
	DisablePlugin(id string) (*model.Response, error)
	GetPlugins() (*model.PluginsResponse, *model.Response, error)
	GetUser(userId, etag string) (*model.User, *model.Response, error)
	GetUserByUsername(userName, etag string) (*model.User, *model.Response, error)
	GetUserByEmail(email, etag string) (*model.User, *model.Response, error)
	PermanentDeleteUser(userId string) (*model.Response, error)
	PermanentDeleteAllUsers() (*model.Response, error)
	CreateUser(user *model.User) (*model.User, *model.Response, error)
	VerifyUserEmailWithoutToken(userId string) (*model.User, *model.Response, error)
	UpdateUserRoles(userId, roles string) (*model.Response, error)
	InviteUsersToTeam(teamId string, userEmails []string) (*model.Response, error)
	SendPasswordResetEmail(email string) (*model.Response, error)
	UpdateUser(user *model.User) (*model.User, *model.Response, error)
	UpdateUserMfa(userId, code string, activate bool) (*model.Response, error)
	UpdateUserPassword(userId, currentPassword, newPassword string) (*model.Response, error)
	UpdateUserHashedPassword(userId, newHashedPassword string) (*model.Response, error)
	CreateUserAccessToken(userId, description string) (*model.UserAccessToken, *model.Response, error)
	RevokeUserAccessToken(tokenId string) (*model.Response, error)
	GetUserAccessTokensForUser(userId string, page, perPage int) ([]*model.UserAccessToken, *model.Response, error)
	ConvertUserToBot(userId string) (*model.Bot, *model.Response, error)
	ConvertBotToUser(userId string, userPatch *model.UserPatch, setSystemAdmin bool) (*model.User, *model.Response, error)
	PromoteGuestToUser(userId string) (*model.Response, error)
	DemoteUserToGuest(guestId string) (*model.Response, error)
	CreateCommand(cmd *model.Command) (*model.Command, *model.Response, error)
	ListCommands(teamId string, customOnly bool) ([]*model.Command, *model.Response, error)
	GetCommandById(cmdId string) (*model.Command, *model.Response, error)
	UpdateCommand(cmd *model.Command) (*model.Command, *model.Response, error)
	MoveCommand(teamId string, commandId string) (*model.Response, error)
	DeleteCommand(commandId string) (*model.Response, error)
	GetConfig() (*model.Config, *model.Response, error)
	UpdateConfig(*model.Config) (*model.Config, *model.Response, error)
	PatchConfig(*model.Config) (*model.Config, *model.Response, error)
	ReloadConfig() (*model.Response, error)
	MigrateConfig(from, to string) (*model.Response, error)
	SyncLdap(includeRemovedMembers bool) (*model.Response, error)
	MigrateIdLdap(toAttribute string) (*model.Response, error)
	GetUsers(page, perPage int, etag string) ([]*model.User, *model.Response, error)
	GetUsersByIds(userIds []string) ([]*model.User, *model.Response, error)
	GetUsersInTeam(teamId string, page, perPage int, etag string) ([]*model.User, *model.Response, error)
	UpdateUserActive(userId string, activate bool) (*model.Response, error)
	UpdateTeam(team *model.Team) (*model.Team, *model.Response, error)
	UpdateChannelPrivacy(channelId string, privacy model.ChannelType) (*model.Channel, *model.Response, error)
	CreateBot(bot *model.Bot) (*model.Bot, *model.Response, error)
	PatchBot(userId string, patch *model.BotPatch) (*model.Bot, *model.Response, error)
	GetBots(page, perPage int, etag string) ([]*model.Bot, *model.Response, error)
	GetBotsIncludeDeleted(page, perPage int, etag string) ([]*model.Bot, *model.Response, error)
	GetBotsOrphaned(page, perPage int, etag string) ([]*model.Bot, *model.Response, error)
	DisableBot(botUserId string) (*model.Bot, *model.Response, error)
	EnableBot(botUserId string) (*model.Bot, *model.Response, error)
	AssignBot(botUserId, newOwnerId string) (*model.Bot, *model.Response, error)
	SetServerBusy(secs int) (*model.Response, error)
	ClearServerBusy() (*model.Response, error)
	GetServerBusy() (*model.ServerBusyState, *model.Response, error)
	CheckIntegrity() ([]model.IntegrityCheckResult, *model.Response, error)
	InstallPluginFromURL(string, bool) (*model.Manifest, *model.Response, error)
	InstallMarketplacePlugin(*model.InstallMarketplacePluginRequest) (*model.Manifest, *model.Response, error)
	GetMarketplacePlugins(*model.MarketplacePluginFilter) ([]*model.MarketplacePlugin, *model.Response, error)
	MigrateAuthToLdap(fromAuthService string, matchField string, force bool) (*model.Response, error)
	MigrateAuthToSaml(fromAuthService string, usersMap map[string]string, auto bool) (*model.Response, error)
	GetPing() (string, *model.Response, error)
	GetPingWithFullServerStatus() (map[string]string, *model.Response, error)
	CreateUpload(us *model.UploadSession) (*model.UploadSession, *model.Response, error)
	GetUpload(uploadId string) (*model.UploadSession, *model.Response, error)
	GetUploadsForUser(userId string) ([]*model.UploadSession, *model.Response, error)
	UploadData(uploadId string, data io.Reader) (*model.FileInfo, *model.Response, error)
	ListImports() ([]string, *model.Response, error)
	GetJob(id string) (*model.Job, *model.Response, error)
	GetJobs(page int, perPage int) ([]*model.Job, *model.Response, error)
	GetJobsByType(jobType string, page int, perPage int) ([]*model.Job, *model.Response, error)
	CreateJob(job *model.Job) (*model.Job, *model.Response, error)
	CancelJob(jobId string) (*model.Response, error)
	CreateIncomingWebhook(hook *model.IncomingWebhook) (*model.IncomingWebhook, *model.Response, error)
	UpdateIncomingWebhook(hook *model.IncomingWebhook) (*model.IncomingWebhook, *model.Response, error)
	GetIncomingWebhooks(page int, perPage int, etag string) ([]*model.IncomingWebhook, *model.Response, error)
	GetIncomingWebhooksForTeam(teamId string, page int, perPage int, etag string) ([]*model.IncomingWebhook, *model.Response, error)
	GetIncomingWebhook(hookID string, etag string) (*model.IncomingWebhook, *model.Response, error)
	DeleteIncomingWebhook(hookID string) (*model.Response, error)
	CreateOutgoingWebhook(hook *model.OutgoingWebhook) (*model.OutgoingWebhook, *model.Response, error)
	UpdateOutgoingWebhook(hook *model.OutgoingWebhook) (*model.OutgoingWebhook, *model.Response, error)
	GetOutgoingWebhooks(page int, perPage int, etag string) ([]*model.OutgoingWebhook, *model.Response, error)
	GetOutgoingWebhook(hookId string) (*model.OutgoingWebhook, *model.Response, error)
	GetOutgoingWebhooksForChannel(channelId string, page int, perPage int, etag string) ([]*model.OutgoingWebhook, *model.Response, error)
	GetOutgoingWebhooksForTeam(teamId string, page int, perPage int, etag string) ([]*model.OutgoingWebhook, *model.Response, error)
	RegenOutgoingHookToken(hookId string) (*model.OutgoingWebhook, *model.Response, error)
	DeleteOutgoingWebhook(hookId string) (*model.Response, error)
	ListExports() ([]string, *model.Response, error)
	DeleteExport(name string) (*model.Response, error)
	DownloadExport(name string, wr io.Writer, offset int64) (int64, *model.Response, error)
	ResetSamlAuthDataToEmail(includeDeleted bool, dryRun bool, userIDs []string) (int64, *model.Response, error)
}
