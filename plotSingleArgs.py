import matplotlib.pyplot as plt
import numpy as np
import scipy.interpolate as intr


args = np.genfromtxt("args.csv")
vals = np.genfromtxt("vals.csv")

rargs = np.array([2.3  ,  2.9  ,  3.2  ,  2.6  ,  3.5  ,  3.8  ,  4.1,    3.6625])
rvals = np.array([0.49100787 ,0.31960135 ,0.24867926 ,0.40236075 ,0.20244296 ,0.19892302,
                  0.24000631, 0.19443274])


rbf = intr.Rbf(rargs,rvals)
rbf_args = np.linspace(args.min(),args.max(),50)
rbf_vals = rbf(rbf_args)

best_arg = np.argmin(rbf_vals)

#print("Best guess is %f"%rbf_args[best_arg])


plt.plot(rbf_args,rbf_vals,alpha=0.4)
plt.plot(args,vals,'--')
plt.plot(rargs,rvals,'*')
plt.show()