import csv, random, pprint, os
import numpy             as np
import matplotlib.pyplot as plt

pp   = pprint.PrettyPrinter(indent=2)
data = {}

os.system('rm datasets/results/bandwidth/*.png')

def rand_color():
    r = lambda: random.randint(0,255)

    for i in range(3):
        value = r()

    return value

for mutation in ['ring']:#, 'early', 'old_ring', 'centralized', 'gossip']:
    bandwidth_results      = mutation + '_bandwidth_usage'
    retransmission_results = mutation + '_retransmission'

    with open('../results/csv-datasets/' + bandwidth_results + '.csv') as csvfile_bandwidth:
        with open('../results/csv-datasets/' + retransmission_results + '.csv') as csvfile_retransmission:
            bandwidth_file         = csv.reader(csvfile_bandwidth, delimiter=',')
            retransmission_file    = csv.reader(csvfile_retransmission, delimiter=',')
            data['bandwidth']      = {}
            data['retransmission'] = {}

            for row in bandwidth_file:
                if not row[0] in data['bandwidth'].keys():
                    data['bandwidth'][row[0]] = []

                data['bandwidth'][row[0]].append(((float(row[1]) / 1000.0), int(row[2])))

            for row in retransmission_file:
                if not row[0] in data['retransmission'].keys():
                    data['retransmission'][row[0]] = []

                data['retransmission'][row[0]].append(float(row[1])/ 1000.0)

    pp.pprint(data)

    for key in data['bandwidth'].keys():
        if key in data['retransmission'].keys():
            time       = [int(value[0])   for value in data['bandwidth'][key]]
            bandwidths = [float(value[1]) for value in data['bandwidth'][key]]
            color      = '#%02X%02X%02X' % (rand_color(), rand_color(), rand_color())

            pp.pprint(bandwidths)

            ## Plot charts
            plt.xlabel('time (s)')
            plt.ylabel('bandwidth (msgs/s)')

            plt.plot(time, bandwidths, color=color, alpha=0.8, label='node ' + key)

            for retrans in data['retransmission'][key]:
                plt.axvline(float(retrans), color='red')

            ## Show results
            plt.legend()
            plt.savefig('datasets/results/bandwidth/' + bandwidth_results + '_' + key + '.png')
            plt.clf()
