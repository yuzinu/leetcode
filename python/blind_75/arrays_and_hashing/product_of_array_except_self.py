from typing import List
import unittest

# 238. Product of Array Except Self (Medium)
class Solution:
  def product_except_self(self, nums: List[int]) -> List[int]:
    res = [1] * (len(nums))

    prefix = 1

    for i in range(len(nums)):
      res[i] = prefix
      prefix *= nums[i]

    postfix = 1

    for i in range(len(nums) - 1, -1, -1):
      res[i] *= postfix
      postfix *= nums[i]

    return res

class Test(unittest.TestCase):
  def test_1(self):
    self.assertEqual(Solution.product_except_self(self, [1,2,3,4]), [24,12,8,6], "Should be [24,12,8,6]")

  def test_2(self):
    self.assertEqual(Solution.product_except_self(self, [-1,1,0,-3,3]), [0,0,9,0,0], "Should be [0,0,9,0,0]")

if __name__ == '__main__':
  unittest.main()
