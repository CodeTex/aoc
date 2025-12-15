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

function update_counter(counter::Dict{Int, Int})::Dict{Int, Int}
  new_counter = Dict{Int, Int}()

  # fish that spawn this day
  spawning = get(counter, 0, 0)

  # shift all other timers down by 1
  for timer in 1:8
    if haskey(counter, timer)
      new_counter[timer-1] = counter[timer]
    end
  end

  # fish that spawned reset to 6 and create new fish at 8
  new_counter[6] = get(new_counter, 6, 0) + spawning
  new_counter[8] = spawning

  return new_counter
end

function main()
  data = read_input(INPUT_FP)
  counter = init_fish_dict(data)
  println("Initial state: $counter")

  x = 80
  for i in 1:x
    counter = update_counter(counter)
  end

  println("End state: $counter")

  total_fish = sum(values(counter))
  println("Total fish after $x days: $total_fish")
end

main()
