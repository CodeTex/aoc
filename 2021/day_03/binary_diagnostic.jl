# credit to Eric Burden for showing me a lot new tools in Julia
# code is slightly modified from https://www.ericburden.work/blog/2021/12/03/advent-of-code-2021-day-3/
using Base


INPUT_FP = joinpath(dirname(Base.source_path()), "input.txt")

process_input(s::AbstractString)::Vector{Bool} = split(s, "") .== "1"

function read_input(fp::String)::Matrix{Bool}
    vector = [process_input(l) for l in readlines(fp)]
    matrix = reduce(hcat, vector)
    transpose(matrix)
end

function most_common_bool(arr)::Bool
    trues = count(arr)
    trues >= (length(arr) - trues)
end

function bitvec_to_int(bv::BitVector)::Int
    powers = (length(bv)-1:-1:0) |> collect |> (x -> x[bv])
    sum(2 .^ powers)
end

function calculate_power_consumation(data::Matrix{Bool})::Int
    gamma = eachcol(data) |> (x -> map(most_common_bool, x)) |> BitVector
    epsilon = .!gamma
    bitvec_to_int(gamma) * bitvec_to_int(epsilon)
end

function find_first_match(data::Matrix{Bool}, discriminator)
    mask = trues(size(data, 1))
    for col in eachcol(data)
        common_value = discriminator(col[mask])
        # Carry forward only mask indices where the common value
        # is found in each column
        mask = mask .& (col .== common_value)
        # Stop looking if mask contains only one `true`. 
        sum(mask) == 1 && break
    end
    # Convert n x 1 BitMatrix to BitVector
    data[mask, :] |> Iterators.flatten |> BitVector
end

function calculate_life_support_rating(data::Matrix{Bool})::Int
    ox_gen_rating = find_first_match(data, most_common_bool)
    co2_sc_rating = find_first_match(data, !most_common_bool)
    bitvec_to_int(ox_gen_rating) * bitvec_to_int(co2_sc_rating)
end

function main()
    data = read_input(INPUT_FP)
    val = calculate_power_consumation(data)
    println("\nPower Consumption:\t", val)
    val = calculate_life_support_rating(data)
    println("\nLife Support Rating:\t", val)
end

main()
