import csv, pprint
import numpy             as np
import matplotlib.pyplot as plt

## Load data
pp = pprint.PrettyPrinter(indent=2)
data = {}
metrics = {
    'edge_case_latency': {
        'default_delta':       2.0,
        'max_tries':           5,
        'percentage_miss':     0.0,
        'percentage_faults':   0.0,
        'probability_to_fail': 0.0,
        'bandwidth':           250,
        'latency':             650.0
    },
    'edge_case_bandwidth': {
        'default_delta':       2.0,
        'max_tries':           5,
        'percentage_miss':     0.0,
        'percentage_faults':   0.0,
        'probability_to_fail': 0.0,
        'bandwidth':           500,
        'latency':             100.0
    },
    'edge_case_faults': {
        'default_delta':       2.0,
        'max_tries':           5,
        'percentage_miss':     8.0,
        'percentage_faults':   40.0,
        'probability_to_fail': 40.0,
        'bandwidth':           250,
        'latency':             100.0
    },
    'edge_case_lost_messages_and_faults': {
        'default_delta':       2.0,
        'max_tries':           5,
        'percentage_miss':     60.0,
        'percentage_faults':   30.0,
        'probability_to_fail': 30.0,
        'bandwidth':           250,
        'latency':             100.0
    },
    'edge_case_high_default_delta': {
        'default_delta':       5.0,
        'max_tries':           5,
        'percentage_miss':     0.0,
        'percentage_faults':   0.0,
        'probability_to_fail': 0.0,
        'bandwidth':           250,
        'latency':             100.0
    },
    'edge_case_low_default_delta': {
        'default_delta':       1.0,
        'max_tries':           5,
        'percentage_miss':     0.0,
        'percentage_faults':   0.0,
        'probability_to_fail': 0.0,
        'bandwidth':           250,
        'latency':             100.0
    }
}

result_type = 'all_edge_cases'
levels      = metrics.keys()

default_delta       = None
max_tries           = None
percentage_miss     = None
percentage_faults   = None
probability_to_fail = None
bandwidth           = None
latency             = None

def mutation_error(mutation):
    return 0

def valid_row(row):
    return (float(row[1]) == default_delta  and
    int(row[2])   == max_tries           and
    float(row[3]) == percentage_miss     and
    float(row[4]) == percentage_faults   and
    float(row[5]) == probability_to_fail and
    float(row[6]) == latency             and
    int(row[7])   == bandwidth)

print(levels)

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

        print(level)

        for row in readCSV:
            if not row[9] in data.keys():
                data[row[9]] = []

            if valid_row(row):
                data[row[9]].append((row[0], row[10]))

        pp.pprint(data)
    if data['ring']:
        mutations = {}
        nodes = [value[0] for value in data['ring']]

        for mutation in data:
            mutations[mutation] = {}
            mutations[mutation]['times'] = [float(value[1]) for value in data[mutation]]
            # mutations[mutation]['error+'] = [float(value[1]) + error for value in data[mutation]]
            # mutations[mutation]['error+2'] = [float(value[1]) + mutation_error(mutation) for value in data[mutation]]
            # mutations[mutation]['error-'] = [float(value[1]) - error for value in data[mutation]]
            # mutations[mutation]['error-2'] = [float(value[1]) - mutation_error(mutation) for value in data[mutation]]

        pp.pprint(mutations)

        ## Plot charts
        plt.xlabel('# nodes')
        plt.ylabel('consensus time (ms)')

        # Early
        # plt.fill_between(nodes, mutations['early']['error-'], mutations['early']['error+'], color="#e5ce70", alpha=0.4)
        # plt.fill_between(nodes, mutations['early']['error-2'], mutations['early']['error+2'], color="#ddc256", alpha=0.6)
        nodes = [value[0] for value in data['early']]
        plt.plot(nodes, mutations['early']['times'], color="#ea9f15", alpha=0.8, label="Early")

        # Ring
        # plt.fill_between(nodes, mutations['ring']['error-'], mutations['ring']['error+'], color="#adf442", alpha=0.4)
        # plt.fill_between(nodes, mutations['ring']['error-2'], mutations['ring']['error+2'], color="#87c924", alpha=0.6)
        nodes = [value[0] for value in data['ring']]
        plt.plot(nodes, mutations['ring']['times'], color="#1cbf42", alpha=0.8, label="Ring")

        # Old Ring
        # plt.fill_between(nodes, mutations['old_ring']['error-'], mutations['old_ring']['error+'], color="#ddf342", alpha=0.4)
        # plt.fill_between(nodes, mutations['old_ring']['error-2'], mutations['old_ring']['error+2'], color="#97c934", alpha=0.6)
        nodes = [value[0] for value in data['old_ring']]
        plt.plot(nodes, mutations['old_ring']['times'], color="#2c3f42", alpha=0.8, label="Old Ring")

        # Centralized
        # plt.fill_between(nodes, mutations['centralized']['error-'], mutations['centralized']['error+'], color="#f48989", alpha=0.4)
        # plt.fill_between(nodes, mutations['centralized']['error-2'], mutations['centralized']['error+2'], color="#d16262", alpha=0.6)
        nodes = [value[0] for value in data['centralized']]
        plt.plot(nodes, mutations['centralized']['times'], color="#ea1515", alpha=0.8, label="Centralized")

        # Gossip
        # plt.fill_between(nodes, mutations['gossip']['error-'], mutations['gossip']['error+'], color="skyblue", alpha=0.4)
        # plt.fill_between(nodes, mutations['gossip']['error-2'], mutations['gossip']['error+2'], color="#3e7df2", alpha=0.6)
        nodes = [value[0] for value in data['gossip']]
        plt.plot(nodes, mutations['gossip']['times'], color="Slateblue", alpha=0.8, label="Gossip")

        ## Show results
        plt.legend()
        # plt.show()
        plt.savefig('datasets/results/edge_cases/' + level + '.png')

        # Clear plot
        plt.clf()
        data.clear()
