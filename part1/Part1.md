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

### 1.02

Create a new deployment using the course project image.

`$ kubectl create deployment project-dep --image=chanioxaris/devops-k8s-mooc-project`

![screenshot](1.02/img/1.02_create_deployment.png)

List available pods.

`$ kubectl get pods`

![screenshot](1.02/img/1.02_list_pods.png)

View logs and verify that server has been started.

`$ kubectl logs -f project-dep-8466769d84-p9kjk`

![screenshot](1.02/img/1.02_log_pod.png)