def detect_anagrams(word, candidates):
  anagrams = []

  # Same word is not an anagram
  candidates = [x for x in candidates if x.lower() != word.lower()]

  letters = sorted(word.lower())

  for candidate in candidates:
    anagram = sorted(candidate.lower())
    # If the sorted list of letters in two words are the same, presto!
    if letters == anagram:
      anagrams.append(candidate)

  return anagrams

