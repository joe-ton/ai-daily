
def main():
    nums: list = [1, 2, 3, 4]
    target: int = 7
    result: list = run(nums, target)

    print(f"Result: {result}")

def run(nums: list, target: int) -> list:
    seen: dict = {} # previous values; int to idx

    for idx, val in enumerate(nums):
        # complement = target - num
        complement = target - val
        if complement in seen:
            comp_idx = seen[complement]
            return [comp_idx, idx]
        seen[val] = idx

    return []

if __name__ == "__main__":
    main()
