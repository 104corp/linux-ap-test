#!/bin/bash

# 獲取命令行參數中指定的文件路徑
FILE_PATH="$1"
cp "$FILE_PATH" ~/new_run.sh

# 如果沒有提供文件路徑，則給出提示並退出
if [[ -z "$FILE_PATH" ]]; then
  echo "請提供文件路徑作為腳本的第一個參數。"
  exit 1
fi

# 使用 grep 檢查是否為 UTF-8 編碼文件
# shellcheck disable=SC2088
if grep -P -n "[\x80-\xBF]" "~/new_run.sh" || grep -P -n "[\xFE-\xFF]" "~/new_run.sh" ; then
  cp new_run.sh new_bigfile_run.sh
  iconv -f BIG5 -t UTF-8//IGNORE ~/new_bigfile_run.sh > ~/new_run.sh
fi
cat ~/new_run.sh
sed -i.bak 's/\/usr\/bin\/sudo -u root \$JAVA_HOME\/bin\///g' ~/new_run.sh
#sed -i '/cd \$AP_PROG\/\$APID/d' ~/new_run.sh
sed -i '1s/^/# this is command\n/' ~/new_run.sh
sed -i '/^$/d' ~/new_run.sh
sed -i 's/\$APID\///g' ~/new_run.sh
sed -i 's/\$APID//g' ~/new_run.sh
