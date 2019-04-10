# Results

In order to have an overview and to be able to analyze all the results, a set of graphs was built. Each of these graphs represents a specific environment.

## Without faults

```
default_delta       = 0.5 s
max_tries           = 3 tries
percentage_miss     = 0.0 %
percentage_faults   = 0.0 %
probability_to_fail = 0.0 %
bandwidth           = 20 msgs/s
latency             = 25.0 ms
```

### Graph

![](results/plot_1_without_faults.png)

<div style="page-break-after: always;"></div>

## Normal Case

```
default_delta       = 1.5 s
max_tries           = 3 tries
percentage_miss     = 4.0 %
percentage_faults   = 10.0 %
probability_to_fail = 5.00.0 %
bandwidth           = 100 msgs/s
latency             = 125.0 ms
```

### Graph

![](results/plot_2_normal_case.png)

<div style="page-break-after: always;"></div>

## Large bandwidth

```
default_delta       = 1.5 s
max_tries           = 3 tries
percentage_miss     = 4.0 %
percentage_faults   = 15.0 %
probability_to_fail = 10.0 %
bandwidth           = 180 msgs/s
latency             = 125.0 ms
```

### Graph

![](results/plot_3_large_bandwidth.png)

<div style="page-break-after: always;"></div>

## High latency

```
default_delta       = 1.5 s
max_tries           = 3 tries
percentage_miss     = 4.0 %
percentage_faults   = 15.0 %
probability_to_fail = 10.0 %
bandwidth           = 100 msgs/s
latency             = 225.0 ms
```

### Graph

![](results/plot_3_high_latency.png)
