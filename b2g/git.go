package b2g

import (
	"io"
	"os"
	"os/exec"

	"github.com/creack/pty"
	"github.com/fatih/color"
)

func DownloadBitbucketRepos(r string) {
	color.Green("Downloading bitbucket repo %s", r)
	wd := "/tmp"
	os.Chdir(wd)
	cmd := "git clone git@bitbucket.org:bestseller-ecom/" + r + ".git"
	out := exec.Command("/bin/sh", "-c", cmd)
	f, err := pty.Start(out)
	CheckIfError(err)
	io.Copy(os.Stdout, f)
}

func AddRemote(r string) {
	color.Green("Adding github remote in local repo %s", r)
	chdirRepo(r)
	cmd := "git remote add github git@github.com:bestseller-ecom/" + r + ".git"
	out := exec.Command("/bin/sh", "-c", cmd)
	f, err := pty.Start(out)
	CheckIfError(err)
	io.Copy(os.Stdout, f)
}

func PushRemote(r string) {
	color.Green("Pushing all content from bitbucket to github in repo %s", r)
	chdirRepo(r)
	cmd := "git push -f --all github"
	out := exec.Command("/bin/sh", "-c", cmd)
	f, err := pty.Start(out)
	CheckIfError(err)
	io.Copy(os.Stdout, f)
}
