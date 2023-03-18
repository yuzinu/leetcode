from typing import List
from collections import defaultdict
import unittest

# 49. Group Anagrams (Medium)
class Solution:
  def group_anagrams(self, strs: List[str]) -> List[List[str]]:
    ans = defaultdict(list)

    for s in strs:
      count = [0] * 26
      for c in s:
        count[ord(c) - ord("a")] += 1
      ans[tuple(count)].append(s)
    print(ans)
    return ans.values()

class Test(unittest.TestCase):
  def test_1(self):
    self.assertEqual(
      list(Solution.group_anagrams(self, ["eat","tea","tan","ate","nat","bat"])),
      [["eat","tea","ate"],["tan","nat"],["bat"]],
      'Should be [["eat","tea","ate"],["tan","nat"],["bat"]]'
    )

  def test_2(self):
    self.assertEqual(list(Solution.group_anagrams(self, [""])), [[""]], 'Should be [[""]]')

  def test_3(self):
    self.assertEqual(list(Solution.group_anagrams(self, ["a"])), [["a"]], 'Should be [["a"]]')

if __name__ == '__main__':
  unittest.main()
