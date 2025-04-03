
def get_two_sum(nums: list, target: int) -> list: 
    seen: dict = {} # int to int
    
    for idx, num in enumerate(nums):
        complement = target - num
        print()

        if complement in seen:
            compIdx = seen[complement]
            return [compIdx, idx]
        seen[num] = idx

    return []

def main():
    nums: list = [1, 2, 3, 4]
    target: int = 7

    resp: list = get_two_sum(nums, target)
    print(resp)
    if resp is None:
        print(f"Error:{resp}")
    print(f"Response:{resp}")

main()
