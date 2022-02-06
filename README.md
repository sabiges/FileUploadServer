
<!-- File Upload server and client -->
## File Upload server and client
-----------------------------

## Author : Abeesh KS
-----------------------------

## Description :
-------------

Writing a simple file store service (HTTP server and a command line client)
that stores plain-text files. Server would receive requests from clients to 
store, update, delete files, and perform operations on files stored in the server.

## Note: Created this package with no 3rd party dependency

### Supporting operations:


```
1. Add files to the store.

E.g: store add file1.txt file2.txt ; should send both files - file1.txt
and file2.txt in the current path to the file store. Add command should fail if the
file already exists in the server.

2. List files in the store
e.g. store ls ; should list the files in the store

3. Remove a file:
E.g:
store rm file.txt ; should remove file.txt from store

4. Update contents of a file in the store:
E.g: store update file.txt ; should update contents of file.txt in
server with the local file.txt or create a new file.txt in server if it is
absent.

5. Support the following operations on files
a. word count: store wc ; returns the number of words in all the files stored
in server
```

## Testing with GO file later made to store go binary : /* go build fileuploadclient.go -o store */  

Run Server in a separate terminal/docker container/kubernetes pod. 

## In Linux Machine , new terminal : go run fileuploadserver.go

Run Client in seperate teminal

## Different commands used.
```
 go run fileuploadclient.go add file2.txt
 
 go run fileuploadclient.go add file2.txt 
 
 go run fileuploadclient.go update file2.txt 
 
 go run fileuploadclient.go wc
 
 go run fileuploadclient.go update file2.txt file2.txt file3.txt file4.txt
 
 go run fileuploadclient.go wc
 
 go run workingfileuploadclient.go rm file2.txt
 
 go run workingfileuploadclient.go ls
```

## Docker file building 
```
Step1 : git clone https://github.com/sabiges/FileUploadServer.git 

Step2: cd FileUploadServer

Step 3: [FileUploadServer]$ docker build -t fileuploadser:v1 -f dockerfiles/Dockerfile   . 
Sending build context to Docker daemon    106kB
Step 1/11 : FROM golang:1.16-alpine
 ---> 4bcb0d501de3
Step 2/11 : ENV http_proxy ""
 ---> Using cache
 ---> 3bac09133c33
Step 3/11 : ENV https_proxy ""
 ---> Using cache
 ---> 993a1f707305
Step 4/11 : WORKDIR /app
 ---> Using cache
 ---> f4dc5d61d6fd
Step 5/11 : COPY src/server/go.mod ./
 ---> Using cache
 ---> 57322e5f098f
Step 6/11 : COPY src/server/go.sum ./
 ---> Using cache
 ---> 2c5cdf05454b
Step 7/11 : RUN go mod download
 ---> Using cache
 ---> b106d8dc325e
Step 8/11 : COPY src/server/*.go ./
 ---> Using cache
 ---> 7a51e68a73bb
Step 9/11 : RUN go build -o /server-fileupload
 ---> Using cache
 ---> f330deaf2fe5
Step 10/11 : EXPOSE 4000
 ---> Using cache
 ---> 0b1970426f96
Step 11/11 : CMD [ "/server-fileupload" ]
 ---> Using cache
 ---> 22ccb24c0a94
Successfully built 22ccb24c0a94
Successfully tagged fileuploadser:v1
```

## Server image execution through different approach
```
1. Invoking go binary

a. cd FileUploadServer/src/server

b. make install

c. ls -lrt  FileUploadServer/bin/
   o/p : store_server_binary

2. Invoking Docker image


a. Option 1 : Build the docker image mentioned above ; Option 2 , Use uploaded image in FileUploadServer/image folder

b. [FileUploadServer]$ docker run -it fileuploadser:v1

 o/p Starting server in 4000 port

# Client binary creation

a.  cd FileUploadServer/src/client

b. make install

c. ls -lrt  FileUploadServer/bin/
   o/p : store
```   
   
## go test for client and server
```
cd FileUploadServer/src/client

b. make test

c. ls -lrt  FileUploadServer/src/client/bin/
   o/p : store_test
   
Log:

[abeeshks@localhost client]$ go test
PASS
ok  	test/FileUploadServer/src/client	0.002s



--------------------------------   
cd FileUploadServer/src/server

b. make test

c. ls -lrt  FileUploadServer/src/server/bin/
   o/p : 
   
Log:
Running tool: /usr/local/go/bin/go test -timeout 30s -run ^TestUploadServer$ test/FileUploadServer/server

ok  	test/FileUploadServer/server	0.003s
```


## kubernetes env preparation:

Prerequiste : kvm2 driver is required.

