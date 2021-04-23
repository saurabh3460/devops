### setup local cluster 

``k3d registry create registry.local --port 5000`` <br>
``k3d cluster create test --config static/local/k3d-test.yml`` <br>
``kubectl get nodes`` 

### on windows
Open ``c:\windows\system32\drivers\etc\hosts``

### on linux
Use command: ``sudo vim /etc/hosts``

Now past the below lines <br>

``127.0.0.1  registry.local ## Add this line ``  <br>
``127.0.0.1  k3d-registry.local ## Add this line ``

### start cluster
``tilt up``



# You can now use the registry like this (example):
# 1. create a new cluster that uses this registry
k3d cluster create --registry-use k3d-registry.local:5000

# 2. tag an existing local image to be pushed to the registry
docker tag nginx:latest k3d-registry.local:5000/mynginx:v0.1

# 3. push that image to the registry
docker push k3d-registry.local:5000/mynginx:v0.1

# 4. run a pod that uses this image
kubectl run mynginx --image k3d-registry.local:5000/mynginx:v0.1