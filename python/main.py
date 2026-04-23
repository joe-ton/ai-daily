def two_sum(nums: list[int], target: int) -> list[int]:
    left, right = 0, len(nums) - 1

    while left < right:
        sum: int = nums[left] + nums[right]

        if sum == target:
            return [left + 1, right + 1]
        elif sum < target:
            left += 1
        else:
            right -= 1


def main():
    nums: list[int] = [1, 2, 3, 4]
    target: int = 7

    response: list[int] = two_sum(nums, target)

    print(f"Response: {response}")


if __name__ == "__main__":
    main()
