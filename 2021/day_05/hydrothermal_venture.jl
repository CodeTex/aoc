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

function get_points_on_segment(segment::Tuple{Int, Int, Int, Int})::Vector{Tuple{Int, Int}}
  x1, y1, x2, y2 = segment
  if x1 == x2
    # x coords are const
    return [(x1, y) for y in min(y1, y2):max(y1, y2)]
  else
    # y coords are const
    return [(x, y1) for x in min(x1, x2):max(x1, x2)]
  end
end

function solve_part1(segments::Vector{Tuple{Int, Int, Int, Int}})::Dict{Tuple{Int, Int}, Int}
  # Only conside horizontal or vertical lines
  filtered = filter(seg -> seg[1] == seg[3] || seg[2] == seg[4], segments)
  println("Filtered out $(length(segments)-length(filtered)) line segments\n")

  overlaps = Dict{Tuple{Int, Int}, Int}()
  for segment in filtered
    points = get_points_on_segment(segment)
    for point in points
      overlaps[point] = get(overlaps, point, 0) + 1
    end
  end

  println("Total coordinates: $(length(overlaps))")

  result = sum(1 for count in values(overlaps) if count >= 2)
  println("Total coordinates with 2+ overlaps: $result")

  return overlaps
end

function main()
  data = read_input(INPUT_FP)
  println("Parsed $(length(data)) line segments\n")

  overlaps = solve_part1(data)

end

main()
