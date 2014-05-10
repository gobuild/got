
**Got 是一个基于gobuild.io的Go二进制包管理工具**

** 当前正在紧张开发中，欢迎贡献代码**

[English(英文版)](https://github.com/lunny/got/blob/master/README.md)

# 为什么我们需要一个二进制的Go包管理工具

对于开发者，其实一直在使用`go install`，那我们还需要一个二进制的Go包管理工具吗？我认为需要。因为开发者也是使用者，当你在使用别人的成熟的工具时，并不需要去查看他的代码，更不需要为此而安装Go，安装git，hg，也不需要编译。这个时候开发者变成了使用者。

因此我基于gobuild.io创建了got。让我们开始吧。

# 安装

你可以从gobuild.io获得 got：

    wget http://gobuild.io/github.com/lunny/got/master/darwin/amd64 -O output.zip
    unzip output.zip
    cp got /usr/local/bin/

或者如果你有安装go的话，那么直接

    go get github.com/lunny/got
    
# 如何使用

使用got非常简单，比如我们需要安装gopm，那么只需要

    got github.com/gpmgo/gopm
    
就好了，对于github.com的包还有更简单的方式，可以省去域名:

    got gpmgo/gopm
    
然后，就可以直接使用 `gopm help`

# 已经经过测试的包如下，从此可以删除掉这些包对应的源码了。

    got nsf/gocode
    got beego/bee
    got gpmgo/gopm
    got bradfitz/goimports
    got mitchellh/gox
    got wendal/gor
    got laher/goxc
    got parkghost/gohttpbench
    got shxsun/fswatch
    got tools/godep
    got mattn/gom
    got codegangsta/gin
    got codeskyblue/gobuild
    got zachlatta/postman
    got coreos/etcd
    got hashicorp/serf
    got FiloSottile/Heartbleed
    got cyfdecyf/cow
    got apcera/gnatsd
    got shenfeng/http-watcher
    got nf/goplayer
    got piranha/goreplace
    got mtourne/gurl

希望大家喜欢。