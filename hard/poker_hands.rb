def card_ranks(cards)
  v = {"2" => 2, "3" => 3, "4" => 4, "5" => 5, "6" => 6, "7" => 7, "8" => 8,
       "9" => 9, "T" => 10, "J" => 11, "Q" => 12, "K" => 13, "A" => 14}
  ranks = cards.map {|x| v[x[0]]}
  ranks.sort!.reverse!
  ranks = [5,4,3,2,1] if ranks == [14,5,4,3,2]
  ranks
end

def straight(ranks)
  ranks == [ranks[0], ranks[0]-1, ranks[0]-2, ranks[0]-3, ranks[0]-4]
end

def flush(hand)
  hand[0][1] == hand[1][1] and hand[0][1] == hand[2][1] and
  hand[0][1] == hand[3][1] and hand[0][1] == hand[4][1]
end

def kind(n, ranks)
  for i in ranks
    return i if ranks.count(i) == n
  end
  nil
end

def two_pair(ranks)
  h = kind(2, ranks)
  return nil unless h
  for i in ranks
    return [h, i] if ranks.count(i) == 2 and i != h
  end
  nil
end

def poker(h0, h1)
  for i in (0..h0.length-1)
    if h0[i].kind_of?(Array)
      for j in (0..h0[i].length-1)
        if h0[i][j] > h1[i][j]
          return "left"
        elsif h0[i][j] < h1[i][j]
          return "right"
        end
      end
    else
      if h0[i] > h1[i]
        return "left"
      elsif h0[i] < h1[i]
        return "right"
      end
    end
  end
  "none"
end

def hand_rank(hand)
  ranks = card_ranks(hand)
  return [8, ranks.max] if straight(ranks) and flush(hand)                        # straight flush
  return [7, kind(4, ranks), kind(1, ranks)] if kind(4, ranks)                    # 4 of a kind
  return [6, kind(3, ranks), kind(2, ranks)] if kind(3, ranks) and kind(2, ranks) # full house
  return [5, ranks] if flush(hand)                                                # flush
  return [4, ranks.max] if straight(ranks)                                        # straight
  return [3, kind(3, ranks), ranks] if kind(3, ranks)                             # 3 of a kind
  return [2, two_pair(ranks), kind(1, ranks)] if two_pair(ranks)                  # 2 pair
  return [1, kind(2, ranks), ranks] if kind(2, ranks)                             # kind
  [0, ranks]                                                                      # high card
end

File.open(ARGV[0]).each_line do |line|
  s = line.split(" ")
  puts poker(hand_rank(s[0..4]), hand_rank(s[5..-1]))
end
