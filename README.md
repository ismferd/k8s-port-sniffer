# k8s-port-sniffer

## TODO

-  options:      
    # this toleration is to have the daemonset runnable on master nodes remove it if your masters can't run pods: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/#writing-a-daemonset-spec
    # bucket name
    # AWS credentials
    # metrics

-  to develop in the interviews:
    # Why not to have prometheus operator
    # Service to monitor through Prometheus
    # te decía que un buen motivo para tenerlo todo junto, es que si van por separado, y cambia el producer, la gente que tenga el cli para consumir, se les rompe en plan tendrías que tener ojo con versiones y tal si va todo junto te olvidas de eso

    source <(kubectl completion bash | sed 's/kubectl/k/g')


    AWS_ENDPOINT https://s3.eu-central-1.amazonaws.com