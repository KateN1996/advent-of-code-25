def findMax(line):
    maxVal = 0
    for i in range(len(line)):
        for j in range(i+ 1, len(line)):
            curVal = int(line[i] + line[j])
            maxVal = max(curVal,maxVal)
    return maxVal

curBankSum = 0
totalOutput = 0
with open('input.txt', 'r') as file:
    for line in file:
        totalOutput += findMax(line)

print(f"total out put {totalOutput}")