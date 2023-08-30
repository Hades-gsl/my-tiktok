#!/bin/bash

dir=$(pwd)
svrs=("service/feed" "service/user" "service/publish" "service/control")

if [[ "$1" == "build" ]]; then
    for svr in "${svrs[@]}"; do
        cd "$dir/$svr" && sh build.sh
        echo build $(basename "$svr") success
    done
elif [[ "$1" == "run" ]]; then
    for svr in "${svrs[@]}"; do
        cd "$dir/$svr" && sh build.sh
        echo build $(basename "$svr") success
    done
    for svr in "${svrs[@]}"; do
        cd "$dir/$svr" && output/bin/$(basename "$svr")&
        echo run $(basename "$svr") success
    done
elif [[ "$1" == "end" ]]; then
    for svr in "${svrs[@]}"; do
        pid=$(ps aux | grep output/bin/$(basename "$svr") | grep -v grep | awk '{print $2}')
        kill $pid
        echo kill $(basename "$svr") success
    done
fi

