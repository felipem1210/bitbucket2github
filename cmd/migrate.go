package cmd

import (
	"strings"

	"github.com/felipem1210/bitbucket2github/b2g"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrates a bitbucket repo",
	Long:  `Migrates a list of repositories from bitbucket to github, It doesn't create github repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		repos, _ := cmd.Flags().GetString("repos")
		team, _ := cmd.Flags().GetString("team")
		reposSlice := strings.Split(repos, ",")
		for _, repo := range reposSlice {
			b2g.DownloadBitbucketRepos(repo)
			b2g.AddRemote(repo)
			b2g.BranchProtectionUpdate(repo, team)
			b2g.PushRemote(repo)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.PersistentFlags().StringP("repos", "r", "", "List of repos, comma separated, to migrate from bitbucket to github")
	migrateCmd.PersistentFlags().StringP("team", "t", "", "Github team that owns the repository")
	migrateCmd.MarkPersistentFlagRequired("repos")
	migrateCmd.MarkPersistentFlagRequired("team")

}
