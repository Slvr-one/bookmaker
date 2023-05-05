#!/usr/bin/env bash
#plugins are at /usr/share/jenkins/ref/plugins.txt
. .env
#wget http://localhost:8080/jnlpJars/jenkins-cli.jar
java -jar jenkins-cli.jar -s http://$JENKINS_ADMIN_ID:$JENKINS_ADMIN_PASSWORD@localhost:8080/ install-plugin \
ace-editor \
cloudbees-credentials \
greenballs \
ant \
antisamy-markup-formatter \
build-timeout \
cloudbees-folder \
configuration-as-code \
credentials-binding \
credentials \
credentials-binding \
display-url-api \
docker-commons \
docker-workflow \
email-ext:2.92 \
folders \
git-changelog:1.45 \
git-server:latest \
git:4.13.0 \
git-client:3.13.0 \
gitlab:1.5.36 \
gitlab-merge-request-jenkins:latest \
gitlab-oauth:latest \
gitlab-plugin:latest \
github-branch-source:latest \
gradle:latest \
junit:1156.vcf492e95a_a_b_0 \
ldap:2.12 \
mailer:438.v02c7f0a_12fa_4 \
matrix-auth:latest \
matrix-authorization-strategy:3.1.5 \
maven-plugin:latest \
pam-auth:latest \
pipeline:590.v6a_d052e5a_a_b_5 \
pipeline-basic-steps:994.vd57e3ca_46d24 \
pipeline-build-step:latest \
pipeline-graph-analysis:latest \
pipeline-input-step:latest \
pipeline-milestone-step:latest \
pipeline-model-api:1.0.2 \
pipeline-model-declarative-agent:1.0.2 \
pipeline-model-definition:1.0.2 \
pipeline-githubnotify-step:latest \
pipeline-github-lib:latest \
pipeline-rest-api:latest \
pipeline-stage-step:latest \
pipeline-stage-tags-metadata:latest \
pipeline-stage-view:latest \
pipeline-utility-steps:latest \
plain-credentials:latest \
ssh-agent:latest \
ssh-credentials:latest \
ssh-slaves:latest \
timestamper:latest \
workflow-aggregator:latest \
ws-cleanup:latest

java -jar jenkins-cli.jar -s http://$JENKINS_ADMIN_ID:$JENKINS_ADMIN_PASSWORD@localhost:8080/ restart