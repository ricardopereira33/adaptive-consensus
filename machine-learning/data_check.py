import csv
import numpy             as np
import matplotlib.pyplot as plt

result_type = 'ring_mutations'

data = {}
i = 1
line = 1
with open('datasets/results/global_results_' + result_type + '.csv') as csvfile:
    readCSV = csv.reader(csvfile, delimiter=',')
    next(readCSV, None)

    for row in readCSV:
        if i % 2 == 0:
            if row[8] != 'old_ring':
                print("OLD RING TIMEOUT - " + str(line))
            else:
                i += 1
        else:
            if row[8] != 'ring':
                print("RING TIMEOUT - " + str(line))
            else:
                i += 1
        line += 1

