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
          sh 'git add ./go_out/*'
          sh 'git commit -m "自动生成pb和micropb"'
          sh 'git push origin dev'
        }
      }
    }
  }