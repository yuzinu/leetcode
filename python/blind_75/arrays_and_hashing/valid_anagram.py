import unittest

# 242. Valid Anagram (Easy)
class Solution:
  def is_anagram(self, s: str, t: str) -> bool:
    if len(s) != len(t):
      return False

    countS, countT = {}, {}

    for i in range(len(s)):
      countS[s[i]] = 1 + countS.get(s[i], 0)
      countT[t[i]] = 1 + countT.get(t[i], 0)
    return countS == countT

class Test(unittest.TestCase):
  def test_1(self):
    self.assertEqual(Solution.is_anagram(self, "anagram", "nagaram"), True, "Should be True")

  def test_2(self):
    self.assertEqual(Solution.is_anagram(self, "rat", "car"), False, "Should be False")

if __name__ == '__main__':
  unittest.main()
