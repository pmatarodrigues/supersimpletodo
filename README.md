# Super Simple Todo

## Installation

- Clone this project

$ `git clone git@github.com:pmatarodrigues/supersimpletodo.git`

- Add folder to PATH
(This next command depends on whether you use bash or zsh or any other shell)

#### Bash
$ `echo "export PATH=\$PATH:$(pwd)" >> ~/.bashrc`

#### Zsh
$ `echo "export PATH=\$PATH:$(pwd)" >> ~/.zshrc`


## Usage

- View todo list:

$ `sst`

- Add item:

$ `sst <project> "<item>" `

#### Example:

$ `sst "personal projects" "add docs to supersimpletodo"`


## Development

- Build:

$ `go build`


- Run:

$ `go run .`