require "minitest/autorun"
require_relative "mrz"

# Fields of the ICAO Doc 9303 specimen passport line.
LINE = "L898902C36UTO7408122F1204159ZE184226B<<<<<10".freeze

class CheckDigitTest < Minitest::Test
  def test_fields_of_the_icao_specimen
    assert_equal 6, Mrz.check_digit("L898902C3") # document number
    assert_equal 2, Mrz.check_digit("740812") # birth date
    assert_equal 9, Mrz.check_digit("120415") # expiry date
    assert_equal 1, Mrz.check_digit("ZE184226B<<<<<") # personal number
  end

  def test_composite_of_the_icao_specimen
    assert_equal 0, Mrz.check_digit(LINE[0, 10] + LINE[13, 7] + LINE[21, 22])
  end

  def test_fillers_count_as_zero
    assert_equal 0, Mrz.check_digit("<<<<<<")
  end

  def test_one_misread_character_changes_the_digit
    refute_equal 6, Mrz.check_digit("L898902C4")
  end

  def test_refuses_a_character_outside_the_mrz_alphabet
    assert_raises(ArgumentError) { Mrz.check_digit("l898902c3") }
  end
end
