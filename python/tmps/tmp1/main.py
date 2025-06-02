
def main():
    nums = [1, 2, 3, 4]
    target = 7
    resp = run(nums, target)
    print(resp)

def run(nums: list, target: int) -> list:
    if len(nums) < 2:
        return []

    seen = {}

    for idx, num in enumerate(nums):
        complement = target - num
        if complement in seen:
            compl_idx = seen[complement]
            return [compl_idx, idx]
        seen[num] = idx
    return []


if __name__ == "__main__":
    main()
