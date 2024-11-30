import os


def create_day_folder(day_number):
    """Creates a folder named day-<number> with the specified content."""

    folder_name = f"day-{day_number:02d}"
    os.makedirs(folder_name, exist_ok=True)

    # Create day-<number>.rb
    rb_file_path = os.path.join(folder_name, f"{folder_name}.rb")
    with open(rb_file_path, "w") as rb_file:
        rb_file.write(
            """# frozen_string_literal: true

input = File.open(File.join(__dir__, 'input.txt')).readlines.map(&:to_s)

def solve_part_1(input)
    # Implement your solution here
    42
end

def solve_part_2(input)
    # Implement your solution here
    42
end


if __FILE__ == $PROGRAM_NAME
  if input.empty?
    puts 'No input found'
    exit 1
  end
  puts "Part 1: #{solve_part_1(input)}"
  puts "Part 2: #{solve_part_2(input)}"
end
"""
        )

    # Create day-<number>_spec.rb
    spec_file_path = os.path.join(folder_name, f"{folder_name}_spec.rb")
    with open(spec_file_path, "w") as spec_file:
        spec_file.write(
            f"""# frozen_string_literal: true

require 'rspec/autorun'
require_relative '{folder_name}'

RSpec.describe 'Day {day_number}' do
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
"""
        )

    # Create an empty input.txt
    input_file_path = os.path.join(folder_name, "input.txt")
    open(input_file_path, "w").close()


if __name__ == "__main__":
    for i in range(1, 26):
        create_day_folder(i)
