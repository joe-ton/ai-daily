

# Given a list of integers and a target
# With the given list of integers, find two values whose together to the give
# target.  Return the indices of the two values

# list = [1, 2, 3, 4]
# target = 7
# return [2, 3]

# list = [1, 2, 6, 10]
# target = 3
# return [0, 1]
def main():
    nums: list = [1, 2, 3, 4]
    target: int = 7
    result: list = run(nums, target)
    print(f"Result: {result}")

def run(nums: list, target: int) -> list:
    seen: dict = {}

    for idx, num in enumerate(nums):
        complement: int = target - num
        if complement in seen:
            comp_idx: int = seen[complement]
            return [comp_idx, idx]
        seen[num] = idx

    return []

if __name__ == "__main__":
    main()
