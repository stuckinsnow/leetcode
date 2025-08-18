s = "anagram"
t = "nagaram"

def is_anagram(s, t):
    if len(s) != len(t):
        return False
    sorted_s = ''.join(sorted(s))
    sorted_t = ''.join(sorted(t))
    return sorted_s == sorted_t

print(is_anagram(s, t))
