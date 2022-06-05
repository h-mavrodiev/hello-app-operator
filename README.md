# hello-app-operator
Kubernetes operator that deploys the hello-app - https://github.com/h-mavrodiev/hello-app

The aim of the operator:
- To extend the Kubernetes API with the defined CRD (Custom Resouce Definition)
- Wil monitor and react to any changes on object of the type defined by the CRD.
- Will utilize validating webhook for controlling the creating and any changes on ojbect from the CRD type (WIP).
- Will install the web application (the hello-app) on demand in a kubernetes cluster.
