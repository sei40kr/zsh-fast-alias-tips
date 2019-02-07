# fast-alias-tips.plugin.zsh
# author: Seong Yong-ju <sei40kr@gmail.com>

FAST_ALIAS_TIPS_PREFIX="💡 $(tput bold)"
FAST_ALIAS_TIPS_SUFFIX="$(tput sgr0)"

if [[ ! -L "${0:a}" ]]; then
    __fast_alias_tips_dir="${0:a:h}"
else
    __fast_alias_tips_dir="$(readlink "${0:a}")"
    __fast_alias_tips_dir="${__fast_alias_tips_dir:h}"
fi

__fast_alias_tips_preexec() {
    local cmd="$1"
    local cmd_expanded="$2"

    local first="$(cut -d' ' -f1 <<<"$cmd")"

    local suggested="$(alias | "${__fast_alias_tips_dir}/build/def-matcher" "$cmd_expanded")"
    if [[ "$suggested" == '' ]]; then
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
