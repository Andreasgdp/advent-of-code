# test the find_digits method

require 'minitest/autorun'

require_relative 'main_1'

class MainTest < Minitest::Test
  def test_find_digits
    assert_equal [5, 8, 3, 5], find_digits('5eighthree5')
    assert_equal [8, 3], find_digits('eighthree')
    assert_equal [7, 9], find_digits('sevenine')
    assert_equal [2, 1, 9], find_digits('two1nine')
    assert_equal [8, 2, 3], find_digits('eightwothree')
    assert_equal [1, 2, 3], find_digits('abcone2threexyz')
    assert_equal [2, 1, 3, 4], find_digits('xtwone3four')
    assert_equal [4, 9, 8, 7, 2], find_digits('4nineeightseven2')
    assert_equal [1, 8, 2, 3, 4], find_digits('zoneight234')
    assert_equal [7, 6], find_digits('7pqrstsixteen')
  end
end
