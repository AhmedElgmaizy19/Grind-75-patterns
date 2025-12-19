
def is_valid_plaindrome(s):
    txt = ''.join(c.lower() for c in s if c.isalnum())
    start = 0
    end = len(txt) -1
    
    while start < end:
        if txt[start] != txt[end]:
            return False
        start += 1
        end -= 1
    return True
