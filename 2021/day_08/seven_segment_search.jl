INPUT_FP = joinpath(dirname(Base.source_path()), "input.txt")

function read_input(fp::String)::Vector{String}
  return readlines(fp)
end

function parse_line(line::String)::Vector{String}
  parts = split(line, "|")
  output_str = strip(parts[2])
  output_values = split(output_str)
  
  return output_values
end

function count_easy_digits(output_values::Vector{<:AbstractString})::Int
  # Digits 1, 4, 7, 8 use unique segment counts: 2, 3, 4, 7
  return count(val -> length(val) in (2, 3, 4, 7), output_values)
end

function main()
  lines = read_input(INPUT_FP)
  
  total = sum(line -> count_easy_digits(parse_line(line)), lines)
  
  println("Part 1: Total count of digits 1, 4, 7, 8 in output values: $total")
end

main()
