pipeline {
  agent any
  stages {
    stage('检出proto') {
      steps {
        checkout([
          $class: 'GitSCM',
          branches: [[name: env.GIT_CODE_DEV_BRANCH_NAME]],
          userRemoteConfigs: [[
            url: env.GIT_REPO_URL,
            credentialsId: env.CREDENTIALS_ID
          ]]])
        }
      }
      stage('生成go代码') {
        steps {
          echo '开始生成go code'
          sh 'ls -al'
          sh 'go env -w GO111MODULE=on'
          sh 'go env -w GOPROXY=https://goproxy.cn,direct'
          sh 'go env'
          sh 'make proto'
          sh 'ls -al ./go_out'
          sh 'rm -rf ./go_out/e.coding.net/*'
          sh 'git add *'
          sh 'git commit -m "自动生成pb和micropb"'
          sh 'git remote -v'
          sh 'git branch -v'
        }
      }
      stage('推送到 CODING') {
        steps {
          sh "git push https://${PROJECT_TOKEN_GK}:${PROJECT_TOKEN}@e.coding.net/mmstudio/blade/proto.git HEAD:dev"
        }
      }
    }
  }