# frozen_string_literal: true

input = File.open(File.join(__dir__, 'input.txt')).readlines.map(&:to_s)
# input = File.open(File.join(__dir__, 'sample_input1.txt')).readlines.map(&:to_s)

def get_2_lists_from_input(input)
  list1 = []
  list2 = []

  input.each do |line|
    list1 << line.split(' ')[0].to_i
    list2 << line.split(' ')[1].to_i
  end

  [list1, list2]
end

def sort_list(list)
  list.sort
end

def abs_diff(a, b)
  (a - b).abs
end

def total_distance(list1, list2)
  total = 0
  list1.each_with_index do |value, index|
    total += abs_diff(value, list2[index])
  end
  total
end

def solve_part1(input)
  list1, list2 = get_2_lists_from_input(input)
  list1 = sort_list(list1)
  list2 = sort_list(list2)
  total_distance(list1, list2)
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
