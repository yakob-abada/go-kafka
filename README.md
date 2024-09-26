# go-kafka

- This is a small demonstration how to run kafka on golang.
- To get a good understanding about kafka and how it works go [read here](https://www.freecodecamp.org/news/build-a-real-time-notification-system-with-go-and-kafka/).
- We get files:
  * `producer.go` is responsible to create a producer and send a message with content `notification_text` to kafka server. Topic has been set too.
  * `consumer.go` is responsible to create a consumer and that listen to any coming message from kafka.
  * `docker-compose.yml` helps to run Apache Kafka server running inside a container.