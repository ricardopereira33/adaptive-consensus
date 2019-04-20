import csv
import numpy             as np
import matplotlib.pyplot as plt

## Load data
data = {}

metrics = {
    'low': {
        'default_delta':       0.5,
        'max_tries':           3,
        'percentage_miss':     0.0,
        'percentage_faults':   0.0,
        'probability_to_fail': 0.0,
        'bandwidth':           20,
        'latency':             25.0
    },
    'medium': {
        'default_delta':       1.5,
        'max_tries':           3,
        'percentage_miss':     4.0,
        'percentage_faults':   10.0,
        'probability_to_fail': 5.0,
        'bandwidth':           100,
        'latency':             125.0
    },
    'large_bandwidth': {
        'default_delta':       1.5,
        'max_tries':           3,
        'percentage_miss':     4.0,
        'percentage_faults':   15.0,
        'probability_to_fail': 10.0,
        'bandwidth':           180,
        'latency':             125.0
    },
    'high_latency': {
        'default_delta':       1.5,
        'max_tries':           3,
        'percentage_miss':     4.0,
        'percentage_faults':   15.0,
        'probability_to_fail': 10.0,
        'bandwidth':           100,
        'latency':             225.0
    },
    'edge_case_latency': {
        'default_delta':       1.0,
        'max_tries':           3,
        'percentage_miss':     0.0,
        'percentage_faults':   0.0,
        'probability_to_fail': 0.0,
        'bandwidth':           100,
        'latency':             650.0
    },
    'edge_case_bandwidth': {
        'default_delta':       1.0,
        'max_tries':           3,
        'percentage_miss':     0.0,
        'percentage_faults':   0.0,
        'probability_to_fail': 0.0,
        'bandwidth':           500,
        'latency':             125.0
    },
    'edge_case_faults': {
        'default_delta':       1.0,
        'max_tries':           3,
        'percentage_miss':     0.0,
        'percentage_faults':   40.0,
        'probability_to_fail': 40.0,
        'bandwidth':           200,
        'latency':             125.0
    },
    'edge_case_lost_messages': {
        'default_delta':       1.0,
        'max_tries':           3,
        'percentage_miss':     60.0,
        'percentage_faults':   0.0,
        'probability_to_fail': 0.0,
        'bandwidth':           200,
        'latency':             125.0
    },
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

result_type = 'ring_mutations'
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

def mutation_error(mutation):
    if 'early' == mutation:
        return 573
    elif 'ring2' == mutation:
        return 564
    elif 'centralized' == mutation:
        return 800
    elif 'gossip' == mutation:
        return 296
    else:
        return 0

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
            if not row[8] in data.keys():
                data[row[8]] = []

            if valid_row(row):
                data[row[8]].append((row[0], row[9]))

    print(data)

    if data['ring']:
        mutations = {}
        nodes = [value[0] for value in data['ring']]

        for mutation in data:
            mutations[mutation] = {}
            mutations[mutation]['times'] = [float(value[1]) for value in data[mutation]]
            mutations[mutation]['error+'] = [float(value[1]) + error for value in data[mutation]]
            mutations[mutation]['error+2'] = [float(value[1]) + mutation_error(mutation) for value in data[mutation]]
            mutations[mutation]['error-'] = [float(value[1]) - error for value in data[mutation]]
            mutations[mutation]['error-2'] = [float(value[1]) - mutation_error(mutation) for value in data[mutation]]

        ## Plot charts
        plt.xlabel('# nodes')
        plt.ylabel('consensus time (ms)')

        # # Early
        # plt.fill_between(nodes, mutations['early']['error-'], mutations['early']['error+'], color="#e5ce70", alpha=0.4)
        # plt.fill_between(nodes, mutations['early']['error-2'], mutations['early']['error+2'], color="#ddc256", alpha=0.6)
        # plt.plot(nodes, mutations['early']['times'], color="#ea9f15", alpha=0.8, label="Early")

        # Ring
        plt.fill_between(nodes, mutations['ring']['error-'], mutations['ring']['error+'], color="#adf442", alpha=0.4)
        plt.fill_between(nodes, mutations['ring']['error-2'], mutations['ring']['error+2'], color="#87c924", alpha=0.6)
        plt.plot(nodes, mutations['ring']['times'], color="#1cbf42", alpha=0.8, label="Ring")

        # Old Ring
        plt.fill_between(nodes, mutations['old_ring']['error-'], mutations['old_ring']['error+'], color="#ddf342", alpha=0.4)
        plt.fill_between(nodes, mutations['old_ring']['error-2'], mutations['old_ring']['error+2'], color="#97c934", alpha=0.6)
        plt.plot(nodes, mutations['old_ring']['times'], color="#2c3f42", alpha=0.8, label="Old Ring")


        # # Centralized
        # plt.fill_between(nodes, mutations['centralized']['error-'], mutations['centralized']['error+'], color="#f48989", alpha=0.4)
        # plt.fill_between(nodes, mutations['centralized']['error-2'], mutations['centralized']['error+2'], color="#d16262", alpha=0.6)
        # plt.plot(nodes, mutations['centralized']['times'], color="#ea1515", alpha=0.8, label="Centralized")

        # # Gossip
        # plt.fill_between(nodes, mutations['gossip']['error-'], mutations['gossip']['error+'], color="skyblue", alpha=0.4)
        # plt.fill_between(nodes, mutations['gossip']['error-2'], mutations['gossip']['error+2'], color="#3e7df2", alpha=0.6)
        # plt.plot(nodes, mutations['gossip']['times'], color="Slateblue", alpha=0.8, label="Gossip")

        ## Show results
        plt.legend()
        # plt.show()
        plt.savefig('datasets/ring_results/' + level + '.png')

        # Clear plot
        plt.clf()

        data.clear()
