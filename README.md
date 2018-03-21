## battlesnake-go

A simple [BattleSnake AI](http://battlesnake.io) written in Go. This was my first year participating in the competition, and the snake did moderately well (I got to the finals of the Beginner category).  

The main issue with the snake currently, is that it doesn't check to make sure a route has a viable exit, before choosing it as it's target.

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

## Setup Instructions

Visit [battlesnake.io/readme](http://battlesnake.io/readme) for API documentation and instructions for running your AI.

To get started, you'll need:
  1. A working Go development environment ([guide](https://golang.org/doc/install)).
  2. Read [Heroku's guide to deploying Go apps](https://devcenter.heroku.com/articles/getting-started-with-go#introduction)

### Running the AI locally

1) [Fork this repo](https://github.com/sendwithus/battlesnake-go/fork).

2) Clone repo to your development environment:
```
git clone git@github.com:USERNAME/battlesnake-go.git $GOPATH/github.com/USERNAME/battlesnake-go
cd $GOPATH/github.com/USERNAME/battlesnake-go
```

3) Compile the battlesnake-go server.
```
go build
```
This will create a `battlesnake-go` executable.

4) Run the server.
```
./battlesnake-go
```

5) Test the client in your browser: [http://127.0.0.1:9000/start](http://127.0.0.1:9000/start)


### Deploying to Heroku

1) Create a new Go Heroku app using Go buildpack.
```
heroku create
```

2) Push code to Heroku servers.
```
git push heroku master
```

3) Open Heroku app in browser.
```
heroku open
```
