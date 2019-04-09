import pdb
import pandas as pd
import numpy  as np

from sklearn.metrics         import mean_absolute_error
from sklearn.model_selection import train_test_split
from sklearn.ensemble        import RandomForestRegressor
from sklearn.metrics         import accuracy_score

DIR      = "../results/"
FILENAME = "global_results.csv"

def convert_mutation_to_code(mutation):
    mutation_codes = {
        'early':       1,
        'ring':        2,
        'centralized': 3,
        'gossip':      4
    }
    return mutation_codes.get(mutation.strip(), "Invalid mutation")

def get_accuracy(max_leaf_nodes, train_X, val_X, train_y, val_y):
    model = RandomForestRegressor(max_leaf_nodes=max_leaf_nodes, random_state=1, n_estimators=100)
    model.fit(train_X, train_y)
    predicted = model.predict(val_X)
    val_mae = mean_absolute_error(predicted, val_y)
    return val_mae

##########
# SCRIPT #
##########

print('> Loading data')
data     = pd.read_csv(DIR + FILENAME)
features = ['nodes', 'defaultDelta', 'maxTries', 'percentageMiss', 'percentageFaults', 'probabilityToFail', 'latency', 'bandwidth', 'mutation']

data.mutation = data.mutation.apply(convert_mutation_to_code)

target_column = data.resultTime
input_columns = data[features]

pdb.set_trace()

print("> Find the best value for max_leaf_nodes")
train_X, val_X, train_y, val_y = train_test_split(input_columns, target_column, random_state=1)

min_error = 0
max_leaf_nodes = 0

for nodes in range(10, 5000, 20):
    val_mae = get_accuracy(nodes, train_X, val_X, train_y, val_y)

    if val_mae < min_error:
        min_error = val_mae
        max_leaf_nodes = nodes

print("Min error: %f | Leaf nodes %d" %(min_error, max_leaf_nodes))

print('> Training')
model = RandomForestRegressor(max_leaf_nodes=max_leaf_nodes, random_state=1, n_estimators=100)

train_X, val_X, train_y, val_y = train_test_split(input_columns, target_column, random_state=1)

model.fit(train_X, train_y)

print("> Compute predictions")
predicted = model.predict(val_X)

val_mae = mean_absolute_error(predicted, val_y)
print("Validation MAE for best value of max_leaf_nodes: {:,.0f}".format(val_mae))
