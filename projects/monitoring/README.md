

###1. Work on docker 
on_docker containes basics of prometheus

We using [bitnami's prometheus image](https://hub.docker.com/r/bitnami/prometheus)

```sh
docker run --rm \
	-v ${PWD}/data:/opt/bitnami/prometheus/data \
	-v ${PWD}/prometheus.yml:/opt/bitnami/prometheus/conf/prometheus.yml \
	-p 9090:9090 \
    bitnami/prometheus:latest 
```

