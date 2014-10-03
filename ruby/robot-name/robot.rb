class Robot

  attr_accessor :name
  @name

  def initialize
    @name = random_string(2) + random_number(3)
  end

  private

  def random_string(length=10)
    select_random(2, 'abcdefghjkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ')
 end

  def random_number(length=10)
    select_random(3, '0123456789')
  end

  def select_random(length=10, chars="")
    length.times.collect { chars[rand(chars.size)] }.join
  end
end
