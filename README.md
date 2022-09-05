# my-social-network

Tell about host variable in db/migration.go:32

## Angular + Golang

This is my **Angular** + **Golang** simple social network web application for my own education

## Table of contents

- [my-social-network](#my-social-network)
  - [Pre-requisites](#pre-requisites)
  - [Getting started](#getting-started)
    - [Run with docker](#run-with-docker)
    - [Run in terminal](#run-in-terminal)
  - [Interface](#interface)

## Pre-requisites

1) [NodeJS](https://nodejs.org/en/) version 16.5.0
2) [PostgreSQL](https://www.postgresql.org/download/) version 14.3
3) [Golang](https://go.dev/) vesion 1.18.3


## Getting started

- Clone the repository:

```[bash]
git clone https://github.com/savukhin/my-social-network.git
```

### Run with docker

- Run docker compose from the my-social-network folder

```[bash]
docker-compose up --build
```

### Run in terminal

- Set up GOPATH to server folder (in PATH). My variant:

```[bash]
export GOPATH=~/github/my-social-network/server
export PATH=$GOPATH/bin:$PATH
```

- Install application dependencies:

```[bash]
cd my-social-network
npm install
npm run clientinstall
```

- Create PostgreSQL database "mySocialNetwork" and change configuration in server/src/api/.env file for your database

- Build and run the project:

```[bash]
npm start
```

## Interface

### Auth

<img src="Readme files/Login page.png" title="Login page" width="400px"/>
<img src="Readme files/Register page.png" title="Register page" width="400px"/>

### User Page

<img src="Readme files/User page 1.png" title="User page 1" width="400px"/>
<img src="Readme files/Status editing.png" title="R/Status editing" width="400px"/>
<img src="Readme files/User page 2.png" title="CUser page 2" width="400px"/>
<img src="Readme files/Edit user page.png" title="Edit user page" width="400px"/>
<img src="Readme files/User page 3.png" title="User page 3" width="400px"/>

### Chat

<img src="Readme files/Friends page.png" title="Friends page" width="400px"/>
<img src="Readme files/Messages page.png" title="Messages page" width="400px"/>
<img src="Readme files/Chat page 1.png" title="Chat page" width="400px"/>
<img src="Readme files/Chat page 2.png" title="Chat page" width="400px"/>
