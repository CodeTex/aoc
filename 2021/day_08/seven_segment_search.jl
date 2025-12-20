INPUT_FP = joinpath(dirname(Base.source_path()), "input.txt")

function read_input(fp::String)::Vector{String}
  return readlines(fp)
end

function parse_line(line::String)::Tuple{Vector{String}, Vector{String}}
  parts = split(line, "|")
  
  # Get signal patterns (left side) and output values (right side)
  patterns = split(strip(parts[1]))
  outputs = split(strip(parts[2]))
  
  return (patterns, outputs)
end

function count_easy_digits(output_values::Vector{<:AbstractString})::Int
  # Digits 1, 4, 7, 8 use unique segment counts: 2, 3, 4, 7
  return count(val -> length(val) in (2, 3, 4, 7), output_values)
end

function decode_patterns(patterns::Vector{<:AbstractString})::Dict{Set{Char}, Int}
  # Convert patterns to sets for easy set operations
  pattern_sets = Set.(collect.(patterns))
  
  # Find the easy digits first (by unique length)
  one = only(filter(p -> length(p) == 2, pattern_sets))
  four = only(filter(p -> length(p) == 4, pattern_sets))
  seven = only(filter(p -> length(p) == 3, pattern_sets))
  eight = only(filter(p -> length(p) == 7, pattern_sets))
  
  # Find digits with 5 segments (2, 3, 5)
  len5 = filter(p -> length(p) == 5, pattern_sets)
  three = only(filter(p -> issubset(one, p), len5))  # 3 contains all of 1
  five = only(filter(p -> p != three && length(intersect(p, four)) == 3, len5))  # 5 shares 3 with 4
  two = only(filter(p -> p != three && p != five, len5))  # 2 is the remaining one
  
  # Find digits with 6 segments (0, 6, 9)
  len6 = filter(p -> length(p) == 6, pattern_sets)
  nine = only(filter(p -> issubset(four, p), len6))  # 9 contains all of 4
  zero = only(filter(p -> p != nine && issubset(one, p), len6))  # 0 contains all of 1
  six = only(filter(p -> p != nine && p != zero, len6))  # 6 is the remaining one
  
  # Build mapping from pattern to digit
  return Dict(
    zero => 0, one => 1, two => 2, three => 3, four => 4,
    five => 5, six => 6, seven => 7, eight => 8, nine => 9
  )
end

function decode_output(outputs::Vector{<:AbstractString}, mapping::Dict{Set{Char}, Int})::Int
  digits = [mapping[Set(collect(output))] for output in outputs]
  
  # Combine digits into a number: [5, 3, 5, 3] -> 5353
  return sum(digit * 10^(length(digits) - i) for (i, digit) in enumerate(digits))
end

function solve_line(line::String)::Int
  patterns, outputs = parse_line(line)
  mapping = decode_patterns(patterns)
  return decode_output(outputs, mapping)
end

function main()
  lines = read_input(INPUT_FP)
  
  # Part 1
  total_part1 = sum(lines) do line
    _, outputs = parse_line(line)
    count_easy_digits(outputs)
  end
  println("Part 1: Total count of digits 1, 4, 7, 8: $total_part1")
  
  # Part 2
  total_part2 = sum(solve_line, lines)
  println("Part 2: Sum of all output values: $total_part2")
end

main()
