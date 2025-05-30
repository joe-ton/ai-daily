import pytest
from main import two_sum

@pytest.mark.parametrize(
    ("nums", "target", "expected"),
    [
        ([2, 7, 11, 15], 7, (0, 1)),
    ],
)
