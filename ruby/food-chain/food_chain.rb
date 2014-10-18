class FoodChainSong

  def initialize()
    song = String.new
    verses = String.new

    song = <<-SONG.gsub(/^ */, '')
      blank
      |I know an old lady who swallowed a fly.
      |I know an old lady who swallowed a spider.\nIt wriggled and jiggled and tickled inside her.
      |I know an old lady who swallowed a bird.\nHow absurd to swallow a bird!
      |I know an old lady who swallowed a cat.\nImagine that, to swallow a cat!
      |I know an old lady who swallowed a dog.\nWhat a hog, to swallow a dog!
      |I know an old lady who swallowed a goat.\nJust opened her throat and swallowed a goat!
      |I know an old lady who swallowed a cow.\nI don't know how she swallowed a cow!
      |I know an old lady who swallowed a horse.\nShe's dead, of course!
    SONG

    @verse = Array.new(song.split('|'))

    verses = <<-VERSE.gsub(/^ */, '')
      blank
      |
      |She swallowed the spider to catch the fly.
      |She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.
      |She swallowed the cat to catch the bird.
      |She swallowed the dog to catch the cat.
      |She swallowed the goat to catch the dog.
      |She swallowed the cow to catch the goat.
    VERSE

    @accum = Array.new(verses.split('|'))

    @ending = "I don't know why she swallowed the fly. Perhaps she'll die.\n"
  end

  def verse(number_of_verses)
    song = ''
    song << @verse[number_of_verses]

    if number_of_verses != 8
      (2..number_of_verses).reverse_each { |n| song << @accum[n] }
      song << @ending
    end

    return song
  end

  def verses(*args)
    song = ''
    args.each do |n|
      song << verse(n) << "\n"
    end

    return song

  end

  def sing()
    verses(1, 8)
  end

end
