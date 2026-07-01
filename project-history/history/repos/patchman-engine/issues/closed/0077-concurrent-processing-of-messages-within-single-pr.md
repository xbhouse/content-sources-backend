---
type: issue
number: 77
title: "Concurrent processing of messages within single process"
state: closed
author: semtexzv
created: 2020-01-27T14:00:45Z
updated: 2020-05-28T08:23:24Z
closed: 2020-05-28T08:23:23Z
labels: []
url: https://github.com/RedHatInsights/patchman-engine/issues/77
---

# Issue #77: Concurrent processing of messages within single process

**Author**: @semtexzv
**Created**: January 27, 2020 at 02:00 PM UTC
**Status**: Closed
**Labels**: None

## Description

## Kafka
In order to understand the problem, you must understand how kafka works. In essence kafka is a set of append-only files, with a distributed key-value store. Each kafka *topic-partition* pair corresponds to single file, each entry in this file being a message. Then, on top of this, kafka uses zookeper to store offsets into these logs, which denote position of last processed message.

Kafka consumer is very simple. On start it does some setup, picks which partitions it will read from, and reads offsets from zookepeer. Then it contacts kafka server, and waits for new messages being appended to chosen partitons. After correctly reading these messages, a consumer will have to at some point commit the fact that it read them. This is done by writing message offset into zookeeper.
In the event of consumer crash, it can read where it left of, and continue processing messages without skipping or double processing the messages.

These commits can be automatic or manual. In automatic commit, the consumer commits messages immediately after receiving them. In the manual case, programmer can chose when to commit the messages.

Note: Each partition can be consumed by at most 1 consumer from single CMSfR application, but each consumer can read from multiple partitions.

###  Previous state
The current implementation of our kafka readers utilized automatic immediate commiting, which 
meant that upon receiving message, this message was immediately committed as processed sucessfully. This posed a problem, when process crashed after receiving messages. The received messages would be committed, but they would be processed asynchronously later.

``` 
| receive m1 | commit m1  | receive m2 | commit m2 | ---- spawn m1 and m2 ----> | handle m1 | handle m2 |
```

and for single message:
```
| receive m1 | commit m1 | ----- spawn m1-----> | handle m1 |
```
This mean't if process crashed:
```
| receive m1 | commit m1 | ~~~ crash ~~~~
``` 
we would have not processed `m1` correctly, but it still would be commited.

This pattern was implemented because it allowed us to spawn a goroutine for each message, and therefore allowed us almost infinite concurrency. It was an easy way to get up and running.

### Manual commits
The solution is to use manual commits, and commit messages only after they have been correctly processed. This however requires us to process messages synchronously, reducing overall throughput. This reduction is also bad because we are bound by waiting on inventory & vmaas and not by actual work done in this component.

### The problem with manual commits
This change would result in following flow for messages `m1` and `m2`
```
 | receive m1 | handle m1 | commit m1 | receive m2 | handle m2 | commit m2 |
```
Since the `handle` phase takes much longer than `receive` and `commit` it means that we incur a lot of lag when handling a lot of messages. This is in some way alleviated by the fact, that we can run multiple independent `listener` component instances. Each listener has own consumer, which can process messages independently. 

If we want to introduce concurrency in one listener, we need to ensure, that messages are committed 
in the same order they were received. If we don't do this, we might create a situation which would allow skipping of messages and double processing of messages. This is caused by [segmentio/kafka-go](
https://github.com/segmentio/kafka-go/blob/16d85b1f54d2396fe4fecf43248d14ac51bec961/reader.go#Le169). The library which we are using is just sending raw message offsets without doing any internal reordering, and if we have following flow:

```
 | receive m1 | receive m2 | handle m2 | commit m2 | ~~~ crash ~~~ 
```
We would have commited m2 before even processing m1, and that means we have missed m1.

### Solution
1. Multiple listeners -  Immediate solution which ensures acceptable performance is spawning multiple listeners, the maximal number of listeners being number of partitions, In this case, we would have N pods, each listener running only 1 consumer
2. Multiple consumers per 1 listener pod - This would reduce overhead, and allow us run smaller number of pods. Concurrency is still limited by number of partitions.
3. Internal queueing within one listener - In this case we would have one consumer per pod, spawning work into work queue. The results from this work queue would have to be ordered, ensuring no messages are skipped (Double processing might occur in rare case).

| Partitions |  Listener pods | Consumers per pod | Goroutines per consumer | +                 | -                                 | Concurrency |
|------------|----------------|-------------------|-------------------------|-------------------|-----------------------------------|-------------|
| P          | N              | 1                 | 1                       |  Simple           |  Slow, Pod overhead               | min(N,P)    |
| P          | N              | M                 | 1                       |  Low pod overhead |  Error handling, complexity       | min(N*M, P) |
| P          | N              | M                 | S                       | Max throughput    | Complexity - work-queue in golang | N*S         |

### Why is concurrency important ?
Throughput. Total throughput of our application can be roughly approximated by how many systems we can process concurrently and how much time takes one system `throughput = concurrency / latency`. ` 10 systems in parallel / 1 second per system = 10 systems per second`

If we implement our system with concurrency of `20`, and slowest stage of the processing pipeline takes `500 ms`, then we will be able to process at most ` 20 / 0.5 = 40` systems per second.

---

## Discussion

### Comment by @beav on January 28, 2020 at 01:29 PM UTC

Thanks for the detailed writeup!

If you used the one consumer per pod solution, could the number of running pods scale based on queue depth?

### Comment by @semtexzv on January 30, 2020 at 03:50 PM UTC

Yes, but the hassle of creating new consumer & re-balancing the partitions could make this not worth it. 

It'd also be difficult to use custom metric with autoscaler.

We have implemented 8x8x1 topology with 8 listeners, each running 8 consumers, and one goroutine per consumer. This should be OK for the time being.

### Comment by @semtexzv on May 28, 2020 at 08:23 AM UTC

I think this part of the project is now well understood and for the most part solved. We have multiple threads & pods providing a high degree of concurrency

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/issues/77*
