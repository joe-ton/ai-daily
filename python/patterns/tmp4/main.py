
def run(nums: list, target: int) -> list:
    seen = {}
    for idx, num in enumerate(nums):
        complement = target - num
        if complement in nums:
            comp_idx = seen[complement]
            return [comp_idx, idx]
        seen[num] = idx
    return []

def main():
    nums = [1, 2, 3, 4]
    target = 7
    resp = run(nums, target)
    print(resp)


if __name__ == "__main__":
    main()
