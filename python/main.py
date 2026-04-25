def two_sum(nums: list[int], target: int) -> list[int]:

    int_to_idx: dict[int, int] = {}

    for idx, num in enumerate(nums):
        complement = target - num

        if complement in int_to_idx:
            return [int_to_idx[complement], idx]

        int_to_idx[num] = idx

    return []
