
def main():
    nums = [1, 2, 3, 4]
    target = 7

    resp: tuple = two_sum(nums, target)
    print(f"Response: {resp}")

def two_sum(nums: list, target: int) -> tuple:
    if len(nums) < 2:
        return ()

    seen = {}

    for idx, val in enumerate(nums):
        complement = target - val
        if complement in seen:
            comp_idx = seen[complement]
            return (comp_idx, idx)
        seen[val] = idx

    return ()

if __name__ == "__main__":
    main()
