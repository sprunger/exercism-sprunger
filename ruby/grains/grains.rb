class Grains

  def initialize
    @TOTAL_SQUARES = 64
  end

  def square(n)
    2 ** (n - 1)
  end

  def total
    2 ** (@TOTAL_SQUARES) - 1
  end

end
