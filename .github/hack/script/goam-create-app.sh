#!/bin/bash

# 檢查是否提供了足夠的參數
if [ "$#" -ne 5 ]; then
    echo "使用方法: $0 <location> <apnum> <env> <team> <repo>"
    exit 1
fi

# 輸入參數
LOCATION=$1
APNUM=$2
ENV=$3
TEAM=$4
REPO=$5

# 從 cronjob-version.yaml 中讀取 version
CRONJOB_VERSION=$(grep 'version:' .github/hack/embedFS/cronjob-version.yaml | awk '{print $2}' | tr -d '"')

# shellcheck disable=SC2164
cd "$REPO"

# 如果環境是 lab，則執行 goam 命令
if [ "$ENV" == "lab" ]; then
    goam app create-appset ap-$APNUM --team $TEAM --repo https://github.com/104corp/$REPO.git --clusters $LOCATION --env $ENV --helm-subchart cronjob:$CRONJOB_VERSION --sync-auto
else
    goam app create-appset ap-$APNUM --team $TEAM --repo https://github.com/104corp/$REPO.git --clusters $LOCATION --env $ENV --helm-subchart cronjob:$CRONJOB_VERSION
fi