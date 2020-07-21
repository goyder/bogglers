import unittest
import bogglers.data as data


class TestNetworkGeneration(unittest.TestCase):
    """
    Tests associated with building the network structure.
    """

    def test_4x4_generated_networks_have_expected_names(self):
        network = data.generate_network(4, 4)

        self.assertIn(
            "A1",
            network.keys(),
            "Could not find A1 tile."
        )

        self.assertIn(
            "D4",
            network.keys(),
            "Could not find A1 tile."
        )

    def test_7x7_generated_networks_have_expected_names(self):
        network = data.generate_network(7, 7)

        self.assertIn(
            "A1",
            network.keys(),
            "Could not find A1 tile."
        )

        self.assertIn(
            "G7",
            network.keys(),
            "Could not find A1 tile."
        )

    def test_topleft_corner_tile_has_correct_vertices(self):
        network = data.generate_network(4, 4)
        tile = "A1"
        expected_vertices = {"B1", "B2", "A2"}

        self.assertEqual(
            expected_vertices,
            set(network[tile]),
            "Vertices did not match."
        )

    def test_bottom_right_corner_tile_has_correct_vertices(self):
        network = data.generate_network(4, 4)
        tile = "D4"
        expected_vertices = {"C3", "C4", "D3"}

        self.assertEqual(
            expected_vertices,
            set(network[tile]),
            "Vertices did not match."
        )

    def test_centre_tile_has_correct_vertices(self):
        network = data.generate_network(4, 4)
        tile = "B2"
        expected_vertices = {"A1", "B1", "C1", "C2", "C3", "B3", "A3", "A2"}

        self.assertEqual(
            expected_vertices,
            set(network[tile]),
            "Vertices did not match."
        )


class TestLetterSelection(unittest.TestCase):
    """
    Tests associated with "rolling" or "shaking" the Boggle board.
    """

    def test_generate_letter_selection_and_get_right_number_of_letters(self):
        """
        Roll the actual Boggle dice selection and get the letters back.
        :return:
        """
        letters = data.generate_letters(data.dice)

        self.assertEqual(
            16,
            len(letters),
            "Did not get expected number of letters back."
        )


class TestNetworkMapping(unittest.TestCase):
    """
    Tests associated with mapping letters to a network.
    """

    def test_mapping_to_network_returns_right_letter(self):
        network = {"A1": [], "B1": [], "C1": []}
        letters = ["A", "B", "C"]

        network_mapping = data.assign_letters(network, letters)

        self.assertEqual(
            "A",
            network_mapping["A1"],
            "Did not have mapping expected."
        )
        self.assertEqual(
            "B",
            network_mapping["B1"],
            "Did not have mapping expected."
        )
        self.assertEqual(
            "C",
            network_mapping["C1"],
            "Did not have mapping expected."
        )


class TestDictionary(unittest.TestCase):
    """
    Tests associated with loading a dictionary.
    """

    def test_open_dictionary(self):
        data.generate_dictionary()

    def test_find_cat_in_dictionary(self):
        dictionary = data.generate_dictionary()
        self.assertIn(
            "CAT",
            dictionary,
            "Could not find word 'CAT'."
        )


if __name__ == '__main__':
    unittest.main()
