
def main():
    nums = [1, 2, 3, 4]
    target = 7
    resp = two_sum(nums, target)
    print(f"Response: {resp}")

"""
Find two indices whose values are equal together to target
nums: ex. [1, 2, 3, 4]
target: 7
response: [2, 3]
"""
def two_sum(nums: list, target: int) -> list:
    if len(nums) < 2:
        return []
    # TODO: empty args

    seen = {} # int to idx

    for idx, num in enumerate(nums):
        complement = target - num
        if complement in seen:
            comp_idx = seen[complement]
            return [comp_idx, idx]
        seen[num] = idx

    return [] # no indices match


if __name__ == "__main__":
    main()
