# Project Setup

I'll be using Ubuntu 20.04 via WSL throughout this project. 

## Installing Go

The latest stable release of go as of the time of writing is 1.19.3, which we will install by:

1. Download and install go binary
```bash
$ mkdir -p ~/downloads && cd ~/downloads && wget https://golang.org/dl/go1.19.3.linux-amd64.tar.gz
$ sudo tar -C /usr/local -xvf go1.19.3.linux-amd64.tar.gz
```

2. We then set the go paths and refresh the profile
```bash
$ sudo echo "export PATH=$PATH:/usr/local/go/bin\n" >> ~/.profile
$ source ~/.profile
```

3. Finally, we check whether the install was successful via
```bash
$ go version
```

