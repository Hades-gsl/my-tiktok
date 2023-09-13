#!/bin/sh

dir=$(pwd)
svrs="service/feed service/user service/publish service/favorite service/control"

if [ "$1" = "build" ]; then
    for svr in $svrs; do
        (cd "$dir/$svr" && sh build.sh)
        echo "build $(basename "$svr") success"
    done
elif [ "$1" = "run" ]; then
    for svr in $svrs; do
        (cd "$dir/$svr" && output/bin/$(basename "$svr")&)
        echo "run $(basename "$svr") success"
    done
elif [ "$1" = "end" ]; then
    for svr in $svrs; do
        pid=$(ps aux | grep "output/bin/$(basename "$svr")" | grep -v grep | awk '{print $2}')
        kill $pid
        echo "kill $(basename "$svr") success"
    done
elif [ "$1" = "clean" ]; then
    for svr in $svrs; do
        (cd "$dir/$svr" && rm -rf output)
    done
elif [ "$1" = "debug" ]; then
    for svr in $svrs; do
        (cd "$dir/$svr" && sh build.sh)
        echo "build $(basename "$svr") success"
    done
    for svr in $svrs; do
        (cd "$dir/$svr" && output/bin/$(basename "$svr")&)
        echo "run $(basename "$svr") success"
    done
fi
