class Prime

  def self.nth(n)
    raise ArgumentError if n < 1
    p = prime_numbers(n)
    p[n-1]
  end

  def self.prime_numbers(n)
    return [2] if n == 1
    prime_numbers = []
    while prime_numbers.length < (n)
      (1..105000).each do |number|
        prime_numbers << number if is_prime(number)
      end
    end

    prime_numbers.to_a
  end

  # Naive primality test
  # Given an input number n, check whether any integer m from 2 to n − 1
  # evenly divides n
  def self.is_prime(n)
    return false if (n % 2) == 0
    (2..(n-1)).each do |x|
      return false if (n % x) == 0
    end
    return true
  end

end



# Even though Prime has been included with Ruby since 1.9.3 that's not the
# point of the exercise.

# require 'prime'

# Prime.class_eval do
#   def self.nth(which)
#     raise ArgumentError if which < 1
#     (Prime.first which)[-1]
#   end
# end
