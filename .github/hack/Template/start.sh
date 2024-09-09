# this is command
#!/bin/sh
TARGET=$GOAL
FILE_PATH="/opt/AP/Patch_program/$APNUMBER/run.sh"
if [[ -z "$FILE_PATH" ]]; then
  echo "FILE_PATH 位置缺失。"
  exit 1
fi
cp "$FILE_PATH" /opt/new_run.sh
if grep -P -n "[\x80-\xBF]" "/opt/new_run.sh" || grep -P -n "[\xFE-\xFF]" "/opt/new_run.sh" ; then
  cp /opt/new_run.sh /opt/new_bigfile_run.sh
  iconv -f BIG5 -t UTF-8//IGNORE /opt/new_bigfile_run.sh > /opt/new_run.sh
fi
sed -i.bak 's/\/usr\/bin\/sudo -u root \$JAVA_HOME\/bin\///g' /opt/new_run.sh
sed -i '/^java/i export JAVA_OPTS="$JAVA_OPTS -DCRYPTO_ADVANCED_ENDPOINT=http://crypto-client.$E-devops-crypto-client.svc.cluster.local"' /opt/new_run.sh
bash /opt/new_run.sh