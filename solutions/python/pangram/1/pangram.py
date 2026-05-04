def is_pangram(sentence):
    s = set()
    for l in sentence:
        if l.isalpha(): 
            s.add(l.lower())
    return len(s) == 26
