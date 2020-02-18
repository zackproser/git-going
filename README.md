# Git Going  
A git repository scaffolding utility that I am writing for myself.

## Overview

Git-going is a simple command line tool that lets you start your next project very quickly. It automates the tedious manual steps of starting a new GitHub hosted project. 

You need only provide a new project name and gitgoing will: 

* Create a local directory, ensuring no filename / project conflicts
* Initializes it as a git repo 
* Creates a remote github repo named accordingly
* Sets the remote origin for the local repository 
* Generates boilerplate for a README, LICENSE, .gitignore, etc
* Creates the initial commit and pushes everything to the remote origin

At this point, you should be able to get coding immediately.

## Getting started 

Install the CLI: 
`go install github.com/zackproser/git-going`

Ensure your PATH is configured correctly: 

`which git-going`

should output something like: 

`/home/zachary/go/bin/git-going`

If it doesn't, you need to check your `GOPATH` env variables and ensure your go installation can find packages installed via `go install`. Check out [this guide on configuring your GOPATH](https://github.com/golang/go/wiki/SettingGOPATH) for more information. 

### Authentication 

You need to generate a new GitHub Personal access token. [Read more about access tokens here](https://help.github.com/en/github/authenticating-to-github/creating-a-personal-access-token-for-the-command-line), or go directly to the interface for managing them at [https://github.com/settings/tokens](https://github.com/settings/tokens).

Your token needs sufficient permissions to create repositories on your behalf. 

Once you have generated your token, you need to export it as the environment variable `GIT_GOING_GITHUB_TOKEN.` The gitgoing CLI will panic if you don't have this variable set. 

## Usage 

Create a new project 

`git-going create -n 'My Awesome New Project'`

should output something like: 

```
git going create -n 'My Amazing New Project'
DEBU[0000] Converted My Amazing New Project to slug: my-amazing-new-project 
DEBU[0000] Successfully created My Amazing New Project 
```

In this example, your remote repository would now be available at `https://github.com/<your-username>/my-amazing-new-project.git`. 
