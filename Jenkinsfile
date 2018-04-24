pipeline {
  agent {
    docker {
      image 'golang:1.9.2-stretch'
    }

  }
  stages {
    stage('') {
      steps {
        sh '''go version
go get -u github.com/golang/dep/cmd/dep
go get -u github.com/golang/lint/golint
go get github.com/tebeka/go2xunit\'
cd $GOPATH/src/cmd/project && dep ensure'''
      }
    }
  }
}