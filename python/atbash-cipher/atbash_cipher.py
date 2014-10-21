def decode(cipher):
  plain = ''
  for c in cipher.lower():
    if ord('a') <= ord(c) <= ord('z'):
      offset = ord(c) - ord('a')
      plain += chr(ord('z') - offset)
    elif ord('0') <= ord(c) <= ord('9'):
      plain += c

  return plain

def encode(text):
  block_size = 5
  cipher = ''
  output = ''
  for c in text.lower():
    if ord('a') <= ord(c) <= ord('z'):
      offset = ord(c) - ord('a')
      cipher += chr(ord('z') - offset)
    elif ord('0') <= ord(c) <= ord('9'):
      cipher += c

  for i in range(0, len(cipher), block_size):
     output += cipher[i:i+block_size] + ' '

  return output.strip()

