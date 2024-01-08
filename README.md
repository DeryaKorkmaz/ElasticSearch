# ElasticSearch

## Create Index

##### PUT {{BASE_URL}}/api/products

```json
{
  "mappings": {
    "properties": {
      "ProductID": { "type": "integer" },
      "ProductName": { "type": "text" },
      "ProductType": { "type": "text" },
      "StockStatus": { "type": "text" }
            }
       }
}
```

## Open Index

##### POST {{BASE_URL}}/api/products/_open


## Close Index

##### POST {{BASE_URL}}/api/products/_close


## All Indexes

##### GET {{BASE_URL}}/_cat/indices

## Create Analyzer

Search sonuçlarından esnek yanıt alabilmek için analyzer indeximize analyzer oluşturuyoruz.

##### PUT  {{BASE_URL}}/api/products/_settings

```json
{
  "settings": {
    "analysis": {
      "analyzer": {
        "txt_analyzer": {
          "type": "custom",
          "tokenizer": "standard",
          "filter": [
            "lowercase",
            "txt_filter"
          ]
        }
      },
      "filter": {
        "txt_filter": {
          "type": "asciifolding",
          "preserve_original": true
        }
      }
    }
  },
 "mappings": {
    "properties": {
      "ProductID": { "type": "integer" },
      "ProductName": { "type": "text" },
      "ProductType": { "type": "text" },
      "StockStatus": { "type": "text" }
            }
       }
  }
```

## Index Settings

##### GET {{BASE_URL}}/api/products/_settings



## Insert Data (Bulk Import)
##### POST {{BASE_URL}}/api/products/_bulk

```json
{"index":{"_index":"products"}}
{"ProductID":1002,"ProductName":"Phone","ProductType":"E","StockStatus":"Yes"}
{"index":{"_index":"products"}}
{"ProductID":1003,"ProductName":"Tv","ProductType":"E","StockStatus":"Yes"}
{"index":{"_index":"products"}}
{"ProductID":1004,"ProductName":"PC","ProductType":"S","StockStatus":"No"}
...
```
##  Search
##### GET {{BASE_URL}}/api/products/_search

```json
{
  "query": {
    "match_phrase": {
      "ProductID":454360
    }
  }
}
```


## Multi Search

##### GET {{BASE_URL}}/api/products/_search

```json
{
    "query": {
        "bool": {
            "must": [
                {
                    "multi_match": {
                        "query": "TV",
                        "type": "bool_prefix",
                        "fields": [
                            "ProductName"
                        ]
                    }
                }
            ]
        }
    },
    "sort": [
        {
            "Rating": {
                "order": "desc"
            }
        }
    ],
    "size": 10
}
```

## All Data

##### GET {{BASE_URL}}/api/products/_search

```json
{
  "query": {
    "match_all": {}
  }
}
```

## Truncate Data

##### POST {{BASE_URL}}/api/products/_delete_by_query

```json
{
  "query": {
    "match_all": {}
  }
}
```
