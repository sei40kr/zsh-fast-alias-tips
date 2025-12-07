# zsh-fast-alias-tips

Helps you remembering the aliases you defined once.

Written in zsh and Rust. Ported from [djui/alias-tips](https://github.com/djui/alias-tips).

## Example

```
$ alias gst='git status'

$ git status
💡  gst
On branch master
Your branch is up to date with 'origin/master'.

nothing to commit, working tree clean
```

## Install

### Requirements

- Rust (cargo)
- zsh

### Install with [zinit](https://github.com/zdharma/zinit)

```sh
zinit ice atclone'cargo build --release' atpull'%atclone'
zinit light sei40kr/zsh-fast-alias-tips
```

The plugin will automatically build the `alias-matcher` binary during installation and updates.

## Customization

| Variable                 | Default value       | Description           |
| :--                      | :--                 | :--                   |
| `FAST_ALIAS_TIPS_PREFIX` | `"💡 $(tput bold)"` | The prefix of the Tips |
| `FAST_ALIAS_TIPS_SUFFIX` | `"$(tput sgr0)"`    | The suffix of the Tips |
