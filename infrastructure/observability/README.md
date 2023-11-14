

### Install Cert-manager
helm upgrade --install cert-manager --set installCRDs=true --namespace cert-manager --version 1.13.2 jetstack/cert-manager

### Monitoring

helm repo add grafana https://grafana.github.io/helm-charts
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update prometheus-community

helm repo add vm https://victoriametrics.github.io/helm-charts/
helm repo update vm

k create ns mon
helm upgrade --install vmstack vm/victoria-metrics-k8s-stack -f values.yaml --namespace mon 
