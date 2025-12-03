def isInvalid(num):
    s = str(num)
    n = len(s)
    
    for patternLen in range(1, n // 2 + 1):
        if n % patternLen == 0:
            pattern = s[:patternLen]
            repeats = n // patternLen
            
            if s == pattern * repeats:
                return True
    
    return False

with open('puzzle.txt', 'r') as f:
    line = f.read().strip()

ranges = [tuple(map(int, r.split('-'))) for r in line.split(',')]

invalidIds = []
for start, end in ranges:
    rangeInvalidIds = [num for num in range(start, end + 1) if isInvalid(num)]
    invalidIds.extend(rangeInvalidIds)


print(f"Sum: {sum(invalidIds)}")