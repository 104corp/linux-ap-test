#!/bin/bash

# GitHub 存儲庫 URL
REPO_URL="https://github.com/104corp/k8s-gitops-infra-rancher"

# issue body 樣板
issue_body_template() {
  # 讀取固定路徑的 Markdown 模板文件
  local template_file="template.md"

  if [[ -f "$template_file" ]]; then
    cat "$template_file"
  else
    echo "找不到模板文件: $template_file"
    return 1
  fi
}

# issue title 樣板
issue_title_template() {
  echo "Create namespaced resources for team application"
}

# 創建 issue
create_issue() {
  issue_body=$(issue_body_template)
  if [[ $? -ne 0 ]]; then
    echo "無法讀取 issue body 模板，創建 issue 失敗"
    return 1
  fi

  issue_title=$(issue_title_template)

  # 使用 GitHub CLI 創建 issue
  gh issue create --repo "$REPO_URL" --title "$issue_title" --label "team,create-team-application" --body "$issue_body"
  if [[ $? -ne 0 ]]; then
    echo "創建 issue 失敗"
    return 1
  fi

  echo "Issue 創建成功"
}

# 主程序入口
create_issue
