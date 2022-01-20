#!/bin/bash 
#this files directory 
base_path=$(cd `dirname $0`; pwd)
echo "current directory are: $base_path"
app_mod="menu"

cd $base_path"/app"

mkdir $app_mod

cd $app_mod


mkdir $app_mod"_config"

mkdir $app_mod"_dao"
echo package $app_mod"_dao" >$base_path"/app/"$app_mod"/"$app_mod"_dao/"dao.go

mkdir $app_mod"_control"
echo package $app_mod"_control" >$base_path"/app/"$app_mod"/"$app_mod"_control/"control.go

mkdir $app_mod"_model"
echo package $app_mod"_model" >$base_path"/app/"$app_mod"/"$app_mod"_model/"model.go

mkdir $app_mod"_cache"
echo package $app_mod"_cache" >$base_path"/app/"$app_mod"/"$app_mod"_cache/"cache.go

mkdir $app_mod"_service"
echo package $app_mod"_service" >$base_path"/app/"$app_mod"/"$app_mod"_service/"service.go

mkdir $app_mod"_router"
echo package $app_mod"_router" >$base_path"/app/"$app_mod"/"$app_mod"_router/"router.go

mkdir $app_mod"_param"
echo package $app_mod"_param" >$base_path"/app/"$app_mod"/"$app_mod"_param/"param.go

mkdir $app_mod"_const"
echo package $app_mod"_const" >$base_path"/app/"$app_mod"/"$app_mod"_const/"const.go

cmd="cmd"
mkdir $cmd
cd $cmd

mkdir "rpc"
mkdir "web"