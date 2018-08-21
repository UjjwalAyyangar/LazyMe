# LazyMe
Script to make repo creation faster for lazy developers like me.

## Installation
```sh
> go get github.com/UjjwalAyyangar/lazyme
```

## Usage

```sh
> lazyme
Enter GitHub Username or Email : shubhodeep9
Enter GitHub Password :
Enter the details for the repository you want to create

Name of the repository : sometest
Do you want it to be private? : [n]
Give a one line description for your Repository : testing
Do you want to initialize this repository with a README? : [n] y
Yohoo! Repo created, visit https://github.com/shubhodeep9/sometest

Do you want to clone the repository? : [y] y
Yay, you may visit the directory
Bye!
```

## Development
Instructions to setup development environment for contribution:

- Clone the repository in `$GOPATH/src/github.com/UjjwalAyyangar/`
- Install dependencies via `curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && dep ensure`

## Contributors
[@UjjwalAyyangar](https://github.com/UjjwalAyyangar) · [@shubhodeep9](https://github.com/shubhodeep9)
