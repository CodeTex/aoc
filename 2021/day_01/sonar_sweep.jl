using Base
using DelimitedFiles
using Pkg
using Printf

Pkg.add("BenchmarkTools")
using BenchmarkTools


INPUT_FP = joinpath(dirname(Base.source_path()), "input.txt")


function read_input(fp::String)::Vector{Int16}
    data = readdlm(fp, '\t', Int16, '\n')
    vec(data)
end


function print_solution(val::Int)
    println("\nSolution:\t", val)
end


function print_solution(val::Int, time::Float64)
    print_solution(val)
    @printf "Duration:\t%.6fs" time
end


function calc_step_one(data::Vector{Int16})::Int
    prev = data[begin]
    inc = 0
    for e in data[begin+1:end]
        if e > prev
            inc += 1
        end
        prev = e
    end
    return inc
end

function calc_step_two(data::Vector{Int16})::Int
    prev = sum(data[1:3])
    inc = 0
    for i in 2:length(data)-2
        new = sum(data[i:i+2])
        if new > prev
            inc += 1
        end
        prev = new
    end
    return inc
end


function main()

    data = read_input(INPUT_FP)

    val = calc_step_one(data)

    t = @benchmark calc_step_one($data)

    print_solution(val, mean(t).time / 1000000)

    val = calc_step_two(data)

    t = @benchmark calc_step_one($data)

    print_solution(val, mean(t).time / 1000000)

end

main()
