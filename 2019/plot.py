import matplotlib.pyplot as plt
import numpy as np

data = np.loadtxt("data.dat",delimiter=',')

x = data[:,0]
y = data[:,1]

indexin = []
indexout = []

for index in range(len(data)):
    if data[index,0]**2+data[index,1]**2 < 1:
        indexin.append(index)
    else:
        indexout.append(index)



xin = data[indexin,0]
yin = data[indexin,1]
xout = data[indexout,0]
yout = data[indexout,1]

f = plt.figure()
plt.plot(xin,yin,'r.')
plt.plot(xout,yout,'b.')
#plt.title("Random Walk")
#plt.show()

f.savefig("foo.pdf", bbox_inches='tight')

exit()
