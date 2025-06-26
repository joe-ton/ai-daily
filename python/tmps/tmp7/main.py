def main():
    nums: list = [1, 2, 3, 4]
    target: int= 7
    result: list = run(nums, target)

    print(f"Result: {result}")

def run(nums: list, target: int) -> list:
    if len(nums) < 2:
        return []

    seen: dict = {}

    for idx, num in enumerate(nums):
        complement: int = target - num
        if complement in seen:
            compIdx: int = seen[complement]
            return [compIdx, idx]
        seen[num] = idx

    return []


if __name__ == "__main__":
    main()
