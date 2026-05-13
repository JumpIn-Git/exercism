def is_valid(isbn):
    isbn = isbn.replace('-','')
    if len(isbn) != 10:
        return False
    sum = 0
    i = 10
    for n in isbn[:9]:
        if not n.isnumeric():
            return False
        sum += int(n) * i
        i -= 1
    if isbn[9] == 'X':
        sum += 10
    elif isbn[9].isnumeric():
            sum += int(isbn[9])
    else: 
        return False
    return sum % 11 == 0