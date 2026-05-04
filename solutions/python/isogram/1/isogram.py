def is_isogram(string):
    seen = set()
    for l in string:
        if not l.isalpha():
            continue
        l = l.lower()
        if l in seen:
            return False
        else: seen.add(l)
    return True