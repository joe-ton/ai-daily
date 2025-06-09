
def main():
    nums: list = [1, 2, 3, 4]
    target: int = 7
    result = run(nums, target)

    print(f"Result: {result}")

"""
Find indices whose values are equal to target
"""
def run(nums: list, target: int) -> list:
    if len(nums) < 2:
        return []

    seen: dict = {} # complements; int to idx

    for idx, num in enumerate(nums):
        complement = target - num
        if complement in seen:
            comp_idx = seen[complement]
            return [comp_idx, idx]
        seen[num] = idx

    return []

if __name__ == "__main__":
    main()
