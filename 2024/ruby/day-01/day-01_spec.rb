# frozen_string_literal: true

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
      expect(solve_part1(sample_input1)).to eq(11)
    end
  end

  describe '#solve_part_2' do
    it 'returns the correct answer for the sample input' do
      expect(solve_part2(sample_input2)).to eq(31)
    end
  end

  # ...
end
