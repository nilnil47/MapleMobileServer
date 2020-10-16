## installation

For the generating the grpc files you will need:
* protoc compiler
* go-grpc plugin
* make

###Install For Windows:

Enable chocolaty (chocolaty is package manager for windows)

From Powershell with **admin** 
```
Set-ExecutionPolicy Bypass -Scope Process -Force; `
  iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))
```

From Powershell with **admin** 
```
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

 
## using the heroku app 

if you don't familiar with heroku checkout [this link](https://devcenter.heroku.com/articles/getting-started-with-go#set-up) 
download the heroku cli from [heroku websitehttps](//devcenter.heroku.com/articles/heroku-cli)

```
# login to heroku
heroku login

# push branch to heroku
git push heroku <branch>
```
