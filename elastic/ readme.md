
# Elasticsearch 常用 API

## 索引操作

### 创建索引
```json
PUT /index_name
{
    "settings": {
        "number_of_shards": 1,
        "number_of_replicas": 0
    }
}
```

### 删除索引
```json
DELETE /index_name
```

### 查看索引
```json
GET /_cat/indices
GET /index_name
```

## 文档操作

### 新增文档
```json
POST /index_name/_doc
{
    "field": "value"
}
```

### 更新文档
```json
PUT /index_name/_doc/1
{
    "field": "new_value"
}
```

### 删除文档
```json
DELETE /index_name/_doc/1
```

### 查询文档
```json
GET /index_name/_doc/1
```

## 搜索操作

### 基础搜索
```json
GET /index_name/_search
{
    "query": {
        "match": {
            "field": "keyword"
        }
    }
}
```

### 聚合查询
```json
GET /index_name/_search
{
    "aggs": {
        "group": {
            "terms": {
                "field": "category"
            }
        }
    }
}
```
