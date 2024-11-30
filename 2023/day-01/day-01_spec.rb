require 'rspec/autorun'
require_relative 'day-01'

RSpec.describe 'Day 1' do
  let(:sample_input1) do
    File.open(File.join(__dir__, 'sample_input1.txt')).readlines.map(&:to_s)
  end

  let(:sample_input2) do
    File.open(File.join(__dir__, 'sample_input2.txt')).readlines.map(&:to_s)
  end

  describe '#solve_part_1' do
    it 'returns the correct answer for the sample input' do
      expect(solve_part_1(sample_input1)).to eq(142)
    end
  end

  describe '#solve_part_2' do
    it 'returns the correct answer for the sample input' do
      expect(solve_part_2(sample_input2)).to eq(281)
    end

    it 'solves corner cases when finding digits' do
      expect(find_digits2('5eighthree5')).to eq([5, 8, 3, 5])
      expect(find_digits2('eighthree')).to eq([8, 3])
      expect(find_digits2('sevenine')).to eq([7, 9])
      expect(find_digits2('two1nine')).to eq([2, 1, 9])
      expect(find_digits2('eightwothree')).to eq([8, 2, 3])
      expect(find_digits2('abcone2threexyz')).to eq([1, 2, 3])
      expect(find_digits2('xtwone3four')).to eq([2, 1, 3, 4])
      expect(find_digits2('4nineeightseven2')).to eq([4, 9, 8, 7, 2])
      expect(find_digits2('zoneight234')).to eq([1, 8, 2, 3, 4])
      expect(find_digits2('7pqrstsixteen')).to eq([7, 6])
    end
  end

  # ...
end
