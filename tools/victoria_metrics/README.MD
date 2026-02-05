VictoriaMetrics 是一个高性能、经济高效的时间序列数据库（TSDB）和监控解决方案:类比Prometheus

## 官网文档:
https://docs.victoriametrics.com/victoriametrics/quick-start/
## git
https://github.com/VictoriaMetrics/VictoriaMetrics

## docker 安装
```
docker pull victoriametrics/victoria-metrics:v1.135.0
docker run -it --rm -v `pwd`/victoria-metrics-data:/victoria-metrics-data -p 8428:8428 \
 victoriametrics/victoria-metrics:v1.135.0 --selfScrapeInterval=5s -storageDataPath=victoria-metrics-data
```

## 访问地址:
http://localhost:8428/

