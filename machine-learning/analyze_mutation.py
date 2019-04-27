import csv, random, pprint
import numpy             as np
import matplotlib.pyplot as plt

pp      = pprint.PrettyPrinter(indent=2)
data    = {}
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

def rand_color():
    r = lambda: random.randint(0,255)

    for i in range(3):
        value = r()

    return value

for mutation in ['early', 'ring', 'old_ring', 'centralized', 'gossip']:
    for level in levels:
        default_delta       = metrics[level]['default_delta']
        max_tries           = metrics[level]['max_tries']
        percentage_miss     = metrics[level]['percentage_miss']
        percentage_faults   = metrics[level]['percentage_faults']
        probability_to_fail = metrics[level]['probability_to_fail']
        bandwidth           = metrics[level]['bandwidth']
        latency             = metrics[level]['latency']

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

        if data[mutation]:
            mutations           = {}
            mutations[mutation] = {}
            mutations[mutation]['times'] = [float(value[1]) for value in data[mutation]]

            ## Plot charts
            plt.xlabel('# nodes')
            plt.ylabel('consensus time (ms)')

            color = '#%02X%02X%02X' % (rand_color(),rand_color(),rand_color())
            nodes = [value[0] for value in data[mutation]]

            plt.plot(nodes, mutations[mutation]['times'], color=color, alpha=0.8, label=level)

            # Clear plot
            # plt.clf()
            data.clear()

    ## Show results
    plt.legend()
    # plt.show()
    plt.savefig('datasets/results/edge_cases/' + mutation + '_' + level + '.png')
    plt.clf()
