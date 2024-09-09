#!/bin/bash

# 輸入參數
LOCATION=$1
APNUM=$2
NAMESPACE=$3

APNUM_CLEAN=$(echo $APNUM | tr '-' '_')

# 創建目錄
mkdir -p /opt/AP/${LOCATION}/Patch_output/${APNUM_CLEAN}
mkdir -p /opt/AP/${LOCATION}/Patch_log/${APNUM_CLEAN}

# 設置 RUN_AS_USER 變量
RUN_AS_USER=$(./.github/hack/script/hash-id.sh ${NAMESPACE})

# 打印創建成功信息
echo "create success RUN_AS_USER: $RUN_AS_USER"
echo $RUN_AS_USER
echo "/opt/AP/${LOCATION}/Patch_output/${APNUM_CLEAN}"

# 修改目錄所有者
sudo chown $RUN_AS_USER:root /opt/AP/${LOCATION}/Patch_output/${APNUM_CLEAN}
sudo chown $RUN_AS_USER:root /opt/AP/${LOCATION}/Patch_log/${APNUM_CLEAN}
