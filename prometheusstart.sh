docker rm -f prometheus
docker run --name prometheus -d -p 0.0.0.0:9090:9090 --volume="/Users/leemike/go/src/go-prometheus-exporter/config":/etc/config prom/prometheus --config.file=/etc/config/prometheus.yml
