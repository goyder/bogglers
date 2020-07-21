import unittest
import bogglers.data as data
import bogglers.search as search


class TestSearch(unittest.TestCase):
    """
    Tests associated with searching networks for words.
    """

    def test_find_expected_words_in_3x3_map(self):
        """
        Test an example 3x3 matrix.
        :return:
        """
        # Generate an example network.
        letters = list("COTASTTIN")
        network = data.generate_network(3, 3)
        mapping = data.assign_letters(network, letters)
        expected_words = ["CAT", "COAT", "STINT", "COT", "COST"]

        words = search.generate_words(
            network,
            mapping
        )

        for expected_word in expected_words:
            self.assertIn(
                expected_word,
                words,
                "Could not find word in generated list."
            )

    def test_find_words_in_4x4_map(self):
        board = data.StandardBoggle()
        print(board)

        words = search.generate_words(board.network,
                                      board.mapping,
                                      max_letters=4)
        print(words)


if __name__ == '__main__':
    unittest.main()
