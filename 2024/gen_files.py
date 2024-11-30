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
  let(:sample_input1) do
    File.open(File.join(__dir__, 'sample_input1.txt')).readlines.map(&:to_s)
  end

  let(:sample_input2) do
    File.open(File.join(__dir__, 'sample_input2.txt')).readlines.map(&:to_s)
  end

  describe '#solve_part_1' do
    it 'returns the correct answer for the sample input' do
      expect(solve_part1(sample_input1)).to eq(42)
    end
  end

  describe '#solve_part_2' do
    it 'returns the correct answer for the sample input' do
      expect(solve_part2(sample_input1)).to eq(42)
    end
  end

  # ...
end
"""
        )

    # Create an empty input.txt
    input_file_path = os.path.join(folder_name, "input.txt")
    open(input_file_path, "w").close()
    # Create empty sample_input1.txt and sample_input2.txt files
    open(os.path.join(folder_name, "sample_input1.txt"), "w").close()
    open(os.path.join(folder_name, "sample_input2.txt"), "w").close()


if __name__ == "__main__":
    for i in range(1, 26):
        create_day_folder(i)
