# frozen_string_literal: false

input = File.open('input.txt').readlines.map(&:to_s)

def find_digits(line)
  digits = []
  i = 0
  while i < line.length
    if line[i].match?(/\d/)
      digits << line[i].to_i
    elsif line[i..].start_with?('one')
      digits << 1
    elsif line[i..].start_with?('two')
      digits << 2
    elsif line[i..].start_with?('three')
      digits << 3
    elsif line[i..].start_with?('four')
      digits << 4
    elsif line[i..].start_with?('five')
      digits << 5
    elsif line[i..].start_with?('six')
      digits << 6
    elsif line[i..].start_with?('seven')
      digits << 7
    elsif line[i..].start_with?('eight')
      digits << 8
    elsif line[i..].start_with?('nine')
      digits << 9
    end
    i += 1
  end
  digits
end

total = 0

input.each do |line|
  digits = find_digits(line)
  next if digits.empty?

  first_digit = digits.first
  last_digit = digits.last
  total += "#{first_digit}#{last_digit}".to_i
end
puts total
