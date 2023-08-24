#!/bin/bash

# for dir in service/*; do
#     sh $dir/build.sh
#     svr=$(basename $dir)
#     dir/output/bin/svr
# done

dir=$(pwd)

cd $dir/service/user && sh build.sh
cd $dir/service/control && sh build.sh

cd $dir
service/user/output/bin/user&
service/control/output/bin/control&
