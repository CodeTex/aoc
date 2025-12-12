using Base

INPUT_FP = joinpath(dirname(Base.source_path()), "input.txt")

function parse_line(line::String)::Tuple{Int, Int, Int, Int}
  line = strip(line)
  parts = split(line, " -> ")
  coords1 = split(strip(parts[1]), ",")
  coords2 = split(strip(parts[2]), ",")
  
  x1::Int = parse(Int, strip(coords1[1]))
  y1::Int = parse(Int, strip(coords1[2]))
  x2::Int = parse(Int, strip(coords2[1]))
  y2::Int = parse(Int, strip(coords2[2]))
  
  return (x1, y1, x2, y2)
end

function read_input(filename::String)::Vector{Tuple{Int, Int, Int, Int}}
  lines::Vector{String} = readlines(filename)
  # Filter out empty lines
  lines = filter(line -> !isempty(strip(line)), lines)
  segments = parse_line.(lines)
  return segments
end

function main()
  data = read_input(INPUT_FP)
  println("Parsed $(length(data)) line segments\n")

  filtered = filter(seg -> seg[1] == seg[3] || seg[2] == seg[4])
  # println("Filtered out $(length(data)-length(filtered)) line segments\n")  
end

main()
