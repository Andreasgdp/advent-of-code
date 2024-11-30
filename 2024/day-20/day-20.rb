# frozen_string_literal: true

input = File.open(File.join(__dir__, 'input.txt')).readlines.map(&:to_s)

def solve_part1(input)
    # Implement your solution here
    42
end

def solve_part2(input)
    # Implement your solution here
    42
end


if __FILE__ == $PROGRAM_NAME
  if input.empty?
    puts 'No input found'
    exit 1
  end
  puts "Part 1: #{solve_part1(input)}"
  puts "Part 2: #{solve_part2(input)}"
end
