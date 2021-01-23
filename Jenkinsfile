pipeline {
  agent any
  stages {
    stage('并行阶段 1') {
      parallel {
        stage('检出proto') {
          steps {
            sh 'rm -rf ./proto'
            echo 'git clone https://bitbucket.org/east-eden/proto.git'
            sh 'git clone -b ${BIT_BUCKET_PROTO_BRANCH_NAME} --depth=1 https://taowan:MTbtTpCfVwaeRdrHdynd@bitbucket.org/east-eden/proto.git'
          }
        }
        stage('检出server') {
          steps {
            sh 'rm -rf ./server'
            echo 'git clone https://bitbucket.org/east-eden/server.git'
            sh 'git clone -b ${BIT_BUCKET_PROTO_BRANCH_NAME} --depth=1 https://taowan:MTbtTpCfVwaeRdrHdynd@bitbucket.org/east-eden/server.git'
          }
        }
      }
    }
    stage('生成go代码') {
      steps {
        echo '开始生成go code'
        sh 'ls -al'
        sh 'pwd'
        dir('proto') {
          sh 'go env -w GO111MODULE=on'
          sh 'go env -w GOPROXY=https://goproxy.cn,direct'
          sh 'make linux_proto'
        }

      }
    }
    stage('推送到 bitbucket') {
      steps {
        dir('server') {
          sh 'git status'
          sh 'git add proto/*'
          sh 'git config --global user.email "tao.wan@centurygame.com"'
          sh 'git config --global user.name "Tao Wan"'
          sh 'git commit -m "自动生成pb和micropb"'
          sh 'git remote -v'
          sh 'git branch -v'
          sh 'git push origin ${BIT_BUCKET_PROTO_BRANCH_NAME}'
        }

      }
    }
  }
}