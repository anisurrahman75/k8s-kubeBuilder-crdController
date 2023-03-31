# k8s-kubeBuilder
## Scaffold new project using following Command:
- ``kubebuilder init --domain DOMAIN_NAME --repo MODULE_NAME``
- ``kubebuilder create api --group GROUP_NAME --version VERSION_NAME --kind KIND_NAME``

## Concrete example: [ in my case ]
- initial cmd: `make`
- make manifests: `make manifests`
- install <Mycrd: AppsCode>: `make install`
- run this project : `make run`

## Generate Some Appscode Object as example: [ run this another terminal ]
- `kubectl apply -f config/samples/a.yaml`
- `kubectl apply -f config/samples/b.yaml`
- `kubectl apply -f config/samples/c.yaml`

## Display all objects: 
- `kubectl get svc`
- `kubectl get deploy`
- `kubectl get pods`