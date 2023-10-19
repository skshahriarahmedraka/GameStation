# AI Crop Disease Detection

### step 1 : Create AKS Cluster

create a kubernetes cluster inside azure kubernetes  service 

### step 2:

authenticate azure kubernetes cluster from your local pc 

here , resource Group = aifarm1

cluster Name = aifarmcluster1

```
 az aks get-credentials --resource-group aifarm1 --name aifarmcluster1
```

### step 4:

install the Ingress-Nginx Controller

https://kubernetes.github.io/ingress-nginx/deploy/#quick-start

```
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx && \
helm repo update && \
helm install ingress-nginx ingress-nginx/ingress-nginx
```

check the ingress service 

```
 kubectl get services ingress-nginx-controller
```


### deploy all the yaml files

```
kubectl apply -f namespace.yaml

```
### create a namespaces
```
kubectl get all --namespace ai-farming

```
### deploy applications 
```
kubectl apply -y ./applications/
kubectl apply -y ./ingress/

```
### install cert manager
```
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.13.0/cert-manager.yaml
```
### install certificate issurer and certificate
```
kubectl apply -y ./letsencrypt/
```
