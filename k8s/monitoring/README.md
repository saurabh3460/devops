## Prometheus

1. Setup and run Prometheus locally.
2. configure it to scrape itself.
3. configure it to scrape an example application.
4. work with queries, rules, and graphs to use collected time series data.

### 1. Setup and run Prometheus locally.

```sh
docker run --name prometheus \
    -v ${PWD}/prometheus.yml:/opt/bitnami/prometheus/conf/prometheus.yml \
    -p 9090:9090 \
    bitnami/prometheus:latest
```

### 2. configure it to scrape itself.
- Prometheus collects metrics from targets by scraping metrics HTTP endpoints.


[getting_started](https://prometheus.io/docs/prometheus/latest/getting_started/)

[configuration](https://prometheus.io/docs/prometheus/latest/configuration/configuration/)

https://www.metricfire.com/blog/understanding-the-prometheus-rate-function/

https://chronosphere.io/learn/an-introduction-to-the-four-primary-types-of-prometheus-metrics/

/home/saurabh/Desktop/devops/k8s/monitoring

