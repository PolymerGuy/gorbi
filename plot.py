import matplotlib.pyplot as plt
import numpy as np

args = np.genfromtxt("args.csv")
vals = np.genfromtxt("vals.csv")

rargs = np.array([2.3  ,  2.9  ,  3.2  ,  2.6  ,  3.5  ,  3.8  ,  4.1,    3.6625])
rvals = np.array([0.49100787 ,0.31960135 ,0.24867926 ,0.40236075 ,0.20244296 ,0.19892302,
                  0.24000631, 0.19443274])

plt.plot(args,vals)
plt.plot(rargs,rvals,'*')
plt.show()