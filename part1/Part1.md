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

The manifest directory is located [here](1.03/manifests).

Create a new deployment by applying the manifest.

`$ kubectl apply -f manifests/deployment.yaml`

![screenshot](1.03/img/1.03_apply_manifest.png)

View logs and verify that running successfully.

`$ kubectl logs hashgenerator-dep-688c764588-2j626`

![screenshot](1.03/img/1.03_log_pod.png)

### 1.04

The manifest directory is located [here](1.04/manifests).

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

The manifest service file is located [here](1.06/manifests/service.yaml).

Apply the new manifest service.

`$ kubectl apply -f manifests/service.yaml`

![screenshot](1.06/img/1.06_apply_service.png)

Check if we can access the server from outside.

`$ curl http://localhost:8082`

![screenshot](1.06/img/1.06_curl.png)

### 1.07

The service file is located [here](1.07/manifests/service.yaml).

Apply the new service.

`$ kubectl apply -f manifests/service.yaml`

![screenshot](1.07/img/1.07_apply_service.png)

The ingress file is located [here](1.07/manifests/ingress.yaml).

Apply the new ingress.

`$ kubectl apply -f manifests/ingress.yaml`

![screenshot](1.07/img/1.07_apply_ingress.png)

Check if we can access the server from outside.

`$ curl http://localhost:8081`

![screenshot](1.07/img/1.07_curl.png)

### 1.08

The service file is located [here](1.08/manifests/service.yaml).

Apply the new service.

`$ kubectl apply -f manifests/service.yaml`

![screenshot](1.08/img/1.08_apply_service.png)

The ingress file is located [here](1.08/manifests/ingress.yaml).

Apply the new ingress.

`$ kubectl apply -f manifests/ingress.yaml`

![screenshot](1.08/img/1.08_apply_ingress.png)

Check if we can access the server from outside.

`$ curl http://localhost:8081`

![screenshot](1.08/img/1.08_curl.png)

### 1.09

The deployment file is located [here](1.09/manifests/deployment.yaml).

Create a new deployment.

`$ kubectl apply -f manifests/deployment.yaml`

![screenshot](1.09/img/1.09_apply_deployment.png)

The service file is located [here](1.09/manifests/service.yaml).

Apply the new service.

`$ kubectl apply -f manifests/service.yaml`

![screenshot](1.09/img/1.09_apply_service.png)

The ingress file is located [here](1.09/manifests/ingress.yaml).

Apply the new ingress.

`$ kubectl apply -f manifests/ingress.yaml`

![screenshot](1.09/img/1.09_apply_ingress.png)

Check if server is working as expected.

`$ curl http://localhost:8081/pingpong`

![screenshot](1.09/img/1.09_curl.png)

### 1.10

The new deployment file is located [here](1.10/manifests/deployment.yaml).

Check if server is working as expected.

`$ curl http://localhost:8081`

![screenshot](1.10/img/1.10_curl.png)

### 1.11

The new persistent volume file is located [here](1.11/manifests/persistentvolume.yaml).

The new persistent volume claim file is located [here](1.11/manifests/persistentvolumeclaim.yaml).

The new deployment file for pingpong app is located [here](1.11/manifests/deployment-ping-pong.yaml).

The new deployment file for main app is located [here](1.11/manifests/deployment-main.yaml).

Apply the persistent volume manifest.

`$ kubectl apply -f manifests/persistentvolume.yaml`

![screenshot](1.11/img/1.11_apply_persistentvolume.png)

Apply the persistent volume claim manifest.

`$ kubectl apply -f manifests/persistentvolumeclaim.yaml`

![screenshot](1.11/img/1.11_apply_persistentvolumeclaim.png)

Check if server is working as expected.

`$ curl http://localhost:8081/pingpong`
`$ curl http://localhost:8081`

![screenshot](1.11/img/1.11_curl.png)

### 1.12

The new persistent volume file is located [here](1.12/manifests/persistentvolume.yaml).

The new persistent volume claim file is located [here](1.12/manifests/persistentvolumeclaim.yaml).

The new deployment file is located [here](1.12/manifests/deployment.yaml).

The new service file is located [here](1.12/manifests/service.yaml).

Apply the persistent volume manifest.

`$ kubectl apply -f manifests/persistentvolume.yaml`

![screenshot](1.12/img/1.12_apply_persistentvolume.png)

Apply the persistent volume claim manifest.

`$ kubectl apply -f manifests/persistentvolumeclaim.yaml`

![screenshot](1.12/img/1.12_apply_persistentvolumeclaim.png)

Check if web browser renders the image.

![screenshot](1.12/img/1.12_webpage.png)

### 1.13

Verify that form and list of todos renders correctly

![screenshot](1.13/img/1.13_webpage.png)