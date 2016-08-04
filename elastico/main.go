package main

import (
	"fmt"
	"gopkg.in/olivere/elastic.v3"
	"log"
	"time"
)

const (
	indexName    = "grid_jobs"
	docType      = "session_history"
	indexMapping = `{
						"mappings" : {
							"session_history" : {
								"properties" : {
									"CLUSTER_NAME" : {
										"type": "string",
										"index": "not_analyzed"
									},
									"SESSION_NAME" : {
										"type": "string",
										"index": "not_analyzed"
									},
									"PRIORITY": {"type": "integer"}
								}
							}
						}
					}`
)

type Session struct {
	ClusterName string `json:"CLUSTER_NAME"`
	SessionName string `json:"SESSION_NAME"`
	Priority    int    `json:"PRIORITY`
}

func main() {
	// Health check identifies dead nodes/connections https://github.com/olivere/elastic/wiki/Healthcheck
	client, err := elastic.NewClient(elastic.SetHealthcheckInterval(5*time.Second),
		elastic.SetURL("http://localhost:9200", "http://localhost:9201"),
	)
	if err != nil {
		log.Fatal("Error connecting to elasticsearch cluster :", err.Error())
	}
	esversion, err := client.ElasticsearchVersion("http://localhost:9200")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

	// Let's create an Index
	exists, err := client.IndexExists(indexName).Do()
	if err != nil {
		panic(err)
	}
	if !exists {
		createIndex, err := client.CreateIndex(indexName).
			Body(indexMapping).
			Do()
		if err != nil {
			panic(err)
		}
		if !createIndex.Acknowledged {
			log.Printf("Creation of Index not Acknowledged, Beware !")
		}
		fmt.Printf("Created Index %s\n", indexName)
	}

	s := Session{SessionName: "Test Job", ClusterName: "UAT Grid", Priority: 5}
	_, err = client.Index().
		Index(indexName).
		Type(docType).
		Id("1").
		BodyJson(s).
		Refresh(true).
		Do()
	if err != nil {
		panic(err)
	}
}
