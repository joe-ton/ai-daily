
def main():
    nums = [1, 2, 3, 4]
    target = 7
    resp = run(nums, target)
    print(f"Response: {resp}")

def run(nums: list, target: int) -> list:
    if len(nums) < 2:
        return []

    seen = {}

    for idx, num in enumerate(nums):
        complement = target - num
        if complement in seen:
            comp_idx = seen[complement]
            return [comp_idx, idx]
        seen[num] = target

if __name__ == "__main__":
    main()
