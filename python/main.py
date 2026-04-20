def sliding_window(s: string) -> int:
    char_set = set()
    left = 0
    max_length = 0

    for right in range(len(s)):
        while s[right] in char_set:




def main():
    nums: list[int] = [1, 2, 3, 4]
    target: int = 7

    response: list[int] = two_sum(nums, target)

    print(f"Response: {response}")


if __name__ == "__main__":
    main()
