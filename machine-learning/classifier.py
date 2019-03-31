import glob, re
import pandas as pd
import numpy as np

from random import shuffle
from sklearn.model_selection import train_test_split
from sklearn.neighbors import KNeighborsClassifier
from sklearn.ensemble import RandomForestClassifier
from sklearn.metrics import accuracy_score
from sklearn.svm import LinearSVC

DIR               = "datasets/results/csv-datasets/"
FILE_EXTENSION    = "*.csv"
MUTATION_TO_TRAIN = ['ring', 'early']
MUTATION_CODE = {
    'early': 1,
    'ring': 2,
    'centralized': 3,
    'gossip': 4
}

# Loading data
print('> Loading data...')
all_files = glob.glob(DIR + FILE_EXTENSION)

datasets = []
labels = []

for i in range(1,5):
    shuffle(all_files)

for file in all_files:
    dataset = pd.read_csv(file, usecols=list(range(100)), header=None)
    mutation_type = re.findall(r'([^/-]+?)_\d', file)[0]
    max_value = dataset.max().max()

    if mutation_type in MUTATION_TO_TRAIN and max_value != 0:
        dataset = dataset.apply(lambda value: value / max_value)
        datasets.append(dataset)
        labels.append(MUTATION_CODE[mutation_type])

# Training
print('> Training...')
# clf = KNeighborsClassifier()
clf = RandomForestClassifier(n_estimators=100)
# clf = LinearSVC()

train_X, val_X, train_y, val_y = train_test_split(datasets, labels, random_state=1)

clf.fit(train_X, train_y)
print('> Testing...')

# Test on the next 100 images:
print("> Compute predictions")
predicted = clf.predict(val_X)

print(1 in predicted)
print(2 in predicted)

print("> Accuracy: ", accuracy_score(val_y, predicted))
