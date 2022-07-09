#! /bin/bash
go install github.com/cosmtrek/air@latest

#检查本地是否存在mysql，若不存在则创建mysql

#检查本地是否存在redis数据库，若不存在则创建redis

# 运行make install命令
make install