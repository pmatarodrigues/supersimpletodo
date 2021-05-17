# Super Simple Todo

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

## Installation

#### Clone this project

$ `git clone git@github.com:pmatarodrigues/supersimpletodo.git`

#### Add folder to PATH
(This next command depends on whether you use bash or zsh or any other shell)

- Bash

&emsp; $ `echo "export PATH=\$PATH:$(pwd)" >> ~/.bashrc`

- Zsh

&emsp; $ `echo "export PATH=\$PATH:$(pwd)" >> ~/.zshrc`


## Usage

#### View todo list:

$ `sst`

#### Add item:

$ `sst <project> "<item>" `

- Example:

&emsp; $ `sst "personal projects" "add docs to supersimpletodo"`


## Development

#### Build:

$ `go build`


#### Run:

$ `go run .`