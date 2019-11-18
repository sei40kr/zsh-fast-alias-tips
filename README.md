# zsh-fast-alias-tips

[![Build Status](https://travis-ci.com/sei40kr/zsh-fast-alias-tips.svg?branch=master)](https://travis-ci.com/sei40kr/zsh-fast-alias-tips)

A zsh plugin to help remembering those aliases you defined once.
Ported from [djui/alias-tips](https://github.com/djui/alias-tips).

## Example

```sh
$ docker
💡  dk
...

$ git checkout
💡  gco
...

$ git checkout master
💡  gcm
...
```

## Install

### Install with zplugin

```sh
zplugin ice from'gh-r' as'program'
zplugin light sei40kr/fast-alias-tips-bin
zplugin light sei40kr/zsh-fast-alias-tips
```
