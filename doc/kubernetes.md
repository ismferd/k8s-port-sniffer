# Kubernetes

k8s-port-scanner will run as a daemonset in Kubernetes.

## Manifests

You can find the terraform code [here](../infrastructure/kubernetes).

### Namespace

We are going to isolete our application in a namespace called sre

### Service

We are going to build a service ClusterIP, in a future we would like to expose metrics to prometheus.

### Daemonset
* Readiness and liveness: Both are waiting for to be opene 8080 port.

### Config Map value

* ITERATION_TIME: The time loop in seconds that k8s-port-scanner will be executed. It must be informed and the value must be upper than 60
* AWS_REGION: Region where your bucket has been deployed. It must be informed or we won't get the ports information.
* AWS_ENDPOINT: Custom endpoint, useful to localstack test. Example s3 AWS endpoint: "https://s3.us-east-2.amazonaws.com"
* BUCKET_NAME: The bucket name that you has created through Terraform.
* PORTS: A whitelisted list of ports what we want to avoid to collect. List have to be informed as string,string. Example PORTS: 9090,22,8080

### Deployment Manifests

using `make deploy_kubernetes` you must inform before of:  
* export AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID} 
* export AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY} 
* export AWS_DEFAULT_REGION=${AWS_DEFAULT_REGION} 
* export KUBECONFIG=${KUBECONFIG} 

or

`AWS_ACCESS_KEY_ID=YOUR_AWS_ACCESS AWS_SECRET_ACCESS_KEY=YOUR_AWS_SECRET AWS_DEFAULT_REGION=CLUSTER_AWS_REGION KUBECONFIG=KUBECONFIG_PATH make deploy_kubernetes`

### Secrets

* Remember to change values of (secret.yaml)[../infrastructure/kubernetes/secret.yaml]

## Watching ports opened

- Currently, we don't have ingress, so you can do a port-forward in order to see the data stored on s3.
