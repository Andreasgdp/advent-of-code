# frozen_string_literal: true

require 'rspec/autorun'
require_relative 'day-11'

RSpec.describe 'Day 11' do
  let(:sample_input) do
    <<~INPUT
      1abc2
      pqr3stu8vwx
      a1b2c3d4e5f
      treb7uchet
    INPUT
  end

  describe '#solve_part_1' do
    it 'returns the correct answer for the sample input' do
      expect(solve_part_1(sample_input)).to eq(42)
    end
  end

  describe '#solve_part_2' do
    it 'returns the correct answer for the sample input' do
      expect(solve_part_2(sample_input)).to eq(42)
    end
  end

  # ...
end
