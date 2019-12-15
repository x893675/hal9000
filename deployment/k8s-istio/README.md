```shell
kubectl apply -f rbac.yaml

kubectl create namespace test

kubectl label namespaces test istio-injection=enabled

kubectl apply -f rolebinding -n test

kubectl apply -f greeter-*.yaml -n test

```
