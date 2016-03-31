import psutil
import os
import time

maxCounter = 1000000
maxRounds = 100
counters = []

min = 8 * 1024 * 1024 * 1024 # 8G
max = 0
memory = []

p = psutil.Process( os.getpid() )

start = time.time()

for round in range(maxRounds):
    for counter in range(maxCounter):
        if round == 0:
            counters[counter] = counter
        else:
            counters.append(counter)

    m = p.memory_info()[1]

    if m < min: min = m
    if m > max: max = m

    memory.append(m)

    # print(str(round) + ': ' + str(m))

print(time.time() - start)
print(min)
print(sum(memory)/len(memory))
print(max)
