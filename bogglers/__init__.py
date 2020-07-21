import os

file_dir = os.path.dirname(os.path.realpath(__file__))
data_dir = os.path.realpath(os.path.join(file_dir, "..", "data"))
dictionary_filepath = os.path.join(data_dir, "scrabble_words.txt")
