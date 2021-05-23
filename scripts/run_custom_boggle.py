from bogglers import data, search
from pprint import pprint
import time


print("Searching...")
t = time.time()
network = data.generate_network(4, 4)
letters = list("LTWNHWCCRJIYRFIE")
mapping = data.assign_letters(network, letters)

words = search.generate_words(
    network,
    mapping,
    max_letters=4
)
elapsed = time.time() - t

print("Search complete.")
print("Time elapsed: {}".format(elapsed))
print("Found {} words.".format(len(words)))
pprint(words)

