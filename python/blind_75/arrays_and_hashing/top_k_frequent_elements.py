from typing import List
import unittest

# 347. Top K Frequent Elements (Medium)
class Solution:
  def top_k_frequent(self, nums: List[int], k: int) -> List[int]:
    count = {}
    freq = [[] for i in range(len(nums) + 1)]

    for n in nums:
      count[n] = 1 + count.get(n, 0)
    for n, c in count.items():
      freq[c].append(n)

    res = []
    for i in range(len(freq) - 1, 0, -1):
      for n in freq[i]:
        res.append(n)
        if len(res) == k:
          return res

class Test(unittest.TestCase):
  def test_1(self):
    self.assertEqual(Solution.top_k_frequent(self, [1,1,1,2,2,3], 2), [1,2], "Should be [1,2]")

  def test_2(self):
    self.assertEqual(Solution.top_k_frequent(self, [1], 1), [1], "Should be [1]")

if __name__ == '__main__':
  unittest.main()
