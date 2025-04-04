
# 3.9+ idiomatic
from typing import Optional

def get_two_sum(nums: List[int], target: int) -> Optional[List[int]]:


def main():
    nums: List[int] = [1, 2, 3, 4]
    target: int = 7

    resp: Optional[List[int]] = get_two_sum(nums, target)



if __name__ == "__main__":
    main()

