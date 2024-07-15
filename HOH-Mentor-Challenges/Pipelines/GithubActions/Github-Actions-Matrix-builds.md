# Github Actions Matrixed Builds 

For this challenge you have been provided the following in the `src` folder:

- Multiple docker images 
- Multiple required files for those images

For this challenge you will implement many components to the build process. 

Here are the requirements:

- The build action must create a matrixed job that builds all docker files at the same time
- The pipeline must use the action [crazy-max/ghaction-container-scan@v3](https://github.com/marketplace/actions/container-scan) to do a tarball scan of your container images.
- The pipeline must be able to bump tag versions of images from a user defined tag schema. 
 
 
ex: docker-tag.json 

```json

name: my-docker-image
tag: v1.0.1

```



All of the docker files have been tested. There is no need to debug.....hopefully. Also these files may shed some light on your Docker Challenge.
