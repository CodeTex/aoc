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

function main()
  line = readline(INPUT_FP)
  positions = parse.(Int, split(strip(line), ","))

  cost, pos = frequency_search(positions)
  println("Minimum fuel required: $cost @ $pos")
end

main()
