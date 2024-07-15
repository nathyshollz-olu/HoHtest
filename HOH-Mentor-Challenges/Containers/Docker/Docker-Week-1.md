# Docker Container Challenge Week 1

For this challenge you will complete some basic level docker container builds



1. Create a hello world docker container imperatively (from the command line)

2. Build a container image the prints hello world to the console 5 times and sleeps for 2000 seconds
** You must build this container from the ubuntu:latest image
** You must use a bash script to execute the commands

1. Build a webserver using nginx as a container

** For this challenge use the nginx:latest tag to build the container
** Scenerio: The development team currently serves a static web page in production. With the organization's recent shift to kubernetes as a platform they need to transfer this work into a container
** The team's work has moved to the /Containers/Docker/src/week-1/ directory
** Since the team does not currently have a pipeline for their application they are requesting that the static pages are added to the contianer image directly
** Before enabling this in production the team wants to see this container run in a local environment. Provide a Dockerfile and the necessary commands to build and run this container exposing the nginx internal port to port 80 on the docker host. 
** The final product should allow the development team to run your container on a local machine and see their webpage from the ip of their machine at port 80


## Notes for this challenge

All work for week-1 will go into the /Containers/Docker/products/week-1 directory. 

Create a PR with your changes to merge into master