*[minikube instalation guide](https://fedoramagazine.org/minikube-kubernetes/)

*[kvm installation guide](https://computingforgeeks.com/how-to-install-kvm-on-fedora/)


Run below commands:

### To Create Kubernetes cluster with minikube, execute below commands:
---------------------------------
```
--
1. echo setting no_proxy
2. export no_proxy=$no_proxy,192.168.39.140,192.168.39.248,localhost,127.0.0.1,10.43.192.2,10.43.192.3,10.43.192.4,10.43.192.4,10.43.192.6,dockerregistry.ims.nokia.com,registry.access.redhat.com,myregistry.local,quay.io,10.96.0.0/12,192.168.0.0/12,192.168.39.0/24,172.17.0.0/12,k8s.gcr.io,gcr.io
3. unset https_proxy
4. unset http_proxy
5. echo starting minikube
6. minikube start --vm-driver=kvm2 --docker-env http_proxy=$http_proxy --docker-env https_proxy=$https_proxy --docker-env no_proxy=$no_proxy 
7. eval $(minikube docker-env)

------------------------
For cleaning the minikube cache files and reinstall freshely, you can try below command

sudo  minikube stop; sudo minikube delete;sudo  iptables -F && sudo iptables -t nat -F && sudo iptables -t mangle -F && sudo iptables -X ;sudo  cd /root/minikube ;sudo  rm -rf .minikube
```

Logs:

sh minkube_start.sh --> Same commands as mentioned above
-------------------
```
setting no_proxy
starting minikube
😄  minikube v1.5.2 on Fedora 31
🔥  Creating kvm2 VM (CPUs=2, Memory=2000MB, Disk=20000MB) ...
🌐  Found network options:
    ▪ no_proxy=localhost,127.0.0.1,10.96.0.0/12,192.168.99.0/24,192.168.39.0/24,172.17.0.0/12,192.168.39.140,192.168.39.248,localhost,127.0.0.1,10.43.192.2,10.43.192.3,10.43.192.4,10.43.192.4,10.43.192.6,dockerregistry.ims.nokia.com,registry.access.redhat.com,myregistry.local,quay.io,10.96.0.0/12,192.168.0.0/12,192.168.39.0/24,172.17.0.0/12,k8s.gcr.io,gcr.io
🐳  Preparing Kubernetes v1.16.2 on Docker '18.09.9' ...
    ▪ env http_proxy=
    ▪ env https_proxy=
    ▪ env no_proxy=localhost,127.0.0.1,10.96.0.0/12,192.168.99.0/24,192.168.39.0/24,172.17.0.0/12,192.168.39.140,192.168.39.248,localhost,127.0.0.1,10.43.192.2,10.43.192.3,10.43.192.4,10.43.192.4,10.43.192.6,dockerregistry.ims.nokia.com,registry.access.redhat.com,myregistry.local,quay.io,10.96.0.0/12,192.168.0.0/12,192.168.39.0/24,172.17.0.0/12,k8s.gcr.io,gcr.io
🚜  Pulling images ...
🚀  Launching Kubernetes ... 
⌛  Waiting for: apiserver
🏄  Done! kubectl is now configured to use "minikube"
```
-----------------------------------------

## kubernetes deployment and file creation
```
Step 1: FileUploadServer/kubernetes

Step 2: Create deployment and service

kubectl create -f deployment_fileupload.yaml

kubectl create -f fileupload_service.yaml


o/p
[abeeshks@localhost kubernetes]$ kubectl get pods
NAME                            READY   STATUS    RESTARTS   AGE
my-fileupload-77d6dc46d-jhl59   1/1     Running   0          78m
[abeeshks@localhost kubernetes]$ kubectl get svc
NAME            TYPE        CLUSTER-IP       EXTERNAL-IP     PORT(S)          AGE
kubernetes      ClusterIP   10.96.0.1        <none>          443/TCP          88m
my-fileupload   NodePort    10.109.131.140   192.168.39.34   4000:30004/TCP   41m
[abeeshks@localhost kubernetes]$ 
```


## How to test the functionality.
```
Use help option

Step 1: 
cd FileUploadServer/bin

$ ./store help
Destination Address : 127.0.0.1
Destination Port : 4000

Wrong input : [./store help]

Tool Excecution commands supported

store -ip <ip> -port <port> add <file1> <file2> ...
store -ip <ip> -port <port> ls
store -ip <ip> -port <port> rm <file1> <file2> ...
store -ip <ip> -port <port> update <file1> <file2> ...
store -ip <ip> -port <port> wc
```
------------------------------------
# Project Directory Structure
---------------------------
```
[abeeshks@localhost FileUploadServer]$ tree
.
├── bin
│   ├── store
│   └── store_server_binary
├── dockerfiles
│   └── Dockerfile
├── image
│   └── fileuploadserverimage.tar >> Due to huge size, this folder is removed
├── kubernetes
│   ├── deployment_fileupload.yaml
│   └── fileupload_service.yaml
├── README.md
├── src
│   ├── client
│   │   ├── bin
│   │   │   └── store_test
│   │   ├── file1.txt
│   │   ├── file2.txt
│   │   ├── file3.txt
│   │   ├── file4.txt
│   │   ├── fileuploadclient.go
│   │   ├── fileuploadclient_test.go
│   │   ├── go.mod
│   │   ├── go.sum
│   │   ├── Makefile
│   │   └── store
│   ├── Makefile.include
│   └── server
│       ├── bin
│       │   └── store_server_test
│       ├── fileuploadserver.go
│       ├── fileuploadserver_test.go
│       ├── go.mod
│       ├── go.sum
│       ├── Makefile
│       └── store_server_binary
└── Tested_Logs
```

------------------------------------
# Tested logs are there in  this link: [Tested_logs](https://github.com/sabiges/FileUploadServer/blob/main/Tested_Logs) -- Please go through it for more details.


Thankyou.



