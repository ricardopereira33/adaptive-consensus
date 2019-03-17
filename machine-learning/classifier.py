import glob, re
import pandas as pd
import numpy as np
from sklearn.neighbors import KNeighborsClassifier
from sklearn.ensemble import RandomForestClassifier
from sklearn.metrics import accuracy_score
from sklearn.svm import LinearSVC

DIR = "../results/csv-datasets/"
FILE_EXTENSION = "*.csv"

# Loading data
print('> Loading data...')
all_files = glob.glob(DIR + FILE_EXTENSION)

datasets = []
labels = []

for file in all_files:
    dataset = pd.read_csv(file, usecols=list(range(100)), header=None)
    mutation_type = re.findall(r'([^/-]+?)_\d', file)[0]

    datasets.append(dataset)
    labels.append(mutation_type)

# Training
print('> Training...')
# clf = KNeighborsClassifier()
# clf = RandomForestClassifier(n_estimators=100)
clf = LinearSVC()

train_x = datasets[:250]
train_y = np.asarray(labels[:250])

clf.fit(train_x, train_y)

# Test on the next 100 images:
print('> Testing...')
test_x = datasets[250:300]
expected = labels[250:300]

print("> Compute predictions")
predicted = clf.predict(test_x)

print("> Accuracy: ", accuracy_score(expected, predicted))