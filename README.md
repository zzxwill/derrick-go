# derrick-go
🐳A tool to help you containerize application in seconds



# Quick demos

## A golang application

- Clone a sample project

Clone this sample project into your Golang path.

```shell
$ git clone git@github.com:zzxwill/golang-web-application.git
$ cd golang-web-application
```

- Compile the application

```shell
$ derrick-go init
? Please input image name with tag (such as "registry.com/user/repo:tag"):  zzxwill/golang-web-application:latest
Successfully detected your platform is Golang and compiled it successfully.
```

- Push the image and deploy it to Kubernetes

```shell
$ derrick-go up -k
#1 [internal] load .dockerignore
#1 sha256:daa4b49e67a2b1678515c23e671c4892e448407d9879e991a96e123d9e26bc08
#1 transferring context: 34B done
#1 DONE 0.0s
...
The application image zzxwill/golang-web-application:latest has been successfully built.
The push refers to repository [docker.io/zzxwill/golang-web-application]
eb5e68ae951b: Preparing
5a91cd45462f: Preparing
c04d1437198b: Preparing
5a91cd45462f: Layer already exists
c04d1437198b: Layer already exists
eb5e68ae951b: Layer already exists
latest: digest: sha256:dbf02a8fccfaab2bdf901e18d3244ef3121108c8bea6dfaaa6429bf3693bd93b size: 946
service/golang-web-application unchanged
deployment.apps/golang-web-application unchanged
Your application has been built and deployed to your Kubernetes cluster! You can run `kubectl get svc` to get exposed ports.
```

`derrick-go up` will just build and push the image.

- Vist the application

```shell
$ kubectl port-forward service/golang-web-application 8080:8080
Forwarding from 127.0.0.1:8080 -> 8080
Forwarding from [::1]:8080 -> 8080
Handling connection for 8080
Handling connection for 8080
```

```shell
$ curl 127.0.0.1:8080/Derrick
Hi there, I love Derrick!
$ curl 127.0.0.1:8080/Golang
Hi there, I love Golang!
```

