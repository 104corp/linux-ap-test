#!/bin/bash
# 應用名稱列表
pvcs=("program" "lib" "conf" "log" "output")
#apps=("wsm")
REPO_URL="https://github.com/104corp/k8s-gitops-infra-rancher"

# issue body 樣板
issue_body_template() {
  # 讀取固定路徑的 Markdown 模板文件
  local template_file="$1.md"

  if [[ -f "$template_file" ]]; then
    cat "$template_file"
  else
    echo "找不到模板文件: $template_file"
    return 1
  fi
}

# issue title 樣板
issue_title_template() {
  echo "Create static pvc for the $1"
}

# 創建 issue
create_issue() {
  pvc=$1
  issue_body=$(issue_body_template "$pvc")
  issue_title=$(issue_title_template "$pvc")

  # 使用 GitHub CLI 創建 issue
  gh issue create --repo "$REPO_URL" --title "$issue_title" --label "team,create-team-app-static-pvc" --body "$issue_body"
  if [[ $? -ne 0 ]]; then
    echo "創建 issue 失敗"
    return 1
  fi

  echo "Issue 創建成功"
}

# 遍歷應用名稱列表創建 issue
echo "Creating issues for static pvc: $1"
create_issue "$1"
