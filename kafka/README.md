# Kafka 

This folder includes a set of YAML file to install and configure the requisites for making use of the kafka consumer developed in this repository.

## Kafka Server

### Prerequisites

- Openshift 4.10+ in AWS (*Because will use loadbalancer as listener type* )
- Red Hat Integration - AMQ Streams

### Setting Up Kafka Server

- Create Kafka server

```$bash
oc new-project kafka
oc apply -f kafka_server.yaml -n kafka
```

- Once the Kafka server has been installed properly, it is time to obtains the public 

```$bash
oc apply -f kafka_topic.yaml -n kafka
```

- Finally, it is required to create a new kafka topic

```$bash
oc apply -f kafka_topic.yaml -n kafka
```

## Kafka Client

First at all, it is required to download the respective 

- Create Topic

```$bash
bin/kafka-topics.sh --create --topic one --bootstrap-server a50e9747c108647d6aa9095ddd503628-1581487172.us-east-2.elb.amazonaws.com:9092
```

- Write event in a topic

```$bash
bin/kafka-console-producer.sh --topic my-topic --bootstrap-server a50e9747c108647d6aa9095ddd503628-1581487172.us-east-2.elb.amazonaws.com:9092
```

- Read events in a topic

```$bash
bin/kafka-console-consumer.sh --topic my-topic --from-beginning --bootstrap-server a50e9747c108647d6aa9095ddd503628-1581487172.us-east-2.elb.amazonaws.com:9092
```

- List consumer groups

```$bash
bin/kafka-consumer-groups.sh --bootstrap-server a50e9747c108647d6aa9095ddd503628-1581487172.us-east-2.elb.amazonaws.com:9092 --list
```

NOTE: Please visit [link](https://kafka.apache.org/quickstart) for more information about kafka client basics

## Author

Asier Cidon (@RedHat)
