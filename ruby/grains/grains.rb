class Grains

  def initialize
  end

  def square(n)
    2 ** (n - 1)
    #1 << ( n - 1 )
  end

  def total
    2 ** (64) - 1
    #(1 << 64) - 1
  end

end
