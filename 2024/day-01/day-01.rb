# frozen_string_literal: true

input = File.open(File.join(__dir__, 'input.txt')).readlines.map(&:to_s)

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

# assume the lists are the same length
def total_distance(list1, list2)
  total = 0
  list1.each_with_index do |value, index|
    total += (value - list2[index]).abs
  end
  total
end

# assume the lists are the same length and sorted
def simularity_score(list1, list2)
  score = 0
  list1.each do |value|
    count = 0
    list2.each do |value2|
      break if value2 > value

      count += 1 if value == value2
    end
    score += value * count
  end
  score
end

def solve_part1(input)
  list1, list2 = get_2_lists_from_input(input)
  list1 = sort_list(list1)
  list2 = sort_list(list2)
  total_distance(list1, list2)
end

def solve_part2(input)
  list1, list2 = get_2_lists_from_input(input)
  list1 = sort_list(list1)
  list2 = sort_list(list2)
  simularity_score(list1, list2)
end

if __FILE__ == $PROGRAM_NAME
  if input.empty?
    puts 'No input found'
    exit 1
  end
  puts "Part 1: #{solve_part1(input)}"
  puts "Part 2: #{solve_part2(input)}"
end
