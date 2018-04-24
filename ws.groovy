node{
ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/"){
    withEnv(["GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"]){
        stage('Checkout') {
            checkout scm
        }
        stage('Pre Test'){
            echo "Pulling Dependencies"

            sh 'go version'
            sh 'go get -u github.com/golang/dep/cmd/dep'
            sh 'go get -u github.com/golang/lint/golint'
            sh 'go get github.com/tebeka/go2xunit'
            sh 'cd $GOPATH/src/cmd/project && dep ensure'
        }

        stage ('Test'){
            sh 'cd $GOPATH && go list ./... | grep -v /vendor/ | grep -v github.com | grep -v golang.org > projectPaths'

            def paths = sh returnStdout: true, script: """awk '\$0=" ./src/"\$0' projectPaths"""

            echo 'Vetting'

            sh """cd $GOPATH && go tool vet ${paths}"""

            echo 'Linting'

            sh """cd $GOPATH && golint ${paths}"""

            echo 'Testing'
        }

        stage ('Build'){
            echo 'Building Executable'

            sh """cd $GOPATH/src/cmd/project/ && go build -ldflags '-s'"""
        }

        // stage('Publish'){
        //     sh 'git rev-parse HEAD > commit'
        //     def commit = readFile('commit').trim()

        //     sh 'git name-rev --name-only HEAD > GIT_BRANCH'
        //     def branch = readFile('GIT_BRANCH').trim()

        //     branch = branch.substring(branch.lastIndexOf('/') + 1)

        //     def archive = "$GOPATH/project-${branch}-${commit}.tar.gz"

        //     echo "Building Archive ${archive}"

        //     sh """tar -cvzf ${archive} $GOPATH/src/cmd/project/project"""

        //     echo "Uploading ${archive} to DockerHub"
        //         withCredentials([string(credentialsId: 'bb-upload-key', variable: 'KEY')]) { 
        //             sh """curl -s -u 'user:${KEY}' -X POST 'Downloads Page URL' --form files=@'${archive}' --fail"""
        //         }
        // }
    }
}
}