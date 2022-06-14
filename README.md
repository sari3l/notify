<h1 align="center"><img src="https://raw.githubusercontent.com/sari3l/notify/main/static/logo.png" alt="Logo"/></h1>

[![license](https://img.shields.io/github/license/sari3l/notify?style=flat-square)](https://github.com/sari3l/notify/LIENCES)
[![Go Report Card](https://goreportcard.com/badge/github.com/sari3l/notify)](https://goreportcard.com/report/github.com/sari3l/notify)
[![CodeFactor](https://www.codefactor.io/repository/github/sari3l/notify/badge)](https://www.codefactor.io/repository/github/sari3l/notify)

## Usage

```shell
> ./notify -h       
Usage:
  notify [OPTIONS]

Application Options:
  -s, --show        显示配置信息
  -c, --config=     指定配置文件, 如: $HOME/.config/notify-config.yaml
  -i, --id=         指定通知id
  -l, --level=      指定通告等级 (default: 0)
  -a, --aboveLevel= 指定最低通告等级
  -n, --noticer=    指定通知模块
  -v, --version     版本信息

Help Options:
  -h, --help        Show this help message
```

## Installation

```shell
go install -v github.com/sari3l/notify/cmd/notify@latest
```

## Features

-  管道符输入
-  指定id、level、模块，或最低level广播

## Licenses

[GPL v3.0](https://github.com/sari3l/nofity/blob/main/LICENSE)