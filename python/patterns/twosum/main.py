
from dataclasses import dataclass

@dataclass
class TwoSumInput:
    nums: list[int]
    target: int

@dataclass
class TwoSumResult:
    indices: list[int] | None

def two_sum(data: TwoSumInput) -> TwoSumResult: 
    seen = {}

    for idx, num in enumerate(data.nums):
        complement = TwoSumInput.target - num
        if complement in seen:
            compIdx = seen[num]
            indices = (compIdx, idx)
            return TwoSumResult(indices=indices)
        seen[num] = idx
    return TwoSumResult(indices=None)

def main():
    nums: list[int] = [1, 2, 3, 4]
    target: int = 7
    res = two_sum(TwoSumInput=(nums=nums, target=target))
    print(res) 

if __name__ == "__main__":
    main()
