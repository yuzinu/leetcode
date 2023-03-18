from typing import List
import unittest

# 1. Two Sum (Easy)
class Solution:
  def two_sum(self, nums: List[int], target: int) -> List[int]:
    prevMap = {}  # val -> index

    for i, n in enumerate(nums):
        diff = target - n
        if diff in prevMap:
            return [prevMap[diff], i]
        prevMap[n] = i

class Test(unittest.TestCase):
  def test_1(self):
    self.assertEqual(Solution.two_sum(self, [2,7,11,15], 9), [0,1], "Should be [0,1]")

  def test_2(self):
    self.assertEqual(Solution.two_sum(self, [3,2,4], 6), [1,2], "Should be [1,2]")

  def test_3(self):
    self.assertEqual(Solution.two_sum(self, [3,3], 6), [0,1], "Should be [0,1]")

if __name__ == '__main__':
  unittest.main()
