import csv, random, pprint, os
import numpy             as np
import matplotlib.pyplot as plt

pp   = pprint.PrettyPrinter(indent=2)
data = {}

os.system('rm ../message_exchange.png')

print('Load data')

with open('../messages.csv') as csvfile:
    bandwidth_file = csv.reader(csvfile, delimiter=',')
    data['time'] = []
    data['messages'] = []

    for row in bandwidth_file:
        data['time'].append(int(row[0]))
        data['messages'].append(int(row[1]))

print('Plot results')
print(len(data['time']))
print(len(data['messages']))

## Plot charts
plt.xlabel('time (ms)')
plt.ylabel('# messages')

plt.plot(data['time'], data['messages'], color='blue', alpha=0.8, label='messages')

print('Plot is finish.')

## Show results
plt.legend()
plt.savefig('message_exchange2.png')
# plt.show()