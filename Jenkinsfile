
def incrementVersion(String version) {
    def parts = version.tokenize('.')
    def major = parts[0].toInteger()
    def minor = parts[1].toInteger()
    def patch = parts[2].toInteger()

    patch++
    if (patch > 99) {
        patch = 0
        minor++
        if (minor > 99) {
            minor = 0
            major++
        }
    }

    return "${major}.${minor}.${patch}"
}

pipeline {
    options {
        timestamps()
        timeout(time:5, unit:'MINUTES')
        gitLabConnection('jenkins_gitlab')
        // buildDiscarder(logRotator(numToKeepStr: '100'))
    }

    agent any
    environment {
        //  // Define environment variables here
        // DOCKER_IMAGE = 'bookmaker'
        // KUBECONFIG = '~/.kube/config'
        // TERRAFORM_DIR = '.'

        VERSION = '1.0.0' // initial version
        MSG = sh (script: "git log -1 --pretty=%B ${GIT_COMMIT}", returnStdout: true).trim()
        // PROVISION = sh(script: "git log -1 | tail -1 | grep '#e2e'", returnStatus: true)
        ECR_URL = "514095112279.dkr.ecr.eu-central-1.amazonaws.com/bookmaker"

        APP_REPO = ""
        INFRA_REPO = "" 
        AWS_REGION = "eu-central-1"
        PROJECT = "bookmaker"
        RET = 5 // wget retries for tests
        BM_PORT = 5004
        BM_PORT_EXTERNAL = 5000
        NGINX_PORT = 80
        BM_CONT_NAME = "app"
        NGINX_CONT_NAME = "static"
    }
    tools { 
        go 'golang'
    }
    triggers {
        // pollSCM "* * * * *"
        gitlab(
        triggerOnPush: true,
        setBuildDescription: true,
        branchFilterType: "All",
        // triggerOnMergeRequest: true, triggerOpenMergeRequestOnPush: "both",
        // triggerOnNoteRequest: true,
        // noteRegex: "Jenkins please ret a build",
        // skipWorkInProgressMergeRequest: false,
        // ciSkip: false,
        // addNoteOnMergeRequest: true,
        // addCiMessage: true,
        // addVoteOnMergeRequest: true,
        // acceptMergeRequestOnSuccess: false,
        // secretToken: "NOTVERYSECRET",
        )
    }
    stages {
        stage("0 - preface") {
            steps {
                deleteDir()
                checkout scm
                // checkout changelog: true, poll: true, scm: [
                //     $class: 'GitSCM',
                //     branches: [[name: "origin/${gitlabSourceBranch}"]],
                //     extensions: [[$class: 'PreBuildMerge', options: [fastForwardMode: 'FF', mergeRemote: 'origin', mergeStrategy: 'DEFAULT', mergeTarget: "${gitlabTargetBranch}"]]],
                //     userRemoteConfigs: [[name: 'origin', url: 'git@gitlab:jenkins/application.git']]
                // ]
            }
        }
        stage("1 - build") {
            // when {
            //     expression { BRANCH_NAME =~ "feature/*" || BRANCH_NAME == "main" }
            // }
            steps {
                dir("app") {
                    echo "---- build / package ----"
                    sh "./dockerize.sh $PROJECT"
                    // script {
                    //     app = docker.build("bookmaker")
                    // }
                }
            }
        }
        stage("2 - run") {
            // when {                  
            //     expression { BRANCH_NAME =~ "feature/*" || BRANCH_NAME == "dev" }
            // }
            steps {
                script { //extra config for ut
                    if (BRANCH_NAME == "prod") {
                        // NGINX_PORT = NGINX_PORT.toInteger() + 1
                        BM_PORT = BM_PORT.toInteger() + 1 }
                    if (BRANCH_NAME == "dev") {
                        // NGINX_PORT = NGINX_PORT.toInteger() + 2
                        BM_PORT = BM_PORT.toInteger() + 2 }
                // for main still 5000 & 80
                sh "yq -i \'.services.bookmaker.ports[0] = \"$BM_PORT:5000\"\' docker-compose.yaml"
                // sh "yq -i \'.services.static.ports[0] = \"$NGINX_PORT:80\"\' docker-compose.yaml"
                // sh "sed -i '/\"proxy_pass\"/s/8080/$BM_PORT/' nginx/nginx.conf"
                sh "docker compose up -d"
                // sh "echo && docker ps -a --format \"table {{.ID}}\t{{.Names}}\t{{.RunningFor}}\t{{.Status}}\" && echo"
                // sh "echo && docker ps -a --format \"table {{.Ports}}\" && echo"
                // sh "docker compose logs $NGINX_CONT_NAME"
                }
                script {
                    // NGINX_CONT_NAME = sh(script: "docker ps -a --format \"{{.Names}}\" | grep \"$BRANCH_NAME-$NGINX_CONT_NAME\"", returnStdout: true).trim()
                    BM_CONT_NAME = sh(script: "docker ps -a --format \"{{.Names}}\" | grep \"$BRANCH_NAME-$BM_CONT_NAME\"", returnStdout: true).trim()
                    // sh "docker cp ./app/src/static $NGINX_CONT_NAME:/usr/share/nginx/html"
                    // sh "docker cp ./app/src/tatic/stylesheets/ $NGINX_CONT_NAME:/usr/share/nginx/html/stylesheets/"
                    // sh "docker cp ./nginx/nginx.conf $NGINX_CONT_NAME:/etc/nginx/nginx.conf"
                    // sh "./nginx/reload.sh $NGINX_CONT_NAME"
                    // sh "docker compose exec $NGINX_CONT_NAME -- nginx -t && nginx -s reload"
                }
            }
        }
        stage("- UT -") {
            steps {
                echo "----- unit testing -----"
                sh "./tests/unit-tests.sh $BM_CONT_NAME $RET $BM_PORT_EXTERNAL"
                // sh "./tests/unit-tests.sh $NGINX_CONT_NAME $RET $NGINX_PORT"
            }
        }
        stage("- status -") {
            steps {
                sh "echo && docker ps -a --format \"table {{.ID}}\t{{.Names}}\t{{.RunningFor}}\t{{.Status}}\" && echo"
                sh "echo && docker ps -a --format \"table {{.Ports}}\" && echo"
            }
        }
        stage("4 - tag") {
            when { branch "release" }
            steps {
                sshagent(['jenkins_gitlab']) {
                    sh "git fetch -t || true"
                }
                script {
                    LATEST = sh(script: "git tag -l | tail -1", returnStdout: true).trim()
                    if (LATEST.isEmpty()) {
                        PATCH = ".0"
                        CURRENT_VERSION = "1.0" + PATCH
                    } else {
                        CURRENT_VERSION = LATEST.split('\\.')
                        CURRENT_VERSION[2] = CURRENT_VERSION[2].toInteger() + 1
                        CURRENT_VERSION = CURRENT_VERSION.join('.') //for future ref
                    }
                    sh """
                        docker tag $PROJECT:latest $PROJECT:$CURRENT_VERSION    
                        git clean -xf && git tag $CURRENT_VERSION
                    """
                    sshagent(['jenkins_gitlab']) {
                        sh "git push --tags"
                    }
                }
            }
        }
        stage("3 - E2E") {// if branch is main, ?release
            when { branch "main" }
                steps {
                    echo "---- E2E tests ----"
                    sh "./tests/E2E-tests.sh $BM_CONT_NAME $RET $BM_PORT_EXTERNAL"
                    // sh "./tests/E2E-tests.sh $NGINX_CONT_NAME $RET $NGINX_PORT"
               
            }
        }
        stage("5 - publish") {
            when { branch "main" }
            steps {
                echo "---- publish to ecr ----"
                script {
                    docker.withRegistry("http://${ECR_URL}", "ecr:${AWS_REGION}:jenkins-access-key") {
                        docker.image("${PROJECT}").push("${CURRENT_VERSION}")  //tag CURRENT_VERSION sem version
                        docker.image("${PROJECT}").push('latest')  // update tag "latest"
                    }
                }              
            }
        }
        stage("6 - deploy") {
            // when { branch "main" }
            steps {
                //argo shold notice a change in app git repo, which he listen for, and apply changes to cd.
                sh "yq -i '.App.imageTag = \"$CURRENT_VERSION\"' ./app/chart/values.yaml"
                // sh "sed -i \'/^\$/d\' ./app/chart/values.yaml" //delete empty lines
            }
        }
    }
    post {
        success {
            echo "its a success!"
        }
        failure {
            echo "its a failure.."
            // emailext (
            //     to:      "${EMAIL_TO}",
            //     subject: "Jenkins - ${JOB_NAME}, build - ${BUILD_DISPLAY_NAME} - ${currentBuild.currentResult}",
            //     body:"""
            //     <p>Jenkins job <a href='${JOB_URL}'>${JOB_NAME}</a> (<a href='${BUILD_URL}'>build ${BUILD_DISPLAY_NAME}</a>) has result <strong>${currentBuild.currentResult}</strong>!
            //     <br>You can view the <a href='${BUILD_URL}console'>console log here</a>.</p>
            //     <br><strong>terraform Workspaces list:</strong></p>
            //     ${TERRAFORM_WORKSPACE}</p>
            //     <br><strong>DeleBM Workspace list:</strong></p>
            //     <br>${DELEBM_WORKSPACE}</p>
            //     <p>Source code from commit: <a href='${GIT_URL}/commit/${GIT_COMMIT}'>${GIT_COMMIT}</a> (of branch <em>${GIT_BRANCH}</em>).</p>
            //     <p><img src='https://www.jenkins.io/images/logos/jenkins/jenkins.png' alt='jenkins logo' width='123' height='170'></p>
            //     """
            // )
        }
        always {
            echo "will get this done"
        }
        cleanup {
            sh "docker compose down || true"
            // sh "terraform apply -auto-approve ${DESTROY_PLAN}" 
            cleanWs() //commenBM for testing

        }
    }
}
