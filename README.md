# redirect-to-https
Docker Golang container to redirect all traffic to https

Motivation for this service is to run inside a Kubernetes cluster and will redirect ALL http ingress traffic to https.

Client -> :80  (ELB) -> :http nodeport  (redirect-to-https)
Client -> :443 (ELB) -> :https nodeport (ingress controller)

Container will listen on port 8080 and will redirect every request to https://HOST[/PATH][?QUERYSTRING]
