#!/bin/bash

echo "Team: $team"
echo "Cluster: $cluster"
echo "app: $app"

if [ -z "$team" ]; then
  echo "team is not set"
  exit 1
fi

if [ -z "$cluster" ]; then
  echo "cluster is not set"
  exit 1
fi

if [ -z "$app" ]; then
  echo "app is not set"
  exit 1
fi

# 宣告關聯陣列
declare -A env

# 為字典添加鍵值對
env["lab"]="lab"
env["stg"]="stg"
env["prod"]="prod"
env["gke31"]="prod"

cd k8s-gitops-infra-rancher
goam infra team create-app-namespace $app --team $team --cluster $cluster --env ${env[$cluster]} --hierarchical-namespace --sealed-secret-placeholder=false