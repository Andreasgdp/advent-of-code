require 'rspec/autorun'
require_relative 'main_1'

RSpec.describe 'find_digits' do
  it 'returns the correct digits for various inputs' do
    expect(find_digits('5eighthree5')).to eq([5, 8, 3, 5])
    expect(find_digits('eighthree')).to eq([8, 3])
    expect(find_digits('sevenine')).to eq([7, 9])
    expect(find_digits('two1nine')).to eq([2, 1, 9])
    expect(find_digits('eightwothree')).to eq([8, 2, 3])
    expect(find_digits('abcone2threexyz')).to eq([1, 2, 3])
    expect(find_digits('xtwone3four')).to eq([2, 1, 3, 4])
    expect(find_digits('4nineeightseven2')).to eq([4, 9, 8, 7, 2])
    expect(find_digits('zoneight234')).to eq([1, 8, 2, 3, 4])
    expect(find_digits('7pqrstsixteen')).to eq([7, 6])
  end
end
