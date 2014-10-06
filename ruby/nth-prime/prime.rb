class Prime

  def self.nth(n)
    raise ArgumentError if n < 1
    prime_numbers(n)[-1]
  end

  def self.prime_numbers(no_of_primes)
    return [2] if no_of_primes == 1

    prime_numbers=[]

    count = 1
    number = 3

    while count < no_of_primes
      square_root = Math.sqrt(number)
      divisor = 2

      (divisor..square_root).each do
        (number % divisor) != 0 ? divisor +=1 : break
      end

      if divisor > square_root
        prime_numbers << number
        count +=1
      end

      number += 2
    end

    return prime_numbers

  end
end



# Even though Prime has been included with Ruby since 1.9.3 that's not the
# point of the exercise.

# require 'prime'

# Prime.class_eval do
#   def self.nth(how_many)
#     raise ArgumentError if how_many < 1
#     (Prime.first how_many)[-1]
#   end
# end
