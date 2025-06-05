from main import run

import pytest

@pytest.mark.parametrize("nums,target,expected", [
    ([1, 2, 3, 4], 7, [2, 3])
])
def test_run(nums, target, expected):
    assert run(nums, target) == expected
