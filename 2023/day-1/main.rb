# frozen_string_literal: false

input = File.open('test_input.txt').readlines.map(&:to_s)

DIGIT_WORDS = {
  'one' => 1, 'two' => 2, 'three' => 3, 'four' => 4,
  'five' => 5, 'six' => 6, 'seven' => 7, 'eight' => 8, 'nine' => 9
}

def find_digits(line)
  line.scan(/\d|one|two|three|four|five|six|seven|eight|nine/).map do |match|
    match.match?(/\d/) ? match.to_i : DIGIT_WORDS[match]
  end
end

total = input.sum do |line|
  digits = find_digits(line)
  puts digits
  digits.empty? ? 0 : "#{digits.first}#{digits.last}".to_i
end

puts total
