

from dataclasses import dataclass

@dataclass
class TwoSumInput:
    nums: list[int]
    target: int

@dataclass
class TwoSumResult:
    indices: tuple[int, int] | None

def two_sum(data: TwoSumInput) -> TwoSumResult:
    seen = {}

    for idx, num in enumerate(data.nums):
        complement = data.target - num
        if complement in seen:
            compIdx = seen[complement]
            indices = (compIdx, idx)
            return TwoSumResult(indices)
        seen[num] = idx

    return TwoSumResult(indices=None)

def main():
    nums = [1, 2, 3, 4]
    target = 7
    data = TwoSumInput(nums=nums, target=target)
    result = two_sum(data)
    print(result)

if __name__ == "__main__":
    main()
