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
    # horizontal: x coords are const
    return [(x1, y) for y in min(y1, y2):max(y1, y2)]
  elseif y1 == y2
    # horizontal: y coords are const
    return [(x, y1) for x in min(x1, x2):max(x1, x2)]
  else
    # diagonal: both x and y change
    dx = sign(x2 - x1)
    dy = sign(y2 - y1)
    steps = abs(x2 - x1)

    return [(x1 + dx*i, y1 + dy*i) for i in 0:steps]
  end
end

function count_overlaps(segments::Vector{Tuple{Int, Int, Int, Int}}, include_diag::Bool)::Int
  if include_diag
    filtered = segments
  else
    filtered = filter(seg -> seg[1] == seg[3] || seg[2] == seg[4], segments)
  end

  # Track overlaps in segment lines
  overlaps = Dict{Tuple{Int, Int}, Int}()
  for segment in filtered
    points = get_points_on_segment(segment)
    for point in points
      overlaps[point] = get(overlaps, point, 0) + 1
    end
  end

  println("Total coordinates: $(length(overlaps))")

  # Count coords with 2+ overlaps
  return sum(1 for count in values(overlaps) if count >= 2)
end

function main()
  data = read_input(INPUT_FP)
  println("Parsed $(length(data)) line segments\n")

  straight_overlaps = count_overlaps(data, false)
  println("Solution #1: $straight_overlaps")

  all_overlaps = count_overlaps(data, true)
  println("Solution #2: $all_overlaps")
end

main()
