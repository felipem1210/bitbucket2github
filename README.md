# bitbucket2github
A repo with Github workflows and a script to migrate from bitbucket to github

## Installation

Change the version for the [version](https://github.com/felipem1210/git-helper/tags) you want (withouth initial v)

### Linux amd64

```sh
export B2G_VERSION=0.1.0
curl -L "https://github.com/bestseller-ecom/bibucket2github_/releases/download/v${B2G_VERSION}/bibucket2github_${B2G_VERSION}_linux_amd64.tar.gz" |tar xzv -C /tmp
sudo mv /tmp/bibucket2github /usr/local/bin/bibucket2github
chmod +x /usr/local/bin/bibucket2github
```

### MacOS amd64

```sh
export B2G_VERSION=0.1.0
curl -L "https://github.com/bestseller-ecom/bibucket2github/releases/download/v${B2G_VERSION}/git-helper_${B2G_VERSION}_darwin_amd64.tar.gz" |tar xzv -C /tmp
sudo mv /tmp/bibucket2github /usr/local/bin/bibucket2github
chmod +x /usr/local/bin/bibucket2github
```

### Envars needed

* Define the envar GITHUB_TOKEN with your Github Personal Access Token [PAT](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token).

```sh
export GITHUB_TOKEN=<your-gh-pat>
```

## Usage

Use the `--help`

```sh
bitbucket2github --help
```