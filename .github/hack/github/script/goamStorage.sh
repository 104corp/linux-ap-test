#!/bin/bash

echo "Team: $team"
echo "Env: $env"
echo "App: $name"
echo "Volume: $volume"
echo "Bu: $bu"
echo "Server: $server"
eval path="$path"
echo "Path: $path"
echo "LinuxAp: $linuxap"

if [ -z "$team" ]; then
  echo "team is not set"
  exit 1
fi

if [ -z "$env" ]; then
  echo "env is not set"
  exit 1
fi

if [ -z "$name" ]; then
  echo "name is not set"
  exit 1
fi

if [ -z "$volume" ]; then
  echo "volume is not set"
  exit 1
fi

if [ -z "$bu" ]; then
  echo "bu is not set"
  exit 1
fi

if [ -z "$server" ]; then
  echo "server is not set"
  exit 1
fi

if [ -z "$path" ]; then
  echo "path is not set"
  exit 1
fi

declare -A clusters
clusters["lab"]="lab"
clusters["stg"]="stg"
clusters["prod"]="prod"
clusters["dr"]="gke31"

# 宣告關聯陣列
declare -A envs

# 為字典添加鍵值對
envs["lab"]="l"
envs["stg"]="s"
envs["prod"]="p"
envs["dr"]="p"
namespace="${envs[$env]}-$team-$name"

declare -A environments
environments["lab"]="lab"
environments["stg"]="stg"
environments["prod"]="prod"
environments["dr"]="prod"

cd k8s-gitops-infra-rancher
goam infra team create-app-static-pvc $volume --namespace $namespace --env ${environments[$env]} --cluster ${clusters[$env]} --business-unit $bu --storage "30Gi" --access-mode "ReadWriteMany" --nfs-server $server --nfs-path "$path"