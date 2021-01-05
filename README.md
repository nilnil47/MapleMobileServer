## installation

For the generating the grpc files you will need:
* protoc compiler
* go-grpc plugin
* make

### Install For Windows:

Enable chocolaty (chocolaty is package manager for windows)

From Powershell with **admin** 
```
Set-ExecutionPolicy Bypass -Scope Process -Force; `
  iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))
```

From Powershell with **admin** 
```

# install go
choch install golang

# install protoc
choco install protoc --pre

# install make
choco install make 

# install go-grpc plugin
go get github.com/golang/protobuf/protoc-gen-go

# install mongodb
choco install mongodb

# install robo3t - a gui tool for mongo db
choco install robo3t
```

To use mongodb from cli you need to add the mongodb bin directory to
the PATH environment variable

### install for ubunto
```
sudo snap install go
```

## download the module dependencies 
```
cd <root dir> 
go mod download
```

## compile proto files
```
make grpc
```

## run the server
```
go run main/main.go
```

## using the heroku app

At the time this readme has written heroku has not have support for http2, so it does not possible to use
heroku for the maple server. I will keep this heroku snippet code just because it has already written.   

if you don't familiar with heroku checkout [this link](https://devcenter.heroku.com/articles/getting-started-with-go#set-up) 
download the heroku cli from [heroku websitehttps](https://devcenter.heroku.com/articles/heroku-cli)

```
# login to heroku
heroku login

# if you did not alredy created heroku app
heroku create

# push branch to heroku
git push heroku <branch>
```

## kill the server on windows 
```
Get-Process -Id (Get-NetTCPConnection -LocalPort 80).OwningProcess
taskkill /PID <PID> /F
```
