![gitgoing logo](./doc/GITGOING.png) 

[![GoDoc](https://godoc.org/github.com/zackproser/git-going?status.svg)](https://godoc.org/github.com/zackproser/git-going)
[![CodeFactor](https://www.codefactor.io/repository/github/zackproser/git-going/badge)](https://www.codefactor.io/repository/github/zackproser/git-going)
[![Go Report Card](https://goreportcard.com/badge/github.com/zackproser/git-going)](https://goreportcard.com/report/github.com/zackproser/git-going)

A Golang project scaffolding utility that I am writing for fun and practice. With a single command, this tool automates all the git setup and project boilerplate generation for a cmd-style Go program.

## Overview

Git-going is a simple command line tool that lets you start your next project very quickly. It automates the tedious manual steps of starting a new GitHub hosted project. 

You need only provide a new project name and git-going will: 

* Create a local directory, ensuring no filename / project conflicts
* Initialize it as a git repo 
* Create a remote github repo named accordingly
* Set the remote origin for the local repository 
* Generate boilerplate for a README, LICENSE, .gitignore, etc
* Create the initial commit
* Push to the remote origin, setting upstream tracking

## Roadmap

Aspirationally, this tool will create the ideal starting point for a Golang project with a single command. This includes: 

* GitHub repository continuous integration configuration
* GitHub actions configuration
* GitHub release automation configuration
* Readme badges for go report card, builds, etc 
* Boilerplate CLI project setup, generated GO, project layout, etc 

## Getting started 

### Install via go get 

Install the CLI: 
`go install github.com/zackproser/git-going`

Ensure your PATH is configured correctly: 

`which git-going`

should output something like: 

`/home/zachary/go/bin/git-going`

If it doesn't, you need to check your `GOPATH` env variables and ensure your go installation can find packages installed via `go install`. Check out [this guide on configuring your GOPATH](https://github.com/golang/go/wiki/SettingGOPATH) for more information. 

### Install binary directly 

Visit [this project's Releases](https://github.com/zackproser/git-going/releases) and download the latest binary for your platform. 

Then extract the archive and move the binary somewhere within your `GOPATH`: 

```
tar xvzf ~/Downloads/git-going_1.1.0_Linux_x86_64.tar.gz -C ~/Downloads/git-going_1.1.0
cp ~/Downloads/git-going_1.1.0/git-going ~/go/bin/
```
Verify your installation: 

```
git-going help
```

### Authentication 

You need to generate a new GitHub Personal access token. [Read more about access tokens here](https://help.github.com/en/github/authenticating-to-github/creating-a-personal-access-token-for-the-command-line), or go directly to the interface for managing them at [https://github.com/settings/tokens](https://github.com/settings/tokens).

Your token needs sufficient permissions to create repositories on your behalf. 

Once you have generated your token, you need to export it as the environment variable `GIT_GOING_GITHUB_TOKEN.` The git-going CLI will panic if you don't have this variable set. 

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
