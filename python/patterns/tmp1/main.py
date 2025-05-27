def main(): 
    nums = [1, 2, 3, 4]
    target = 7
    resp = two_sum(nums, target)
    print(resp)

def two_sum(nums: list, target: int) -> list:
    seen = {}

    for idx, num in enumerate(nums):
        complement = target - num
        if complement in seen:
            comp_idx = seen[complement]
            return [comp_idx, idx]
        seen[num] = idx
    return []

if __name__ == "__main__":
    main()

