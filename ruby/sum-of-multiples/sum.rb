class SumOfMultiples

  attr_reader :divisors

  def initialize(*divisors)
    @divisors = divisors
  end

  # Class Method
  def self.to(limit)
    new(3,5).to(limit)
  end

  # Instance Method
  def to(limit)
    sum = 0
    (1...limit).each do |i|
      sum += i if divisors.any? { |d| i % d == 0 }
    end
    sum
  end

end
