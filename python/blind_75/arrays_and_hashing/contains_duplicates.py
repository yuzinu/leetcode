from typing import List
import unittest

# 217. Contains Duplicate (Easy)
class Solution:
  def contains_duplicate(self, nums: List[int]) -> bool:
    hashset = set()

    for n in nums:
        if n in hashset:
            return True
        hashset.add(n)
    return False

class TestSum(unittest.TestCase):
  def test_1(self):
    self.assertEqual(Solution.contains_duplicate(self, [1,2,3,1]), True, "Should be True")

  def test_2(self):
    self.assertEqual(Solution.contains_duplicate(self, [1,2,3,4]), False, "Should be False")

  def test_3(self):
    self.assertEqual(Solution.contains_duplicate(self, [1,1,1,3,3,4,3,2,4,2]), True, "Should be True")

if __name__ == '__main__':
  unittest.main()
