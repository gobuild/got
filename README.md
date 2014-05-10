
**Got is A binary install tool for Go executable package via gobuild.io.**

** Heavily development, contributings are welcome!**

[中文版(Chinese)](https://github.com/lunny/got/blob/master/README_CN.md)

# Why we need a binary install tool

Since we have `go install`, do we need a binary install tool? I think YES, because I think we are BOTH consumer AND developer. When we decide to test or use some command, we just a consumer, why we need install git or hg, why we need download the source codes and build? Our requirement is only install and use it.

So I make got. Let's got it.

# Install

You can download got from gobuild.io

    wget http://gobuild.io/github.com/lunny/got/master/darwin/amd64 -O output.zip
    unzip output.zip
    cp got /usr/local/bin/

or if you have installed go tool

    go get github.com/lunny/got
    
# How to use

Use got is simple, say we need gopm tool, then

    got github.com/gpmgo/gopm
    
more simpler, for github.com go package we can ignore the domain name

    got gpmgo/gopm
    
and then type `gopm help`

# Well known Go Packages tested

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

Wish you like.