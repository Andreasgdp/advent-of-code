# frozen_string_literal: true
# typed: true

require 'sorbet-runtime'
extend T::Sig

input = File.open(File.join(__dir__, 'input.txt')).readlines.map(&:to_s)

sig { params(input: T::Array[String]).returns(T::Array[T::Array[Integer]]) }
def get_reports(input)
  reports = T.let([], T::Array[T::Array[Integer]])
  input.each do |line|
    report = line.split(' ').map(&:to_i)
    reports << report
  end
  reports
end

# the report is safe if levels are all increasing or all decreasing
# and any two adjacent levels differ by at least one and at most 3
sig { params(report: T::Array[Integer]).returns(T::Boolean) }
def is_safe_report(report)
  is_increasing = T.let(false, T::Boolean)
  is_decreasing = T.let(false, T::Boolean)
  report.each_with_index do |curr_level, index|
    next if index == 0

    return false if (T.must(report[index - 1]) - curr_level).abs > 3

    return false if T.must(report[index - 1]) == curr_level

    if T.must(report[index - 1]) < curr_level
      is_increasing = true
    elsif T.must(report[index - 1]) > curr_level
      is_decreasing = true
    end
  end

  return false if is_increasing && is_decreasing || !is_increasing && !is_decreasing

  true
end

sig { params(reports: T::Array[T::Array[Integer]], dampener: T::Boolean).returns(Integer) }
def count_safe_reports(reports, dampener = false)
  count_safe_reports = T.let(0, Integer)
  reports.each do |report|
    if is_safe_report(report)
      count_safe_reports += 1
    elsif dampener
      report.each_with_index do |level, index|
        new_report = report.dup
        new_report.delete_at(index)
        if is_safe_report(new_report)
          count_safe_reports += 1
          break
        end
      end
    end
  end
  count_safe_reports
end

sig { params(input: T::Array[String]).returns(Integer) }
def solve_part1(input)
  reports = get_reports(input)

  count_safe_reports(reports)
end

sig { params(input: T::Array[String]).returns(Integer) }
def solve_part2(input)
  reports = get_reports(input)

  count_safe_reports(reports, true)
end

if __FILE__ == $PROGRAM_NAME
  if input.empty?
    puts 'No input found'
    exit 1
  end
  puts "Part 1: #{solve_part1(input)}"
  puts "Part 2: #{solve_part2(input)}"
end
