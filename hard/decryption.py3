import random
mess = "012222 1114142503 0313012513 03141418192102 0113 2419182119021713 06131715070119"
key = "BHISOECRTMGWYVALUZDNFJKPQX"

fwords = ['ALL', 'AT', 'ATHINAI', 'BE', 'BEEN', 'BITTORRENT', 'BLIZZARD',
          'CAN', 'CALL', 'COME', 'COULD', 'DID', 'DO', 'DOWN', 'EACH',
          'FIRST', 'FOR', 'GET', 'GO', 'HAVE', 'HAD', 'LINUX', 'LOL',
          'LOLCAT', 'MIDNIGHT', 'MORE', 'KTHXBAI', 'NUMBER', 'PEERS',
          'PROTOCOL', 'ROTFL', 'SCNR', 'SEEDING', 'START', 'YOUTUBE']

def wfit(wx, w, f, alpha):
    for ix, i in enumerate(w):
        if i in alpha:
            if alpha[i] != f[ix]:
                return False
        elif f[ix] in seen:
            return False
        else:
            alpha[i] = f[ix]
            seen.append(f[ix])
    if wx == len(words)-1:
        return True
    wfound = False
    for ix, i in enumerate(words):
        if ix <= wx:
            continue
        for j in fwords:
            if wstruct(i) == wstruct(j):
                if wfit(ix, i, j, alpha):
                    wfound = True
                    break
    return wfound

def wstruct(word):
    ret = []
    curr = 0
    strd = {}
    for i in word:
        if i in strd:
            ret.append(strd[i])
        else:
            strd[i] = curr
            ret.append(curr)
            curr += 1
    return ret

mess = mess.split()
for ix, i in enumerate(mess):
    wrd = []
    for j in range(len(i)//2):
        wrd.append(int(i[2*j:2*j+2]))
    mess[ix] = wrd
words = []
for i in mess:
    wrd = ""
    for j in i:
        wrd += chr(64+j)
    words.append(wrd)

while True:
    alpha = {}
    seen = []
    for ix, i in enumerate(words):
        wfound = False
        for j in fwords:
            if wstruct(i) == wstruct(j): 
                if wfit(ix, i, j, alpha):
                    wfound = True
                    break
        if not wfound:
            break
    if wfound:
        break
    random.shuffle(fwords)

print(*[''.join(alpha[j] if j in alpha else '?' for j in i) for i in words])
