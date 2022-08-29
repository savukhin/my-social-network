# my-social-network

Tell about host variable in db/migration.go:32

## Angular + Golang

This is my <b>Angular</b> + <b>Golang</b> simple social network web application for my own education

## Table of contents
- [my-social-network](#my-social-network)
    - [Pre-requisites](#pre-requisites)
    - [Getting started](#getting-started)
    - [Interface](#interface)

## Pre-requisites
1) [NodeJS](https://nodejs.org/en/) version 16.5.0
2) [PostgreSQL](https://www.postgresql.org/download/) version 14.3
3) [Golang](https://go.dev/) vesion 1.18.3


## Getting started
- Clone the repository:
```
git clone https://github.com/savukhin/my-social-network.git
```
- Set up GOPATH to server folder (in PATH). My variant:
```
export GOPATH=~/github/my-social-network/server
export PATH=$GOPATH/bin:$PATH
```
- Install application dependencies:
```
cd my-social-network
npm install
npm run clientinstall
```
- Create PostgreSQL database "mySocialNetwork" and change configuration in server/src/api/db/connection.go:32. My variant:
```
host := "127.0.0.1"
port := "5432"
user := "postgres"
password := "admin"
dbname := "socialNetwork"
```
- Build and run the project:
```
npm start
```

# Run with docker
- Cange host configuration in server/src/api/db/connection.go:32 to 
```
host := "database"
```
- Run docker compose from the my-social-network folder
```
docker-compose up --build
```

# Interface
## Auth
<img src="Readme files/Login page.png" title="Login page" width="400px"/>
<img src="Readme files/Register page.png" title="Register page" width="400px"/>

## User Page
<img src="Readme files/User page 1.png" title="User page 1" width="400px"/>
<img src="Readme files/Status editing.png" title="R/Status editing" width="400px"/>
<img src="Readme files/User page 2.png" title="CUser page 2" width="400px"/>
<img src="Readme files/Edit user page.png" title="Edit user page" width="400px"/>
<img src="Readme files/User page 3.png" title="User page 3" width="400px"/>

## Chat
<img src="Readme files/Friends page.png" title="Friends page" width="400px"/>
<img src="Readme files/Messages page.png" title="Messages page" width="400px"/>
<img src="Readme files/Chat page 1.png" title="Chat page" width="400px"/>
<img src="Readme files/Chat page 2.png" title="Chat page" width="400px"/>
