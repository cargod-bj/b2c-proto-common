
<div align=center>
<h1>b2c-proto-common</h1>
</div>
<div align=center>
<img src="https://img.shields.io/badge/golang-1.14-blue"/>
<img src="https://img.shields.io/badge/protoc-v*.*-red"/>
</div>

# Project Guidelines
- Framework：[protobuf](https://github.com/protocolbuffers/protobuf)
## 1. Getting started
### 1.1 根据操作系统下载对应版本protoc 安装文件
- 下载地址：[Protocol Buffers](https://github.com/protocolbuffers/protobuf/releases)
### 1.2 阅读安装文件中的readme 按照各操作系统安装指南进行安装，例如linux mac：
```
Protocol Buffers - Google's data interchange format
Copyright 2008 Google Inc.
https://developers.google.com/protocol-buffers/

This package contains a precompiled binary version of the protocol buffer
compiler (protoc). This binary is intended for users who want to use Protocol
Buffers in languages other than C++ but do not want to compile protoc
themselves. To install, simply place this binary somewhere in your PATH.

If you intend to use the included well known types then don't forget to
copy the contents of the 'include' directory somewhere as well, for example
into '/usr/local/include/'.

Please refer to our official github site for more installation instructions:
https://github.com/protocolbuffers/protobuf
```
### 1.3 其他
- protobuf
- protoc-gen-go
- protoc-gen-micro
- Getting Started：[Go-Micro v2](https://micro.mu/getting-started)

## 踩坑指南
- prto 方法名必须大写，不然不识别
- protobuf一定要注意版本问题
- 运行项目出现 undefined: balancer.PickOptions 可以运行go get google.golang.org/grpc@v1.26.0 指定grpc版本



