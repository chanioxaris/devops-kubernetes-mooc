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

### 1.03

The manifest directory is located [here](../src/main-application/manifests).

Create a new deployment by applying the manifest.

`$ kubectl apply -f manifests/deployment.yaml`

![screenshot](1.03/img/1.03_apply_manifest.png)

View logs and verify that running successfully.

`$ kubectl logs hashgenerator-dep-688c764588-2j626`

![screenshot](1.03/img/1.03_log_pod.png)

### 1.04

The manifest directory is located [here](../src/project/manifests).

Create a new deployment by applying the manifest.

`$ kubectl apply -f manifests/deployment.yaml`

![screenshot](1.04/img/1.04_apply_manifest.png)

View logs and verify that the server is running.

`$ kubectl logs project-dep-67f848d568-4dgkx`

![screenshot](1.04/img/1.04_log_pod.png)

### 1.05

Forward the server port 8080 into 8080 local port.

`$ kubectl port-forward project-dep-67f848d568-qs47c 8080:8080`

![screenshot](1.05/img/1.05_port_forward.png)

Check if port forwarding is working.

`$ curl http://localhost:8080`

![screenshot](1.05/img/1.05_curl.png)

### 1.06

The manifest service file is located [here](../src/project/manifests/service.yaml).

Apply the new manifest service.

`$ kubectl apply -f manifests/service.yaml`

![screenshot](1.06/img/1.06_apply_service.png)

Check if we can access the server from outside.

`$ curl http://localhost:8082`

![screenshot](1.06/img/1.06_curl.png)

### 1.07

The service file is located [here](../src/main-application/manifests/service.yaml).

Apply the new service.

`$ kubectl apply -f manifests/service.yaml`

![screenshot](1.07/img/1.07_apply_service.png)

The ingress file is located [here](../src/main-application/manifests/ingress.yaml).

Apply the new ingress.

`$ kubectl apply -f manifests/ingress.yaml`

![screenshot](1.07/img/1.07_apply_ingress.png)

Check if we can access the server from outside.

`$ curl http://localhost:8081`

![screenshot](1.07/img/1.07_curl.png)
