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


if __name__ == '__main__':
    unittest.main()
