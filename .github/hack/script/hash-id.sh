#!/bin/bash

# 設定命名空間名稱
namespace="$1"

# 檢查是否有提供命名空間名稱
if [ -z "$namespace" ]; then
    echo "請提供命名空間名稱作為參數"
    exit 1
fi

# 計算命名空間的 SHA1 哈希並截取前7位
sha_hex=$(echo -n $namespace | sha1sum | cut -c 1-7)

# 將截取的十六進制數轉換為十進制數
run_as_user=$((16#$sha_hex))

# 將 runAsUser 的值設定為環境變數
echo $run_as_user
