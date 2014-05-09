
**Got is A binary install tool for Go executable package via gobuild.io.**

** Heavily development, contributings are welcome!**

# Why we need a binary install tool

Since we have `go install`, do we need a binary install tool? I think YES, because I think we are BOTH consumer AND developer. When we decide to test or use some command, we just a consumer, why we need install git or hg, why we need download the source codes and build? Our requirement is only install and use it.

So I make got. Let's got it.

# Install

You can download got from gobuild.io

    wget http://gobuild.io/github.com/lunny/got/master/darwin/amd64 -O output.zip
    unzpi output.zip
    cp got /usr/local/bin/

or if you have installed go tool

    go install github.com/lunny/got
    
# How to use

Use got is simple, say we need xorm tool, then

    got github.com/go-xorm/cmd/xorm
    
more simpler, for github.com go package we can ignore the domain name

    got go-xorm/cmd/xorm
    
and then type `xorm help`

Wish you like.