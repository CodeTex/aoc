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

function main()
  lines = read_input(INPUT_FP)
  
  println("Total lines read: $(length(lines))")
  println("\nExample of first 3 lines:")
  
  for i in 1:3
    output_values = parse_line(lines[i])
    println("\nLine $i:")
    println("  Output values: $output_values")
    println("  Lengths: $(length.(output_values))")
  end
end

main()
