import unittest

from mrz import check_digit

# Fields of the ICAO Doc 9303 specimen passport line:
# L898902C36UTO7408122F1204159ZE184226B<<<<<10
LINE = "L898902C36UTO7408122F1204159ZE184226B<<<<<10"


class CheckDigitTest(unittest.TestCase):
    def test_fields_of_the_icao_specimen(self):
        self.assertEqual(check_digit("L898902C3"), 6)  # document number
        self.assertEqual(check_digit("740812"), 2)  # birth date
        self.assertEqual(check_digit("120415"), 9)  # expiry date
        self.assertEqual(check_digit("ZE184226B<<<<<"), 1)  # personal number

    def test_composite_of_the_icao_specimen(self):
        self.assertEqual(check_digit(LINE[0:10] + LINE[13:20] + LINE[21:43]), 0)

    def test_fillers_count_as_zero(self):
        self.assertEqual(check_digit("<<<<<<"), 0)

    def test_one_misread_character_changes_the_digit(self):
        self.assertNotEqual(check_digit("L898902C4"), 6)

    def test_refuses_a_character_outside_the_mrz_alphabet(self):
        with self.assertRaises(ValueError):
            check_digit("l898902c3")


if __name__ == "__main__":
    unittest.main()
