from bogglers import data, search
from pprint import pprint

board = data.StandardBoggle()
print("Boggle board to search:")

print("============")
print(board)
print("============")
print("Searching...")
words = search.generate_words(board.network,
                              board.mapping,
                              max_letters=4)

print("Search complete.")
print("Found {} words.".format(len(words)))
pprint(words)

