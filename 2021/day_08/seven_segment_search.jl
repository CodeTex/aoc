INPUT_FP = joinpath(dirname(Base.source_path()), "input.txt")

function read_input(fp::String)::Vector{String}
  return readlines(fp)
end

function main()
  lines = read_input(INPUT_FP)
end

main()
