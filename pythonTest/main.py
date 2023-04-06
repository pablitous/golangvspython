import os.path, time
start = time.time()
filename = "f:/desarrollo/golangvspython/test.dat"
newfile ='testPython.txt'

if os.path.exists(newfile):
    os.remove(newfile)
    print("The file already exists and has been deleted successfully")

testFile = open(newfile, 'w')

with open(filename) as f:
    content = f.read().splitlines()

numline = 0
for line in content:
    numline += 1
    if numline != 1:
        newline = line[438:]
        times = len(newline)/90
        for i in range(0, int(times)):
            print(newline[i*90:(i+1)*90] , end="\n", file = testFile)
testFile.close()
print("Execution time: %sms" % ((time.time() - start)*1000))