#!/bin/bash

# Author      : Balaji Pothula <balan.pothula@gmail.com>,
# Date        : Wednesday, 02 March 2022,
# Description : Installing Jenkins Server.

# exit immediately if a command exits with a non-zero exit status.
# set -e

# debugging shell script.
# set -x

# upgrade installed packages to latest.
yum -y -q upgrade

# install Network File System utilities.
# install Docker.
yum -y -q install docker nfs-utils

# enable Docker service at boot time.
systemctl enable docker.service

# start the Docker service.
systemctl start docker.service

# create root directory for Tomcat, Jenkins and Maven.
mkdir -p /home/ec2-user/root/.jenkins /home/ec2-user/root/.m2

# mount EFS to root directory.
mount -t nfs -o nfsvers=4.1,rsize=1048576,wsize=1048576,hard,timeo=600,retrans=2,noresvport fs-0858aeb075dc7a734.efs.eu-central-1.amazonaws.com:/ /home/ec2-user/root

# change permissions of root directory.
chmod 777 /home/ec2-user/root

# pull Jenkins supported Docker image.
docker pull balajipothula/jenkins:openjdk11

# run jenkins container.
docker run --name jenkins -d -i --privileged -p 8080:8080 -v $HOME/root/.jenkins:/root/.jenkins -v $HOME/root/.m2:/root/.m2 balajipothula/jenkins:openjdk11 sh

# execute Jenkins in container.
docker exec -i jenkins java -jar /root/.jenkins/jenkins.war --httpPort=8080 &
