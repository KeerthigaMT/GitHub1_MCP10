package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow string `json:"workflow"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Inputs map[string]interface{} `json:"inputs"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Ref string `json:"ref"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Received_events_url string `json:"received_events_url"`
	Avatar_url string `json:"avatar_url"`
	Events_url string `json:"events_url"`
	Organizations_url string `json:"organizations_url"`
	Gravatar_id string `json:"gravatar_id"`
	Url string `json:"url"`
	Following_url string `json:"following_url"`
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Login string `json:"login"`
	Starred_url string `json:"starred_url"`
	Starred_at string `json:"starred_at,omitempty"`
	Email string `json:"email,omitempty"`
	Node_id string `json:"node_id"`
	Repos_url string `json:"repos_url"`
	Followers_url string `json:"followers_url"`
	Name string `json:"name,omitempty"`
	TypeField string `json:"type"`
	Gists_url string `json:"gists_url"`
	Site_admin bool `json:"site_admin"`
	Subscriptions_url string `json:"subscriptions_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Requester interface{} `json:"requester,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment map[string]interface{} `json:"comment"` // The [comment](https://docs.github.com/rest/reference/issues#comments) itself.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) the comment belongs to.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Limit int `json:"limit"`
	Remaining int `json:"remaining"`
	Reset int `json:"reset"`
	Used int `json:"used"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Event string `json:"event,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Comments []GeneratedType `json:"comments,omitempty"`
	Commit_id string `json:"commit_id,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Permission string `json:"permission"`
	User GeneratedType `json:"user"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
}

