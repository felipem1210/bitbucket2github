package b2g

import (
	"context"

	"github.com/google/go-github/v47/github"
	"golang.org/x/oauth2"
)

// Authenticate with Github
func githubInitClient() (*github.Client, context.Context) {
	gh_token := getEnvValue("GITHUB_TOKEN")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: gh_token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return client, ctx
}

func BranchProtectionUpdate(r string, t string) {
	client, ctx := githubInitClient()
	var teams, users []string
	users = append(users, "bse-sys")
	teams = append(teams, t)
	//repoService := &github.RepositoriesService{}
	protectionOptions := &github.ProtectionRequest{
		// RequiredStatusChecks: &github.RequiredStatusChecks{
		// 	Strict: false,
		// },
		RequiredPullRequestReviews: &github.PullRequestReviewsEnforcementRequest{
			BypassPullRequestAllowancesRequest: &github.BypassPullRequestAllowancesRequest{
				Users: users,
				Teams: teams,
				Apps:  users,
			},
			DismissStaleReviews:          false,
			RequireCodeOwnerReviews:      false,
			RequiredApprovingReviewCount: 2,
		},
		Restrictions: &github.BranchRestrictionsRequest{
			Users: users,
			Teams: teams,
		},
		EnforceAdmins:                  false,
		RequireLinearHistory:           github.Bool(false),
		AllowForcePushes:               github.Bool(true),
		AllowDeletions:                 github.Bool(false),
		RequiredConversationResolution: github.Bool(false),
	}
	_, _, err := client.Repositories.UpdateBranchProtection(ctx, "bestseller-ecom", r, "master", protectionOptions)
	CheckIfError(err)
}
