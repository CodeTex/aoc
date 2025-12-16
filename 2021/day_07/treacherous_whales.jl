INPUT_FP = joinpath(dirname(Base.source_path()), "input.txt")

function frequency_search(positions::Vector{Int})::Tuple{Int, Int}
  freq::Dict{Int, Int} = Dict()
  for pos in positions
    freq[pos] = get(freq, pos, 0) + 1
  end

  min_pos, max_pos = extrema(positions)
  min_cost::Int = typemax(Int)
  best_pos::Int = min_pos

  for target in min_pos:max_pos
    cost::Int = sum(count * abs(pos - target) for (pos, count) in freq)
    if cost < min_cost
      min_cost = cost
      best_pos = target
    end
  end

  return (min_cost, best_pos)
end

function brute_force_search(positions::Vector{Int})::Tuple{Int, Int}
  min_pos, max_pos = extrema(positions)
  min_cost::Int = typemax(Int)
  best_pos::Int = min_pos  

  for target in min_pos:max_pos
    cost::Int = sum(abs(pos - target) for pos in positions)
    if cost < min_cost
      min_cost = cost
      best_pos = target
    end
  end

  return (min_cost, best_pos)
end

function solve_part2(positions::Vector{Int})::Tuple{Int, Int}
  min_pos, max_pos = extrema(positions)
  min_cost::Int = typemax(Int)
  best_post::Int = min_pos

  for target in min_pos:max_pos
    cost::Int = 0
    for pos in positions
      dist::Int = abs(pos - target)
      cost += div(dist * (dist + 1), 2)
    end
    if cost < min_cost
      min_cost = cost
      best_pos = target
    end
  end

  return (min_cost, min_pos)
end

function main()
  line = readline(INPUT_FP)
  positions = parse.(Int, split(strip(line), ","))

  # Warm up JIT
  brute_force_search(positions)
  frequency_search(positions)

  # Measure brute force approach
  t1::Float64 = @elapsed cost1, pos1 = brute_force_search(positions)

  # Measure frequency approach
  t2::Float64 = @elapsed cost2, pos2 = frequency_search(positions)

  println("Brute force: cost=$cost1, position=$pos1, time=$(round(t1 * 1000, digits=3)) ms")
  println("Frequency: cost=$cost2, position=$pos2, time=$(round(t2 * 1000, digits=3)) ms")
  println("Speedup: $(round(t2 / t1, digits=2))x")

  cost_p2, pos_p2 = solve_part2(positions)
  println("\nPart 2: cost=$cost_p2, position=$pos_p2")
end

main()
