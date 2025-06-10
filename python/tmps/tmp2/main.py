
def main():
    nums: list = [1, 2, 3, 4]
    result: bool = run(nums)
    
    print(f"Result: {result}")

def run(nums: list[int]) -> bool:
    # hashset
    seen: set[int] = set()

    # iterate given nums
    for val in nums:
        # happy path: if current val in seen, return true
        if val in seen:
            return True
        # update values seen
        seen.add(val)

    return False

if __name__ == "__main__":
    main()
