# Results about ring mutations

In order to have an overview and to be able to analyze all the results, a set of graphs was built. Each of these graphs represents a specific environment.

## New ring mutation

A new implementation of it was made, in order to improve the performance of the ring mutation. Two things were changed:

- The `delta0` function only allows it to send a message to the next or previous peer.

- The `delta` function, responsible for the retransmission, allows the retransmission for the next or previous peer as well. However, when a peer didn't send a response, the message is sent to the next (or previous) and skip this one. The number of skips/steps is independent for each side.

## Messages exchange (overview)

### Old Ring

Sent                                   |  Received
:-------------------------------------:|:----------------------------------------:
![](results/ring/sent_old_ring.png)    |  ![](results/ring/received_old_ring.png)

### New Ring

Sent                                   |  Received
:-------------------------------------:|:------------------------------------:
![](results/ring/sent_ring.png)        |  ![](results/ring/received_ring.png)

<div style="page-break-after: always;"></div>

## Environment without faults

### Low default delta

```
default_delta       = 1 s
max_tries           = 3 tries
percentage_miss     = 0.0 %
percentage_faults   = 0.0 %
probability_to_fail = 0.0 %
bandwidth           = 200 msgs/s
latency             = 125.0 ms
```

#### Graph

![](results/ring/ring_without_faults_low_default_delta.png)

<div style="page-break-after: always;"></div>

### Normal Case

```
default_delta       = 3 s
max_tries           = 3 tries
percentage_miss     = 0.0 %
percentage_faults   = 0.0 %
probability_to_fail = 0.0 %
bandwidth           = 100 msgs/s
latency             = 125.0 ms
```

#### Graph

![](results/ring/ring_without_faults_normal.png)

<div style="page-break-after: always;"></div>

### Large bandwidth

```
default_delta       = 3 s
max_tries           = 3 tries
percentage_miss     = 0.0 %
percentage_faults   = 0.0 %
probability_to_fail = 0.0 %
bandwidth           = 300 msgs/s
latency             = 125.0 ms
```

#### Graph

![](results/ring/ring_without_faults_large_bandwidth.png)

<div style="page-break-after: always;"></div>

### High latency

```
default_delta       = 3 s
max_tries           = 3 tries
percentage_miss     = 0.0 %
percentage_faults   = 0.0 %
probability_to_fail = 0.0 %
bandwidth           = 100 msgs/s
latency             = 375.0 ms
```

#### Graph

![](results/ring/ring_without_faults_high_latency.png)

<div style="page-break-after: always;"></div>

## Environment with faults

### Low default delta

```
default_delta       = 1 s
max_tries           = 3 tries
percentage_miss     = 8.0 %
percentage_faults   = 15.0 %
probability_to_fail = 15.0 %
bandwidth           = 200 msgs/s
latency             = 125.0 ms
```

#### Graph

![](results/ring/ring_with_faults_low_default_delta.png)

<div style="page-break-after: always;"></div>

### Normal Case

```
default_delta       = 3 s
max_tries           = 3 tries
percentage_miss     = 8.0 %
percentage_faults   = 15.0 %
probability_to_fail = 15.0 %
bandwidth           = 100 msgs/s
latency             = 125.0 ms
```

#### Graph

![](results/ring/ring_with_faults_normal.png)

<div style="page-break-after: always;"></div>

### Large bandwidth

```
default_delta       = 3 s
max_tries           = 3 tries
percentage_miss     = 8.0 %
percentage_faults   = 15.0 %
probability_to_fail = 15.0 %
bandwidth           = 300 msgs/s
latency             = 125.0 ms
```

#### Graph

![](results/ring/ring_with_faults_large_bandwidth.png)

<div style="page-break-after: always;"></div>

### High latency

```
default_delta       = 3 s
max_tries           = 3 tries
percentage_miss     = 8.0 %
percentage_faults   = 15.0 %
probability_to_fail = 15.0 %
bandwidth           = 100 msgs/s
latency             = 375.0 ms
```

#### Graph

![](results/ring/ring_with_faults_high_latency.png)

<div style="page-break-after: always;"></div>

### A lot of faults

```
default_delta       = 3 s
max_tries           = 3 tries
percentage_miss     = 16.0 %
percentage_faults   = 30.0 %
probability_to_fail = 30.0 %
bandwidth           = 200 msgs/s
latency             = 125.0 ms
```

#### Graph

![](results/ring/ring_with_a_lot_of_faults.png)

# Conclusion

The previous graphs show that the new ring mutation brings better performance in a faulty environment compared to the old ring mutation. However, for a non-faulty environment, the old ring mutation has a better performance.