def is_anagram(word, candidate):
  word = word.lower()
  candidate = candidate.lower()

  return word != candidate and sorted(word) == sorted(candidate)

def detect_anagrams(word, candidates):
  anagrams = []

  for candidate in candidates:
    if is_anagram(word, candidate):
      anagrams.append(candidate)

  return anagrams

