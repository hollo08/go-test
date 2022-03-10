package main

import (
    "fmt"
    "github.com/Shopify/sarama"
    "sync"
)

var (
    wg sync.WaitGroup
)

func main() {
    consumer, err := sarama.NewConsumer([]string{"10.1.1.210:9092"}, nil)
    
    if err != nil {
        panic(err)
    }
    
    partitionList, err := consumer.Partitions("testGo")
    
    if err != nil {
        panic(err)
    }
    
    for partition := range partitionList {
        pc, err := consumer.ConsumePartition("testGo", int32(partition), sarama.OffsetNewest)
        if err != nil {
            panic(err)
        }
        
        defer pc.AsyncClose()
        
        wg.Add(1)
        
        go func(sarama.PartitionConsumer) {
            defer wg.Done()
            for msg := range pc.Messages() {
                fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
            }
        }(pc)
        wg.Wait()
        consumer.Close()
    }
}
