import csv
import numpy             as np
import matplotlib.pyplot as plt

## Load data
data = {}

metrics = {
    'ring_without_faults_low_default_delta': {
        'default_delta':       1.0,
        'max_tries':           3,
        'percentage_miss':     0.0,
        'percentage_faults':   0.0,
        'probability_to_fail': 0.0,
        'bandwidth':           200,
        'latency':             125.0
    },
    'ring_without_faults_normal': {
        'default_delta':       3.0,
        'max_tries':           3,
        'percentage_miss':     0.0,
        'percentage_faults':   0.0,
        'probability_to_fail': 0.0,
        'bandwidth':           100,
        'latency':             125.0
    },
    'ring_without_faults_high_latency': {
        'default_delta':       3.0,
        'max_tries':           3,
        'percentage_miss':     0.0,
        'percentage_faults':   0.0,
        'probability_to_fail': 0.0,
        'bandwidth':           100,
        'latency':             375.0
    },
     'ring_without_faults_large_bandwidth': {
        'default_delta':       3.0,
        'max_tries':           3,
        'percentage_miss':     0.0,
        'percentage_faults':   0.0,
        'probability_to_fail': 0.0,
        'bandwidth':           300,
        'latency':             125.0
    },
    'ring_with_faults_low_default_delta': {
        'default_delta':       1.0,
        'max_tries':           3,
        'percentage_miss':     8.0,
        'percentage_faults':   15.0,
        'probability_to_fail': 15.0,
        'bandwidth':           200,
        'latency':             125.0
    },
    'ring_with_faults_normal': {
        'default_delta':       3.0,
        'max_tries':           3,
        'percentage_miss':     8.0,
        'percentage_faults':   15.0,
        'probability_to_fail': 15.0,
        'bandwidth':           100,
        'latency':             125.0
    },
    'ring_with_a_lot_of_faults': {
        'default_delta':       3.0,
        'max_tries':           3,
        'percentage_miss':     16.0,
        'percentage_faults':   30.0,
        'probability_to_fail': 30.0,
        'bandwidth':           200,
        'latency':             125.0
    },
    'ring_with_faults_high_latency': {
        'default_delta':       3.0,
        'max_tries':           3,
        'percentage_miss':     8.0,
        'percentage_faults':   15.0,
        'probability_to_fail': 15.0,
        'bandwidth':           100,
        'latency':             375.0
    },
    'ring_with_faults_large_bandwidth': {
        'default_delta':       3.0,
        'max_tries':           3,
        'percentage_miss':     8.0,
        'percentage_faults':   15.0,
        'probability_to_fail': 15.0,
        'bandwidth':           300,
        'latency':             125.0
    }
}

result_type = 'ring'
levels      = [
    'ring_without_faults_low_default_delta',
    'ring_without_faults_normal',
    'ring_without_faults_high_latency',
    'ring_without_faults_large_bandwidth',
    'ring_with_faults_low_default_delta',
    'ring_with_faults_normal',
    'ring_with_a_lot_of_faults',
    'ring_with_faults_high_latency',
    'ring_with_faults_large_bandwidth'
]

default_delta       = None
max_tries           = None
percentage_miss     = None
percentage_faults   = None
probability_to_fail = None
bandwidth           = None
latency             = None

def valid_row(row):
    return (float(row[1]) == default_delta  and
    int(row[2])   == max_tries           and
    float(row[3]) == percentage_miss     and
    float(row[4]) == percentage_faults   and
    float(row[5]) == probability_to_fail and
    float(row[6]) == latency             and
    int(row[7])   == bandwidth)

for level in levels:
    default_delta       = metrics[level]['default_delta']
    max_tries           = metrics[level]['max_tries']
    percentage_miss     = metrics[level]['percentage_miss']
    percentage_faults   = metrics[level]['percentage_faults']
    probability_to_fail = metrics[level]['probability_to_fail']
    bandwidth           = metrics[level]['bandwidth']
    latency             = metrics[level]['latency']

    nodes = 50
    # error = 571.0
    error = 0

    with open('datasets/results/global_results_' + result_type + '.csv') as csvfile:
        readCSV = csv.reader(csvfile, delimiter=',')
        next(readCSV, None)

        for row in readCSV:
            if not row[9] in data.keys():
                data[row[9]] = []

            if valid_row(row):
                data[row[9]].append((row[0], row[10]))

    print(data)

    if data['ring']:
        mutations = {}

        for mutation in data:
            mutations[mutation] = {}
            mutations[mutation]['times'] = [float(value[1]) for value in data[mutation]]

        ## Plot charts
        plt.xlabel('# nodes')
        plt.ylabel('consensus time (ms)')

        # Ring
        nodes = [value[0] for value in data['ring']]
        plt.plot(nodes, mutations['ring']['times'], color="#1cbf42", alpha=0.8, label="Ring")

        # Old Ring
        nodes = [value[0] for value in data['old_ring']]
        plt.plot(nodes, mutations['old_ring']['times'], color="#2c3f42", alpha=0.8, label="Old Ring")

        ## Show results
        plt.legend()
        # plt.show()
        plt.savefig('datasets/results/ring/' + level + '.png')

        # Clear plot
        plt.clf()
        data.clear()
