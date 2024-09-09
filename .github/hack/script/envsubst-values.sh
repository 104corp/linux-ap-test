#!/bin/bash
# 如果沒有提供文件路徑，則給出提示並退出
if [[ -z "$1" ]]; then
  echo "linuxap host 缺失。"
  exit 1
fi
if [[ -z "$2" ]]; then
  echo "AP 號碼缺失。"
  exit 1
fi
if [[ -z "$3" ]]; then
  echo "repo 位置缺失。"
  exit 1
fi
if [[ -z "$4" ]]; then
  echo "環境缺失。"
  exit 1
fi
if [[ -z "$5" ]]; then
  echo "cluster 缺失。"
  exit 1
fi
# 獲取命令行參數中指定的文件路徑
HOST="$1"
APNUM="$2"
REPO="$3"
ENV="$4"
CLUSTER="$5"
APNUM_CLEAN=$(echo $APNUM | tr '-' '_')

cat output.yaml
.github/hack/script/bigfile.sh /opt/AP/Patch_program/"$HOST"/Patch_program/"$APNUM_CLEAN"/run.sh
export SCRIPT_CONTENT="$(sed 's/^/        /' ~/new_run.sh)"
envsubst < output.yaml > "$REPO"/ap-"$APNUM"/overlays/"$ENV"/"$CLUSTER".values.yaml
sed -i '/# this is command/d' "$REPO"/ap-"$APNUM"/overlays/"$ENV"/"$CLUSTER".values.yaml
echo "create success"
cat "$REPO"/ap-"$APNUM"/overlays/"$ENV"/"$CLUSTER".values.yaml