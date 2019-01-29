[![CircleCI](https://circleci.com/gh/nesheep5/checo.svg?style=svg)](https://circleci.com/gh/nesheep5/checo)
# checo
Checo is CLI to check if an SNS account exists.
  
Check SNS:   
 Twitter / Facebook / Instagram  

## Installation or Download
### Installation
```
$ go get github.com/nesheep5/checo/cmd/checo
```

### Download
You can also download and use the executable file.  
Available OS : Linux / Mac / Windows  
  
https://github.com/nesheep5/checo/releases/tag/lastest

## Usage
```
NAME:
   checo - Function to check if an SNS account exists.

USAGE:
   checo [global options] command [command options] [arguments...]

VERSION:
   v1.0.0

COMMANDS:
     list, l   Shows checking SNS list
     check, c  Checking SNS account
     help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## Example
```
$ checo check nesheep5
Search Account: nesheep5

Twitter : Oops! Alrady Exists...
Facebook : OK! No Exists!
Instagram : OK! No Exists!
```
