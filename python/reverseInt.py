

def reverseInt(z:int) -> int:
    zStr = list(str(z))

    i = 0
    j = len(zStr) -1
    for char in zStr:
        zStr[i] , zStr[j] = zStr[j], zStr[i]
        
        if  i < j:
            i += 1
            j -= 1
        else:
            break

    return "".join(zStr)  
        


if __name__ == "__main__":
    print(reverseInt(100))