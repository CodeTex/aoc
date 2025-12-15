INPUT_FP = joinpath(dirname(Base.source_path()), "input.txt")

function read_input(fp::String)::Vector{Int}
  line = readline(fp)
  return parse.(Int, split(strip(line), ","))
end

function init_fish_dict(timers::Vector{Int})::Dict{Int, Int}
  fish_counts = Dict{Int, Int}()
  for timer in timers
    fish_counts[timer] = get(fish_counts, timer, 0) + 1
  end
  return fish_counts
end

function main()
  data = read_input(INPUT_FP)
  counts = init_fish_dict(data)
end

main()
