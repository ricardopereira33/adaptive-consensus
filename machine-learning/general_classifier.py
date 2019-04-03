import pdb
import pandas as pd
import numpy  as np

from sklearn.model_selection import train_test_split
from sklearn.ensemble        import RandomForestClassifier
from sklearn.metrics         import accuracy_score

DIR      = "../results/csv-datasets/"
FILENAME = "global_results.csv"

def convert_mutation_to_code(mutation):
    mutation_codes = {
        'early':       1,
        'ring':        2,
        'centralized': 3,
        'gossip':      4
    }
    return mutation_codes.get(mutation.strip(), "Invalid mutation")

def convert_bool_to_binary(value):
    bool_binary_value = {
        'true':  1,
        'false': 0
    }
    return bool_binary_value.get(value.strip(), "Invalid bool")

print('> Loading data')
data     = pd.read_csv(DIR + FILENAME)
features = ['nodes', 'defaultDelta', 'maxTries', 'percentageMiss', 'faulty', 'resultTime']

data.faulty = data.faulty.apply(convert_bool_to_binary)

target_column = data.mutation
input_columns = data[features]

pdb.set_trace()

print('> Training')
clf = RandomForestClassifier(n_estimators=100)

train_X, val_X, train_y, val_y = train_test_split(input_columns, target_column, random_state=1)

clf.fit(train_X, train_y)

print("> Compute predictions")
predicted = clf.predict(val_X)

print("> Accuracy: ", accuracy_score(val_y, predicted))
