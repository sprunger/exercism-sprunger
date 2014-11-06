require 'prime'

class PrimeFactors
  def self.for(n)
    n.prime_division.map { |base, exp| [base]*exp }.flatten

    # --- The long way
    # factors = []
    # divisor = 2

    # while number > 1
    #   while (number % divisor) == 0
    #     factors << divisor
    #     number /= divisor
    #   end
    #   divisor += 1
    # end

    # factors
  end

end
