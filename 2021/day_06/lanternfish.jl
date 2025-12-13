using Base

INPUT_FP = joinpath(dirname(Base.source_path()), "input.txt")

function read_input(fp::String)::Vector{Int}
  line = readline(fp)
  return parse.(Int, split(strip(line), ","))
end

function main()
  data = read_input(INPUT_FP)
  println("Length: $(length(data))")
end

main()
