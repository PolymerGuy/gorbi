import matplotlib.pyplot as plt
import numpy as np

args = np.genfromtxt("args.csv")
vals = np.genfromtxt("vals.csv")


plt.plot(args,vals)
plt.show()