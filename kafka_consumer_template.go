package main

import (
    "fmt"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
    // Setup consumer configuration
    consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
        "group.id":          "go-consumer-group",
        "auto.offset.reset": "earliest", // Change to "latest" for production
        "enable.auto.commit": false,     // We manually commit after processing
    })
    if err != nil {
        log.Fatalf("Failed to create consumer: %v", err)
    }

    // Subscribe to topics
    err = consumer.Subscribe("demo-topic", nil)
    if err != nil {
        log.Fatalf("Failed to subscribe to topic: %v", err)
    }

    // Handle graceful shutdown
    sigchan := make(chan os.Signal, 1)
    signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

    run := true

    log.Println("Kafka consumer started...")
    for run {
        select {
        case sig := <-sigchan:
            log.Printf("Caught signal %v: terminating", sig)
            run = false
        default:
            msg, err := consumer.ReadMessage(1 * time.Second)
            if err != nil {
                // Timeout or temporary error
                if kafkaErr, ok := err.(kafka.Error); ok && kafkaErr.Code() == kafka.ErrTimedOut {
                    continue
                }
                log.Printf("Consumer error: %v\n", err)
                continue
            }

            // ✅ Process the message
            fmt.Printf("Received message: %s [topic: %s, partition: %d, offset: %v]\n",
                string(msg.Value), *msg.TopicPartition.Topic, msg.TopicPartition.Partition, msg.TopicPartition.Offset)

            // Simulate processing success (add retry or error handling as needed)

            // ✅ Commit offset manually
            _, err = consumer.CommitMessage(msg)
            if err != nil {
                log.Printf("Commit error: %v\n", err)
            }
        }
    }

    // ✅ Close consumer safely
    log.Println("Closing consumer...")
    err = consumer.Close()
    if err != nil {
        log.Fatalf("Failed to close consumer: %v", err)
    }

    log.Println("Consumer shutdown complete.")
}