// Package represents the Package schema from the OpenAPI specification
type Package struct {
	Owner GeneratedType `json:"owner,omitempty"` // A GitHub user.
	Id int `json:"id"` // Unique identifier of the package.
	Package_type string `json:"package_type"`
	Created_at string `json:"created_at"`
	Repository GeneratedType `json:"repository,omitempty"` // Minimal Repository
	Updated_at string `json:"updated_at"`
	Visibility string `json:"visibility"`
	Html_url string `json:"html_url"`
	Url string `json:"url"`
	Version_count int `json:"version_count"` // The number of versions of the package.
	Name string `json:"name"` // The name of the package.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Event string `json:"event"`
	Id int `json:"id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Commit_id string `json:"commit_id"`
	Node_id string `json:"node_id"`
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Created_at string `json:"created_at"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Url string `json:"url"`
	Commit_url string `json:"commit_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes,omitempty"` // The changes to the project if the action was `edited`.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project map[string]interface{} `json:"project"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The generated name of the release
	Body string `json:"body"` // The generated body describing the contents of the release supporting markdown formatting
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	State string `json:"state"`
	Description string `json:"description"`
	Domains []string `json:"domains"` // Array of the domain set and its alternate name (if it is configured)
	Expires_at string `json:"expires_at,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Reason string `json:"reason,omitempty"`
	Number int `json:"number"`
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Title string `json:"title"`
}

// Thread represents the Thread schema from the OpenAPI specification
type Thread struct {
	Last_read_at string `json:"last_read_at"`
	Reason string `json:"reason"`
	Unread bool `json:"unread"`
	Updated_at string `json:"updated_at"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Id string `json:"id"`
	Subject map[string]interface{} `json:"subject"`
	Url string `json:"url"`
	Subscription_url string `json:"subscription_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Body_html string `json:"body_html,omitempty"`
	Line int `json:"line,omitempty"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	User GeneratedType `json:"user"` // A GitHub user.
	Start_line int `json:"start_line,omitempty"` // The first line of the range for a multi-line comment.
	Links map[string]interface{} `json:"_links"`
	In_reply_to_id int `json:"in_reply_to_id,omitempty"`
	Original_line int `json:"original_line,omitempty"` // The original line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Original_position int `json:"original_position"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Diff_hunk string `json:"diff_hunk"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Pull_request_review_id int `json:"pull_request_review_id"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body string `json:"body"`
	Id int `json:"id"`
	Original_start_line int `json:"original_start_line,omitempty"` // The original first line of the range for a multi-line comment.
	Start_side string `json:"start_side,omitempty"` // The side of the first line of the range for a multi-line comment.
	Original_commit_id string `json:"original_commit_id"`
	Path string `json:"path"`
	Pull_request_url string `json:"pull_request_url"`
	Side string `json:"side,omitempty"` // The side of the first line of the range for a multi-line comment.
	Node_id string `json:"node_id"`
	Position int `json:"position"`
	Body_text string `json:"body_text,omitempty"`
	Commit_id string `json:"commit_id"`
	Updated_at string `json:"updated_at"`
	Html_url string `json:"html_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Url string `json:"url"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Commit_url string `json:"commit_url"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Teams []Team `json:"teams"`
	Users []GeneratedType `json:"users"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Week int `json:"week"`
	Days []int `json:"days"`
	Total int `json:"total"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Member map[string]interface{} `json:"member"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// Feed represents the Feed schema from the OpenAPI specification
type Feed struct {
	Current_user_public_url string `json:"current_user_public_url,omitempty"`
	User_url string `json:"user_url"`
	Repository_discussions_url string `json:"repository_discussions_url,omitempty"` // A feed of discussions for a given repository.
	Security_advisories_url string `json:"security_advisories_url,omitempty"`
	Links map[string]interface{} `json:"_links"`
	Current_user_url string `json:"current_user_url,omitempty"`
	Timeline_url string `json:"timeline_url"`
	Current_user_actor_url string `json:"current_user_actor_url,omitempty"`
	Current_user_organization_url string `json:"current_user_organization_url,omitempty"`
	Current_user_organization_urls []string `json:"current_user_organization_urls,omitempty"`
	Repository_discussions_category_url string `json:"repository_discussions_category_url,omitempty"` // A feed of discussions for a given repository and category.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Status string `json:"status"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name,omitempty"`
	Events_url string `json:"events_url"`
	Followers int `json:"followers"`
	Owned_private_repos int `json:"owned_private_repos,omitempty"`
	Public_gists int `json:"public_gists"`
	Public_repos int `json:"public_repos"`
	Default_repository_permission string `json:"default_repository_permission,omitempty"`
	Twitter_username string `json:"twitter_username,omitempty"`
	Issues_url string `json:"issues_url"`
	Location string `json:"location,omitempty"`
	Avatar_url string `json:"avatar_url"`
	Login string `json:"login"`
	Members_url string `json:"members_url"`
	Html_url string `json:"html_url"`
	Repos_url string `json:"repos_url"`
	Description string `json:"description"`
	Members_can_create_public_repositories bool `json:"members_can_create_public_repositories,omitempty"`
	Created_at string `json:"created_at"`
	Dependabot_alerts_enabled_for_new_repositories bool `json:"dependabot_alerts_enabled_for_new_repositories,omitempty"` // Whether GitHub Advanced Security is automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Dependency_graph_enabled_for_new_repositories bool `json:"dependency_graph_enabled_for_new_repositories,omitempty"` // Whether dependency graph is automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Members_can_create_internal_repositories bool `json:"members_can_create_internal_repositories,omitempty"`
	Total_private_repos int `json:"total_private_repos,omitempty"`
	Secret_scanning_push_protection_custom_link_enabled bool `json:"secret_scanning_push_protection_custom_link_enabled,omitempty"` // Whether a custom link is shown to contributors who are blocked from pushing a secret by push protection.
	Has_repository_projects bool `json:"has_repository_projects"`
	Public_members_url string `json:"public_members_url"`
	Updated_at string `json:"updated_at"`
	Billing_email string `json:"billing_email,omitempty"`
	Members_can_fork_private_repositories bool `json:"members_can_fork_private_repositories,omitempty"`
	Secret_scanning_push_protection_custom_link string `json:"secret_scanning_push_protection_custom_link,omitempty"` // An optional URL string to display to contributors who are blocked from pushing a secret.
	Id int `json:"id"`
	Url string `json:"url"`
	Secret_scanning_push_protection_enabled_for_new_repositories bool `json:"secret_scanning_push_protection_enabled_for_new_repositories,omitempty"` // Whether secret scanning push protection is automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Blog string `json:"blog,omitempty"`
	Hooks_url string `json:"hooks_url"`
	Members_can_create_public_pages bool `json:"members_can_create_public_pages,omitempty"`
	Members_can_create_repositories bool `json:"members_can_create_repositories,omitempty"`
	Two_factor_requirement_enabled bool `json:"two_factor_requirement_enabled,omitempty"`
	Email string `json:"email,omitempty"`
	Members_can_create_private_pages bool `json:"members_can_create_private_pages,omitempty"`
	Members_can_create_private_repositories bool `json:"members_can_create_private_repositories,omitempty"`
	Collaborators int `json:"collaborators,omitempty"`
	Following int `json:"following"`
	Members_can_create_pages bool `json:"members_can_create_pages,omitempty"`
	Has_organization_projects bool `json:"has_organization_projects"`
	Members_allowed_repository_creation_type string `json:"members_allowed_repository_creation_type,omitempty"`
	Node_id string `json:"node_id"`
	Advanced_security_enabled_for_new_repositories bool `json:"advanced_security_enabled_for_new_repositories,omitempty"` // Whether GitHub Advanced Security is enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Secret_scanning_enabled_for_new_repositories bool `json:"secret_scanning_enabled_for_new_repositories,omitempty"` // Whether secret scanning is automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Plan map[string]interface{} `json:"plan,omitempty"`
	Disk_usage int `json:"disk_usage,omitempty"`
	Company string `json:"company,omitempty"`
	TypeField string `json:"type"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Dependabot_security_updates_enabled_for_new_repositories bool `json:"dependabot_security_updates_enabled_for_new_repositories,omitempty"` // Whether dependabot security updates are automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Is_verified bool `json:"is_verified,omitempty"`
	Private_gists int `json:"private_gists,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Download_url string `json:"download_url"`
	Filename string `json:"filename"`
	Os string `json:"os"`
	Sha256_checksum string `json:"sha256_checksum,omitempty"`
	Temp_download_token string `json:"temp_download_token,omitempty"` // A short lived bearer token used to download the runner, if needed.
	Architecture string `json:"architecture"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ref_type string `json:"ref_type"` // The type of Git ref object deleted in the repository.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pusher_type string `json:"pusher_type"` // The pusher type for the event. Can be either `user` or a deploy key.
	Ref string `json:"ref"` // The [`git ref`](https://docs.github.com/rest/reference/git#get-a-reference) resource.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Release interface{} `json:"release"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// License represents the License schema from the OpenAPI specification
type License struct {
	Description string `json:"description"`
	Limitations []string `json:"limitations"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
	Permissions []string `json:"permissions"`
	Featured bool `json:"featured"`
	Html_url string `json:"html_url"`
	Implementation string `json:"implementation"`
	Key string `json:"key"`
	Spdx_id string `json:"spdx_id"`
	Url string `json:"url"`
	Body string `json:"body"`
	Conditions []string `json:"conditions"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Guid string `json:"guid,omitempty"` // The GUID of the tool used to generate the code scanning analysis, if provided in the uploaded SARIF data.
	Name string `json:"name,omitempty"` // The name of the tool used to generate the code scanning analysis.
	Version string `json:"version,omitempty"` // The version of the tool used to generate the code scanning analysis.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Billable map[string]interface{} `json:"billable"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Comment string `json:"comment"` // The comment submitted with the deployment review
	Environments []map[string]interface{} `json:"environments"` // The list of environments that were approved or rejected
	State string `json:"state"` // Whether deployment to the environment(s) was approved or rejected or pending (with comments)
	User GeneratedType `json:"user"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Requested_teams []Team `json:"requested_teams,omitempty"`
	Body string `json:"body"`
	Diff_url string `json:"diff_url"`
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Base map[string]interface{} `json:"base"`
	Closed_at string `json:"closed_at"`
	Draft bool `json:"draft,omitempty"` // Indicates whether or not the pull request is a draft.
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	User GeneratedType `json:"user"` // A GitHub user.
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Merge_commit_sha string `json:"merge_commit_sha"`
	Statuses_url string `json:"statuses_url"`
	State string `json:"state"`
	Node_id string `json:"node_id"`
	Patch_url string `json:"patch_url"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Locked bool `json:"locked"`
	Id int `json:"id"`
	Url string `json:"url"`
	Head map[string]interface{} `json:"head"`
	Created_at string `json:"created_at"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Requested_reviewers []GeneratedType `json:"requested_reviewers,omitempty"`
	Review_comment_url string `json:"review_comment_url"`
	Updated_at string `json:"updated_at"`
	Auto_merge GeneratedType `json:"auto_merge"` // The status of auto merging a pull request.
	Issue_url string `json:"issue_url"`
	Links map[string]interface{} `json:"_links"`
	Commits_url string `json:"commits_url"`
	Review_comments_url string `json:"review_comments_url"`
	Title string `json:"title"`
	Comments_url string `json:"comments_url"`
	Merged_at string `json:"merged_at"`
	Html_url string `json:"html_url"`
	Labels []map[string]interface{} `json:"labels"`
	Number int `json:"number"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Account GeneratedType `json:"account"` // A GitHub user.
	Has_multiple_single_files bool `json:"has_multiple_single_files,omitempty"`
	Permissions GeneratedType `json:"permissions"` // The permissions granted to the user-to-server access token.
	Repositories_url string `json:"repositories_url"`
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Single_file_name string `json:"single_file_name"`
	Single_file_paths []string `json:"single_file_paths,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Help_uri string `json:"help_uri,omitempty"` // A link to the documentation for the rule used to detect the alert.
	Severity string `json:"severity,omitempty"` // The severity of the alert.
	Description string `json:"description,omitempty"` // A short description of the rule used to detect the alert.
	Full_description string `json:"full_description,omitempty"` // description of the rule used to detect the alert.
	Help string `json:"help,omitempty"` // Detailed documentation for the rule as GitHub Flavored Markdown.
	Name string `json:"name,omitempty"` // The name of the rule used to detect the alert.
	Security_severity_level string `json:"security_severity_level,omitempty"` // The security severity of the alert.
	Id string `json:"id,omitempty"` // A unique identifier for the rule used to detect the alert.
	Tags []string `json:"tags,omitempty"` // A set of tags applicable for the rule.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Page_url string `json:"page_url"` // The URI to the deployed GitHub Pages.
	Preview_url string `json:"preview_url,omitempty"` // The URI to the deployed GitHub Pages preview.
	Status_url string `json:"status_url"` // The URI to monitor GitHub Pages deployment status.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Sha string `json:"sha"` // The Commit SHA.
	Updated_at string `json:"updated_at"`
	Name string `json:"name"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Branches []map[string]interface{} `json:"branches"` // An array of branch objects containing the status' SHA. Each branch contains the given SHA, but the SHA may or may not be the head of the branch. The array includes a maximum of 10 branches.
	State string `json:"state"` // The new state. Can be `pending`, `success`, `failure`, or `error`.
	Context string `json:"context"`
	Id int `json:"id"` // The unique identifier of the status.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Avatar_url string `json:"avatar_url,omitempty"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Target_url string `json:"target_url"` // The optional link added to the status.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Commit map[string]interface{} `json:"commit"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Description string `json:"description"` // The optional human-readable description added to the status.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Uniques int `json:"uniques"`
	Clones []Traffic `json:"clones"`
	Count int `json:"count"`
}

// Migration represents the Migration schema from the OpenAPI specification
type Migration struct {
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Exclude_releases bool `json:"exclude_releases"`
	Lock_repositories bool `json:"lock_repositories"`
	Created_at string `json:"created_at"`
	Exclude_attachments bool `json:"exclude_attachments"`
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Exclude_git_data bool `json:"exclude_git_data"`
	Url string `json:"url"`
	Exclude []interface{} `json:"exclude,omitempty"`
	Guid string `json:"guid"`
	Repositories []Repository `json:"repositories"` // The repositories included in the migration. Only returned for export migrations.
	State string `json:"state"`
	Exclude_metadata bool `json:"exclude_metadata"`
	Exclude_owner_projects bool `json:"exclude_owner_projects"`
	Org_metadata_only bool `json:"org_metadata_only"`
	Archive_url string `json:"archive_url,omitempty"`
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Membership map[string]interface{} `json:"membership,omitempty"` // The membership between the user and the organization. Not present when the action is `member_invited`.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Team map[string]interface{} `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Member map[string]interface{} `json:"member"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Scope string `json:"scope"` // The scope of the membership. Currently, can only be `team`.
	Sender map[string]interface{} `json:"sender"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Zen string `json:"zen,omitempty"` // Random string of GitHub zen.
	Hook map[string]interface{} `json:"hook,omitempty"` // The webhook that is being pinged
	Hook_id int `json:"hook_id,omitempty"` // The ID of the webhook that triggered the ping.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Sponsorship map[string]interface{} `json:"sponsorship"`
	Action string `json:"action"`
}

// Hook represents the Hook schema from the OpenAPI specification
type Hook struct {
	Updated_at string `json:"updated_at"`
	Deliveries_url string `json:"deliveries_url,omitempty"`
	Id int `json:"id"` // Unique identifier of the webhook.
	TypeField string `json:"type"`
	Url string `json:"url"`
	Active bool `json:"active"` // Determines whether the hook is actually triggered on pushes.
	Config map[string]interface{} `json:"config"`
	Created_at string `json:"created_at"`
	Test_url string `json:"test_url"`
	Events []string `json:"events"` // Determines what events the hook is triggered for. Default: ['push'].
	Last_response GeneratedType `json:"last_response"`
	Name string `json:"name"` // The name of a valid service, use 'web' for a webhook.
	Ping_url string `json:"ping_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id float64 `json:"id"`
	Project_node_id string `json:"project_node_id,omitempty"`
	Content_node_id string `json:"content_node_id"`
	Content_type string `json:"content_type"` // The type of content tracked in a project item
	Creator GeneratedType `json:"creator,omitempty"` // A GitHub user.
	Node_id string `json:"node_id,omitempty"`
	Archived_at string `json:"archived_at"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Can_encrypt_comms bool `json:"can_encrypt_comms"`
	Can_sign bool `json:"can_sign"`
	Public_key string `json:"public_key"`
	Revoked bool `json:"revoked"`
	Can_encrypt_storage bool `json:"can_encrypt_storage"`
	Expires_at string `json:"expires_at"`
	Raw_key string `json:"raw_key"`
	Emails []map[string]interface{} `json:"emails"`
	Name string `json:"name,omitempty"`
	Subkeys []map[string]interface{} `json:"subkeys"`
	Created_at string `json:"created_at"`
	Can_certify bool `json:"can_certify"`
	Primary_key_id int `json:"primary_key_id"`
	Key_id string `json:"key_id"`
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	PackageField map[string]interface{} `json:"package"` // Information about the package.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Team map[string]interface{} `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// Blob represents the Blob schema from the OpenAPI specification
type Blob struct {
	Encoding string `json:"encoding"`
	Highlighted_content string `json:"highlighted_content,omitempty"`
	Node_id string `json:"node_id"`
	Sha string `json:"sha"`
	Size int `json:"size"`
	Url string `json:"url"`
	Content string `json:"content"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Requester map[string]interface{} `json:"requester"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Installation Installation `json:"installation"` // Installation
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Repositories_added []map[string]interface{} `json:"repositories_added"` // An array of repository objects, which were added to the installation.
	Repositories_removed []map[string]interface{} `json:"repositories_removed"` // An array of repository objects, which were removed from the installation.
}

// Issue represents the Issue schema from the OpenAPI specification
type Issue struct {
	Repository_url string `json:"repository_url"`
	Locked bool `json:"locked"`
	Comments int `json:"comments"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Timeline_url string `json:"timeline_url,omitempty"`
	Labels_url string `json:"labels_url"`
	State string `json:"state"` // State of the issue; either 'open' or 'closed'
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Created_at string `json:"created_at"`
	Title string `json:"title"` // Title of the issue
	Url string `json:"url"` // URL for the issue
	Node_id string `json:"node_id"`
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	Body string `json:"body,omitempty"` // Contents of the issue
	Reactions GeneratedType `json:"reactions,omitempty"`
	Draft bool `json:"draft,omitempty"`
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	State_reason string `json:"state_reason,omitempty"` // The reason for the current state
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Updated_at string `json:"updated_at"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Closed_by GeneratedType `json:"closed_by,omitempty"` // A GitHub user.
	Html_url string `json:"html_url"`
	Body_text string `json:"body_text,omitempty"`
	User GeneratedType `json:"user"` // A GitHub user.
	Closed_at string `json:"closed_at"`
	Events_url string `json:"events_url"`
	Body_html string `json:"body_html,omitempty"`
	Labels []interface{} `json:"labels"` // Labels to associate with this issue; pass one or more label names to replace the set of labels on this issue; send an empty array to clear all labels from the issue; note that the labels are silently dropped for users without push access to the repository
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Comments_url string `json:"comments_url"`
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Id int `json:"id"`
	Number int `json:"number"` // Number uniquely identifying the issue within its repository
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Answer map[string]interface{} `json:"answer"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Has_downloads bool `json:"has_downloads,omitempty"`
	Tags_url string `json:"tags_url"`
	Allow_forking bool `json:"allow_forking,omitempty"`
	Has_wiki bool `json:"has_wiki,omitempty"`
	Private bool `json:"private"`
	Pushed_at string `json:"pushed_at,omitempty"`
	Id int `json:"id"`
	Issue_comment_url string `json:"issue_comment_url"`
	Stargazers_url string `json:"stargazers_url"`
	Mirror_url string `json:"mirror_url,omitempty"`
	Downloads_url string `json:"downloads_url"`
	Teams_url string `json:"teams_url"`
	Default_branch string `json:"default_branch,omitempty"`
	Language string `json:"language,omitempty"`
	Subscription_url string `json:"subscription_url"`
	Disabled bool `json:"disabled,omitempty"`
	Contents_url string `json:"contents_url"`
	Notifications_url string `json:"notifications_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Homepage string `json:"homepage,omitempty"`
	Url string `json:"url"`
	Watchers int `json:"watchers,omitempty"`
	Archived bool `json:"archived,omitempty"`
	Subscribers_url string `json:"subscribers_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Branches_url string `json:"branches_url"`
	Has_projects bool `json:"has_projects,omitempty"`
	Keys_url string `json:"keys_url"`
	Size int `json:"size,omitempty"` // The size of the repository. Size is calculated hourly. When a repository is initially created, the size is 0.
	Git_tags_url string `json:"git_tags_url"`
	Topics []string `json:"topics,omitempty"`
	Trees_url string `json:"trees_url"`
	Assignees_url string `json:"assignees_url"`
	Collaborators_url string `json:"collaborators_url"`
	Is_template bool `json:"is_template,omitempty"`
	Name string `json:"name"`
	Statuses_url string `json:"statuses_url"`
	Full_name string `json:"full_name"`
	Network_count int `json:"network_count,omitempty"`
	Forks int `json:"forks,omitempty"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Comments_url string `json:"comments_url"`
	Blobs_url string `json:"blobs_url"`
	Description string `json:"description"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Events_url string `json:"events_url"`
	Has_pages bool `json:"has_pages,omitempty"`
	Git_refs_url string `json:"git_refs_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Visibility string `json:"visibility,omitempty"`
	Stargazers_count int `json:"stargazers_count,omitempty"`
	Has_issues bool `json:"has_issues,omitempty"`
	Code_of_conduct GeneratedType `json:"code_of_conduct,omitempty"` // Code Of Conduct
	Pulls_url string `json:"pulls_url"`
	Html_url string `json:"html_url"`
	Watchers_count int `json:"watchers_count,omitempty"`
	Deployments_url string `json:"deployments_url"`
	Merges_url string `json:"merges_url"`
	Open_issues_count int `json:"open_issues_count,omitempty"`
	Issue_events_url string `json:"issue_events_url"`
	Forks_url string `json:"forks_url"`
	Languages_url string `json:"languages_url"`
	Open_issues int `json:"open_issues,omitempty"`
	Labels_url string `json:"labels_url"`
	Node_id string `json:"node_id"`
	Commits_url string `json:"commits_url"`
	Updated_at string `json:"updated_at,omitempty"`
	Contributors_url string `json:"contributors_url"`
	Svn_url string `json:"svn_url,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Compare_url string `json:"compare_url"`
	Archive_url string `json:"archive_url"`
	Fork bool `json:"fork"`
	Git_url string `json:"git_url,omitempty"`
	License map[string]interface{} `json:"license,omitempty"`
	Forks_count int `json:"forks_count,omitempty"`
	Git_commits_url string `json:"git_commits_url"`
	Ssh_url string `json:"ssh_url,omitempty"`
	Hooks_url string `json:"hooks_url"`
	Issues_url string `json:"issues_url"`
	Clone_url string `json:"clone_url,omitempty"`
	Has_discussions bool `json:"has_discussions,omitempty"`
	Role_name string `json:"role_name,omitempty"`
	Releases_url string `json:"releases_url"`
	Security_and_analysis GeneratedType `json:"security_and_analysis,omitempty"`
	Milestones_url string `json:"milestones_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	State string `json:"state,omitempty"` // State of a code scanning alert.
	Commit_sha string `json:"commit_sha,omitempty"`
	Environment string `json:"environment,omitempty"` // Identifies the variable values associated with the environment in which the analysis that generated this alert instance was performed, such as the language that was analyzed.
	Html_url string `json:"html_url,omitempty"`
	Location GeneratedType `json:"location,omitempty"` // Describe a region within a file for the alert.
	Ref string `json:"ref,omitempty"` // The full Git reference, formatted as `refs/heads/<branch name>`, `refs/pull/<number>/merge`, or `refs/pull/<number>/head`.
	Analysis_key string `json:"analysis_key,omitempty"` // Identifies the configuration under which the analysis was executed. For example, in GitHub Actions this includes the workflow filename and job name.
	Category string `json:"category,omitempty"` // Identifies the configuration under which the analysis was executed. Used to distinguish between multiple analyses for the same tool and commit, but performed on different languages or different parts of the code.
	Classifications []string `json:"classifications,omitempty"` // Classifications that have been applied to the file that triggered the alert. For example identifying it as documentation, or a generated file.
	Message map[string]interface{} `json:"message,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// Key represents the Key schema from the OpenAPI specification
type Key struct {
	Title string `json:"title"`
	Url string `json:"url"`
	Verified bool `json:"verified"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Key string `json:"key"`
	Read_only bool `json:"read_only"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commits_url string `json:"commits_url,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
	Public bool `json:"public,omitempty"`
	History []GeneratedType `json:"history,omitempty"`
	Description string `json:"description,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Comments int `json:"comments,omitempty"`
	User string `json:"user,omitempty"`
	Git_push_url string `json:"git_push_url,omitempty"`
	Url string `json:"url,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Fork_of map[string]interface{} `json:"fork_of,omitempty"` // Gist
	Truncated bool `json:"truncated,omitempty"`
	Id string `json:"id,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Comments_url string `json:"comments_url,omitempty"`
	Forks_url string `json:"forks_url,omitempty"`
	Files map[string]interface{} `json:"files,omitempty"`
	Forks []map[string]interface{} `json:"forks,omitempty"`
	Owner GeneratedType `json:"owner,omitempty"` // A GitHub user.
	Git_pull_url string `json:"git_pull_url,omitempty"`
}

// Autolink represents the Autolink schema from the OpenAPI specification
type Autolink struct {
	Key_prefix string `json:"key_prefix"` // The prefix of a key that is linkified.
	Url_template string `json:"url_template"` // A template for the target URL that is generated if a key was found.
	Id int `json:"id"`
	Is_alphanumeric bool `json:"is_alphanumeric"` // Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Merge_group map[string]interface{} `json:"merge_group"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// Release represents the Release schema from the OpenAPI specification
type Release struct {
	Published_at string `json:"published_at"`
	Discussion_url string `json:"discussion_url,omitempty"` // The URL of the release discussion.
	Name string `json:"name"`
	Tag_name string `json:"tag_name"` // The name of the tag.
	Tarball_url string `json:"tarball_url"`
	Upload_url string `json:"upload_url"`
	Draft bool `json:"draft"` // true to create a draft (unpublished) release, false to create a published one.
	Body string `json:"body,omitempty"`
	Target_commitish string `json:"target_commitish"` // Specifies the commitish value that determines where the Git tag is created from.
	Zipball_url string `json:"zipball_url"`
	Body_text string `json:"body_text,omitempty"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Mentions_count int `json:"mentions_count,omitempty"`
	Prerelease bool `json:"prerelease"` // Whether to identify the release as a prerelease or a full release.
	Html_url string `json:"html_url"`
	Assets []GeneratedType `json:"assets"`
	Body_html string `json:"body_html,omitempty"`
	Created_at string `json:"created_at"`
	Assets_url string `json:"assets_url"`
	Url string `json:"url"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Author GeneratedType `json:"author"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Visibility string `json:"visibility"` // Visibility of a secret
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Member map[string]interface{} `json:"member"`
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Sender map[string]interface{} `json:"sender"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Scope string `json:"scope"` // The scope of the membership. Currently, can only be `team`.
	Team map[string]interface{} `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Bio string `json:"bio,omitempty"`
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Blog string `json:"blog,omitempty"`
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"`
	Score float64 `json:"score"`
	Public_repos int `json:"public_repos,omitempty"`
	Following_url string `json:"following_url"`
	Email string `json:"email,omitempty"`
	Received_events_url string `json:"received_events_url"`
	Following int `json:"following,omitempty"`
	Organizations_url string `json:"organizations_url"`
	Followers_url string `json:"followers_url"`
	Gravatar_id string `json:"gravatar_id"`
	Starred_url string `json:"starred_url"`
	Location string `json:"location,omitempty"`
	Login string `json:"login"`
	Updated_at string `json:"updated_at,omitempty"`
	Followers int `json:"followers,omitempty"`
	Repos_url string `json:"repos_url"`
	Suspended_at string `json:"suspended_at,omitempty"`
	Gists_url string `json:"gists_url"`
	Created_at string `json:"created_at,omitempty"`
	Hireable bool `json:"hireable,omitempty"`
	Url string `json:"url"`
	Company string `json:"company,omitempty"`
	Events_url string `json:"events_url"`
	Id int `json:"id"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Public_gists int `json:"public_gists,omitempty"`
	Subscriptions_url string `json:"subscriptions_url"`
	Avatar_url string `json:"avatar_url"`
	Site_admin bool `json:"site_admin"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repositories []Repository `json:"repositories,omitempty"` // The repositories this token has access to
	Repository_selection string `json:"repository_selection,omitempty"` // Describe whether all repositories have been selected or there's a selection involved
	Single_file string `json:"single_file,omitempty"`
	Token string `json:"token"` // The token used for authentication
	Expires_at string `json:"expires_at"` // The time this token expires
	Permissions map[string]interface{} `json:"permissions,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Event string `json:"event"`
	Milestone map[string]interface{} `json:"milestone"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Commit_id string `json:"commit_id"`
	Id int `json:"id"`
	Url string `json:"url"`
	Commit_url string `json:"commit_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Starred_url string `json:"starred_url"`
	Id int `json:"id"`
	Name string `json:"name,omitempty"`
	Node_id string `json:"node_id"`
	Gists_url string `json:"gists_url"`
	Login string `json:"login"`
	Html_url string `json:"html_url"`
	Events_url string `json:"events_url"`
	Organizations_url string `json:"organizations_url"`
	Gravatar_id string `json:"gravatar_id"`
	Followers_url string `json:"followers_url"`
	Received_events_url string `json:"received_events_url"`
	Site_admin bool `json:"site_admin"`
	TypeField string `json:"type"`
	Following_url string `json:"following_url"`
	Repos_url string `json:"repos_url"`
	Email string `json:"email,omitempty"`
	Subscriptions_url string `json:"subscriptions_url"`
	Avatar_url string `json:"avatar_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Url string `json:"url"`
	Role_name string `json:"role_name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Count int `json:"count"`
	Path string `json:"path"`
	Title string `json:"title"`
	Uniques int `json:"uniques"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.requested_action JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Tarball_url string `json:"tarball_url"`
	Zipball_url string `json:"zipball_url"`
	Commit map[string]interface{} `json:"commit"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit map[string]interface{} `json:"commit"`
	Name string `json:"name"`
	Protected bool `json:"protected"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Status string `json:"status"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Base_branch string `json:"base_branch,omitempty"`
	Merge_type string `json:"merge_type,omitempty"`
	Message string `json:"message,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Change_status map[string]interface{} `json:"change_status,omitempty"`
	Committed_at string `json:"committed_at,omitempty"`
	Url string `json:"url,omitempty"`
	User GeneratedType `json:"user,omitempty"` // A GitHub user.
	Version string `json:"version,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Dismissed_reason string `json:"dismissed_reason"` // **Required when the state is dismissed.** The reason for dismissing or closing the alert.
	Url string `json:"url"` // The REST API URL of the alert resource.
	Dismissed_comment string `json:"dismissed_comment,omitempty"` // The dismissal comment associated with the dismissal of the alert.
	Tool GeneratedType `json:"tool"`
	Number int `json:"number"` // The security alert number.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Fixed_at string `json:"fixed_at,omitempty"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	State string `json:"state"` // State of a code scanning alert.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
	Most_recent_instance GeneratedType `json:"most_recent_instance"`
	Rule GeneratedType `json:"rule"`
	Instances_url string `json:"instances_url"` // The REST API URL for fetching the list of instances for an alert.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion map[string]interface{} `json:"discussion"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Label map[string]interface{} `json:"label"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Include_claim_keys []string `json:"include_claim_keys,omitempty"` // Array of unique strings. Each claim key can only contain alphanumeric characters and underscores.
	Use_default bool `json:"use_default"` // Whether to use the default template or not. If `true`, the `include_claim_keys` field is ignored.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Requested_team Team `json:"requested_team,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Review_requester GeneratedType `json:"review_requester"` // A GitHub user.
	Id int `json:"id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Requested_reviewer GeneratedType `json:"requested_reviewer,omitempty"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Commit_url string `json:"commit_url"`
	Event string `json:"event"`
	Url string `json:"url"`
	Node_id string `json:"node_id"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Forks_url string `json:"forks_url"`
	Updated_at string `json:"updated_at"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	License GeneratedType `json:"license"` // License Simple
	Watchers int `json:"watchers"`
	Trees_url string `json:"trees_url"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Subscription_url string `json:"subscription_url"`
	Downloads_url string `json:"downloads_url"`
	Git_commits_url string `json:"git_commits_url"`
	Description string `json:"description"`
	Keys_url string `json:"keys_url"`
	Url string `json:"url"`
	Archive_url string `json:"archive_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Watchers_count int `json:"watchers_count"`
	Languages_url string `json:"languages_url"`
	Labels_url string `json:"labels_url"`
	Ssh_url string `json:"ssh_url"`
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Assignees_url string `json:"assignees_url"`
	Subscribers_url string `json:"subscribers_url"`
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Pushed_at string `json:"pushed_at"`
	Forks_count int `json:"forks_count"`
	Git_tags_url string `json:"git_tags_url"`
	Archived bool `json:"archived"` // Whether the repository is archived.
	Language string `json:"language"`
	Open_issues int `json:"open_issues"`
	Collaborators_url string `json:"collaborators_url"`
	Issues_url string `json:"issues_url"`
	Milestones_url string `json:"milestones_url"`
	Homepage string `json:"homepage"`
	Releases_url string `json:"releases_url"`
	Has_pages bool `json:"has_pages"`
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Role_name string `json:"role_name,omitempty"`
	Merges_url string `json:"merges_url"`
	Master_branch string `json:"master_branch,omitempty"`
	Comments_url string `json:"comments_url"`
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Hooks_url string `json:"hooks_url"`
	Events_url string `json:"events_url"`
	Notifications_url string `json:"notifications_url"`
	Stargazers_url string `json:"stargazers_url"`
	Commits_url string `json:"commits_url"`
	Html_url string `json:"html_url"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Tags_url string `json:"tags_url"`
	Open_issues_count int `json:"open_issues_count"`
	Git_refs_url string `json:"git_refs_url"`
	Clone_url string `json:"clone_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Stargazers_count int `json:"stargazers_count"`
	Created_at string `json:"created_at"`
	Id int `json:"id"` // Unique identifier of the repository
	Issue_comment_url string `json:"issue_comment_url"`
	Blobs_url string `json:"blobs_url"`
	Full_name string `json:"full_name"`
	Compare_url string `json:"compare_url"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Forks int `json:"forks"`
	Network_count int `json:"network_count,omitempty"`
	Contributors_url string `json:"contributors_url"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Statuses_url string `json:"statuses_url"`
	Contents_url string `json:"contents_url"`
	Teams_url string `json:"teams_url"`
	Deployments_url string `json:"deployments_url"`
	Svn_url string `json:"svn_url"`
	Branches_url string `json:"branches_url"`
	Issue_events_url string `json:"issue_events_url"`
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Fork bool `json:"fork"`
	Pulls_url string `json:"pulls_url"`
	Node_id string `json:"node_id"`
	Git_url string `json:"git_url"`
	Mirror_url string `json:"mirror_url"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Template_repository GeneratedType `json:"template_repository,omitempty"` // A repository on GitHub.
	Private bool `json:"private"` // Whether the repository is private or public.
	Topics []string `json:"topics,omitempty"`
	Size int `json:"size"`
	Name string `json:"name"` // The name of the repository.
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Analyses_url string `json:"analyses_url,omitempty"` // The REST API URL for getting the analyses associated with the upload.
	Errors []string `json:"errors,omitempty"` // Any errors that ocurred during processing of the delivery.
	Processing_status string `json:"processing_status,omitempty"` // `pending` files have not yet been processed, while `complete` means results from the SARIF have been stored. `failed` files have either not been processed at all, or could only be partially processed.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"`
	Updated_at string `json:"updated_at"` // The date and time at which the variable was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Value string `json:"value"` // The value of the variable.
	Visibility string `json:"visibility"` // Visibility of a variable
	Created_at string `json:"created_at"` // The date and time at which the variable was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Name string `json:"name"` // The name of the variable.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pages []map[string]interface{} `json:"pages"` // The pages that were updated.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Release map[string]interface{} `json:"release"` // The [release](https://docs.github.com/rest/reference/repos/#get-a-release) object.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Event string `json:"event"`
	Requested_reviewer GeneratedType `json:"requested_reviewer,omitempty"` // A GitHub user.
	Requested_team Team `json:"requested_team,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Review_requester GeneratedType `json:"review_requester"` // A GitHub user.
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Id int `json:"id"`
	Commit_id string `json:"commit_id"`
}

// Label represents the Label schema from the OpenAPI specification
type Label struct {
	Description string `json:"description"`
	Id int64 `json:"id"`
	Name string `json:"name"` // The name of the label.
	Node_id string `json:"node_id"`
	Url string `json:"url"` // URL for the label
	Color string `json:"color"` // 6-character hex code, without the leading #, identifying the color
	DefaultField bool `json:"default"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// Artifact represents the Artifact schema from the OpenAPI specification
type Artifact struct {
	Expired bool `json:"expired"` // Whether or not the artifact has expired.
	Size_in_bytes int `json:"size_in_bytes"` // The size in bytes of the artifact.
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Archive_download_url string `json:"archive_download_url"`
	Id int `json:"id"`
	Name string `json:"name"` // The name of the artifact.
	Node_id string `json:"node_id"`
	Workflow_run map[string]interface{} `json:"workflow_run,omitempty"`
	Created_at string `json:"created_at"`
	Expires_at string `json:"expires_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Path string `json:"path"`
	Raw_details string `json:"raw_details"`
	Start_column int `json:"start_column"`
	Start_line int `json:"start_line"`
	Title string `json:"title"`
	Annotation_level string `json:"annotation_level"`
	Blob_href string `json:"blob_href"`
	End_column int `json:"end_column"`
	End_line int `json:"end_line"`
	Message string `json:"message"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Event string `json:"event"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Commit_id string `json:"commit_id"`
	Label map[string]interface{} `json:"label"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Label map[string]interface{} `json:"label,omitempty"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Url string `json:"url"` // URL for the repository invitation
	Expired bool `json:"expired,omitempty"` // Whether or not the invitation has expired
	Html_url string `json:"html_url"`
	Invitee GeneratedType `json:"invitee"` // A GitHub user.
	Permissions string `json:"permissions"` // The permission associated with the invitation.
	Id int `json:"id"` // Unique identifier of the repository invitation.
	Inviter GeneratedType `json:"inviter"` // A GitHub user.
	Node_id string `json:"node_id"`
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository map[string]interface{} `json:"repository"` // A git repository
	Head_commit map[string]interface{} `json:"head_commit"`
	Deleted bool `json:"deleted"` // Whether this push deleted the `ref`.
	Pusher map[string]interface{} `json:"pusher"` // Metaproperties for Git author/committer information.
	Commits []map[string]interface{} `json:"commits"` // An array of commit objects describing the pushed commits. (Pushed commits are all commits that are included in the `compare` between the `before` commit and the `after` commit.) The array includes a maximum of 20 commits. If necessary, you can use the [Commits API](https://docs.github.com/rest/reference/repos#commits) to fetch additional commits. This limit is applied to timeline events only and isn't applied to webhook deliveries.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Compare string `json:"compare"` // URL that shows the changes in this `ref` update, from the `before` commit to the `after` commit. For a newly created `ref` that is directly based on the default branch, this is the comparison between the head of the default branch and the `after` commit. Otherwise, this shows all commits until the `after` commit.
	Forced bool `json:"forced"` // Whether this push was a force push of the `ref`.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Before string `json:"before"` // The SHA of the most recent commit on `ref` before the push.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Created bool `json:"created"` // Whether this push created the `ref`.
	After string `json:"after"` // The SHA of the most recent commit on `ref` after the push.
	Base_ref string `json:"base_ref"`
	Ref string `json:"ref"` // The full git ref that was pushed. Example: `refs/heads/main` or `refs/tags/v3.14.1`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Limit string `json:"limit"` // The type of GitHub user that can comment, open issues, or create pull requests while the interaction limit is in effect.
	Origin string `json:"origin"`
	Expires_at string `json:"expires_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Documentation_url string `json:"documentation_url"`
	Errors []string `json:"errors,omitempty"`
	Message string `json:"message"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the ping JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Requester interface{} `json:"requester,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url,omitempty"` // The REST API URL for checking the status of the upload.
	Id string `json:"id,omitempty"` // An identifier for the upload.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Documentation_url string `json:"documentation_url,omitempty"`
	Message string `json:"message,omitempty"`
	Status string `json:"status,omitempty"`
	Url string `json:"url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Check_run GeneratedType `json:"check_run"` // A check performed on the code of a given code change
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Message string `json:"message"`
	Documentation_url string `json:"documentation_url"`
	Errors []map[string]interface{} `json:"errors,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	From string `json:"from"`
	To string `json:"to"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Forkee interface{} `json:"forkee"` // The created [`repository`](https://docs.github.com/rest/reference/repos#get-a-repository) resource.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Id int `json:"id"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Build map[string]interface{} `json:"build"` // The [List GitHub Pages builds](https://docs.github.com/rest/reference/repos#list-github-pages-builds) itself.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Blocked_user map[string]interface{} `json:"blocked_user"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sponsorship map[string]interface{} `json:"sponsorship"`
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	PackageField map[string]interface{} `json:"package"` // Information about the package.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Details interface{} `json:"details"`
	TypeField string `json:"type"` // The location type. Because secrets may be found in different types of resources (ie. code, comments, issues), this field identifies the type of resource where the secret was found.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request map[string]interface{} `json:"pull_request"`
	Action string `json:"action"`
	Assignee map[string]interface{} `json:"assignee,omitempty"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Number int `json:"number"` // The pull request number.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Key map[string]interface{} `json:"key"` // The [`deploy key`](https://docs.github.com/rest/reference/deployments#get-a-deploy-key) resource.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action,omitempty"`
	Check_run GeneratedType `json:"check_run"` // A check performed on the code of a given code change
}

// Dependency represents the Dependency schema from the OpenAPI specification
type Dependency struct {
	Package_url string `json:"package_url,omitempty"` // Package-url (PURL) of dependency. See https://github.com/package-url/purl-spec for more details.
	Relationship string `json:"relationship,omitempty"` // A notation of whether a dependency is requested directly by this manifest or is a dependency of another dependency.
	Scope string `json:"scope,omitempty"` // A notation of whether the dependency is required for the primary build artifact (runtime) or is only used for development. Future versions of this specification may allow for more granular scopes.
	Dependencies []string `json:"dependencies,omitempty"` // Array of package-url (PURLs) of direct child dependencies.
	Metadata Metadata `json:"metadata,omitempty"` // User-defined metadata to store domain-specific information limited to 8 keys with scalar values.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Href string `json:"href"`
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Milestone map[string]interface{} `json:"milestone"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Release map[string]interface{} `json:"release"` // The [release](https://docs.github.com/rest/reference/repos/#get-a-release) object.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// Traffic represents the Traffic schema from the OpenAPI specification
type Traffic struct {
	Timestamp string `json:"timestamp"`
	Uniques int `json:"uniques"`
	Count int `json:"count"`
}

// Commit represents the Commit schema from the OpenAPI specification
type Commit struct {
	Parents []map[string]interface{} `json:"parents"`
	Author GeneratedType `json:"author"` // A GitHub user.
	Comments_url string `json:"comments_url"`
	Commit map[string]interface{} `json:"commit"`
	Sha string `json:"sha"`
	Stats map[string]interface{} `json:"stats,omitempty"`
	Url string `json:"url"`
	Committer GeneratedType `json:"committer"` // A GitHub user.
	Files []GeneratedType `json:"files,omitempty"`
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Object map[string]interface{} `json:"object"`
	Ref string `json:"ref"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Links map[string]interface{} `json:"_links"`
	Download_url string `json:"download_url"`
	Git_url string `json:"git_url"`
	Html_url string `json:"html_url"`
	Path string `json:"path"`
	Sha string `json:"sha"`
	TypeField string `json:"type"`
	Url string `json:"url"`
	Name string `json:"name"`
	Size int `json:"size"`
	Submodule_git_url string `json:"submodule_git_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Actions_meta map[string]interface{} `json:"actions_meta,omitempty"`
	Check_suite map[string]interface{} `json:"check_suite"` // The [check_suite](https://docs.github.com/rest/reference/checks#suites).
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Active_caches_count int `json:"active_caches_count"` // The number of active caches in the repository.
	Active_caches_size_in_bytes int `json:"active_caches_size_in_bytes"` // The sum of the size in bytes of all the active cache items in the repository.
	Full_name string `json:"full_name"` // The repository owner and name for the cache usage being shown.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Projects_v2 GeneratedType `json:"projects_v2"` // A projects v2 project
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes"` // The changes to the collaborator permissions
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Member map[string]interface{} `json:"member"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Team map[string]interface{} `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id string `json:"id,omitempty"` // Id for the export details
	Sha string `json:"sha,omitempty"` // Git commit SHA of the exported branch
	State string `json:"state,omitempty"` // State of the latest export
	Branch string `json:"branch,omitempty"` // Name of the exported branch
	Completed_at string `json:"completed_at,omitempty"` // Completion time of the last export operation
	Export_url string `json:"export_url,omitempty"` // Url for fetching export details
	Html_url string `json:"html_url,omitempty"` // Web url for the exported branch
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id string `json:"id,omitempty"` // A unique identifier for the rule used to detect the alert.
	Name string `json:"name,omitempty"` // The name of the rule used to detect the alert.
	Severity string `json:"severity,omitempty"` // The severity of the alert.
	Tags []string `json:"tags,omitempty"` // A set of tags applicable for the rule.
	Description string `json:"description,omitempty"` // A short description of the rule used to detect the alert.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Key map[string]interface{} `json:"key"` // The [`deploy key`](https://docs.github.com/rest/reference/deployments#get-a-deploy-key) resource.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue interface{} `json:"issue"`
	Milestone map[string]interface{} `json:"milestone,omitempty"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Parents []map[string]interface{} `json:"parents"`
	Sha string `json:"sha"` // SHA for the commit
	Tree map[string]interface{} `json:"tree"`
	Message string `json:"message"` // Message describing the purpose of the commit
	Committer map[string]interface{} `json:"committer"` // Identifying information for the git-user
	Url string `json:"url"`
	Verification map[string]interface{} `json:"verification"`
	Author map[string]interface{} `json:"author"` // Identifying information for the git-user
	Html_url string `json:"html_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name,omitempty"`
	Date string `json:"date,omitempty"`
	Email string `json:"email,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sha string `json:"sha"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Id int `json:"id"` // The project card's ID
	Project_url string `json:"project_url"`
	Updated_at string `json:"updated_at"`
	Content_url string `json:"content_url,omitempty"`
	Note string `json:"note"`
	Project_id string `json:"project_id,omitempty"`
	Url string `json:"url"`
	Column_name string `json:"column_name,omitempty"`
	Column_url string `json:"column_url"`
	Node_id string `json:"node_id"`
	Archived bool `json:"archived,omitempty"` // Whether or not the card is archived
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Release map[string]interface{} `json:"release"` // The [release](https://docs.github.com/rest/reference/repos/#get-a-release) object.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// Link represents the Link schema from the OpenAPI specification
type Link struct {
	Href string `json:"href"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Total_active_caches_count int `json:"total_active_caches_count"` // The count of active caches across all repositories of an enterprise or an organization.
	Total_active_caches_size_in_bytes int `json:"total_active_caches_size_in_bytes"` // The total size in bytes of all active cache items across all repositories of an enterprise or an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Title string `json:"title,omitempty"`
	Url string `json:"url,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Id int `json:"id,omitempty"`
	Key string `json:"key"` // The Base64 encoded public key.
	Key_id string `json:"key_id"` // The identifier for the key.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Short_description string `json:"short_description"`
	Title string `json:"title"`
	Updated_at string `json:"updated_at"`
	Closed_at string `json:"closed_at"`
	Id float64 `json:"id"`
	Node_id string `json:"node_id"`
	Public bool `json:"public"`
	Created_at string `json:"created_at"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Deleted_at string `json:"deleted_at"`
	Deleted_by GeneratedType `json:"deleted_by"` // A GitHub user.
	Description string `json:"description"`
	Number int `json:"number"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Effective_date string `json:"effective_date"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Action string `json:"action"`
	Previous_marketplace_purchase map[string]interface{} `json:"previous_marketplace_purchase,omitempty"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Marketplace_purchase interface{} `json:"marketplace_purchase"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Included_gigabytes_bandwidth int `json:"included_gigabytes_bandwidth"` // Free storage space (GB) for GitHub Packages.
	Total_gigabytes_bandwidth_used int `json:"total_gigabytes_bandwidth_used"` // Sum of the free and paid storage space (GB) for GitHuub Packages.
	Total_paid_gigabytes_bandwidth_used int `json:"total_paid_gigabytes_bandwidth_used"` // Total paid storage space (GB) for GitHuub Packages.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow map[string]interface{} `json:"workflow"`
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Guid string `json:"guid"` // Unique identifier for the event (shared with all deliveries for all webhooks that subscribe to this event).
	Redelivery bool `json:"redelivery"` // Whether the webhook delivery is a redelivery.
	Status string `json:"status"` // Describes the response returned after attempting the delivery.
	Duration float64 `json:"duration"` // Time spent delivering.
	Status_code int `json:"status_code"` // Status code received when delivery was made.
	Delivered_at string `json:"delivered_at"` // Time when the webhook delivery occurred.
	Id int `json:"id"` // Unique identifier of the webhook delivery.
	Installation_id int `json:"installation_id"` // The id of the GitHub App installation associated with this event.
	Repository_id int `json:"repository_id"` // The id of the repository associated with this event.
	Event string `json:"event"` // The event that triggered the delivery.
	Action string `json:"action"` // The type of activity for the event that triggered the delivery.
}

// Discussion represents the Discussion schema from the OpenAPI specification
type Discussion struct {
	Answer_chosen_by map[string]interface{} `json:"answer_chosen_by"`
	Repository_url string `json:"repository_url"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Comments int `json:"comments"`
	Answer_chosen_at string `json:"answer_chosen_at"`
	Locked bool `json:"locked"`
	Answer_html_url string `json:"answer_html_url"`
	Created_at string `json:"created_at"`
	State string `json:"state"` // The current state of the discussion. `converting` means that the discussion is being converted from an issue. `transferring` means that the discussion is being transferred from another repository.
	Updated_at string `json:"updated_at"`
	User map[string]interface{} `json:"user"`
	Active_lock_reason string `json:"active_lock_reason"`
	Body string `json:"body"`
	Title string `json:"title"`
	Id int `json:"id"`
	Timeline_url string `json:"timeline_url,omitempty"`
	Reactions map[string]interface{} `json:"reactions,omitempty"`
	Category map[string]interface{} `json:"category"`
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Number int `json:"number"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Preferences map[string]interface{} `json:"preferences"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Views []Traffic `json:"views"`
	Count int `json:"count"`
	Uniques int `json:"uniques"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Author GeneratedType `json:"author"` // A GitHub user.
	Total int `json:"total"`
	Weeks []map[string]interface{} `json:"weeks"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Links map[string]interface{} `json:"_links"`
	Body string `json:"body"` // The text of the review.
	Id int `json:"id"` // Unique identifier of the review
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Commit_id string `json:"commit_id"` // A commit SHA for the review. If the commit object was garbage collected or forcibly deleted, then it no longer exists in Git and this value will be `null`.
	Pull_request_url string `json:"pull_request_url"`
	State string `json:"state"`
	Body_text string `json:"body_text,omitempty"`
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	User GeneratedType `json:"user"` // A GitHub user.
	Body_html string `json:"body_html,omitempty"`
	Submitted_at string `json:"submitted_at,omitempty"`
}

// Collaborator represents the Collaborator schema from the OpenAPI specification
type Collaborator struct {
	Gravatar_id string `json:"gravatar_id"`
	TypeField string `json:"type"`
	Login string `json:"login"`
	Organizations_url string `json:"organizations_url"`
	Role_name string `json:"role_name"`
	Gists_url string `json:"gists_url"`
	Node_id string `json:"node_id"`
	Starred_url string `json:"starred_url"`
	Html_url string `json:"html_url"`
	Followers_url string `json:"followers_url"`
	Following_url string `json:"following_url"`
	Repos_url string `json:"repos_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Name string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Id int `json:"id"`
	Received_events_url string `json:"received_events_url"`
	Site_admin bool `json:"site_admin"`
	Events_url string `json:"events_url"`
	Url string `json:"url"`
	Subscriptions_url string `json:"subscriptions_url"`
	Avatar_url string `json:"avatar_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Requester interface{} `json:"requester,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Owned_private_repos int `json:"owned_private_repos"`
	Business_plus bool `json:"business_plus,omitempty"`
	Gists_url string `json:"gists_url"`
	Site_admin bool `json:"site_admin"`
	Url string `json:"url"`
	Ldap_dn string `json:"ldap_dn,omitempty"`
	Following int `json:"following"`
	Created_at string `json:"created_at"`
	Company string `json:"company"`
	Events_url string `json:"events_url"`
	Id int `json:"id"`
	Public_gists int `json:"public_gists"`
	Organizations_url string `json:"organizations_url"`
	Public_repos int `json:"public_repos"`
	Email string `json:"email"`
	Avatar_url string `json:"avatar_url"`
	TypeField string `json:"type"`
	Blog string `json:"blog"`
	Disk_usage int `json:"disk_usage"`
	Received_events_url string `json:"received_events_url"`
	Updated_at string `json:"updated_at"`
	Twitter_username string `json:"twitter_username,omitempty"`
	Gravatar_id string `json:"gravatar_id"`
	Subscriptions_url string `json:"subscriptions_url"`
	Login string `json:"login"`
	Followers int `json:"followers"`
	Private_gists int `json:"private_gists"`
	Name string `json:"name"`
	Location string `json:"location"`
	Followers_url string `json:"followers_url"`
	Suspended_at string `json:"suspended_at,omitempty"`
	Following_url string `json:"following_url"`
	Plan map[string]interface{} `json:"plan,omitempty"`
	Total_private_repos int `json:"total_private_repos"`
	Collaborators int `json:"collaborators"`
	Repos_url string `json:"repos_url"`
	Html_url string `json:"html_url"`
	Starred_url string `json:"starred_url"`
	Bio string `json:"bio"`
	Hireable bool `json:"hireable"`
	Two_factor_authentication bool `json:"two_factor_authentication"`
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	ErrorField map[string]interface{} `json:"error"`
	Pusher GeneratedType `json:"pusher"` // A GitHub user.
	Status string `json:"status"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Commit string `json:"commit"`
	Created_at string `json:"created_at"`
	Duration int `json:"duration"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Old_answer map[string]interface{} `json:"old_answer"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Hook map[string]interface{} `json:"hook"` // The modified webhook. This will contain different keys based on the type of webhook it is: repository, organization, business, app, or GitHub Marketplace.
	Hook_id int `json:"hook_id"` // The id of the modified webhook.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action,omitempty"`
	Alert GeneratedType `json:"alert"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Location GeneratedType `json:"location"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project map[string]interface{} `json:"project"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Start_column float64 `json:"start_column"` // The column at which the secret starts within the start line when the file is interpreted as 8BIT ASCII
	Start_line float64 `json:"start_line"` // Line number at which the secret starts in the file
	Blob_sha string `json:"blob_sha"` // SHA-1 hash ID of the associated blob
	Commit_sha string `json:"commit_sha"` // SHA-1 hash ID of the associated commit
	End_column float64 `json:"end_column"` // The column at which the secret ends within the end line when the file is interpreted as 8BIT ASCII
	Blob_url string `json:"blob_url"` // The API URL to get the associated blob resource
	Commit_url string `json:"commit_url"` // The API URL to get the associated commit resource
	End_line float64 `json:"end_line"` // Line number at which the secret ends in the file
	Path string `json:"path"` // The file path in the repository
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Current_user_can_approve bool `json:"current_user_can_approve"` // Whether the currently authenticated user can approve the deployment
	Environment map[string]interface{} `json:"environment"`
	Reviewers []map[string]interface{} `json:"reviewers"` // The people or teams that may approve jobs that reference the environment. You can list up to six users or teams as reviewers. The reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	Wait_timer int `json:"wait_timer"` // The set duration of the wait timer
	Wait_timer_started_at string `json:"wait_timer_started_at"` // The time that the wait timer began.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id,omitempty"` // The unique identifier of the branch policy.
	Name string `json:"name,omitempty"` // The name pattern that branches must match in order to deploy to the environment.
	Node_id string `json:"node_id,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion interface{} `json:"discussion"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Limit string `json:"limit"` // The type of GitHub user that can comment, open issues, or create pull requests while the interaction limit is in effect.
	Expiry string `json:"expiry,omitempty"` // The duration of the interaction restriction. Default: `one_day`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project_card map[string]interface{} `json:"project_card"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Action string `json:"action"`
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Path string `json:"path"` // Path of the workflow file
	Ref string `json:"ref"` // Ref at which the workflow file will be selected
	Scope string `json:"scope"` // Scope of the required workflow
	State string `json:"state"` // State of the required workflow
	Created_at string `json:"created_at"`
	Id float64 `json:"id"` // Unique identifier for a required workflow
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"`
	Updated_at string `json:"updated_at"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Name string `json:"name"` // Name present in the workflow file
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository GeneratedType `json:"repository"` // Full Repository
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Verified bool `json:"verified"`
	Created_at string `json:"created_at"`
	Read_only bool `json:"read_only"`
	Id int `json:"id"`
	Key string `json:"key"`
	Title string `json:"title"`
	Url string `json:"url"`
	Added_by string `json:"added_by,omitempty"`
	Last_used string `json:"last_used,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Dismissal_restrictions map[string]interface{} `json:"dismissal_restrictions,omitempty"`
	Require_code_owner_reviews bool `json:"require_code_owner_reviews"`
	Require_last_push_approval bool `json:"require_last_push_approval,omitempty"` // Whether the most recent push must be approved by someone other than the person who pushed it.
	Required_approving_review_count int `json:"required_approving_review_count,omitempty"`
	Url string `json:"url,omitempty"`
	Bypass_pull_request_allowances map[string]interface{} `json:"bypass_pull_request_allowances,omitempty"` // Allow specific users, teams, or apps to bypass pull request requirements.
	Dismiss_stale_reviews bool `json:"dismiss_stale_reviews"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Start_column int `json:"start_column,omitempty"`
	Start_line int `json:"start_line,omitempty"`
	End_column int `json:"end_column,omitempty"`
	End_line int `json:"end_line,omitempty"`
	Path string `json:"path,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Label map[string]interface{} `json:"label,omitempty"`
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Action string `json:"action"`
	Pull_request map[string]interface{} `json:"pull_request"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url,omitempty"` // The URL to which the payloads will be delivered.
	Content_type string `json:"content_type,omitempty"` // The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`.
	Insecure_ssl GeneratedType `json:"insecure_ssl,omitempty"`
	Secret string `json:"secret,omitempty"` // If provided, the `secret` will be used as the `key` to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/webhooks/event-payloads/#delivery-headers).
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // Name of the project column
	Node_id string `json:"node_id"`
	Project_url string `json:"project_url"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Cards_url string `json:"cards_url"`
	Created_at string `json:"created_at"`
	Id int `json:"id"` // The unique identifier of the project column
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Security_advisory map[string]interface{} `json:"security_advisory"` // The details of the security advisory, including summary, description, and severity.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Membership map[string]interface{} `json:"membership"` // The membership between the user and the organization. Not present when the action is `member_invited`.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Workflow map[string]interface{} `json:"workflow"`
	Workflow_run interface{} `json:"workflow_run"`
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// Codespace represents the Codespace schema from the OpenAPI specification
type Codespace struct {
	Git_status map[string]interface{} `json:"git_status"` // Details about the codespace's git repository.
	Environment_id string `json:"environment_id"` // UUID identifying this codespace's environment.
	Machines_url string `json:"machines_url"` // API URL to access available alternate machine types for this codespace.
	Idle_timeout_notice string `json:"idle_timeout_notice,omitempty"` // Text to show user when codespace idle timeout minutes has been overriden by an organization policy
	Stop_url string `json:"stop_url"` // API URL to stop this codespace.
	Idle_timeout_minutes int `json:"idle_timeout_minutes"` // The number of minutes of inactivity after which this codespace will be automatically stopped.
	Last_used_at string `json:"last_used_at"` // Last known time this codespace was started.
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Runtime_constraints map[string]interface{} `json:"runtime_constraints,omitempty"`
	Devcontainer_path string `json:"devcontainer_path,omitempty"` // Path to devcontainer.json from repo root used to create Codespace.
	Recent_folders []string `json:"recent_folders"`
	Publish_url string `json:"publish_url,omitempty"` // API URL to publish this codespace to a new repository.
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Pending_operation bool `json:"pending_operation,omitempty"` // Whether or not a codespace has a pending async operation. This would mean that the codespace is temporarily unavailable. The only thing that you can do with a codespace in this state is delete it.
	Web_url string `json:"web_url"` // URL to access this codespace on the web.
	Pending_operation_disabled_reason string `json:"pending_operation_disabled_reason,omitempty"` // Text to show user when codespace is disabled by a pending operation
	Machine GeneratedType `json:"machine"` // A description of the machine powering a codespace.
	Pulls_url string `json:"pulls_url"` // API URL for the Pull Request associated with this codespace, if any.
	Location string `json:"location"` // The Azure region where this codespace is located.
	Id int `json:"id"`
	Prebuild bool `json:"prebuild"` // Whether the codespace was created from a prebuild.
	Name string `json:"name"` // Automatically generated name of this codespace.
	Last_known_stop_notice string `json:"last_known_stop_notice,omitempty"` // The text to display to a user when a codespace has been stopped for a potentially actionable reason.
	State string `json:"state"` // State of this codespace.
	Display_name string `json:"display_name,omitempty"` // Display name for this codespace.
	Retention_expires_at string `json:"retention_expires_at,omitempty"` // When a codespace will be auto-deleted based on the "retention_period_minutes" and "last_used_at"
	Billable_owner GeneratedType `json:"billable_owner"` // A GitHub user.
	Retention_period_minutes int `json:"retention_period_minutes,omitempty"` // Duration in minutes after codespace has gone idle in which it will be deleted. Must be integer minutes between 0 and 43200 (30 days).
	Url string `json:"url"` // API URL for this codespace.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Start_url string `json:"start_url"` // API URL to start this codespace.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Registry_package map[string]interface{} `json:"registry_package"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Html_url string `json:"html_url"`
	Id int `json:"id"` // Unique identifier of the team
	Node_id string `json:"node_id"`
	Permission string `json:"permission"` // Permission that the team will have for its repositories
	Privacy string `json:"privacy,omitempty"` // The level of privacy this team should have
	Description string `json:"description"` // Description of the team
	Ldap_dn string `json:"ldap_dn,omitempty"` // Distinguished Name (DN) that team maps to within LDAP environment
	Repositories_url string `json:"repositories_url"`
	Slug string `json:"slug"`
	Members_url string `json:"members_url"`
	Name string `json:"name"` // Name of the team
	Url string `json:"url"` // URL for the team
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Label map[string]interface{} `json:"label"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Confirm_delete_url string `json:"confirm_delete_url"` // Next deletable analysis in chain, with last analysis deletion confirmation
	Next_analysis_url string `json:"next_analysis_url"` // Next deletable analysis in chain, without last analysis deletion confirmation
}

// Manifest represents the Manifest schema from the OpenAPI specification
type Manifest struct {
	File map[string]interface{} `json:"file,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"` // User-defined metadata to store domain-specific information limited to 8 keys with scalar values.
	Name string `json:"name"` // The name of the manifest.
	Resolved map[string]interface{} `json:"resolved,omitempty"` // A collection of resolved package dependencies.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Action string `json:"action"`
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Comment map[string]interface{} `json:"comment"` // The [comment](https://docs.github.com/rest/reference/pulls#comments) itself.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Strict bool `json:"strict"`
	Url string `json:"url"`
	Checks []map[string]interface{} `json:"checks"`
	Contexts []string `json:"contexts"`
	Contexts_url string `json:"contexts_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Event string `json:"event,omitempty"`
	Message string `json:"message"` // Message describing the purpose of the commit
	Committer map[string]interface{} `json:"committer"` // Identifying information for the git-user
	Html_url string `json:"html_url"`
	Parents []map[string]interface{} `json:"parents"`
	Sha string `json:"sha"` // SHA for the commit
	Tree map[string]interface{} `json:"tree"`
	Verification map[string]interface{} `json:"verification"`
	Node_id string `json:"node_id"`
	Author map[string]interface{} `json:"author"` // Identifying information for the git-user
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Number int `json:"number"` // The security alert number.
	Dependency map[string]interface{} `json:"dependency"` // Details for the vulnerable dependency.
	Dismissed_reason string `json:"dismissed_reason"` // The reason that the alert was dismissed.
	Fixed_at string `json:"fixed_at"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	State string `json:"state"` // The state of the Dependabot alert.
	Url string `json:"url"` // The REST API URL of the alert resource.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Repository GeneratedType `json:"repository"` // A GitHub repository.
	Security_vulnerability GeneratedType `json:"security_vulnerability"` // Details pertaining to one vulnerable version range for the advisory.
	Updated_at string `json:"updated_at"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Security_advisory GeneratedType `json:"security_advisory"` // Details for the GitHub Security Advisory.
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
	Dismissed_comment string `json:"dismissed_comment"` // An optional comment associated with the alert's dismissal.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Workflow_job interface{} `json:"workflow_job"`
	Action string `json:"action"`
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue_title_url string `json:"issue_title_url"` // The API URL to get the issue where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"`
	Url string `json:"url"`
	Body string `json:"body,omitempty"`
	Html_url string `json:"html_url"`
	Key string `json:"key"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Url string `json:"url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Id int `json:"id"`
	Commit_id string `json:"commit_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Prebuild_availability string `json:"prebuild_availability"` // Whether a prebuild is currently available when creating a codespace for this machine and repository. If a branch was not specified as a ref, the default branch will be assumed. Value will be "null" if prebuilds are not supported or prebuild availability could not be determined. Value will be "none" if no prebuild is available. Latest values "ready" and "in_progress" indicate the prebuild availability status.
	Storage_in_bytes int `json:"storage_in_bytes"` // How much storage is available to the codespace.
	Cpus int `json:"cpus"` // How many cores are available to the codespace.
	Display_name string `json:"display_name"` // The display name of the machine includes cores, memory, and storage.
	Memory_in_bytes int `json:"memory_in_bytes"` // How much memory is available to the codespace.
	Name string `json:"name"` // The name of the machine.
	Operating_system string `json:"operating_system"` // The operating system of the machine.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"` // The ID of the installation.
	Node_id string `json:"node_id"` // The global node ID of the installation.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Open_issues int `json:"open_issues"`
	Description string `json:"description"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Blobs_url string `json:"blobs_url"`
	Updated_at string `json:"updated_at"`
	Clone_url string `json:"clone_url"`
	Events_url string `json:"events_url"`
	Id int `json:"id"` // Unique identifier of the repository
	Url string `json:"url"`
	Subscription_url string `json:"subscription_url"`
	Open_issues_count int `json:"open_issues_count"`
	Commits_url string `json:"commits_url"`
	Branches_url string `json:"branches_url"`
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Mirror_url string `json:"mirror_url"`
	Downloads_url string `json:"downloads_url"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Private bool `json:"private"` // Whether the repository is private or public.
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Pulls_url string `json:"pulls_url"`
	Trees_url string `json:"trees_url"`
	Compare_url string `json:"compare_url"`
	Languages_url string `json:"languages_url"`
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property has been deprecated. Please use `squash_merge_commit_title` instead.
	Comments_url string `json:"comments_url"`
	Has_pages bool `json:"has_pages"`
	Archive_url string `json:"archive_url"`
	Tags_url string `json:"tags_url"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Issue_comment_url string `json:"issue_comment_url"`
	Homepage string `json:"homepage"`
	Hooks_url string `json:"hooks_url"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Issue_events_url string `json:"issue_events_url"`
	Pushed_at string `json:"pushed_at"`
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Fork bool `json:"fork"`
	Keys_url string `json:"keys_url"`
	Template_repository map[string]interface{} `json:"template_repository,omitempty"`
	Full_name string `json:"full_name"`
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Network_count int `json:"network_count,omitempty"`
	Ssh_url string `json:"ssh_url"`
	License GeneratedType `json:"license"` // License Simple
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Git_url string `json:"git_url"`
	Teams_url string `json:"teams_url"`
	Stargazers_count int `json:"stargazers_count"`
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Name string `json:"name"` // The name of the repository.
	Topics []string `json:"topics,omitempty"`
	Language string `json:"language"`
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Size int `json:"size"` // The size of the repository. Size is calculated hourly. When a repository is initially created, the size is 0.
	Forks int `json:"forks"`
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Forks_count int `json:"forks_count"`
	Git_commits_url string `json:"git_commits_url"`
	Contents_url string `json:"contents_url"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is enabled for this repository
	Statuses_url string `json:"statuses_url"`
	Collaborators_url string `json:"collaborators_url"`
	Svn_url string `json:"svn_url"`
	Html_url string `json:"html_url"`
	Has_discussions bool `json:"has_discussions,omitempty"` // Whether discussions are enabled.
	Releases_url string `json:"releases_url"`
	Created_at string `json:"created_at"`
	Merges_url string `json:"merges_url"`
	Labels_url string `json:"labels_url"`
	Node_id string `json:"node_id"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Starred_at string `json:"starred_at,omitempty"`
	Milestones_url string `json:"milestones_url"`
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Forks_url string `json:"forks_url"`
	Assignees_url string `json:"assignees_url"`
	Issues_url string `json:"issues_url"`
	Git_tags_url string `json:"git_tags_url"`
	Watchers int `json:"watchers"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub user.
	Watchers_count int `json:"watchers_count"`
	Master_branch string `json:"master_branch,omitempty"`
	Subscribers_url string `json:"subscribers_url"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Git_refs_url string `json:"git_refs_url"`
	Notifications_url string `json:"notifications_url"`
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether or not a pull request head branch that is behind its base branch can always be updated even if it is not required to be up to date before merging.
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Contributors_url string `json:"contributors_url"`
	Archived bool `json:"archived"` // Whether the repository is archived.
	Deployments_url string `json:"deployments_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Stargazers_url string `json:"stargazers_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Key string `json:"key"`
	Name string `json:"name"`
	Url string `json:"url"`
	Html_url string `json:"html_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Comments []GeneratedType `json:"comments,omitempty"`
	Event string `json:"event,omitempty"`
	Node_id string `json:"node_id,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment map[string]interface{} `json:"comment"` // The [comment](https://docs.github.com/rest/reference/issues#comments) itself.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) the comment belongs to.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Target_url string `json:"target_url"`
	Updated_at string `json:"updated_at"`
	Avatar_url string `json:"avatar_url"`
	Url string `json:"url"`
	Description string `json:"description"`
	Context string `json:"context"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Required bool `json:"required,omitempty"`
	State string `json:"state"`
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Size int `json:"size"`
	Url string `json:"url"`
	Browser_download_url string `json:"browser_download_url"`
	Download_count int `json:"download_count"`
	Label string `json:"label"`
	Uploader GeneratedType `json:"uploader"` // A GitHub user.
	Name string `json:"name"` // The file name of the asset.
	Node_id string `json:"node_id"`
	State string `json:"state"` // State of the release asset.
	Updated_at string `json:"updated_at"`
	Content_type string `json:"content_type"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Actions_meta map[string]interface{} `json:"actions_meta,omitempty"`
	Check_suite map[string]interface{} `json:"check_suite"` // The [check_suite](https://docs.github.com/rest/reference/checks#suites).
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Body_text string `json:"body_text,omitempty"`
	Body string `json:"body,omitempty"` // Contents of the issue comment
	Html_url string `json:"html_url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Event string `json:"event"`
	Updated_at string `json:"updated_at"`
	Id int `json:"id"` // Unique identifier of the issue comment
	Url string `json:"url"` // URL for the issue comment
	Issue_url string `json:"issue_url"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body_html string `json:"body_html,omitempty"`
	User GeneratedType `json:"user"` // A GitHub user.
	Created_at string `json:"created_at"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Head_branch string `json:"head_branch"`
	Pull_requests []GeneratedType `json:"pull_requests"`
	App GeneratedType `json:"app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Before string `json:"before"`
	Head_commit GeneratedType `json:"head_commit"` // A commit.
	Node_id string `json:"node_id"`
	Head_sha string `json:"head_sha"` // The SHA of the head commit that is being checked.
	Runs_rerequestable bool `json:"runs_rerequestable,omitempty"`
	Status string `json:"status"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Latest_check_runs_count int `json:"latest_check_runs_count"`
	Rerequestable bool `json:"rerequestable,omitempty"`
	After string `json:"after"`
	Conclusion string `json:"conclusion"`
	Check_runs_url string `json:"check_runs_url"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Selected_actions_url string `json:"selected_actions_url,omitempty"` // The API URL to use to get or set the actions and reusable workflows that are allowed to run, when `allowed_actions` is set to `selected`.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"` // The API URL to use to get or set the selected repositories that are allowed to run GitHub Actions, when `enabled_repositories` is set to `selected`.
	Allowed_actions string `json:"allowed_actions,omitempty"` // The permissions policy that controls the actions and reusable workflows that are allowed to run.
	Enabled_repositories string `json:"enabled_repositories"` // The policy that controls the repositories in the organization that are allowed to run GitHub Actions.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue interface{} `json:"issue"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Content map[string]interface{} `json:"content"`
	Commit map[string]interface{} `json:"commit"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment map[string]interface{} `json:"comment"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Html_url string `json:"html_url,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Name string `json:"name"` // The name of the package version.
	Package_html_url string `json:"package_html_url"`
	Created_at string `json:"created_at"`
	Description string `json:"description,omitempty"`
	License string `json:"license,omitempty"`
	Deleted_at string `json:"deleted_at,omitempty"`
	Id int `json:"id"` // Unique identifier of the package version.
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Label map[string]interface{} `json:"label"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"` // The changes to the label if the action was `edited`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository_url string `json:"repository_url"`
	Created_at string `json:"created_at"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Description string `json:"description"` // A short description of the status.
	State string `json:"state"` // The state of the status.
	Deployment_url string `json:"deployment_url"`
	Log_url string `json:"log_url,omitempty"` // The URL to associate with this status.
	Target_url string `json:"target_url"` // Deprecated: the URL to associate with this status.
	Url string `json:"url"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Environment string `json:"environment,omitempty"` // The environment of the deployment that the status is for.
	Environment_url string `json:"environment_url,omitempty"` // The URL for accessing your environment.
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the secret_scanning_alert_location.created JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Security_advisory map[string]interface{} `json:"security_advisory"` // The details of the security advisory, including summary, description, and severity.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_id string `json:"commit_id"`
	Url string `json:"url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Node_id string `json:"node_id"`
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Status string `json:"status"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Team map[string]interface{} `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue interface{} `json:"issue"`
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// Snapshot represents the Snapshot schema from the OpenAPI specification
type Snapshot struct {
	Metadata Metadata `json:"metadata,omitempty"` // User-defined metadata to store domain-specific information limited to 8 keys with scalar values.
	Ref string `json:"ref"` // The repository branch that triggered this snapshot.
	Scanned string `json:"scanned"` // The time at which the snapshot was scanned.
	Sha string `json:"sha"` // The commit SHA associated with this dependency snapshot. Maximum length: 40 characters.
	Version int `json:"version"` // The version of the repository snapshot submission.
	Detector map[string]interface{} `json:"detector"` // A description of the detector used.
	Job map[string]interface{} `json:"job"`
	Manifests map[string]interface{} `json:"manifests,omitempty"` // A collection of package manifests, which are a collection of related dependencies declared in a file or representing a logical group of dependencies.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Git_push_url string `json:"git_push_url"`
	Owner GeneratedType `json:"owner,omitempty"` // A GitHub user.
	Truncated bool `json:"truncated,omitempty"`
	Files map[string]interface{} `json:"files"`
	Forks []interface{} `json:"forks,omitempty"`
	Public bool `json:"public"`
	Commits_url string `json:"commits_url"`
	Created_at string `json:"created_at"`
	Description string `json:"description"`
	Updated_at string `json:"updated_at"`
	Html_url string `json:"html_url"`
	Comments_url string `json:"comments_url"`
	History []interface{} `json:"history,omitempty"`
	Forks_url string `json:"forks_url"`
	Url string `json:"url"`
	Git_pull_url string `json:"git_pull_url"`
	Comments int `json:"comments"`
	User GeneratedType `json:"user"` // A GitHub user.
	Id string `json:"id"`
}

// Team represents the Team schema from the OpenAPI specification
type Team struct {
	Name string `json:"name"`
	Parent GeneratedType `json:"parent"` // Groups of organization members that gives permissions on specified repositories.
	Permission string `json:"permission"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Members_url string `json:"members_url"`
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Privacy string `json:"privacy,omitempty"`
	Repositories_url string `json:"repositories_url"`
	Slug string `json:"slug"`
	Url string `json:"url"`
	Description string `json:"description"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Can_approve_pull_request_reviews bool `json:"can_approve_pull_request_reviews"` // Whether GitHub Actions can approve pull requests. Enabling this can be a security risk.
	Default_workflow_permissions string `json:"default_workflow_permissions"` // The default workflow permissions granted to the GITHUB_TOKEN when running workflows.
}

// Job represents the Job schema from the OpenAPI specification
type Job struct {
	Conclusion string `json:"conclusion"` // The outcome of the job.
	Runner_group_name string `json:"runner_group_name"` // The name of the runner group to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Html_url string `json:"html_url"`
	Labels []string `json:"labels"` // Labels for the workflow job. Specified by the "runs_on" attribute in the action's workflow file.
	Node_id string `json:"node_id"`
	Run_id int `json:"run_id"` // The id of the associated workflow run.
	Started_at string `json:"started_at"` // The time that the job started, in ISO 8601 format.
	Run_attempt int `json:"run_attempt,omitempty"` // Attempt number of the associated workflow run, 1 for first attempt and higher if the workflow was re-run.
	Check_run_url string `json:"check_run_url"`
	Status string `json:"status"` // The phase of the lifecycle that the job is currently in.
	Runner_group_id int `json:"runner_group_id"` // The ID of the runner group to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Completed_at string `json:"completed_at"` // The time that the job finished, in ISO 8601 format.
	Steps []map[string]interface{} `json:"steps,omitempty"` // Steps in this job.
	Runner_id int `json:"runner_id"` // The ID of the runner to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Head_sha string `json:"head_sha"` // The SHA of the commit that is being run.
	Run_url string `json:"run_url"`
	Workflow_name string `json:"workflow_name"` // The name of the workflow.
	Name string `json:"name"` // The name of the job.
	Id int `json:"id"` // The id of the job.
	Runner_name string `json:"runner_name"` // The name of the runner to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Head_branch string `json:"head_branch"` // The name of the current branch.
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Html_url string `json:"html_url"`
	Key string `json:"key"`
	Name string `json:"name"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Completed_at string `json:"completed_at"`
	Head_sha string `json:"head_sha"` // The SHA of the commit that is being checked.
	Status string `json:"status"` // The phase of the lifecycle that the check is currently in.
	Deployment GeneratedType `json:"deployment,omitempty"` // A deployment created as the result of an Actions check run from a workflow that references an environment
	External_id string `json:"external_id"`
	Pull_requests []GeneratedType `json:"pull_requests"`
	App GeneratedType `json:"app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Started_at string `json:"started_at"`
	Node_id string `json:"node_id"`
	Check_suite map[string]interface{} `json:"check_suite"`
	Html_url string `json:"html_url"`
	Details_url string `json:"details_url"`
	Id int `json:"id"` // The id of the check.
	Name string `json:"name"` // The name of the check.
	Url string `json:"url"`
	Conclusion string `json:"conclusion"`
	Output map[string]interface{} `json:"output"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Files map[string]interface{} `json:"files"`
	Health_percentage int `json:"health_percentage"`
	Updated_at string `json:"updated_at"`
	Content_reports_enabled bool `json:"content_reports_enabled,omitempty"`
	Description string `json:"description"`
	Documentation string `json:"documentation"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Milestone Milestone `json:"milestone,omitempty"` // A collection of related issues and pull requests.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Members_allowed_repository_creation_type string `json:"members_allowed_repository_creation_type,omitempty"`
	Hooks_url string `json:"hooks_url"`
	Login string `json:"login"`
	Company string `json:"company,omitempty"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Description string `json:"description"`
	Two_factor_requirement_enabled bool `json:"two_factor_requirement_enabled,omitempty"`
	Members_can_create_private_pages bool `json:"members_can_create_private_pages,omitempty"`
	Followers int `json:"followers"`
	Members_can_create_public_pages bool `json:"members_can_create_public_pages,omitempty"`
	Plan map[string]interface{} `json:"plan,omitempty"`
	Members_can_create_pages bool `json:"members_can_create_pages,omitempty"`
	Owned_private_repos int `json:"owned_private_repos,omitempty"`
	Avatar_url string `json:"avatar_url"`
	Twitter_username string `json:"twitter_username,omitempty"`
	TypeField string `json:"type"`
	Public_repos int `json:"public_repos"`
	Location string `json:"location,omitempty"`
	Url string `json:"url"`
	Blog string `json:"blog,omitempty"`
	Default_repository_permission string `json:"default_repository_permission,omitempty"`
	Created_at string `json:"created_at"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Issues_url string `json:"issues_url"`
	Repos_url string `json:"repos_url"`
	Members_can_create_repositories bool `json:"members_can_create_repositories,omitempty"`
	Members_can_create_public_repositories bool `json:"members_can_create_public_repositories,omitempty"`
	Members_url string `json:"members_url"`
	Has_repository_projects bool `json:"has_repository_projects"`
	Billing_email string `json:"billing_email,omitempty"`
	Public_members_url string `json:"public_members_url"`
	Html_url string `json:"html_url"`
	Email string `json:"email,omitempty"`
	Name string `json:"name,omitempty"`
	Events_url string `json:"events_url"`
	Members_can_fork_private_repositories bool `json:"members_can_fork_private_repositories,omitempty"`
	Is_verified bool `json:"is_verified,omitempty"`
	Public_gists int `json:"public_gists"`
	Following int `json:"following"`
	Private_gists int `json:"private_gists,omitempty"`
	Total_private_repos int `json:"total_private_repos,omitempty"`
	Collaborators int `json:"collaborators,omitempty"`
	Updated_at string `json:"updated_at"`
	Members_can_create_private_repositories bool `json:"members_can_create_private_repositories,omitempty"`
	Members_can_create_internal_repositories bool `json:"members_can_create_internal_repositories,omitempty"`
	Disk_usage int `json:"disk_usage,omitempty"`
	Has_organization_projects bool `json:"has_organization_projects"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Change_status map[string]interface{} `json:"change_status"`
	Committed_at string `json:"committed_at"`
	Url string `json:"url"`
	User GeneratedType `json:"user"` // A GitHub user.
	Version string `json:"version"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project_card map[string]interface{} `json:"project_card"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Thread map[string]interface{} `json:"thread"`
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Effective_date string `json:"effective_date"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Marketplace_purchase interface{} `json:"marketplace_purchase"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Previous_marketplace_purchase map[string]interface{} `json:"previous_marketplace_purchase,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Pull_request map[string]interface{} `json:"pull_request"`
	Action string `json:"action"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Number int `json:"number"`
	Reason string `json:"reason"`
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Permission string `json:"permission"`
	Role_name string `json:"role_name"`
	User GeneratedType `json:"user"` // Collaborator
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Avatar_url string `json:"avatar_url"`
	Hooks_url string `json:"hooks_url"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Events_url string `json:"events_url"`
	Issues_url string `json:"issues_url"`
	Members_url string `json:"members_url"`
	Repos_url string `json:"repos_url"`
	Id int `json:"id"`
	Login string `json:"login"`
	Public_members_url string `json:"public_members_url"`
	Description string `json:"description"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Commit_id string `json:"commit_id"`
	Event string `json:"event"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Commit_url string `json:"commit_url"`
	Dismissed_review map[string]interface{} `json:"dismissed_review"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Path string `json:"path"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Sha string `json:"sha"`
	Git_url string `json:"git_url"`
	Language string `json:"language,omitempty"`
	Name string `json:"name"`
	Url string `json:"url"`
	File_size int `json:"file_size,omitempty"`
	Html_url string `json:"html_url"`
	Score float64 `json:"score"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Last_modified_at string `json:"last_modified_at,omitempty"`
	Line_numbers []string `json:"line_numbers,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Idle_timeout_minutes int `json:"idle_timeout_minutes"` // The number of minutes of inactivity after which this codespace will be automatically stopped.
	Last_used_at string `json:"last_used_at"` // Last known time this codespace was started.
	Pending_operation_disabled_reason string `json:"pending_operation_disabled_reason,omitempty"` // Text to show user when codespace is disabled by a pending operation
	Git_status map[string]interface{} `json:"git_status"` // Details about the codespace's git repository.
	Devcontainer_path string `json:"devcontainer_path,omitempty"` // Path to devcontainer.json from repo root used to create Codespace.
	Created_at string `json:"created_at"`
	Prebuild bool `json:"prebuild"` // Whether the codespace was created from a prebuild.
	Stop_url string `json:"stop_url"` // API URL to stop this codespace.
	Name string `json:"name"` // Automatically generated name of this codespace.
	Idle_timeout_notice string `json:"idle_timeout_notice,omitempty"` // Text to show user when codespace idle timeout minutes has been overriden by an organization policy
	Retention_period_minutes int `json:"retention_period_minutes,omitempty"` // Duration in minutes after codespace has gone idle in which it will be deleted. Must be integer minutes between 0 and 43200 (30 days).
	Pending_operation bool `json:"pending_operation,omitempty"` // Whether or not a codespace has a pending async operation. This would mean that the codespace is temporarily unavailable. The only thing that you can do with a codespace in this state is delete it.
	Start_url string `json:"start_url"` // API URL to start this codespace.
	Updated_at string `json:"updated_at"`
	Location string `json:"location"` // The Azure region where this codespace is located.
	State string `json:"state"` // State of this codespace.
	Display_name string `json:"display_name,omitempty"` // Display name for this codespace.
	Machine GeneratedType `json:"machine"` // A description of the machine powering a codespace.
	Url string `json:"url"` // API URL for this codespace.
	Billable_owner GeneratedType `json:"billable_owner"` // A GitHub user.
	Publish_url string `json:"publish_url,omitempty"` // API URL to publish this codespace to a new repository.
	Retention_expires_at string `json:"retention_expires_at,omitempty"` // When a codespace will be auto-deleted based on the "retention_period_minutes" and "last_used_at"
	Recent_folders []string `json:"recent_folders"`
	Repository GeneratedType `json:"repository"` // Full Repository
	Runtime_constraints map[string]interface{} `json:"runtime_constraints,omitempty"`
	Web_url string `json:"web_url"` // URL to access this codespace on the web.
	Machines_url string `json:"machines_url"` // API URL to access available alternate machine types for this codespace.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Environment_id string `json:"environment_id"` // UUID identifying this codespace's environment.
	Pulls_url string `json:"pulls_url"` // API URL for the Pull Request associated with this codespace, if any.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
	Updated_at string `json:"updated_at"`
}

// Email represents the Email schema from the OpenAPI specification
type Email struct {
	Verified bool `json:"verified"`
	Visibility string `json:"visibility"`
	Email string `json:"email"`
	Primary bool `json:"primary"`
}

// Import represents the Import schema from the OpenAPI specification
type Import struct {
	Use_lfs bool `json:"use_lfs,omitempty"`
	Authors_url string `json:"authors_url"`
	Status string `json:"status"`
	Status_text string `json:"status_text,omitempty"`
	Message string `json:"message,omitempty"`
	Project_choices []map[string]interface{} `json:"project_choices,omitempty"`
	Repository_url string `json:"repository_url"`
	Tfvc_project string `json:"tfvc_project,omitempty"`
	Large_files_count int `json:"large_files_count,omitempty"`
	Svn_root string `json:"svn_root,omitempty"`
	Failed_step string `json:"failed_step,omitempty"`
	Svc_root string `json:"svc_root,omitempty"`
	Vcs_url string `json:"vcs_url"` // The URL of the originating repository.
	Vcs string `json:"vcs"`
	Import_percent int `json:"import_percent,omitempty"`
	Error_message string `json:"error_message,omitempty"`
	Commit_count int `json:"commit_count,omitempty"`
	Large_files_size int `json:"large_files_size,omitempty"`
	Has_large_files bool `json:"has_large_files,omitempty"`
	Authors_count int `json:"authors_count,omitempty"`
	Url string `json:"url"`
	Html_url string `json:"html_url"`
	Push_percent int `json:"push_percent,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project_column map[string]interface{} `json:"project_column"`
}

// Metadata represents the Metadata schema from the OpenAPI specification
type Metadata struct {
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Following_url string `json:"following_url"`
	Avatar_url string `json:"avatar_url"`
	Email string `json:"email,omitempty"`
	Received_events_url string `json:"received_events_url"`
	Repos_url string `json:"repos_url"`
	Site_admin bool `json:"site_admin"`
	Html_url string `json:"html_url"`
	Login string `json:"login"`
	TypeField string `json:"type"`
	Gravatar_id string `json:"gravatar_id"`
	Followers_url string `json:"followers_url"`
	Starred_at string `json:"starred_at,omitempty"`
	Subscriptions_url string `json:"subscriptions_url"`
	Gists_url string `json:"gists_url"`
	Starred_url string `json:"starred_url"`
	Events_url string `json:"events_url"`
	Organizations_url string `json:"organizations_url"`
	Id int `json:"id"`
	Url string `json:"url"`
	Name string `json:"name,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Target_type string `json:"target_type"`
	Changes map[string]interface{} `json:"changes"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Account map[string]interface{} `json:"account"`
	Action string `json:"action"`
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Installation GeneratedType `json:"installation"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Review map[string]interface{} `json:"review"` // The review that was affected.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Assignee map[string]interface{} `json:"assignee,omitempty"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"` // The action that was performed.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Protection_url string `json:"protection_url,omitempty"`
	Commit map[string]interface{} `json:"commit"`
	Name string `json:"name"`
	Protected bool `json:"protected"`
	Protection GeneratedType `json:"protection,omitempty"` // Branch Protection
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Action string `json:"action"`
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// Root represents the Root schema from the OpenAPI specification
type Root struct {
	Gists_url string `json:"gists_url"`
	Commit_search_url string `json:"commit_search_url"`
	User_url string `json:"user_url"`
	Events_url string `json:"events_url"`
	Organization_teams_url string `json:"organization_teams_url"`
	Notifications_url string `json:"notifications_url"`
	User_organizations_url string `json:"user_organizations_url"`
	Issue_search_url string `json:"issue_search_url"`
	User_repositories_url string `json:"user_repositories_url"`
	Current_user_authorizations_html_url string `json:"current_user_authorizations_html_url"`
	Label_search_url string `json:"label_search_url"`
	Issues_url string `json:"issues_url"`
	Authorizations_url string `json:"authorizations_url"`
	Code_search_url string `json:"code_search_url"`
	Starred_gists_url string `json:"starred_gists_url"`
	Followers_url string `json:"followers_url"`
	Rate_limit_url string `json:"rate_limit_url"`
	Keys_url string `json:"keys_url"`
	Emails_url string `json:"emails_url"`
	Topic_search_url string `json:"topic_search_url,omitempty"`
	Organization_url string `json:"organization_url"`
	Current_user_repositories_url string `json:"current_user_repositories_url"`
	Public_gists_url string `json:"public_gists_url"`
	User_search_url string `json:"user_search_url"`
	Repository_search_url string `json:"repository_search_url"`
	Organization_repositories_url string `json:"organization_repositories_url"`
	Current_user_url string `json:"current_user_url"`
	Repository_url string `json:"repository_url"`
	Feeds_url string `json:"feeds_url"`
	Hub_url string `json:"hub_url"`
	Starred_url string `json:"starred_url"`
	Emojis_url string `json:"emojis_url"`
	Following_url string `json:"following_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Tree_id string `json:"tree_id"`
	Author map[string]interface{} `json:"author"`
	Committer map[string]interface{} `json:"committer"`
	Id string `json:"id"`
	Message string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert interface{} `json:"alert"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ref string `json:"ref"` // The [`git ref`](https://docs.github.com/rest/reference/git#get-a-reference) resource.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Description string `json:"description"` // The repository's current description.
	Pusher_type string `json:"pusher_type"` // The pusher type for the event. Can be either `user` or a deploy key.
	Ref_type string `json:"ref_type"` // The type of Git ref object created in the repository.
	Master_branch string `json:"master_branch"` // The name of the repository's default branch (usually `main`).
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes"` // The changes to the milestone if the action was `edited`.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"` // The action that was performed.
	Assignee map[string]interface{} `json:"assignee,omitempty"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Author_association string `json:"author_association,omitempty"` // How the author is associated with the repository.
	Dismissed_review GeneratedType `json:"dismissed_review,omitempty"`
	Issue GeneratedType `json:"issue,omitempty"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Commit_id string `json:"commit_id"`
	Url string `json:"url"`
	Label GeneratedType `json:"label,omitempty"` // Issue Event Label
	Milestone GeneratedType `json:"milestone,omitempty"` // Issue Event Milestone
	Requested_team Team `json:"requested_team,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Commit_url string `json:"commit_url"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Created_at string `json:"created_at"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Project_card GeneratedType `json:"project_card,omitempty"` // Issue Event Project Card
	Assigner GeneratedType `json:"assigner,omitempty"` // A GitHub user.
	Event string `json:"event"`
	Rename GeneratedType `json:"rename,omitempty"` // Issue Event Rename
	Assignee GeneratedType `json:"assignee,omitempty"` // A GitHub user.
	Review_requester GeneratedType `json:"review_requester,omitempty"` // A GitHub user.
	Requested_reviewer GeneratedType `json:"requested_reviewer,omitempty"` // A GitHub user.
	Lock_reason string `json:"lock_reason,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Dismissal_commit_id string `json:"dismissal_commit_id,omitempty"`
	Dismissal_message string `json:"dismissal_message"`
	Review_id int `json:"review_id"`
	State string `json:"state"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Verifiable_password_authentication bool `json:"verifiable_password_authentication"`
	Actions []string `json:"actions,omitempty"`
	Dependabot []string `json:"dependabot,omitempty"`
	Hooks []string `json:"hooks,omitempty"`
	Git []string `json:"git,omitempty"`
	Ssh_key_fingerprints map[string]interface{} `json:"ssh_key_fingerprints,omitempty"`
	Importer []string `json:"importer,omitempty"`
	Web []string `json:"web,omitempty"`
	Api []string `json:"api,omitempty"`
	Packages []string `json:"packages,omitempty"`
	Pages []string `json:"pages,omitempty"`
	Ssh_keys []string `json:"ssh_keys,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Html_url string `json:"html_url"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Html_url string `json:"html_url"`
	Details_url string `json:"details_url"`
	Pull_requests []GeneratedType `json:"pull_requests"`
	Url string `json:"url"`
	Output map[string]interface{} `json:"output"`
	App GeneratedType `json:"app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	External_id string `json:"external_id"`
	Head_sha string `json:"head_sha"` // The SHA of the commit that is being checked.
	Id int `json:"id"` // The id of the check.
	Check_suite GeneratedType `json:"check_suite"` // A suite of checks performed on the code of a given code change
	Deployment GeneratedType `json:"deployment,omitempty"` // A deployment created as the result of an Actions check run from a workflow that references an environment
	Conclusion string `json:"conclusion"`
	Completed_at string `json:"completed_at"`
	Name string `json:"name"` // The name of the check.
	Started_at string `json:"started_at"`
	Status string `json:"status"` // The phase of the lifecycle that the check is currently in.
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Html_url string `json:"html_url"` // HTML URL for the pull request review comment.
	Id int `json:"id"` // The ID of the pull request review comment.
	Pull_request_review_id int `json:"pull_request_review_id"` // The ID of the pull request review to which the comment belongs.
	Original_commit_id string `json:"original_commit_id"` // The SHA of the original commit to which the comment applies.
	Original_position int `json:"original_position"` // The index of the original line in the diff to which the comment applies. This field is deprecated; use `original_line` instead.
	Diff_hunk string `json:"diff_hunk"` // The diff of the line that the comment refers to.
	In_reply_to_id int `json:"in_reply_to_id,omitempty"` // The comment ID to reply to.
	Line int `json:"line,omitempty"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Position int `json:"position"` // The line index in the diff to which the comment applies. This field is deprecated; use `line` instead.
	Reactions GeneratedType `json:"reactions,omitempty"`
	Side string `json:"side,omitempty"` // The side of the diff to which the comment applies. The side of the last line of the range for a multi-line comment
	Commit_id string `json:"commit_id"` // The SHA of the commit to which the comment applies.
	Original_start_line int `json:"original_start_line,omitempty"` // The first line of the range for a multi-line comment.
	Start_side string `json:"start_side,omitempty"` // The side of the first line of the range for a multi-line comment.
	Body_text string `json:"body_text,omitempty"`
	Pull_request_url string `json:"pull_request_url"` // URL for the pull request that the review comment belongs to.
	Original_line int `json:"original_line,omitempty"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Start_line int `json:"start_line,omitempty"` // The first line of the range for a multi-line comment.
	Url string `json:"url"` // URL for the pull request review comment
	Created_at string `json:"created_at"`
	Path string `json:"path"` // The relative path of the file to which the comment applies.
	Updated_at string `json:"updated_at"`
	Links map[string]interface{} `json:"_links"`
	Node_id string `json:"node_id"` // The node ID of the pull request review comment.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	User GeneratedType `json:"user"` // A GitHub user.
	Body_html string `json:"body_html,omitempty"`
	Body string `json:"body"` // The text of the comment.
}

// Verification represents the Verification schema from the OpenAPI specification
type Verification struct {
	Payload string `json:"payload"`
	Reason string `json:"reason"`
	Signature string `json:"signature"`
	Verified bool `json:"verified"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	DefaultField bool `json:"default"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Color string `json:"color"`
	Score float64 `json:"score"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Description string `json:"description"`
	Name string `json:"name"`
}

// Authorization represents the Authorization schema from the OpenAPI specification
type Authorization struct {
	Id int `json:"id"`
	Scopes []string `json:"scopes"` // A list of scopes that this authorization is in.
	User GeneratedType `json:"user,omitempty"` // A GitHub user.
	Fingerprint string `json:"fingerprint"`
	Installation GeneratedType `json:"installation,omitempty"`
	Url string `json:"url"`
	Hashed_token string `json:"hashed_token"`
	Note_url string `json:"note_url"`
	Token string `json:"token"`
	Updated_at string `json:"updated_at"`
	App map[string]interface{} `json:"app"`
	Note string `json:"note"`
	Token_last_eight string `json:"token_last_eight"`
	Created_at string `json:"created_at"`
	Expires_at string `json:"expires_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert interface{} `json:"alert"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Comment map[string]interface{} `json:"comment"` // The [comment](https://docs.github.com/rest/reference/pulls#comments) itself.
	Pull_request map[string]interface{} `json:"pull_request"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"` // The changes to the comment.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes interface{} `json:"changes,omitempty"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Previous_marketplace_purchase map[string]interface{} `json:"previous_marketplace_purchase,omitempty"`
	Marketplace_purchase interface{} `json:"marketplace_purchase"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Effective_date string `json:"effective_date"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Resolution_comment string `json:"resolution_comment,omitempty"` // An optional comment to resolve an alert.
	Resolution string `json:"resolution,omitempty"` // **Required when the `state` is `resolved`.** The reason for resolving the alert.
	Resolved_at string `json:"resolved_at,omitempty"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Number int `json:"number,omitempty"` // The security alert number.
	Push_protection_bypassed bool `json:"push_protection_bypassed,omitempty"` // Whether push protection was bypassed for the detected secret.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Created_at string `json:"created_at,omitempty"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Locations_url string `json:"locations_url,omitempty"` // The REST API URL of the code locations for this alert.
	Push_protection_bypassed_at string `json:"push_protection_bypassed_at,omitempty"` // The time that push protection was bypassed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Push_protection_bypassed_by GeneratedType `json:"push_protection_bypassed_by,omitempty"` // A GitHub user.
	Secret string `json:"secret,omitempty"` // The secret that was detected.
	Secret_type string `json:"secret_type,omitempty"` // The type of secret that secret scanning detected.
	Resolved_by GeneratedType `json:"resolved_by,omitempty"` // A GitHub user.
	State string `json:"state,omitempty"` // Sets the state of the secret scanning alert. You must provide `resolution` when you set the state to `resolved`.
	Secret_type_display_name string `json:"secret_type_display_name,omitempty"` // User-friendly name for the detected secret, matching the `secret_type`. For a list of built-in patterns, see "[Secret scanning patterns](https://docs.github.com/code-security/secret-scanning/secret-scanning-patterns#supported-secrets-for-advanced-security)."
	Url string `json:"url,omitempty"` // The REST API URL of the alert resource.
	Html_url string `json:"html_url,omitempty"` // The GitHub URL of the alert resource.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ecosystem string `json:"ecosystem"` // The package's language or package management ecosystem.
	Name string `json:"name"` // The unique package name within its ecosystem.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Node_id string `json:"node_id"`
	Performed_via_github_app Integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
	Commit_id string `json:"commit_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Tree_id string `json:"tree_id"`
	Author map[string]interface{} `json:"author"`
	Committer map[string]interface{} `json:"committer"`
	Id string `json:"id"`
	Message string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Release interface{} `json:"release"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url,omitempty"`
	Checks []map[string]interface{} `json:"checks"`
	Contexts []string `json:"contexts"`
	Contexts_url string `json:"contexts_url,omitempty"`
	Enforcement_level string `json:"enforcement_level,omitempty"`
	Strict bool `json:"strict,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Github_owned_allowed bool `json:"github_owned_allowed,omitempty"` // Whether GitHub-owned actions are allowed. For example, this includes the actions in the `actions` organization.
	Patterns_allowed []string `json:"patterns_allowed,omitempty"` // Specifies a list of string-matching patterns to allow specific action(s) and reusable workflow(s). Wildcards, tags, and SHAs are allowed. For example, `monalisa/octocat@*`, `monalisa/octocat@v2`, `monalisa/*`. **Note**: The `patterns_allowed` setting only applies to public repositories.
	Verified_allowed bool `json:"verified_allowed,omitempty"` // Whether actions from GitHub Marketplace verified creators are allowed. Set to `true` to allow all actions by GitHub Marketplace verified creators.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Label map[string]interface{} `json:"label"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.created JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Requester interface{} `json:"requester,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation Installation `json:"installation"` // Installation
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Branch string `json:"branch"`
	Path string `json:"path"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Reason string `json:"reason"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Number int `json:"number"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Pull_request map[string]interface{} `json:"pull_request"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Node_id string `json:"node_id"`
	Id int `json:"id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
	Commit_id string `json:"commit_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Review map[string]interface{} `json:"review"` // The review that was affected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"`
	Description string `json:"description"`
	Organization GeneratedType `json:"organization"` // Team Organization
	Privacy string `json:"privacy,omitempty"` // The level of privacy this team should have
	Node_id string `json:"node_id"`
	Members_count int `json:"members_count"`
	Repositories_url string `json:"repositories_url"`
	Html_url string `json:"html_url"`
	Id int `json:"id"` // Unique identifier of the team
	Permission string `json:"permission"` // Permission that the team will have for its repositories
	Slug string `json:"slug"`
	Members_url string `json:"members_url"`
	Created_at string `json:"created_at"`
	Ldap_dn string `json:"ldap_dn,omitempty"` // Distinguished Name (DN) that team maps to within LDAP environment
	Repos_count int `json:"repos_count"`
	Name string `json:"name"` // Name of the team
	Parent GeneratedType `json:"parent,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Url string `json:"url"` // URL for the team
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"` // The action performed. Can be `created`.
	Comment map[string]interface{} `json:"comment"` // The [commit comment](https://docs.github.com/rest/reference/repos#get-a-commit-comment) resource.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Workflow map[string]interface{} `json:"workflow,omitempty"`
	Workflow_run map[string]interface{} `json:"workflow_run,omitempty"`
	Deployment_status map[string]interface{} `json:"deployment_status"` // The [deployment status](https://docs.github.com/rest/reference/deployments#list-deployment-statuses).
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Action string `json:"action"`
	Deployment map[string]interface{} `json:"deployment"` // The [deployment](https://docs.github.com/rest/reference/deployments#list-deployments).
	Check_run map[string]interface{} `json:"check_run,omitempty"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Security_advisory map[string]interface{} `json:"security_advisory"` // The details of the security advisory, including summary, description, and severity.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Before string `json:"before"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	After string `json:"after"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Action string `json:"action"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Effective_date string `json:"effective_date,omitempty"` // The `pending_cancellation` and `pending_tier_change` event types will include the date the cancellation or tier change will take effect.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Sponsorship map[string]interface{} `json:"sponsorship"`
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Sponsorship map[string]interface{} `json:"sponsorship"`
}

// Topic represents the Topic schema from the OpenAPI specification
type Topic struct {
	Names []string `json:"names"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Starred_at interface{} `json:"starred_at"` // The time the star was created. This is a timestamp in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`. Will be `null` for the `deleted` action.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Key string `json:"key"`
	Title string `json:"title"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Id int `json:"id"`
	Commit_url string `json:"commit_url"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Assigner GeneratedType `json:"assigner"` // A GitHub user.
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Commit_id string `json:"commit_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow_job map[string]interface{} `json:"workflow_job"`
	Action string `json:"action"`
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Can_approve_pull_request_reviews bool `json:"can_approve_pull_request_reviews,omitempty"` // Whether GitHub Actions can approve pull requests. Enabling this can be a security risk.
	Default_workflow_permissions string `json:"default_workflow_permissions,omitempty"` // The default workflow permissions granted to the GITHUB_TOKEN when running workflows.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Alert map[string]interface{} `json:"alert"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Import_url string `json:"import_url"`
	Name string `json:"name"`
	Remote_id string `json:"remote_id"`
	Remote_name string `json:"remote_name"`
	Url string `json:"url"`
	Email string `json:"email"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow_job map[string]interface{} `json:"workflow_job"`
	Action string `json:"action"`
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Label map[string]interface{} `json:"label,omitempty"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"` // The changes to the issue.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Secret_scanning map[string]interface{} `json:"secret_scanning,omitempty"`
	Secret_scanning_push_protection map[string]interface{} `json:"secret_scanning_push_protection,omitempty"`
	Advanced_security map[string]interface{} `json:"advanced_security,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow map[string]interface{} `json:"workflow"`
	Workflow_run interface{} `json:"workflow_run"`
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Dismissed_comment string `json:"dismissed_comment,omitempty"` // The dismissal comment associated with the dismissal of the alert.
	State string `json:"state"` // State of a code scanning alert.
	Most_recent_instance GeneratedType `json:"most_recent_instance"`
	Rule GeneratedType `json:"rule"`
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Url string `json:"url"` // The REST API URL of the alert resource.
	Dismissed_reason string `json:"dismissed_reason"` // **Required when the state is dismissed.** The reason for dismissing or closing the alert.
	Fixed_at string `json:"fixed_at,omitempty"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Tool GeneratedType `json:"tool"`
	Instances_url string `json:"instances_url"` // The REST API URL for fetching the list of instances for an alert.
	Number int `json:"number"` // The security alert number.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Path string `json:"path"`
	Name string `json:"name"`
	State string `json:"state"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Node_id string `json:"node_id"`
	Source_repository GeneratedType `json:"source_repository"` // Minimal Repository
	Badge_url string `json:"badge_url"`
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Errors []map[string]interface{} `json:"errors"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Following_url string `json:"following_url"`
	Followers int `json:"followers"`
	Gists_url string `json:"gists_url"`
	Hireable bool `json:"hireable"`
	Name string `json:"name"`
	Blog string `json:"blog"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Following int `json:"following"`
	Location string `json:"location"`
	Gravatar_id string `json:"gravatar_id"`
	Company string `json:"company"`
	Events_url string `json:"events_url"`
	Private_gists int `json:"private_gists,omitempty"`
	Public_repos int `json:"public_repos"`
	Updated_at string `json:"updated_at"`
	Repos_url string `json:"repos_url"`
	Suspended_at string `json:"suspended_at,omitempty"`
	Public_gists int `json:"public_gists"`
	Starred_url string `json:"starred_url"`
	Twitter_username string `json:"twitter_username,omitempty"`
	Total_private_repos int `json:"total_private_repos,omitempty"`
	Subscriptions_url string `json:"subscriptions_url"`
	TypeField string `json:"type"`
	Avatar_url string `json:"avatar_url"`
	Bio string `json:"bio"`
	Html_url string `json:"html_url"`
	Organizations_url string `json:"organizations_url"`
	Site_admin bool `json:"site_admin"`
	Email string `json:"email"`
	Login string `json:"login"`
	Disk_usage int `json:"disk_usage,omitempty"`
	Plan map[string]interface{} `json:"plan,omitempty"`
	Url string `json:"url"`
	Collaborators int `json:"collaborators,omitempty"`
	Owned_private_repos int `json:"owned_private_repos,omitempty"`
	Received_events_url string `json:"received_events_url"`
	Followers_url string `json:"followers_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Referrer string `json:"referrer"`
	Uniques int `json:"uniques"`
	Count int `json:"count"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Key string `json:"key"`
	Id int `json:"id"`
}

// Installation represents the Installation schema from the OpenAPI specification
type Installation struct {
	Id int `json:"id"` // The ID of the installation.
	Suspended_by GeneratedType `json:"suspended_by"` // A GitHub user.
	Target_type string `json:"target_type"`
	Created_at string `json:"created_at"`
	App_id int `json:"app_id"`
	Single_file_name string `json:"single_file_name"`
	Contact_email string `json:"contact_email,omitempty"`
	Events []string `json:"events"`
	Permissions GeneratedType `json:"permissions"` // The permissions granted to the user-to-server access token.
	Repositories_url string `json:"repositories_url"`
	Updated_at string `json:"updated_at"`
	Account interface{} `json:"account"`
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Single_file_paths []string `json:"single_file_paths,omitempty"`
	Suspended_at string `json:"suspended_at"`
	App_slug string `json:"app_slug"`
	Has_multiple_single_files bool `json:"has_multiple_single_files,omitempty"`
	Access_tokens_url string `json:"access_tokens_url"`
	Html_url string `json:"html_url"`
	Target_id int `json:"target_id"` // The ID of the user or organization this token is being scoped to.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Spdx_id string `json:"spdx_id"`
	Url string `json:"url"`
	Html_url string `json:"html_url,omitempty"`
	Key string `json:"key"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Task string `json:"task"` // Parameter to specify a task to execute
	Transient_environment bool `json:"transient_environment,omitempty"` // Specifies if the given environment is will no longer exist at some point in the future. Default: false.
	Environment string `json:"environment"` // Name for the target deployment environment.
	Id int `json:"id"` // Unique identifier of the deployment
	Node_id string `json:"node_id"`
	Original_environment string `json:"original_environment,omitempty"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Repository_url string `json:"repository_url"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Description string `json:"description"`
	Production_environment bool `json:"production_environment,omitempty"` // Specifies if the given environment is one that end-users directly interact with. Default: false.
	Statuses_url string `json:"statuses_url"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Delivered_at string `json:"delivered_at"` // Time when the delivery was delivered.
	Duration float64 `json:"duration"` // Time spent delivering.
	Id int `json:"id"` // Unique identifier of the delivery.
	Status_code int `json:"status_code"` // Status code received when delivery was made.
	Url string `json:"url,omitempty"` // The URL target of the delivery.
	Event string `json:"event"` // The event that triggered the delivery.
	Guid string `json:"guid"` // Unique identifier for the event (shared with all deliveries for all webhooks that subscribe to this event).
	Response map[string]interface{} `json:"response"`
	Status string `json:"status"` // Description of the status of the attempted delivery
	Action string `json:"action"` // The type of activity for the event that triggered the delivery.
	Redelivery bool `json:"redelivery"` // Whether the delivery is a redelivery.
	Request map[string]interface{} `json:"request"`
	Installation_id int `json:"installation_id"` // The id of the GitHub App installation associated with this event.
	Repository_id int `json:"repository_id"` // The id of the repository associated with this event.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment map[string]interface{} `json:"comment"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	State string `json:"state"`
	Statuses []GeneratedType `json:"statuses"`
	Total_count int `json:"total_count"`
	Url string `json:"url"`
	Commit_url string `json:"commit_url"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Sha string `json:"sha"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Previous_filename string `json:"previous_filename,omitempty"`
	Status string `json:"status"`
	Contents_url string `json:"contents_url"`
	Deletions int `json:"deletions"`
	Raw_url string `json:"raw_url"`
	Sha string `json:"sha"`
	Additions int `json:"additions"`
	Filename string `json:"filename"`
	Patch string `json:"patch,omitempty"`
	Blob_url string `json:"blob_url"`
	Changes int `json:"changes"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Expires_at string `json:"expires_at"`
	Has_multiple_single_files bool `json:"has_multiple_single_files,omitempty"`
	Permissions GeneratedType `json:"permissions,omitempty"` // The permissions granted to the user-to-server access token.
	Repositories []Repository `json:"repositories,omitempty"`
	Repository_selection string `json:"repository_selection,omitempty"`
	Single_file string `json:"single_file,omitempty"`
	Single_file_paths []string `json:"single_file_paths,omitempty"`
	Token string `json:"token"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Value string `json:"value"` // The value of the variable.
	Created_at string `json:"created_at"` // The date and time at which the variable was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Name string `json:"name"` // The name of the variable.
	Updated_at string `json:"updated_at"` // The date and time at which the variable was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Label map[string]interface{} `json:"label,omitempty"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue_comment_url string `json:"issue_comment_url"` // The API URL to get the issue comment where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The name pattern that branches must match in order to deploy to the environment. Wildcard characters will not match `/`. For example, to match branches that begin with `release/` and contain an additional single slash, use `release/*/*`. For more information about pattern matching syntax, see the [Ruby File.fnmatch documentation](https://ruby-doc.org/core-2.5.1/File.html#method-c-fnmatch).
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Summary string `json:"summary"` // A short, plain text summary of the advisory.
	Withdrawn_at string `json:"withdrawn_at"` // The time that the advisory was withdrawn in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Description string `json:"description"` // A long-form Markdown-supported description of the advisory.
	Cve_id string `json:"cve_id"` // The unique CVE ID assigned to the advisory.
	Cvss map[string]interface{} `json:"cvss"` // Details for the advisory pertaining to the Common Vulnerability Scoring System.
	Cwes []map[string]interface{} `json:"cwes"` // Details for the advisory pertaining to Common Weakness Enumeration.
	Ghsa_id string `json:"ghsa_id"` // The unique GitHub Security Advisory ID assigned to the advisory.
	Identifiers []map[string]interface{} `json:"identifiers"` // Values that identify this advisory among security information sources.
	References []map[string]interface{} `json:"references"` // Links to additional advisory information.
	Severity string `json:"severity"` // The severity of the advisory.
	Updated_at string `json:"updated_at"` // The time that the advisory was last modified in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Vulnerabilities []GeneratedType `json:"vulnerabilities"` // Vulnerable version range information for the advisory.
	Published_at string `json:"published_at"` // The time that the advisory was published in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField string `json:"type"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Active bool `json:"active"`
	Config map[string]interface{} `json:"config"`
	Events []string `json:"events"`
	Name string `json:"name"`
	Id int `json:"id"`
	Created_at string `json:"created_at"`
	Deliveries_url string `json:"deliveries_url,omitempty"`
	Ping_url string `json:"ping_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Marketplace_purchase map[string]interface{} `json:"marketplace_purchase"`
	Organization_billing_email string `json:"organization_billing_email,omitempty"`
	TypeField string `json:"type"`
	Url string `json:"url"`
	Email string `json:"email,omitempty"`
	Id int `json:"id"`
	Login string `json:"login"`
	Marketplace_pending_change map[string]interface{} `json:"marketplace_pending_change,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Subscription_url string `json:"subscription_url"`
	Compare_url string `json:"compare_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Hooks_url string `json:"hooks_url"`
	Private bool `json:"private"`
	Milestones_url string `json:"milestones_url"`
	Id int `json:"id"`
	Contents_url string `json:"contents_url"`
	Html_url string `json:"html_url"`
	Svn_url string `json:"svn_url,omitempty"`
	Pushed_at string `json:"pushed_at,omitempty"`
	Subscribers_url string `json:"subscribers_url"`
	Has_discussions bool `json:"has_discussions,omitempty"`
	Ssh_url string `json:"ssh_url,omitempty"`
	Mirror_url string `json:"mirror_url,omitempty"`
	Watchers_count int `json:"watchers_count,omitempty"`
	Pulls_url string `json:"pulls_url"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Created_at string `json:"created_at,omitempty"`
	Open_issues_count int `json:"open_issues_count,omitempty"`
	Topics []string `json:"topics,omitempty"`
	Stargazers_url string `json:"stargazers_url"`
	Network_count int `json:"network_count,omitempty"`
	Archive_url string `json:"archive_url"`
	Forks int `json:"forks,omitempty"`
	Open_issues int `json:"open_issues,omitempty"`
	Node_id string `json:"node_id"`
	Has_issues bool `json:"has_issues,omitempty"`
	Languages_url string `json:"languages_url"`
	License map[string]interface{} `json:"license,omitempty"`
	Downloads_url string `json:"downloads_url"`
	Role_name string `json:"role_name,omitempty"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Forks_url string `json:"forks_url"`
	Git_refs_url string `json:"git_refs_url"`
	Url string `json:"url"`
	Merges_url string `json:"merges_url"`
	Commits_url string `json:"commits_url"`
	Stargazers_count int `json:"stargazers_count,omitempty"`
	Deployments_url string `json:"deployments_url"`
	Assignees_url string `json:"assignees_url"`
	Language string `json:"language,omitempty"`
	Has_pages bool `json:"has_pages,omitempty"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Trees_url string `json:"trees_url"`
	Statuses_url string `json:"statuses_url"`
	Has_projects bool `json:"has_projects,omitempty"`
	Name string `json:"name"`
	Is_template bool `json:"is_template,omitempty"`
	Keys_url string `json:"keys_url"`
	Has_wiki bool `json:"has_wiki,omitempty"`
	Homepage string `json:"homepage,omitempty"`
	Git_commits_url string `json:"git_commits_url"`
	Has_downloads bool `json:"has_downloads,omitempty"`
	Security_and_analysis GeneratedType `json:"security_and_analysis,omitempty"`
	Forks_count int `json:"forks_count,omitempty"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Description string `json:"description"`
	Issue_comment_url string `json:"issue_comment_url"`
	Comments_url string `json:"comments_url"`
	Git_url string `json:"git_url,omitempty"`
	Visibility string `json:"visibility,omitempty"`
	Issues_url string `json:"issues_url"`
	Full_name string `json:"full_name"`
	Blobs_url string `json:"blobs_url"`
	Releases_url string `json:"releases_url"`
	Issue_events_url string `json:"issue_events_url"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Watchers int `json:"watchers,omitempty"`
	Contributors_url string `json:"contributors_url"`
	Size int `json:"size,omitempty"` // The size of the repository. Size is calculated hourly. When a repository is initially created, the size is 0.
	Updated_at string `json:"updated_at,omitempty"`
	Teams_url string `json:"teams_url"`
	Labels_url string `json:"labels_url"`
	Fork bool `json:"fork"`
	Tags_url string `json:"tags_url"`
	Clone_url string `json:"clone_url,omitempty"`
	Notifications_url string `json:"notifications_url"`
	Archived bool `json:"archived,omitempty"`
	Disabled bool `json:"disabled,omitempty"`
	Code_of_conduct GeneratedType `json:"code_of_conduct,omitempty"` // Code Of Conduct
	Branches_url string `json:"branches_url"`
	Collaborators_url string `json:"collaborators_url"`
	Git_tags_url string `json:"git_tags_url"`
	Default_branch string `json:"default_branch,omitempty"`
	Allow_forking bool `json:"allow_forking,omitempty"`
	Events_url string `json:"events_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Estimated_paid_storage_for_month int `json:"estimated_paid_storage_for_month"` // Estimated storage space (GB) used in billing cycle.
	Estimated_storage_for_month int `json:"estimated_storage_for_month"` // Estimated sum of free and paid storage space (GB) used in billing cycle.
	Days_left_in_billing_cycle int `json:"days_left_in_billing_cycle"` // Numbers of days left in billing cycle.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Url string `json:"url"`
	Requested_reviewers []GeneratedType `json:"requested_reviewers,omitempty"`
	Changed_files int `json:"changed_files"`
	Diff_url string `json:"diff_url"`
	Merged_by GeneratedType `json:"merged_by"` // A GitHub user.
	Review_comments int `json:"review_comments"`
	Comments int `json:"comments"`
	Review_comment_url string `json:"review_comment_url"`
	Locked bool `json:"locked"`
	Auto_merge GeneratedType `json:"auto_merge"` // The status of auto merging a pull request.
	Links map[string]interface{} `json:"_links"`
	Base map[string]interface{} `json:"base"`
	Body string `json:"body"`
	Draft bool `json:"draft,omitempty"` // Indicates whether or not the pull request is a draft.
	Labels []map[string]interface{} `json:"labels"`
	Deletions int `json:"deletions"`
	Rebaseable bool `json:"rebaseable,omitempty"`
	Statuses_url string `json:"statuses_url"`
	Node_id string `json:"node_id"`
	Merged_at string `json:"merged_at"`
	Updated_at string `json:"updated_at"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Comments_url string `json:"comments_url"`
	Html_url string `json:"html_url"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Additions int `json:"additions"`
	Closed_at string `json:"closed_at"`
	Review_comments_url string `json:"review_comments_url"`
	Title string `json:"title"` // The title of the pull request.
	Head map[string]interface{} `json:"head"`
	Requested_teams []GeneratedType `json:"requested_teams,omitempty"`
	Mergeable bool `json:"mergeable"`
	Merged bool `json:"merged"`
	State string `json:"state"` // State of this Pull Request. Either `open` or `closed`.
	Created_at string `json:"created_at"`
	User GeneratedType `json:"user"` // A GitHub user.
	Number int `json:"number"` // Number uniquely identifying the pull request within its repository.
	Commits int `json:"commits"`
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Maintainer_can_modify bool `json:"maintainer_can_modify"` // Indicates whether maintainers can modify the pull request.
	Id int `json:"id"`
	Commits_url string `json:"commits_url"`
	Issue_url string `json:"issue_url"`
	Patch_url string `json:"patch_url"`
	Mergeable_state string `json:"mergeable_state"`
	Merge_commit_sha string `json:"merge_commit_sha"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Release map[string]interface{} `json:"release"` // The [release](https://docs.github.com/rest/reference/repos/#get-a-release) object.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Projects_v2 GeneratedType `json:"projects_v2"` // A projects v2 project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Organization GeneratedType `json:"organization"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Ignored bool `json:"ignored"` // Determines if all notifications should be blocked from this repository.
	Reason string `json:"reason"`
	Repository_url string `json:"repository_url"`
	Subscribed bool `json:"subscribed"` // Determines if notifications should be received from this repository.
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action,omitempty"`
	Check_run GeneratedType `json:"check_run"` // A check performed on the code of a given code change
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Changes map[string]interface{} `json:"changes"` // The changes to the comment if the action was `edited`.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Number int `json:"number"` // The pull request number.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sha string `json:"sha"`
	Tree []map[string]interface{} `json:"tree"` // Objects specifying a tree structure
	Truncated bool `json:"truncated"`
	Url string `json:"url"`
}

// Runner represents the Runner schema from the OpenAPI specification
type Runner struct {
	Labels []GeneratedType `json:"labels"`
	Name string `json:"name"` // The name of the runner.
	Os string `json:"os"` // The Operating System of the runner.
	Status string `json:"status"` // The status of the runner.
	Busy bool `json:"busy"`
	Id int `json:"id"` // The id of the runner.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id,omitempty"` // Unique identifier of the label.
	Name string `json:"name"` // Name of the label.
	TypeField string `json:"type,omitempty"` // The type of label. Read-only labels are applied automatically when the runner is configured.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	After string `json:"after,omitempty"`
	Conclusion string `json:"conclusion,omitempty"`
	Head_sha string `json:"head_sha,omitempty"` // The SHA of the head commit that is being checked.
	Updated_at string `json:"updated_at,omitempty"`
	App Integration `json:"app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Before string `json:"before,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Id int `json:"id,omitempty"`
	Pull_requests []GeneratedType `json:"pull_requests,omitempty"`
	Url string `json:"url,omitempty"`
	Repository GeneratedType `json:"repository,omitempty"` // Minimal Repository
	Status string `json:"status,omitempty"`
	Head_branch string `json:"head_branch,omitempty"`
	Node_id string `json:"node_id,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Milestone Milestone `json:"milestone,omitempty"` // A collection of related issues and pull requests.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Sponsorship map[string]interface{} `json:"sponsorship"`
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	User GeneratedType `json:"user"` // A GitHub user.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body string `json:"body"` // The comment text.
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Key string `json:"key"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
	Spdx_id string `json:"spdx_id"`
	Url string `json:"url"`
	Html_url string `json:"html_url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Thread map[string]interface{} `json:"thread"`
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// Contributor represents the Contributor schema from the OpenAPI specification
type Contributor struct {
	Subscriptions_url string `json:"subscriptions_url,omitempty"`
	Id int `json:"id,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Received_events_url string `json:"received_events_url,omitempty"`
	Starred_url string `json:"starred_url,omitempty"`
	Site_admin bool `json:"site_admin,omitempty"`
	Url string `json:"url,omitempty"`
	Avatar_url string `json:"avatar_url,omitempty"`
	Contributions int `json:"contributions"`
	Email string `json:"email,omitempty"`
	TypeField string `json:"type"`
	Login string `json:"login,omitempty"`
	Name string `json:"name,omitempty"`
	Organizations_url string `json:"organizations_url,omitempty"`
	Followers_url string `json:"followers_url,omitempty"`
	Repos_url string `json:"repos_url,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Gravatar_id string `json:"gravatar_id,omitempty"`
	Following_url string `json:"following_url,omitempty"`
	Gists_url string `json:"gists_url,omitempty"`
	Events_url string `json:"events_url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository_url string `json:"repository_url,omitempty"`
	Subscribed bool `json:"subscribed"`
	Thread_url string `json:"thread_url,omitempty"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Ignored bool `json:"ignored"`
	Reason string `json:"reason"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Privacy string `json:"privacy,omitempty"` // The level of privacy this team should have
	Repositories_url string `json:"repositories_url"`
	Slug string `json:"slug"`
	Ldap_dn string `json:"ldap_dn,omitempty"` // Distinguished Name (DN) that team maps to within LDAP environment
	Id int `json:"id"` // Unique identifier of the team
	Members_url string `json:"members_url"`
	Url string `json:"url"` // URL for the team
	Description string `json:"description"` // Description of the team
	Html_url string `json:"html_url"`
	Name string `json:"name"` // Name of the team
	Permission string `json:"permission"` // Permission that the team will have for its repositories
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Comment map[string]interface{} `json:"comment"` // The [comment](https://docs.github.com/rest/reference/issues#comments) itself.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) the comment belongs to.
	Changes map[string]interface{} `json:"changes"` // The changes to the comment.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	User map[string]interface{} `json:"user,omitempty"`
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Invitation map[string]interface{} `json:"invitation"` // The invitation for the user or email if the action is `member_invited`.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Email string `json:"email,omitempty"`
	Id int `json:"id"`
	Login string `json:"login"`
	Node_id string `json:"node_id,omitempty"`
	Organization_billing_email string `json:"organization_billing_email,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Billable map[string]interface{} `json:"billable"`
	Run_duration_ms int `json:"run_duration_ms,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Workflow map[string]interface{} `json:"workflow"`
	Deployment map[string]interface{} `json:"deployment"` // The [deployment](https://docs.github.com/rest/reference/deployments#list-deployments).
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Registry_package map[string]interface{} `json:"registry_package"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Custom_branch_policies bool `json:"custom_branch_policies"` // Whether only branches that match the specified name patterns can deploy to this environment. If `custom_branch_policies` is `true`, `protected_branches` must be `false`; if `custom_branch_policies` is `false`, `protected_branches` must be `true`.
	Protected_branches bool `json:"protected_branches"` // Whether only branches with branch protection rules can deploy to this environment. If `protected_branches` is `true`, `custom_branch_policies` must be `false`; if `protected_branches` is `false`, `custom_branch_policies` must be `true`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Resolution string `json:"resolution,omitempty"` // **Required when the `state` is `resolved`.** The reason for resolving the alert.
	Secret_type_display_name string `json:"secret_type_display_name,omitempty"` // User-friendly name for the detected secret, matching the `secret_type`. For a list of built-in patterns, see "[Secret scanning patterns](https://docs.github.com/code-security/secret-scanning/secret-scanning-patterns#supported-secrets-for-advanced-security)."
	Url string `json:"url,omitempty"` // The REST API URL of the alert resource.
	Repository GeneratedType `json:"repository,omitempty"` // A GitHub repository.
	Resolution_comment string `json:"resolution_comment,omitempty"` // The comment that was optionally added when this alert was closed
	Locations_url string `json:"locations_url,omitempty"` // The REST API URL of the code locations for this alert.
	Push_protection_bypassed_by GeneratedType `json:"push_protection_bypassed_by,omitempty"` // A GitHub user.
	Resolved_at string `json:"resolved_at,omitempty"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	State string `json:"state,omitempty"` // Sets the state of the secret scanning alert. You must provide `resolution` when you set the state to `resolved`.
	Html_url string `json:"html_url,omitempty"` // The GitHub URL of the alert resource.
	Secret_type string `json:"secret_type,omitempty"` // The type of secret that secret scanning detected.
	Number int `json:"number,omitempty"` // The security alert number.
	Created_at string `json:"created_at,omitempty"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Secret string `json:"secret,omitempty"` // The secret that was detected.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Push_protection_bypassed_at string `json:"push_protection_bypassed_at,omitempty"` // The time that push protection was bypassed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Push_protection_bypassed bool `json:"push_protection_bypassed,omitempty"` // Whether push protection was bypassed for the detected secret.
	Resolved_by GeneratedType `json:"resolved_by,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Selected_actions_url string `json:"selected_actions_url,omitempty"` // The API URL to use to get or set the actions and reusable workflows that are allowed to run, when `allowed_actions` is set to `selected`.
	Allowed_actions string `json:"allowed_actions,omitempty"` // The permissions policy that controls the actions and reusable workflows that are allowed to run.
	Enabled bool `json:"enabled"` // Whether GitHub Actions is enabled on the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Account GeneratedType `json:"account"`
	Billing_cycle string `json:"billing_cycle"`
	Free_trial_ends_on string `json:"free_trial_ends_on"`
	Next_billing_date string `json:"next_billing_date"`
	On_free_trial bool `json:"on_free_trial"`
	Plan GeneratedType `json:"plan"` // Marketplace Listing Plan
	Unit_count int `json:"unit_count"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Failed_at string `json:"failed_at,omitempty"`
	Failed_reason string `json:"failed_reason,omitempty"`
	Invitation_teams_url string `json:"invitation_teams_url"`
	Inviter GeneratedType `json:"inviter"` // A GitHub user.
	Login string `json:"login"`
	Team_count int `json:"team_count"`
	Id int `json:"id"`
	Invitation_source string `json:"invitation_source,omitempty"`
	Created_at string `json:"created_at"`
	Email string `json:"email"`
	Node_id string `json:"node_id"`
	Role string `json:"role"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Restrictions GeneratedType `json:"restrictions,omitempty"` // Branch Restriction Policy
	Required_linear_history map[string]interface{} `json:"required_linear_history,omitempty"`
	Required_pull_request_reviews GeneratedType `json:"required_pull_request_reviews,omitempty"` // Protected Branch Pull Request Review
	Required_status_checks GeneratedType `json:"required_status_checks,omitempty"` // Protected Branch Required Status Check
	Protection_url string `json:"protection_url,omitempty"`
	Allow_deletions map[string]interface{} `json:"allow_deletions,omitempty"`
	Required_signatures map[string]interface{} `json:"required_signatures,omitempty"`
	Allow_force_pushes map[string]interface{} `json:"allow_force_pushes,omitempty"`
	Block_creations map[string]interface{} `json:"block_creations,omitempty"`
	Allow_fork_syncing map[string]interface{} `json:"allow_fork_syncing,omitempty"` // Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow fork syncing. Set to `false` to prevent fork syncing.
	Name string `json:"name,omitempty"`
	Required_conversation_resolution map[string]interface{} `json:"required_conversation_resolution,omitempty"`
	Enforce_admins GeneratedType `json:"enforce_admins,omitempty"` // Protected Branch Admin Enforced
	Lock_branch map[string]interface{} `json:"lock_branch,omitempty"` // Whether to set the branch as read-only. If this is true, users will not be able to push to the branch.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Diff_url string `json:"diff_url"`
	Permalink_url string `json:"permalink_url"`
	Total_commits int `json:"total_commits"`
	Ahead_by int `json:"ahead_by"`
	Base_commit Commit `json:"base_commit"` // Commit
	Commits []Commit `json:"commits"`
	Html_url string `json:"html_url"`
	Merge_base_commit Commit `json:"merge_base_commit"` // Commit
	Patch_url string `json:"patch_url"`
	Status string `json:"status"`
	Url string `json:"url"`
	Behind_by int `json:"behind_by"`
	Files []GeneratedType `json:"files,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Source string `json:"source"`
	Name string `json:"name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Team map[string]interface{} `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"` // The changes to the team if the action was `edited`.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
}

// Project represents the Project schema from the OpenAPI specification
type Project struct {
	Url string `json:"url"`
	Private bool `json:"private,omitempty"` // Whether or not this project can be seen by everyone. Only present if owner is an organization.
	Owner_url string `json:"owner_url"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Html_url string `json:"html_url"`
	Name string `json:"name"` // Name of the project
	Created_at string `json:"created_at"`
	Columns_url string `json:"columns_url"`
	State string `json:"state"` // State of the project; either 'open' or 'closed'
	Number int `json:"number"`
	Updated_at string `json:"updated_at"`
	Node_id string `json:"node_id"`
	Body string `json:"body"` // Body of the project
	Id int `json:"id"`
	Organization_permission string `json:"organization_permission,omitempty"` // The baseline permission that all organization members have on this project. Only present if owner is an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Size int `json:"size"`
	Content string `json:"content"`
	Download_url string `json:"download_url"`
	TypeField string `json:"type"`
	Url string `json:"url"`
	Encoding string `json:"encoding"`
	Git_url string `json:"git_url"`
	License GeneratedType `json:"license"` // License Simple
	Links map[string]interface{} `json:"_links"`
	Html_url string `json:"html_url"`
	Name string `json:"name"`
	Path string `json:"path"`
	Sha string `json:"sha"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"` // If the action was `edited`, the changes to the rule.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Rule map[string]interface{} `json:"rule"` // The branch protection rule. Includes a `name` and all the [branch protection settings](https://docs.github.com/github/administering-a-repository/defining-the-mergeability-of-pull-requests/about-protected-branches#about-branch-protection-settings) applied to branches that match the name. Binary settings are boolean. Multi-level configurations are one of `off`, `non_admins`, or `everyone`. Actor and build lists are arrays of strings.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Include_claim_keys []string `json:"include_claim_keys"` // Array of unique strings. Each claim key can only contain alphanumeric characters and underscores.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Rocket int `json:"rocket"`
	Confused int `json:"confused"`
	Field1 int `json:"-1"`
	Heart int `json:"heart"`
	Total_count int `json:"total_count"`
	Field1 int `json:"+1"`
	Eyes int `json:"eyes"`
	Laugh int `json:"laugh"`
	Url string `json:"url"`
	Hooray int `json:"hooray"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Body string `json:"body"` // The main text of the discussion.
	Body_html string `json:"body_html"`
	Last_edited_at string `json:"last_edited_at"`
	Team_url string `json:"team_url"`
	Html_url string `json:"html_url"`
	Private bool `json:"private"` // Whether or not this discussion should be restricted to team members and organization administrators.
	Title string `json:"title"` // The title of the discussion.
	Updated_at string `json:"updated_at"`
	Author GeneratedType `json:"author"` // A GitHub user.
	Body_version string `json:"body_version"` // The current version of the body content. If provided, this update operation will be rejected if the given version does not match the latest version on the server.
	Number int `json:"number"` // The unique sequence number of a team discussion.
	Reactions GeneratedType `json:"reactions,omitempty"`
	Node_id string `json:"node_id"`
	Pinned bool `json:"pinned"` // Whether or not this discussion should be pinned for easy retrieval.
	Comments_url string `json:"comments_url"`
	Comments_count int `json:"comments_count"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ref string `json:"ref,omitempty"`
	Sha string `json:"sha"`
	Path string `json:"path"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Lock_branch map[string]interface{} `json:"lock_branch,omitempty"` // Whether to set the branch as read-only. If this is true, users will not be able to push to the branch.
	Allow_fork_syncing map[string]interface{} `json:"allow_fork_syncing,omitempty"` // Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow fork syncing. Set to `false` to prevent fork syncing.
	Required_linear_history map[string]interface{} `json:"required_linear_history,omitempty"`
	Restrictions GeneratedType `json:"restrictions,omitempty"` // Branch Restriction Policy
	Block_creations map[string]interface{} `json:"block_creations,omitempty"`
	Enforce_admins map[string]interface{} `json:"enforce_admins,omitempty"`
	Required_conversation_resolution map[string]interface{} `json:"required_conversation_resolution,omitempty"`
	Required_pull_request_reviews map[string]interface{} `json:"required_pull_request_reviews,omitempty"`
	Required_signatures map[string]interface{} `json:"required_signatures,omitempty"`
	Allow_deletions map[string]interface{} `json:"allow_deletions,omitempty"`
	Allow_force_pushes map[string]interface{} `json:"allow_force_pushes,omitempty"`
	Required_status_checks GeneratedType `json:"required_status_checks,omitempty"` // Status Check Policy
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Label map[string]interface{} `json:"label"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Marketplace_purchase interface{} `json:"marketplace_purchase"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Effective_date string `json:"effective_date"`
	Previous_marketplace_purchase map[string]interface{} `json:"previous_marketplace_purchase,omitempty"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Private bool `json:"private,omitempty"` // Whether the project is private or not. Only present when owner is an organization.
	Html_url string `json:"html_url"`
	Organization_permission string `json:"organization_permission,omitempty"` // The organization permission for this project. Only present when owner is an organization.
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Node_id string `json:"node_id"`
	Permissions map[string]interface{} `json:"permissions"`
	Body string `json:"body"`
	Id int `json:"id"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Columns_url string `json:"columns_url"`
	Owner_url string `json:"owner_url"`
	Name string `json:"name"`
	State string `json:"state"`
	Created_at string `json:"created_at"`
	Number int `json:"number"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Pull_request map[string]interface{} `json:"pull_request"`
	Review map[string]interface{} `json:"review"` // The review that was affected.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// Milestone represents the Milestone schema from the OpenAPI specification
type Milestone struct {
	Due_on string `json:"due_on"`
	Node_id string `json:"node_id"`
	State string `json:"state"` // The state of the milestone.
	Description string `json:"description"`
	Html_url string `json:"html_url"`
	Url string `json:"url"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Id int `json:"id"`
	Closed_issues int `json:"closed_issues"`
	Open_issues int `json:"open_issues"`
	Labels_url string `json:"labels_url"`
	Created_at string `json:"created_at"`
	Closed_at string `json:"closed_at"`
	Number int `json:"number"` // The number of the milestone.
	Title string `json:"title"` // The title of the milestone.
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Actor GeneratedType `json:"actor,omitempty"` // A GitHub user.
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Source map[string]interface{} `json:"source"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Key string `json:"key"` // The Base64 encoded public key.
	Key_id string `json:"key_id"` // The identifier for the key.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Detail string `json:"detail,omitempty"`
	Documentation_url string `json:"documentation_url,omitempty"`
	Message string `json:"message,omitempty"`
	Schemas []string `json:"schemas,omitempty"`
	Scimtype string `json:"scimType,omitempty"`
	Status int `json:"status,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request map[string]interface{} `json:"pull_request"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Number int `json:"number"` // The pull request number.
	Action string `json:"action"`
	Assignee map[string]interface{} `json:"assignee"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Labels_url string `json:"labels_url"`
	Updated_at string `json:"updated_at"`
	Closed_at string `json:"closed_at"`
	Title string `json:"title"` // The title of the milestone.
	Description string `json:"description"`
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
	Url string `json:"url"`
	Node_id string `json:"node_id"`
	Open_issues int `json:"open_issues"`
	Due_on string `json:"due_on"`
	Closed_issues int `json:"closed_issues"`
	State string `json:"state"` // The state of the milestone.
	Number int `json:"number"` // The number of the milestone.
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"` // URL for the tag
	Verification Verification `json:"verification,omitempty"`
	Message string `json:"message"` // Message describing the purpose of the tag
	Node_id string `json:"node_id"`
	Object map[string]interface{} `json:"object"`
	Sha string `json:"sha"`
	Tag string `json:"tag"` // Name of the tag
	Tagger map[string]interface{} `json:"tagger"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Role string `json:"role"` // The user's membership type in the organization.
	State string `json:"state"` // The state of the member in the organization. The `pending` state indicates the user has not yet accepted an invitation.
	Url string `json:"url"`
	User GeneratedType `json:"user"` // A GitHub user.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Organization_url string `json:"organization_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Id int `json:"id,omitempty"`
	Pattern string `json:"pattern"`
}

// Repository represents the Repository schema from the OpenAPI specification
type Repository struct {
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub user.
	Keys_url string `json:"keys_url"`
	Comments_url string `json:"comments_url"`
	Archived bool `json:"archived"` // Whether the repository is archived.
	Size int `json:"size"` // The size of the repository. Size is calculated hourly. When a repository is initially created, the size is 0.
	Issue_comment_url string `json:"issue_comment_url"`
	Clone_url string `json:"clone_url"`
	Fork bool `json:"fork"`
	Name string `json:"name"` // The name of the repository.
	Contributors_url string `json:"contributors_url"`
	Events_url string `json:"events_url"`
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is enabled for this repository
	Template_repository map[string]interface{} `json:"template_repository,omitempty"`
	Forks int `json:"forks"`
	Open_issues int `json:"open_issues"`
	Compare_url string `json:"compare_url"`
	Downloads_url string `json:"downloads_url"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Teams_url string `json:"teams_url"`
	Watchers int `json:"watchers"`
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Svn_url string `json:"svn_url"`
	Git_refs_url string `json:"git_refs_url"`
	Deployments_url string `json:"deployments_url"`
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Merges_url string `json:"merges_url"`
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Updated_at string `json:"updated_at"`
	Subscription_url string `json:"subscription_url"`
	Git_tags_url string `json:"git_tags_url"`
	Labels_url string `json:"labels_url"`
	Html_url string `json:"html_url"`
	Master_branch string `json:"master_branch,omitempty"`
	Ssh_url string `json:"ssh_url"`
	Url string `json:"url"`
	Pulls_url string `json:"pulls_url"`
	Collaborators_url string `json:"collaborators_url"`
	Topics []string `json:"topics,omitempty"`
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Issue_events_url string `json:"issue_events_url"`
	Stargazers_count int `json:"stargazers_count"`
	Starred_at string `json:"starred_at,omitempty"`
	Statuses_url string `json:"statuses_url"`
	Releases_url string `json:"releases_url"`
	Forks_count int `json:"forks_count"`
	Git_commits_url string `json:"git_commits_url"`
	Forks_url string `json:"forks_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Language string `json:"language"`
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Homepage string `json:"homepage"`
	Private bool `json:"private"` // Whether the repository is private or public.
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Commits_url string `json:"commits_url"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether or not a pull request head branch that is behind its base branch can always be updated even if it is not required to be up to date before merging.
	Issues_url string `json:"issues_url"`
	Archive_url string `json:"archive_url"`
	Subscribers_url string `json:"subscribers_url"`
	Contents_url string `json:"contents_url"`
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property has been deprecated. Please use `squash_merge_commit_title` instead.
	Milestones_url string `json:"milestones_url"`
	Notifications_url string `json:"notifications_url"`
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Full_name string `json:"full_name"`
	Pushed_at string `json:"pushed_at"`
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Hooks_url string `json:"hooks_url"`
	Open_issues_count int `json:"open_issues_count"`
	Created_at string `json:"created_at"`
	Id int `json:"id"` // Unique identifier of the repository
	Mirror_url string `json:"mirror_url"`
	Stargazers_url string `json:"stargazers_url"`
	Node_id string `json:"node_id"`
	Tags_url string `json:"tags_url"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Has_discussions bool `json:"has_discussions,omitempty"` // Whether discussions are enabled.
	Git_url string `json:"git_url"`
	License GeneratedType `json:"license"` // License Simple
	Languages_url string `json:"languages_url"`
	Has_pages bool `json:"has_pages"`
	Watchers_count int `json:"watchers_count"`
	Trees_url string `json:"trees_url"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Assignees_url string `json:"assignees_url"`
	Branches_url string `json:"branches_url"`
	Network_count int `json:"network_count,omitempty"`
	Description string `json:"description"`
	Blobs_url string `json:"blobs_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Unit_name string `json:"unit_name"`
	Accounts_url string `json:"accounts_url"`
	Monthly_price_in_cents int `json:"monthly_price_in_cents"`
	Number int `json:"number"`
	Price_model string `json:"price_model"`
	State string `json:"state"`
	Yearly_price_in_cents int `json:"yearly_price_in_cents"`
	Bullets []string `json:"bullets"`
	Description string `json:"description"`
	Has_free_trial bool `json:"has_free_trial"`
	Name string `json:"name"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Projects_v2 GeneratedType `json:"projects_v2"` // A projects v2 project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Organization GeneratedType `json:"organization"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Statuses_url string `json:"statuses_url"`
	Open_issues_count int `json:"open_issues_count"`
	Network_count int `json:"network_count"`
	Pulls_url string `json:"pulls_url"`
	Fork bool `json:"fork"`
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"`
	Languages_url string `json:"languages_url"`
	Events_url string `json:"events_url"`
	Updated_at string `json:"updated_at"`
	Forks_url string `json:"forks_url"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Subscription_url string `json:"subscription_url"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"`
	Homepage string `json:"homepage"`
	Contributors_url string `json:"contributors_url"`
	Source Repository `json:"source,omitempty"` // A repository on GitHub.
	Stargazers_url string `json:"stargazers_url"`
	Private bool `json:"private"`
	Default_branch string `json:"default_branch"`
	Archived bool `json:"archived"`
	Has_discussions bool `json:"has_discussions"`
	Has_issues bool `json:"has_issues"`
	Has_wiki bool `json:"has_wiki"`
	Full_name string `json:"full_name"`
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Ssh_url string `json:"ssh_url"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Git_tags_url string `json:"git_tags_url"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub user.
	Master_branch string `json:"master_branch,omitempty"`
	License GeneratedType `json:"license"` // License Simple
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Merges_url string `json:"merges_url"`
	Template_repository GeneratedType `json:"template_repository,omitempty"` // A repository on GitHub.
	Size int `json:"size"` // The size of the repository. Size is calculated hourly. When a repository is initially created, the size is 0.
	Keys_url string `json:"keys_url"`
	Downloads_url string `json:"downloads_url"`
	Is_template bool `json:"is_template,omitempty"`
	Forks_count int `json:"forks_count"`
	Git_refs_url string `json:"git_refs_url"`
	Blobs_url string `json:"blobs_url"`
	Watchers_count int `json:"watchers_count"`
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is allowed.
	Svn_url string `json:"svn_url"`
	Hooks_url string `json:"hooks_url"`
	Language string `json:"language"`
	Open_issues int `json:"open_issues"`
	Milestones_url string `json:"milestones_url"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"`
	Contents_url string `json:"contents_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Deployments_url string `json:"deployments_url"`
	Node_id string `json:"node_id"`
	Id int `json:"id"`
	Compare_url string `json:"compare_url"`
	Subscribers_count int `json:"subscribers_count"`
	Parent Repository `json:"parent,omitempty"` // A repository on GitHub.
	Name string `json:"name"`
	Html_url string `json:"html_url"`
	Description string `json:"description"`
	Has_projects bool `json:"has_projects"`
	Issues_url string `json:"issues_url"`
	Allow_update_branch bool `json:"allow_update_branch,omitempty"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Tags_url string `json:"tags_url"`
	Mirror_url string `json:"mirror_url"`
	Topics []string `json:"topics,omitempty"`
	Comments_url string `json:"comments_url"`
	Assignees_url string `json:"assignees_url"`
	Security_and_analysis GeneratedType `json:"security_and_analysis,omitempty"`
	Issue_comment_url string `json:"issue_comment_url"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Teams_url string `json:"teams_url"`
	Trees_url string `json:"trees_url"`
	Has_downloads bool `json:"has_downloads"`
	Archive_url string `json:"archive_url"`
	Url string `json:"url"`
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Created_at string `json:"created_at"`
	Code_of_conduct GeneratedType `json:"code_of_conduct,omitempty"` // Code of Conduct Simple
	Commits_url string `json:"commits_url"`
	Collaborators_url string `json:"collaborators_url"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"`
	Pushed_at string `json:"pushed_at"`
	Subscribers_url string `json:"subscribers_url"`
	Branches_url string `json:"branches_url"`
	Stargazers_count int `json:"stargazers_count"`
	Has_pages bool `json:"has_pages"`
	Watchers int `json:"watchers"`
	Git_url string `json:"git_url"`
	Git_commits_url string `json:"git_commits_url"`
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Releases_url string `json:"releases_url"`
	Notifications_url string `json:"notifications_url"`
	Issue_events_url string `json:"issue_events_url"`
	Forks int `json:"forks"`
	Allow_forking bool `json:"allow_forking,omitempty"`
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"`
	Labels_url string `json:"labels_url"`
	Clone_url string `json:"clone_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
}

// Stargazer represents the Stargazer schema from the OpenAPI specification
type Stargazer struct {
	Starred_at string `json:"starred_at"`
	User GeneratedType `json:"user"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Team map[string]interface{} `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
}

// Workflow represents the Workflow schema from the OpenAPI specification
type Workflow struct {
	Url string `json:"url"`
	Badge_url string `json:"badge_url"`
	Created_at string `json:"created_at"`
	Deleted_at string `json:"deleted_at,omitempty"`
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Path string `json:"path"`
	Updated_at string `json:"updated_at"`
	Id int `json:"id"`
	Name string `json:"name"`
	State string `json:"state"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField string `json:"type"`
	Links map[string]interface{} `json:"_links"`
	Git_url string `json:"git_url"`
	Size int `json:"size"`
	Target string `json:"target"`
	Download_url string `json:"download_url"`
	Html_url string `json:"html_url"`
	Path string `json:"path"`
	Sha string `json:"sha"`
	Url string `json:"url"`
	Name string `json:"name"`
}

// Hovercard represents the Hovercard schema from the OpenAPI specification
type Hovercard struct {
	Contexts []map[string]interface{} `json:"contexts"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"` // The action that was performed.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Comments_url string `json:"comments_url"` // A template for the API URL to get information about comments on the repository.
	Git_commits_url string `json:"git_commits_url"` // A template for the API URL to get information about Git commits of the repository.
	Forks_url string `json:"forks_url"` // The API URL to list the forks of the repository.
	Languages_url string `json:"languages_url"` // The API URL to get information about the languages of the repository.
	Events_url string `json:"events_url"` // The API URL to list the events of the repository.
	Assignees_url string `json:"assignees_url"` // A template for the API URL to list the available assignees for issues in the repository.
	Releases_url string `json:"releases_url"` // A template for the API URL to get information about releases on the repository.
	Html_url string `json:"html_url"` // The URL to view the repository on GitHub.com.
	Fork bool `json:"fork"` // Whether the repository is a fork.
	Node_id string `json:"node_id"` // The GraphQL identifier of the repository.
	Private bool `json:"private"` // Whether the repository is private.
	Teams_url string `json:"teams_url"` // The API URL to list the teams on the repository.
	Description string `json:"description"` // The repository description.
	Id int `json:"id"` // A unique identifier of the repository.
	Collaborators_url string `json:"collaborators_url"` // A template for the API URL to get information about collaborators of the repository.
	Downloads_url string `json:"downloads_url"` // The API URL to list the downloads on the repository.
	Merges_url string `json:"merges_url"` // The API URL to merge branches in the repository.
	Deployments_url string `json:"deployments_url"` // The API URL to list the deployments of the repository.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Stargazers_url string `json:"stargazers_url"` // The API URL to list the stargazers on the repository.
	Milestones_url string `json:"milestones_url"` // A template for the API URL to get information about milestones of the repository.
	Commits_url string `json:"commits_url"` // A template for the API URL to get information about commits on the repository.
	Contents_url string `json:"contents_url"` // A template for the API URL to get the contents of the repository.
	Contributors_url string `json:"contributors_url"` // A template for the API URL to list the contributors to the repository.
	Archive_url string `json:"archive_url"` // A template for the API URL to download the repository as an archive.
	Name string `json:"name"` // The name of the repository.
	Statuses_url string `json:"statuses_url"` // A template for the API URL to get information about statuses of a commit.
	Trees_url string `json:"trees_url"` // A template for the API URL to create or retrieve a raw Git tree of the repository.
	Issue_events_url string `json:"issue_events_url"` // A template for the API URL to get information about issue events on the repository.
	Notifications_url string `json:"notifications_url"` // A template for the API URL to get information about notifications on the repository.
	Tags_url string `json:"tags_url"` // The API URL to get information about tags on the repository.
	Subscription_url string `json:"subscription_url"` // The API URL to subscribe to notifications for this repository.
	Blobs_url string `json:"blobs_url"` // A template for the API URL to create or retrieve a raw Git blob in the repository.
	Compare_url string `json:"compare_url"` // A template for the API URL to compare two commits or refs.
	Issue_comment_url string `json:"issue_comment_url"` // A template for the API URL to get information about issue comments on the repository.
	Url string `json:"url"` // The URL to get more information about the repository from the GitHub API.
	Issues_url string `json:"issues_url"` // A template for the API URL to get information about issues on the repository.
	Labels_url string `json:"labels_url"` // A template for the API URL to get information about labels of the repository.
	Hooks_url string `json:"hooks_url"` // The API URL to list the hooks on the repository.
	Branches_url string `json:"branches_url"` // A template for the API URL to get information about branches in the repository.
	Git_tags_url string `json:"git_tags_url"` // A template for the API URL to get information about Git tags of the repository.
	Keys_url string `json:"keys_url"` // A template for the API URL to get information about deploy keys on the repository.
	Subscribers_url string `json:"subscribers_url"` // The API URL to list the subscribers on the repository.
	Full_name string `json:"full_name"` // The full, globally unique, name of the repository.
	Git_refs_url string `json:"git_refs_url"` // A template for the API URL to get information about Git refs of the repository.
	Pulls_url string `json:"pulls_url"` // A template for the API URL to get information about pull requests on the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Rule map[string]interface{} `json:"rule"` // The branch protection rule. Includes a `name` and all the [branch protection settings](https://docs.github.com/github/administering-a-repository/defining-the-mergeability-of-pull-requests/about-protected-branches#about-branch-protection-settings) applied to branches that match the name. Binary settings are boolean. Multi-level configurations are one of `off`, `non_admins`, or `everyone`. Actor and build lists are arrays of strings.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert interface{} `json:"alert"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Key_id string `json:"key_id"` // The identifier for the key.
	Key string `json:"key"` // The Base64 encoded public key.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project_column map[string]interface{} `json:"project_column"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"`
	Updated_at string `json:"updated_at"`
	Visibility string `json:"visibility"` // Visibility of a secret
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project_card map[string]interface{} `json:"project_card"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Included_minutes int `json:"included_minutes"` // The amount of free GitHub Actions minutes available.
	Minutes_used_breakdown map[string]interface{} `json:"minutes_used_breakdown"`
	Total_minutes_used int `json:"total_minutes_used"` // The sum of the free and paid GitHub Actions minutes used.
	Total_paid_minutes_used int `json:"total_paid_minutes_used"` // The total paid GitHub Actions minutes used.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Membership map[string]interface{} `json:"membership"` // The membership between the user and the organization. Not present when the action is `member_invited`.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Oid string `json:"oid"`
	Path string `json:"path"`
	Ref_name string `json:"ref_name"`
	Size int `json:"size"`
}

// Reaction represents the Reaction schema from the OpenAPI specification
type Reaction struct {
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	User GeneratedType `json:"user"` // A GitHub user.
	Content string `json:"content"` // The reaction to use
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project_card interface{} `json:"project_card"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Label map[string]interface{} `json:"label,omitempty"`
	Number int `json:"number"` // The pull request number.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Comments int `json:"comments"`
	Labels []interface{} `json:"labels"` // Labels to associate with this issue; pass one or more label names to replace the set of labels on this issue; send an empty array to clear all labels from the issue; note that the labels are silently dropped for users without push access to the repository
	Timeline_url string `json:"timeline_url,omitempty"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Number int `json:"number"` // Number uniquely identifying the issue within its repository
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Closed_by GeneratedType `json:"closed_by,omitempty"` // A GitHub user.
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Updated_at string `json:"updated_at"`
	Body string `json:"body,omitempty"` // Contents of the issue
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Id int `json:"id"`
	Events_url string `json:"events_url"`
	Html_url string `json:"html_url"`
	Url string `json:"url"` // URL for the issue
	State string `json:"state"` // State of the issue; either 'open' or 'closed'
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	User GeneratedType `json:"user"` // A GitHub user.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Draft bool `json:"draft,omitempty"`
	Created_at string `json:"created_at"`
	Body_html string `json:"body_html,omitempty"`
	State_reason string `json:"state_reason,omitempty"` // The reason for the current state
	Repository_url string `json:"repository_url"`
	Title string `json:"title"` // Title of the issue
	Comments_url string `json:"comments_url"`
	Node_id string `json:"node_id"`
	Locked bool `json:"locked"`
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	Body_text string `json:"body_text,omitempty"`
	Closed_at string `json:"closed_at"`
	Labels_url string `json:"labels_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Severity string `json:"severity"` // The severity of the vulnerability.
	Vulnerable_version_range string `json:"vulnerable_version_range"` // Conditions that identify vulnerable versions of this vulnerability's package.
	First_patched_version map[string]interface{} `json:"first_patched_version"` // Details pertaining to the package version that patches this vulnerability.
	PackageField GeneratedType `json:"package"` // Details for the vulnerable package.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project_card map[string]interface{} `json:"project_card"`
	Repository GeneratedType `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Access_level string `json:"access_level"` // Defines the level of access that workflows outside of the repository have to actions and reusable workflows within the repository. `none` means the access is only possible from workflows in this repository. `user` level access allows sharing across user owned private repos only. `organization` level access allows sharing across the organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Projects_v2 GeneratedType `json:"projects_v2"` // A projects v2 project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Rule map[string]interface{} `json:"rule"` // The branch protection rule. Includes a `name` and all the [branch protection settings](https://docs.github.com/github/administering-a-repository/defining-the-mergeability-of-pull-requests/about-protected-branches#about-branch-protection-settings) applied to branches that match the name. Binary settings are boolean. Multi-level configurations are one of `off`, `non_admins`, or `everyone`. Actor and build lists are arrays of strings.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Actions_caches []map[string]interface{} `json:"actions_caches"` // Array of caches
	Total_count int `json:"total_count"` // Total number of caches
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Action string `json:"action"`
	Effective_date string `json:"effective_date,omitempty"` // The `pending_cancellation` and `pending_tier_change` event types will include the date the cancellation or tier change will take effect.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Sponsorship map[string]interface{} `json:"sponsorship"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
	Updated_at string `json:"updated_at"`
}

// Environment represents the Environment schema from the OpenAPI specification
type Environment struct {
	Node_id string `json:"node_id"`
	Created_at string `json:"created_at"` // The time that the environment was created, in ISO 8601 format.
	Name string `json:"name"` // The name of the environment.
	Url string `json:"url"`
	Deployment_branch_policy GeneratedType `json:"deployment_branch_policy,omitempty"` // The type of deployment branch policy for this environment. To allow all branches to deploy, set to `null`.
	Protection_rules []interface{} `json:"protection_rules,omitempty"`
	Updated_at string `json:"updated_at"` // The time that the environment was last updated, in ISO 8601 format.
	Html_url string `json:"html_url"`
	Id int `json:"id"` // The id of the environment.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	State string `json:"state"` // The state of the Dependabot alert.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
	Dismissed_reason string `json:"dismissed_reason"` // The reason that the alert was dismissed.
	Number int `json:"number"` // The security alert number.
	Updated_at string `json:"updated_at"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dependency map[string]interface{} `json:"dependency"` // Details for the vulnerable dependency.
	Fixed_at string `json:"fixed_at"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Url string `json:"url"` // The REST API URL of the alert resource.
	Dismissed_comment string `json:"dismissed_comment"` // An optional comment associated with the alert's dismissal.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Security_advisory GeneratedType `json:"security_advisory"` // Details for the GitHub Security Advisory.
	Security_vulnerability GeneratedType `json:"security_vulnerability"` // Details pertaining to one vulnerable version range for the advisory.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project map[string]interface{} `json:"project"`
	Repository GeneratedType `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Color string `json:"color"`
	Name string `json:"name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Release interface{} `json:"release"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Branch string `json:"branch"`
	Client_payload map[string]interface{} `json:"client_payload"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repo Repository `json:"repo"` // A repository on GitHub.
	Starred_at string `json:"starred_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Selected_repositories_url string `json:"selected_repositories_url"` // The API URL at which the list of repositories this secret is visible to can be retrieved
	Updated_at string `json:"updated_at"` // The date and time at which the secret was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Visibility string `json:"visibility"` // The type of repositories in the organization that the secret is visible to
	Created_at string `json:"created_at"` // The date and time at which the secret was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Name string `json:"name"` // The name of the secret
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Project map[string]interface{} `json:"project"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Resources map[string]interface{} `json:"resources"`
	Rate GeneratedType `json:"rate"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_title string `json:"commit_title"` // Title for the merge commit message.
	Enabled_by GeneratedType `json:"enabled_by"` // A GitHub user.
	Merge_method string `json:"merge_method"` // The merge method to use.
	Commit_message string `json:"commit_message"` // Commit message for the merge commit.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue interface{} `json:"issue"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// Event represents the Event schema from the OpenAPI specification
type Event struct {
	Org Actor `json:"org,omitempty"` // Actor
	Payload map[string]interface{} `json:"payload"`
	Public bool `json:"public"`
	Repo map[string]interface{} `json:"repo"`
	TypeField string `json:"type"`
	Actor Actor `json:"actor"` // Actor
	Created_at string `json:"created_at"`
	Id string `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.rerequested JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Team map[string]interface{} `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Event string `json:"event"`
	Id int `json:"id"`
	Commit_url string `json:"commit_url"`
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	State_reason string `json:"state_reason,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Comment map[string]interface{} `json:"comment"` // The [comment](https://docs.github.com/rest/reference/pulls#comments) itself.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Comment map[string]interface{} `json:"comment"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Changes map[string]interface{} `json:"changes"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// Status represents the Status schema from the OpenAPI specification
type Status struct {
	State string `json:"state"`
	Target_url string `json:"target_url"`
	Updated_at string `json:"updated_at"`
	Context string `json:"context"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Avatar_url string `json:"avatar_url"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Description string `json:"description"`
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_id string `json:"commit_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Commit_url string `json:"commit_url"`
	Lock_reason string `json:"lock_reason"`
	Url string `json:"url"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Client_id string `json:"client_id,omitempty"`
	Description string `json:"description"`
	Id int `json:"id"` // Unique identifier of the GitHub app
	Permissions map[string]interface{} `json:"permissions"` // The set of permissions for the GitHub app
	Webhook_secret string `json:"webhook_secret,omitempty"`
	Html_url string `json:"html_url"`
	External_url string `json:"external_url"`
	Slug string `json:"slug,omitempty"` // The slug name of the GitHub app
	Node_id string `json:"node_id"`
	Client_secret string `json:"client_secret,omitempty"`
	Events []string `json:"events"` // The list of events for the GitHub app
	Name string `json:"name"` // The name of the GitHub app
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Updated_at string `json:"updated_at"`
	Installations_count int `json:"installations_count,omitempty"` // The number of installations associated with the GitHub app
	Pem string `json:"pem,omitempty"`
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Check_run GeneratedType `json:"check_run"` // A check performed on the code of a given code change
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Requested_action map[string]interface{} `json:"requested_action,omitempty"` // The action requested by the user.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Teams_url string `json:"teams_url"`
	Url string `json:"url"`
	Users []map[string]interface{} `json:"users"`
	Users_url string `json:"users_url"`
	Apps []map[string]interface{} `json:"apps"`
	Apps_url string `json:"apps_url"`
	Teams []map[string]interface{} `json:"teams"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
}

// Integration represents the Integration schema from the OpenAPI specification
type Integration struct {
	Name string `json:"name"` // The name of the GitHub app
	Client_id string `json:"client_id,omitempty"`
	Description string `json:"description"`
	Updated_at string `json:"updated_at"`
	Installations_count int `json:"installations_count,omitempty"` // The number of installations associated with the GitHub app
	Created_at string `json:"created_at"`
	External_url string `json:"external_url"`
	Id int `json:"id"` // Unique identifier of the GitHub app
	Slug string `json:"slug,omitempty"` // The slug name of the GitHub app
	Node_id string `json:"node_id"`
	Permissions map[string]interface{} `json:"permissions"` // The set of permissions for the GitHub app
	Webhook_secret string `json:"webhook_secret,omitempty"`
	Client_secret string `json:"client_secret,omitempty"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Html_url string `json:"html_url"`
	Pem string `json:"pem,omitempty"`
	Events []string `json:"events"` // The list of events for the GitHub app
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Project_url string `json:"project_url"`
	Url string `json:"url"`
	Column_name string `json:"column_name"`
	Id int `json:"id"`
	Previous_column_name string `json:"previous_column_name,omitempty"`
	Project_id int `json:"project_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Marketplace_purchase interface{} `json:"marketplace_purchase"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Effective_date string `json:"effective_date"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Previous_marketplace_purchase map[string]interface{} `json:"previous_marketplace_purchase,omitempty"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sha string `json:"sha"`
	Size int `json:"size"`
	Links map[string]interface{} `json:"_links"`
	Path string `json:"path"`
	Entries []map[string]interface{} `json:"entries,omitempty"`
	Html_url string `json:"html_url"`
	Name string `json:"name"`
	TypeField string `json:"type"`
	Download_url string `json:"download_url"`
	Git_url string `json:"git_url"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project_column map[string]interface{} `json:"project_column"`
	Repository GeneratedType `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Draft bool `json:"draft,omitempty"`
	Url string `json:"url"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Events_url string `json:"events_url"`
	Number int `json:"number"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	User GeneratedType `json:"user"` // A GitHub user.
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Labels_url string `json:"labels_url"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Body string `json:"body,omitempty"`
	Id int `json:"id"`
	Comments_url string `json:"comments_url"`
	State_reason string `json:"state_reason,omitempty"`
	Created_at string `json:"created_at"`
	Locked bool `json:"locked"`
	Html_url string `json:"html_url"`
	State string `json:"state"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Repository_url string `json:"repository_url"`
	Updated_at string `json:"updated_at"`
	Body_text string `json:"body_text,omitempty"`
	Timeline_url string `json:"timeline_url,omitempty"`
	Comments int `json:"comments"`
	Title string `json:"title"`
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Node_id string `json:"node_id"`
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	Score float64 `json:"score"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body_html string `json:"body_html,omitempty"`
	Labels []map[string]interface{} `json:"labels"`
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Closed_at string `json:"closed_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Blocked_user map[string]interface{} `json:"blocked_user"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Rename map[string]interface{} `json:"rename"`
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Commit_url string `json:"commit_url"`
	Event string `json:"event"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue_body_url string `json:"issue_body_url"` // The API URL to get the issue where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Html_url string `json:"html_url"`
	Name string `json:"name"`
	Sha string `json:"sha"`
	Size int `json:"size"`
	Submodule_git_url string `json:"submodule_git_url,omitempty"`
	Encoding string `json:"encoding"`
	Path string `json:"path"`
	Target string `json:"target,omitempty"`
	TypeField string `json:"type"`
	Url string `json:"url"`
	Links map[string]interface{} `json:"_links"`
	Content string `json:"content"`
	Git_url string `json:"git_url"`
	Download_url string `json:"download_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Html_url string `json:"html_url"`
	Body string `json:"body"` // The text of the review.
	Id int `json:"id"` // Unique identifier of the review
	State string `json:"state"`
	Submitted_at string `json:"submitted_at,omitempty"`
	User GeneratedType `json:"user"` // A GitHub user.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body_text string `json:"body_text,omitempty"`
	Commit_id string `json:"commit_id"` // A commit SHA for the review.
	Event string `json:"event"`
	Node_id string `json:"node_id"`
	Pull_request_url string `json:"pull_request_url"`
	Links map[string]interface{} `json:"_links"`
	Body_html string `json:"body_html,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The name of the secret
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"` // The API URL at which the list of repositories this secret is visible to can be retrieved
	Updated_at string `json:"updated_at"` // The date and time at which the secret was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Visibility string `json:"visibility"` // The type of repositories in the organization that the secret is visible to
	Created_at string `json:"created_at"` // The date and time at which the secret was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	State string `json:"state"` // The state of the user's membership in the team.
	Url string `json:"url"`
	Role string `json:"role"` // The role of the user in the team.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enabled bool `json:"enabled"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Protection GeneratedType `json:"protection"` // Branch Protection
	Protection_url string `json:"protection_url"`
	Required_approving_review_count int `json:"required_approving_review_count,omitempty"`
	Links map[string]interface{} `json:"_links"`
	Commit Commit `json:"commit"` // Commit
	Name string `json:"name"`
	Pattern string `json:"pattern,omitempty"`
	Protected bool `json:"protected"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Number int `json:"number"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repositories_removed []map[string]interface{} `json:"repositories_removed"` // An array of repository objects, which were removed from the installation.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Installation Installation `json:"installation"` // Installation
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Action string `json:"action"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repositories_added []map[string]interface{} `json:"repositories_added"` // An array of repository objects, which were added to the installation.
	Requester map[string]interface{} `json:"requester"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Display_name string `json:"display_name"` // The display name of the machine includes cores, memory, and storage.
	Memory_in_bytes int `json:"memory_in_bytes"` // How much memory is available to the codespace.
	Name string `json:"name"` // The name of the machine.
	Operating_system string `json:"operating_system"` // The operating system of the machine.
	Prebuild_availability string `json:"prebuild_availability"` // Whether a prebuild is currently available when creating a codespace for this machine and repository. If a branch was not specified as a ref, the default branch will be assumed. Value will be "null" if prebuilds are not supported or prebuild availability could not be determined. Value will be "none" if no prebuild is available. Latest values "ready" and "in_progress" indicate the prebuild availability status.
	Storage_in_bytes int `json:"storage_in_bytes"` // How much storage is available to the codespace.
	Cpus int `json:"cpus"` // How many cores are available to the codespace.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pages string `json:"pages,omitempty"` // The level of permission to grant the access token to retrieve Pages statuses, configuration, and builds, as well as create new builds.
	Repository_projects string `json:"repository_projects,omitempty"` // The level of permission to grant the access token to manage repository projects, columns, and cards.
	Statuses string `json:"statuses,omitempty"` // The level of permission to grant the access token for commit statuses.
	Organization_projects string `json:"organization_projects,omitempty"` // The level of permission to grant the access token to manage organization projects and projects beta (where available).
	Organization_user_blocking string `json:"organization_user_blocking,omitempty"` // The level of permission to grant the access token to view and manage users blocked by the organization.
	Pull_requests string `json:"pull_requests,omitempty"` // The level of permission to grant the access token for pull requests and related comments, assignees, labels, milestones, and merges.
	Checks string `json:"checks,omitempty"` // The level of permission to grant the access token for checks on code.
	Organization_hooks string `json:"organization_hooks,omitempty"` // The level of permission to grant the access token to manage the post-receive hooks for an organization.
	Issues string `json:"issues,omitempty"` // The level of permission to grant the access token for issues and related comments, assignees, labels, and milestones.
	Metadata string `json:"metadata,omitempty"` // The level of permission to grant the access token to search repositories, list collaborators, and access repository metadata.
	Single_file string `json:"single_file,omitempty"` // The level of permission to grant the access token to manage just a single file.
	Organization_administration string `json:"organization_administration,omitempty"` // The level of permission to grant the access token to manage access to an organization.
	Contents string `json:"contents,omitempty"` // The level of permission to grant the access token for repository contents, commits, branches, downloads, releases, and merges.
	Team_discussions string `json:"team_discussions,omitempty"` // The level of permission to grant the access token to manage team discussions and related comments.
	Organization_custom_roles string `json:"organization_custom_roles,omitempty"` // The level of permission to grant the access token for custom repository roles management. This property is in beta and is subject to change.
	Members string `json:"members,omitempty"` // The level of permission to grant the access token for organization teams and members.
	Vulnerability_alerts string `json:"vulnerability_alerts,omitempty"` // The level of permission to grant the access token to manage Dependabot alerts.
	Workflows string `json:"workflows,omitempty"` // The level of permission to grant the access token to update GitHub Actions workflow files.
	Organization_announcement_banners string `json:"organization_announcement_banners,omitempty"` // The level of permission to grant the access token to view and manage announcement banners for an organization.
	Organization_packages string `json:"organization_packages,omitempty"` // The level of permission to grant the access token for organization packages published to GitHub Packages.
	Organization_self_hosted_runners string `json:"organization_self_hosted_runners,omitempty"` // The level of permission to grant the access token to view and manage GitHub Actions self-hosted runners available to an organization.
	Actions string `json:"actions,omitempty"` // The level of permission to grant the access token for GitHub Actions workflows, workflow runs, and artifacts.
	Deployments string `json:"deployments,omitempty"` // The level of permission to grant the access token for deployments and deployment statuses.
	Administration string `json:"administration,omitempty"` // The level of permission to grant the access token for repository creation, deletion, settings, teams, and collaborators creation.
	Environments string `json:"environments,omitempty"` // The level of permission to grant the access token for managing repository environments.
	Repository_hooks string `json:"repository_hooks,omitempty"` // The level of permission to grant the access token to manage the post-receive hooks for a repository.
	Repository_announcement_banners string `json:"repository_announcement_banners,omitempty"` // The level of permission to grant the access token to view and manage announcement banners for a repository.
	Security_events string `json:"security_events,omitempty"` // The level of permission to grant the access token to view and manage security events like code scanning alerts.
	Secrets string `json:"secrets,omitempty"` // The level of permission to grant the access token to manage repository secrets.
	Organization_secrets string `json:"organization_secrets,omitempty"` // The level of permission to grant the access token to manage organization secrets.
	Organization_plan string `json:"organization_plan,omitempty"` // The level of permission to grant the access token for viewing an organization's plan.
	Packages string `json:"packages,omitempty"` // The level of permission to grant the access token for packages published to GitHub Packages.
	Secret_scanning_alerts string `json:"secret_scanning_alerts,omitempty"` // The level of permission to grant the access token to view and manage secret scanning alerts.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_id string `json:"commit_id"`
	Id int `json:"id"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Created_at string `json:"created_at"`
	Label map[string]interface{} `json:"label"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Commit_url string `json:"commit_url"`
	Event string `json:"event"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Key_id string `json:"key_id"` // The identifier for the key.
	Title string `json:"title,omitempty"`
	Url string `json:"url,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Id int `json:"id,omitempty"`
	Key string `json:"key"` // The Base64 encoded public key.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.completed JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	All []int `json:"all"`
	Owner []int `json:"owner"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project map[string]interface{} `json:"project"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Head_repository GeneratedType `json:"head_repository"` // Minimal Repository
	Rerun_url string `json:"rerun_url"` // The URL to rerun the workflow run.
	Check_suite_id int `json:"check_suite_id,omitempty"` // The ID of the associated check suite.
	Updated_at string `json:"updated_at"`
	Check_suite_node_id string `json:"check_suite_node_id,omitempty"` // The node ID of the associated check suite.
	Head_branch string `json:"head_branch"`
	Node_id string `json:"node_id"`
	Status string `json:"status"`
	Workflow_id int `json:"workflow_id"` // The ID of the parent workflow.
	Id int `json:"id"` // The ID of the workflow run.
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Logs_url string `json:"logs_url"` // The URL to download the logs for the workflow run.
	Url string `json:"url"` // The URL to the workflow run.
	Head_sha string `json:"head_sha"` // The SHA of the head commit that points to the version of the workflow being run.
	Jobs_url string `json:"jobs_url"` // The URL to the jobs for the workflow run.
	Event string `json:"event"`
	Triggering_actor GeneratedType `json:"triggering_actor,omitempty"` // A GitHub user.
	Workflow_url string `json:"workflow_url"` // The URL to the workflow.
	Run_attempt int `json:"run_attempt,omitempty"` // Attempt number of the run, 1 for first attempt and higher if the workflow was re-run.
	Referenced_workflows []GeneratedType `json:"referenced_workflows,omitempty"`
	Run_number int `json:"run_number"` // The auto incrementing run number for the workflow run.
	Head_repository_id int `json:"head_repository_id,omitempty"`
	Created_at string `json:"created_at"`
	Pull_requests []GeneratedType `json:"pull_requests"`
	Name string `json:"name,omitempty"` // The name of the workflow run.
	Cancel_url string `json:"cancel_url"` // The URL to cancel the workflow run.
	Run_started_at string `json:"run_started_at,omitempty"` // The start time of the latest run. Resets on re-run.
	Path string `json:"path"` // The full path of the workflow
	Conclusion string `json:"conclusion"`
	Html_url string `json:"html_url"`
	Previous_attempt_url string `json:"previous_attempt_url,omitempty"` // The URL to the previous attempted run of this workflow, if one exists.
	Actor GeneratedType `json:"actor,omitempty"` // A GitHub user.
	Artifacts_url string `json:"artifacts_url"` // The URL to the artifacts for the workflow run.
	Head_commit GeneratedType `json:"head_commit"` // A commit.
	Check_suite_url string `json:"check_suite_url"` // The URL to the associated check suite.
	Display_title string `json:"display_title"` // The event-specific title associated with the run or the run-name if set, or the value of `run-name` if it is set in the workflow.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Comments_url string `json:"comments_url"`
	Score float64 `json:"score"`
	Author GeneratedType `json:"author"` // A GitHub user.
	Commit map[string]interface{} `json:"commit"`
	Html_url string `json:"html_url"`
	Url string `json:"url"`
	Committer GeneratedType `json:"committer"` // Metaproperties for Git author/committer information.
	Node_id string `json:"node_id"`
	Parents []map[string]interface{} `json:"parents"`
	Sha string `json:"sha"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Starred_at string `json:"starred_at"` // The time the star was created. This is a timestamp in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`. Will be `null` for the `deleted` action.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Number int `json:"number"`
	Url string `json:"url"`
	Base map[string]interface{} `json:"base"`
	Head map[string]interface{} `json:"head"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Member map[string]interface{} `json:"member"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Alt_domain map[string]interface{} `json:"alt_domain,omitempty"`
	Domain map[string]interface{} `json:"domain,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow_job interface{} `json:"workflow_job"`
	Action string `json:"action"`
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Instances_url string `json:"instances_url"` // The REST API URL for fetching the list of instances for an alert.
	Repository GeneratedType `json:"repository"` // A GitHub repository.
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
	Fixed_at string `json:"fixed_at,omitempty"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Most_recent_instance GeneratedType `json:"most_recent_instance"`
	Dismissed_reason string `json:"dismissed_reason"` // **Required when the state is dismissed.** The reason for dismissing or closing the alert.
	Tool GeneratedType `json:"tool"`
	State string `json:"state"` // State of a code scanning alert.
	Number int `json:"number"` // The security alert number.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Rule GeneratedType `json:"rule"`
	Dismissed_comment string `json:"dismissed_comment,omitempty"` // The dismissal comment associated with the dismissal of the alert.
	Url string `json:"url"` // The REST API URL of the alert resource.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
}

// Page represents the Page schema from the OpenAPI specification
type Page struct {
	Pending_domain_unverified_at string `json:"pending_domain_unverified_at,omitempty"` // The timestamp when a pending domain becomes unverified.
	Url string `json:"url"` // The API address for accessing this Page resource.
	Build_type string `json:"build_type,omitempty"` // The process in which the Page will be built.
	Https_enforced bool `json:"https_enforced,omitempty"` // Whether https is enabled on the domain
	Protected_domain_state string `json:"protected_domain_state,omitempty"` // The state if the domain is verified
	Status string `json:"status"` // The status of the most recent build of the Page.
	Html_url string `json:"html_url,omitempty"` // The web address the Page can be accessed from.
	Public bool `json:"public"` // Whether the GitHub Pages site is publicly visible. If set to `true`, the site is accessible to anyone on the internet. If set to `false`, the site will only be accessible to users who have at least `read` access to the repository that published the site.
	Source GeneratedType `json:"source,omitempty"`
	Cname string `json:"cname"` // The Pages site's custom domain
	Custom_404 bool `json:"custom_404"` // Whether the Page has a custom 404 page.
	Https_certificate GeneratedType `json:"https_certificate,omitempty"`
}

// Actor represents the Actor schema from the OpenAPI specification
type Actor struct {
	Url string `json:"url"`
	Avatar_url string `json:"avatar_url"`
	Display_login string `json:"display_login,omitempty"`
	Gravatar_id string `json:"gravatar_id"`
	Id int `json:"id"`
	Login string `json:"login"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The name of the CodeQL database.
	Size int `json:"size"` // The size of the CodeQL database file in bytes.
	Updated_at string `json:"updated_at"` // The date and time at which the CodeQL database was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Url string `json:"url"` // The URL at which to download the CodeQL database. The `Accept` header must be set to the value of the `content_type` property.
	Content_type string `json:"content_type"` // The MIME type of the CodeQL database file.
	Id int `json:"id"` // The ID of the CodeQL database.
	Language string `json:"language"` // The language of the CodeQL database.
	Uploader GeneratedType `json:"uploader"` // A GitHub user.
	Created_at string `json:"created_at"` // The date and time at which the CodeQL database was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue interface{} `json:"issue"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Number int `json:"number"` // The unique sequence number of a team discussion comment.
	Updated_at string `json:"updated_at"`
	Author GeneratedType `json:"author"` // A GitHub user.
	Discussion_url string `json:"discussion_url"`
	Last_edited_at string `json:"last_edited_at"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Body_html string `json:"body_html"`
	Created_at string `json:"created_at"`
	Url string `json:"url"`
	Body string `json:"body"` // The main text of the comment.
	Body_version string `json:"body_version"` // The current version of the body content. If provided, this update operation will be rejected if the given version does not match the latest version on the server.
	Html_url string `json:"html_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Message string `json:"message"`
	Sha string `json:"sha"`
	Merged bool `json:"merged"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Description string `json:"description"`
	Curated bool `json:"curated"`
	Display_name string `json:"display_name"`
	Related []map[string]interface{} `json:"related,omitempty"`
	Name string `json:"name"`
	Logo_url string `json:"logo_url,omitempty"`
	Created_at string `json:"created_at"`
	Featured bool `json:"featured"`
	Aliases []map[string]interface{} `json:"aliases,omitempty"`
	Updated_at string `json:"updated_at"`
	Created_by string `json:"created_by"`
	Score float64 `json:"score"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Short_description string `json:"short_description"`
	Released string `json:"released"`
	Repository_count int `json:"repository_count,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	User GeneratedType `json:"user"` // A GitHub user.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
	Path string `json:"path"`
	Line int `json:"line"`
	Node_id string `json:"node_id"`
	Position int `json:"position"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Body string `json:"body"`
	Id int `json:"id"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Actions_meta map[string]interface{} `json:"actions_meta,omitempty"`
	Check_suite map[string]interface{} `json:"check_suite"` // The [check_suite](https://docs.github.com/rest/reference/checks#suites).
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// Enterprise represents the Enterprise schema from the OpenAPI specification
type Enterprise struct {
	Node_id string `json:"node_id"`
	Slug string `json:"slug"` // The slug url identifier for the enterprise.
	Avatar_url string `json:"avatar_url"`
	Html_url string `json:"html_url"`
	Updated_at string `json:"updated_at"`
	Website_url string `json:"website_url,omitempty"` // The enterprise's website URL.
	Created_at string `json:"created_at"`
	Description string `json:"description,omitempty"` // A short description of the enterprise.
	Name string `json:"name"` // The name of the enterprise.
	Id int `json:"id"` // Unique identifier of the enterprise
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Languages_url string `json:"languages_url"`
	Issue_comment_url string `json:"issue_comment_url"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Private bool `json:"private"`
	Created_at string `json:"created_at"`
	Subscribers_url string `json:"subscribers_url"`
	Watchers_count int `json:"watchers_count"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"`
	Size int `json:"size"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Git_refs_url string `json:"git_refs_url"`
	Has_projects bool `json:"has_projects"`
	Fork bool `json:"fork"`
	Pulls_url string `json:"pulls_url"`
	Clone_url string `json:"clone_url"`
	Svn_url string `json:"svn_url"`
	Collaborators_url string `json:"collaborators_url"`
	Url string `json:"url"`
	Forks int `json:"forks"`
	Score float64 `json:"score"`
	Language string `json:"language"`
	Compare_url string `json:"compare_url"`
	Hooks_url string `json:"hooks_url"`
	Branches_url string `json:"branches_url"`
	Blobs_url string `json:"blobs_url"`
	Default_branch string `json:"default_branch"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Stargazers_count int `json:"stargazers_count"`
	Id int `json:"id"`
	Contents_url string `json:"contents_url"`
	Open_issues int `json:"open_issues"`
	Open_issues_count int `json:"open_issues_count"`
	Events_url string `json:"events_url"`
	Tags_url string `json:"tags_url"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"`
	Topics []string `json:"topics,omitempty"`
	Milestones_url string `json:"milestones_url"`
	Node_id string `json:"node_id"`
	Forks_count int `json:"forks_count"`
	Has_discussions bool `json:"has_discussions,omitempty"`
	Watchers int `json:"watchers"`
	Notifications_url string `json:"notifications_url"`
	Ssh_url string `json:"ssh_url"`
	Deployments_url string `json:"deployments_url"`
	Issue_events_url string `json:"issue_events_url"`
	Statuses_url string `json:"statuses_url"`
	Master_branch string `json:"master_branch,omitempty"`
	Subscription_url string `json:"subscription_url"`
	License GeneratedType `json:"license"` // License Simple
	Archived bool `json:"archived"`
	Has_downloads bool `json:"has_downloads"`
	Issues_url string `json:"issues_url"`
	Has_issues bool `json:"has_issues"`
	Has_pages bool `json:"has_pages"`
	Homepage string `json:"homepage"`
	Git_commits_url string `json:"git_commits_url"`
	Updated_at string `json:"updated_at"`
	Keys_url string `json:"keys_url"`
	Name string `json:"name"`
	Is_template bool `json:"is_template,omitempty"`
	Description string `json:"description"`
	Forks_url string `json:"forks_url"`
	Releases_url string `json:"releases_url"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Downloads_url string `json:"downloads_url"`
	Pushed_at string `json:"pushed_at"`
	Full_name string `json:"full_name"`
	Allow_forking bool `json:"allow_forking,omitempty"`
	Archive_url string `json:"archive_url"`
	Mirror_url string `json:"mirror_url"`
	Commits_url string `json:"commits_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Git_url string `json:"git_url"`
	Contributors_url string `json:"contributors_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Comments_url string `json:"comments_url"`
	Trees_url string `json:"trees_url"`
	Html_url string `json:"html_url"`
	Merges_url string `json:"merges_url"`
	Teams_url string `json:"teams_url"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"`
	Labels_url string `json:"labels_url"`
	Has_wiki bool `json:"has_wiki"`
	Assignees_url string `json:"assignees_url"`
	Git_tags_url string `json:"git_tags_url"`
	Stargazers_url string `json:"stargazers_url"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Membership map[string]interface{} `json:"membership,omitempty"` // The membership between the user and the organization. Not present when the action is `member_invited`.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Body string `json:"body,omitempty"` // Contents of the issue comment
	Body_html string `json:"body_html,omitempty"`
	Body_text string `json:"body_text,omitempty"`
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"` // URL for the issue comment
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Html_url string `json:"html_url"`
	Id int `json:"id"` // Unique identifier of the issue comment
	Issue_url string `json:"issue_url"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	User GeneratedType `json:"user"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert interface{} `json:"alert"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project_column map[string]interface{} `json:"project_column"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// Language represents the Language schema from the OpenAPI specification
type Language struct {
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Warning string `json:"warning"` // Warning generated when processing the analysis
	Analysis_key string `json:"analysis_key"` // Identifies the configuration under which the analysis was executed. For example, in GitHub Actions this includes the workflow filename and job name.
	Commit_sha string `json:"commit_sha"` // The SHA of the commit to which the analysis you are uploading relates.
	Environment string `json:"environment"` // Identifies the variable values associated with the environment in which this analysis was performed.
	Rules_count int `json:"rules_count"` // The total number of rules used in the analysis.
	Url string `json:"url"` // The REST API URL of the analysis resource.
	Ref string `json:"ref"` // The full Git reference, formatted as `refs/heads/<branch name>`, `refs/pull/<number>/merge`, or `refs/pull/<number>/head`.
	Sarif_id string `json:"sarif_id"` // An identifier for the upload.
	Tool GeneratedType `json:"tool"`
	Created_at string `json:"created_at"` // The time that the analysis was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	ErrorField string `json:"error"`
	Deletable bool `json:"deletable"`
	Category string `json:"category,omitempty"` // Identifies the configuration under which the analysis was executed. Used to distinguish between multiple analyses for the same tool and commit, but performed on different languages or different parts of the code.
	Id int `json:"id"` // Unique identifier for this analysis.
	Results_count int `json:"results_count"` // The total number of results in the analysis.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Performed_via_github_app Integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Commit_url string `json:"commit_url"`
	Id int `json:"id"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Url string `json:"url"`
	Assigner GeneratedType `json:"assigner"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Event string `json:"event"`
}

// Deployment represents the Deployment schema from the OpenAPI specification
type Deployment struct {
	Original_environment string `json:"original_environment,omitempty"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Ref string `json:"ref"` // The ref to deploy. This can be a branch, tag, or sha.
	Repository_url string `json:"repository_url"`
	Environment string `json:"environment"` // Name for the target deployment environment.
	Transient_environment bool `json:"transient_environment,omitempty"` // Specifies if the given environment is will no longer exist at some point in the future. Default: false.
	Description string `json:"description"`
	Payload interface{} `json:"payload"`
	Statuses_url string `json:"statuses_url"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Id int `json:"id"` // Unique identifier of the deployment
	Task string `json:"task"` // Parameter to specify a task to execute
	Production_environment bool `json:"production_environment,omitempty"` // Specifies if the given environment is one that end-users directly interact with. Default: false.
	Sha string `json:"sha"`
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Requester map[string]interface{} `json:"requester,omitempty"`
}
