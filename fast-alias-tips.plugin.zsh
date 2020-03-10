# fast-alias-tips.plugin.zsh
# author: Seong Yong-ju <sei40kr@gmail.com>

: ${FAST_ALIAS_TIPS_PREFIX:="💡 $(tput bold)"}
: ${FAST_ALIAS_TIPS_SUFFIX:="$(tput sgr0)"}

__fast_alias_tips_preexec() {
    local cmd="$1"
    local cmd_expanded="$2"

    local first="$(cut -d' ' -f1 <<<"$cmd")"

    local suggested="$(alias | "def-matcher" "$cmd_expanded")"
    if [[ ${#suggested} -gt ${#first} ]]; then
        return
    fi

    local suggested_first="$(cut -d' ' -f1 <<<"$suggested")"
    if [[ "$suggested_first" == "$first" ]]; then
        return
    fi

    echo "${FAST_ALIAS_TIPS_PREFIX}${suggested}${FAST_ALIAS_TIPS_SUFFIX}"
}

autoload -Uz add-zsh-hook
add-zsh-hook preexec  __fast_alias_tips_preexec
