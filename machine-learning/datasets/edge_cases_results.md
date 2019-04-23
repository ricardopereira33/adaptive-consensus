# Results about edge cases

In order to have an overview and to be able to analyze all the results, a set of graphs was built. Each of these graphs represents a specific environment.

### High latency

```
default_delta       = 2 s
max_tries           = 5 tries
percentage_miss     = 0.0 %
percentage_faults   = 0.0 %
probability_to_fail = 0.0 %
bandwidth           = 250 msgs/s
latency             = 650.0 ms
```

#### Graph

![](results/edge_cases/edge_case_latency.png)

<div style="page-break-after: always;"></div>

### Large bandwitdh

```
default_delta       = 2 s
max_tries           = 5 tries
percentage_miss     = 0.0 %
percentage_faults   = 0.0 %
probability_to_fail = 0.0 %
bandwidth           = 500 msgs/s
latency             = 100.0 ms
```

#### Graph

![](results/edge_cases/edge_case_bandwidth.png)

<div style="page-break-after: always;"></div>

### Faulty environment

```
default_delta       = 2 s
max_tries           = 5 tries
percentage_miss     = 8.0 %
percentage_faults   = 40.0 %
probability_to_fail = 40.0 %
bandwidth           = 250 msgs/s
latency             = 125.0 ms
```

#### Graph

![](results/edge_cases/edge_case_faults.png)

<div style="page-break-after: always;"></div>

### High probability of losing messages

```
default_delta       = 2 s
max_tries           = 5 tries
percentage_miss     = 60.0 %
percentage_faults   = 30.0 %
probability_to_fail = 30.0 %
bandwidth           = 250 msgs/s
latency             = 100.0 ms
```

#### Graph

![](results/edge_cases/edge_case_lost_messages_and_faults.png)

<div style="page-break-after: always;"></div>

### High default delta

```
default_delta       = 5 s
max_tries           = 5 tries
percentage_miss     = 0.0 %
percentage_faults   = 0.0 %
probability_to_fail = 0.0 %
bandwidth           = 250 msgs/s
latency             = 100.0 ms
```

#### Graph

![](results/edge_cases/edge_case_high_default_delta.png)

<div style="page-break-after: always;"></div>

### Low default delta

```
default_delta       = 1 s
max_tries           = 5 tries
percentage_miss     = 0.0 %
percentage_faults   = 0.0 %
probability_to_fail = 0.0 %
bandwidth           = 250 msgs/s
latency             = 100.0 ms
```

#### Graph

![](results/edge_cases/edge_case_low_default_delta.png)

# Conclusion

All mutation was submitted to several environments and, as we can see, some mutations got better performance in a specific environment. However, the gossip mutation got always the best perfomance (unquestionably) in all environments.