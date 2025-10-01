#include <gtest/gtest.h>

int Add(int a, int b) { return a + b; }

TEST(AddTest, PositiveNumbers) {
  EXPECT_EQ(5, Add(2, 3));
  EXPECT_EQ(10, Add(4, 6));
}

int main(int argc, char **argv) {
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}
