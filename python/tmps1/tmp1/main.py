
def main():
    nums = [1, 2, 3, 4]
    target = 7
    result = find_two_sum(nums, target)

    print(f"Result: {result}")

def find_two_sum(nums: list, target: int) -> list:
    if len(nums) < 2:
        return []

    seen = {}

    for idx, num in enumerate(nums):
        complement = target - num
        if complement in seen:
            compIdx = seen[complement]
            return [compIdx, idx]
        seen[num] = idx

    return []

if __name__ == "__main__":
    main()
