#!/bin/bash

# 如果沒有提供文件路徑，則給出提示並退出
if [[ -z "$1" ]]; then
  echo "AP 號碼缺失。"
  exit 1
fi
if [[ -z "$2" ]]; then
  echo "repo 位置缺失。"
  exit 1
fi
if [[ -z "$3" ]]; then
  echo "環境缺失。"
  exit 1
fi
if [[ -z "$4" ]]; then
  echo "cluster 缺失。"
  exit 1
fi
# 獲取命令行參數中指定的文件路徑
APNUM="$1"
REPO="$2"
ENV="$3"
CLUSTER="$4"

export GOAL="$CLUSTER"
export JAVA_HOME="\$JAVA_HOME"
export FILE_PATH="\$FILE_PATH"
export JAVA_OPTS="\$JAVA_OPTS"
export TARGET="\$TARGET"
export check_dr_ap="\$check_dr_ap"
export APNUMBER=$(echo $APNUM | tr '-' '_')

if [ "$ENV" == "lab" ]; then
  export E="l"
elif [ "$ENV" == "stg" ]; then
  export E="s"
elif [ "$ENV" == "prod" ]; then
  export E="p"
fi

envsubst < .github/hack/Template/start.sh > ~/content.sh

export SCRIPT_CONTENT="$(sed 's/^/        /' ~/content.sh)"
cat ~/content.sh
envsubst < output.yaml > "$REPO"/ap-"$APNUM"/overlays/"$ENV"/"$CLUSTER".values.yaml
cat output.yaml
cat "$REPO"/ap-"$APNUM"/overlays/"$ENV"/"$CLUSTER".values.yaml
sed -i '/# this is command/d' "$REPO"/ap-"$APNUM"/overlays/"$ENV"/"$CLUSTER".values.yaml
echo "create success"