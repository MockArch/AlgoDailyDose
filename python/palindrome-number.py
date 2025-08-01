
def isNumberPalindrome(x : int) -> bool:
    
    num, rev = x, 0
    while num > 0:
        digit = int(num % 10)
        print(digit)
        rev = int(rev * 10) + digit
        print(rev)
        num = int(num / 10)
    
    return rev == x

if __name__ == "__main__":
    print(isNumberPalindrome(-121))