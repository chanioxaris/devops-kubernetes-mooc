## [Part 1](https://devopswithkubernetes.com/part-1)

### 1.01

Create a new deployment using the course main application image.

`$ kubectl create deployment hashgenerator-dep --image=chanioxaris/devops-k8s-mooc-main-app`

![screenshot](1.01/img/1.01_create_deployment.png)

List available pods.

`$ kubectl get pods`

![screenshot](1.01/img/1.01_list_pods.png)

View logs of the newly created pod.

`$ kubectl logs -f hashgenerator-dep-bb955cc76-zj6v7`

![screenshot](1.01/img/1.01_log_pod.png)