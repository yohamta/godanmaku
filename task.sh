#!/bin/bash

SCRIPTS_DIR=$(dirname "$0")/scripts

function install {
    echo "install task not implemented"
}

function build {
    echo "build task not implemented"
}

function build-ios {
    ebitenmobile bind -target ios -o ./mobile/ios/Mobile.framework ./mobile
}

function update {
    bash build-resources.sh
}

function start {
    go run main.go
}

function default {
    start
}

function help {
    echo "$0 <task> <args>"
    echo "Tasks:"
    compgen -A function | cat -n
}

TIMEFORMAT="Task completed in %3lR"
time ${@:-default}
