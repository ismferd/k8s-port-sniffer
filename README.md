# k8s-port-scanner

Tool to scan ports opened at kubernetes nodes.

## Requirements

* Docker compose
* Go 1.15
* Terraform v1.0.5

## How it works

k8s-port-scanner runs over any k8s cluster as daemonset looking for opened ports and saving the results on a s3 bucket.
If you take a look into your bucket you should see as object as nodes you have, we are storing one file per node.

k8s-port-scanner is listening for connections at port 8080. It has a subprocess to show you in a browser the relation node -> ports.
* In order achieve it you should do a `kubernetes port-forward pod 8080:8008`.

## Deployment

[Makefile](Makefile) is the deployment manager, it receives diferent instruccions and params in order to release code, create s3 infra and create k8s infra.

You will find more information reading the Makefile or in the [aws](doc/aws.md) and [kubernetes](doc/kubernetes.md) section.

## Infrastructure

Under the folder [infrastructure](infrastructure) you will find 2 folders more, once related with kubernetes and other with AWS.

### AWS

Basically, you will find terraform files which build a s3 bucket where k8s-port-scanner store the information.

Documentation about [aws](doc/aws.md)

### Kubernetes

Here, you will find all manifest to deploy k8s-port-scanner over Kubernetes.

* Note: Modify [secret.yaml](infrastructure/kubernetes/secret.yaml) before launch it and don't commit your secrets!.

Documentation about [kubernetes](doc/kubernetes.md)

## Run locally

As you kwow, k8s-port-scanner runs over a k8s cluster but you can also run it in your local machine doing next steps:

```
cd src; go build -o k8s-port-scanner
AWS_ACCESS_KEY_ID=YOUR_KEY AWS_SECRET_ACCESS_KEY=YOUR_SECRET ITERATION_TIME=60 AWS_REGION=region AWS_ENDPOINT=https://s3.us-east-2.amazonaws.com BUCKET_NAME=bucket_name ./k8s-port-scanner
```

or running a container:
```
docker run -e AWS_ACCESS_KEY_ID=YOUR_KEY -e AWS_SECRET_ACCESS_KEY=YOUR_SECRET -e ITERATION_TIME=60 -e AWS_REGION=region -e AWS_ENDPOINT=https://s3.us-east-2.amazonaws.com -e BUCKET_NAME=bucket_name ismaelfm/node-port-scanner:latest
```

