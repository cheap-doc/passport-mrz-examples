# MRZ check digits as ICAO Doc 9303 defines them. No dependencies.
module Mrz
  WEIGHTS = [7, 3, 1].freeze

  # The value of one MRZ character: 0-9 as themselves, A-Z as 10-35, "<" as 0.
  def self.char_value(c)
    case c
    when "0".."9" then c.ord - "0".ord
    when "A".."Z" then c.ord - "A".ord + 10
    when "<" then 0
    else raise ArgumentError, "not an MRZ character: #{c.inspect}"
    end
  end

  # The check digit of a field: each value times 7, 3, 1 repeating, summed, mod 10.
  def self.check_digit(data)
    data.each_char.with_index.sum { |c, i| char_value(c) * WEIGHTS[i % 3] } % 10
  end
end
