# Notes on Ingress and service

## Pod port
  pod-ip:8080 (containerPort, the same with the exposed port in Docker file)

## Service port
service:
  cluster-ip:80
  node-ip:nodePort

kube-system/default-backend:
  node-ip:nodePort

## ingress port

loadbalancer-address:80

```
Ingress rules:
  \* -> defaultbackend: nodePort
  \ -> service:80 (type muste be nodePort)
```

ingress controller (defaultbackend) will create a GCP load balancer
with an external ip
