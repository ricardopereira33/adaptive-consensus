# Results about ring mutations

In order to have an overview and to be able to analyze all the results, a set of graphs was built. Each of these graphs represents a specific environment.

## Messages exchange

### Old Ring

Sent                                                |  Received
:--------------------------------------------------:|:-----------------------------------------------------:
![](ring_results/compare_rings/sent_old_ring.png)  |  ![](ring_results/compare_rings/received_old_ring.png)

### New Ring

Sent                                                |  Received
:--------------------------------------------------:|:-------------------------------------------------:
![](ring_results/compare_rings/sent_new_ring.png)  |  ![](ring_results/compare_rings/received_ring.png)

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

![](ring_results/ring_without_faults_low_default_delta.png)

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

![](ring_results/ring_without_faults_normal.png)

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

![](ring_results/ring_without_faults_large_bandwidth.png)

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

![](ring_results/ring_without_faults_high_latency.png)

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

![](ring_results/ring_with_faults_low_default_delta.png)

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

![](ring_results/ring_with_faults_normal.png)

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

![](ring_results/ring_with_faults_large_bandwidth.png)

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

![](ring_results/ring_with_faults_high_latency.png)

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

![](ring_results/ring_with_a_lot_of_faults.png)

