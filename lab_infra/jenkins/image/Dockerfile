FROM jenkins/jenkins:2.382-jdk11

ENV JENKINS_HOME /var/jenkins_home

USER root
RUN apt-get update -qqy
RUN apt-get install -qqy \
    apt-transport-https \
    ca-certificates \
    curl gnupg2 \
    software-properties-common \
    bzip2 sudo wget \
    xclip vim tree lsb-release
    
RUN wget https://github.com/mikefarah/yq/releases/latest/download/yq_linux_amd64 \
  -O /usr/bin/yq && chmod +x /usr/bin/yq
    
#install awscli:
RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
RUN unzip awscliv2.zip
RUN rm -rf awscliv2.zip
RUN ./aws/install

# #installation terraform
# RUN curl -fsSL https://apt.releases.hashicorp.com/gpg | apt-key add - && apt-add-repository "deb [arch=amd64] https://apt.releases.hashicorp.com $(lsb_release -cs) main" && apt-get update && apt-get install terraform

# #install ansible
# RUN apt-get install python3-pip -y && \
#   pip3 install ansible 
# #install boto3 for inventory dynamic
# RUN pip3 install boto3
# #Install LVM for lvol and lvg volumes
# RUN ansible-galaxy collection install community.general
# # Load plugins w jenkins-plugin-cli
# COPY --chown=jenkins:jenkins jenkins/conf/plugins.txt /usr/share/jenkins/ref/plugins.txt
# RUN jenkins-plugin-cli -f /usr/share/jenkins/ref/plugins.txt


#install docker:
RUN mkdir -p /etc/apt/keyrings
RUN curl -fsSL https://download.docker.com/linux/$(. /etc/os-release; echo "$ID")/gpg > /tmp/dkey; apt-key add /tmp/dkey
RUN add-apt-repository \
"deb [arch=amd64] https://download.docker.com/linux/$(. /etc/os-release; echo "$ID") \
$(lsb_release -cs) \
stable"
RUN apt-get update  -qqy \
  && apt-get install -qqy docker-ce docker-ce-cli containerd.io docker-compose-plugin

# RUN useradd -d "$JENKINS_HOME" -u 1000 -ms /bin/bash jenkins || true

RUN usermod -aG docker jenkins 
RUN newgrp docker

USER jenkins 

RUN docker --version
RUN aws --version
RUN docker compose version
RUN echo alias ll="ls -alF" > ~/.bashrc

#some modifications:
#ENV JAVA_OPTS -Djenkins.install.runSetupWizard=false
ENV CASC_JENKINS_CONFIG /var/jenkins_home/casc.yaml
ENV JENKINS_ADMIN_ID dviross
ENV JENKINS_ADMIN_PASSWORD Aa123456

COPY casc.yaml /var/jenkins_home/casc.yaml
# COPY keygen.sh /var/jenkins_home/dvir-scripts/

COPY plugins.txt /usr/share/jenkins/plugins.txt
#RUN /usr/local/bin/plugins.sh /usr/share/jenkins/plugins.txt

#check "./plugins.txt"
# RUN jenkins-plugin-cli --plugins "blueocean:1.25.8 docker-workflow:521.v1a_a_dd2073b_2e"
EXPOSE 8080 50000
