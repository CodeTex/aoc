using Base


INPUT_FP = joinpath(dirname(Base.source_path()), "input.txt")


function step_one(fp::String)::Int
    position = 0
    depth = 0
    for line in eachline(fp)
        cmd, val = split(line)
        val = parse(Int, val)
        if cmd == "forward"
            position += val
        elseif cmd == "down"
            depth += val
        elseif cmd == "up"
            depth -= val
        end
    end
    return position * depth
end

function step_two(fp::String)::Int
    position = 0
    depth = 0
    aim = 0
    for line in eachline(fp)
        cmd, val = split(line)
        val = parse(Int, val)
        if cmd == "forward"
            position += val
            depth += aim * val
        elseif cmd == "down"
            aim += val
        elseif cmd == "up"
            aim -= val
        end
    end
    return position * depth
end

function main()

    sol1 = step_one(INPUT_FP)
    println("\nSolution Step 1:\t", sol1)

    sol2 = step_two(INPUT_FP)
    println("Solution Step 2:\t", sol2)

end

main()
