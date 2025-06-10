
def main(): 
    nums: list = [1, 2, 3, 3]
    result: bool = run(nums)

    print(f"Result: {result}")

def run(nums: list[int]) -> bool:
    seen: set[int] = set()

    for val in nums:
        if val in seen:
            return True
        seen.add(val)

    return False


if __name__ == "__main__":
    main()
