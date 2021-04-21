




### Set values in new config file (dev-config.yaml):
``
kubectl config --kubeconfig=dev-config set-cluster development --server=https://212.2.242.251:6443 --embed-certs --certificate-authority=./keys/ca.crt
``
``
kubectl config --kubeconfig=dev-config set-credentials developer --embed-certs --client-certificate=./keys/john.crt --embed-certs --client-key=./keys/john.key
``

``
kubectl config --kubeconfig=dev-config set-context dev --cluster=development --namespace=finance --user=developer
``

Next step is to create ROLE and ROLE-BINDINGS