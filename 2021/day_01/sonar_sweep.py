import timeit
from pathlib import Path
from typing import Callable


INPUT_FP = Path(__file__).parent / "input.txt"


def read_input(fp: Path) -> list[int]:
    with open(fp, "r") as f:
        data = list(map(int, f.read().splitlines()))
    return data


def print_solution(val: str, time: float | None = None) -> None:
    print(f"\nSolution:\t{val}")
    if time is not None:
        print(f"Duration:\t{time:.6f}s")


def execute(func: Callable, *args) -> tuple[int, float]:
    t = timeit.Timer(lambda: func(*args))
    t_avg = t.timeit(10_000) / 10_000
    return func(*args), t_avg


def calc_step_one(data: list[int]) -> int:
    inc = 0
    prev = data[0]
    for e in data[1:]:
        if e > prev:
            inc += 1
        prev = e
    return inc


def calc_step_two(data: list[int]) -> int:
    inc = 0
    prev = sum(data[:3])
    for i in range(1, len(data) - 2):
        new = sum(data[i : i + 3])
        if new > prev:
            inc += 1
        prev = new
    return inc


def main() -> None:
    data = read_input(INPUT_FP)

    sol, t_avg = execute(calc_step_one, data)

    print_solution(sol, t_avg)

    sol, t_avg = execute(calc_step_two, data)

    print_solution(sol, t_avg)


if __name__ == "__main__":
    main()
