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
      sq_rt_of_num = Math.sqrt(number)
      number_divisible_by = 2

      while number_divisible_by <= sq_rt_of_num
        break if number % number_divisible_by == 0
        number_divisible_by = number_divisible_by + 1
      end

      if number_divisible_by > sq_rt_of_num
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
