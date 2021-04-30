import os, winshell
import requests
directory = "4RDM"
path = os.path.join(os.environ["ProgramFiles"], "4RDM")
try:
    os.mkdir(path)
except:
    pass
url = 'https://4rdm.cf/4rdmlogo.ico'
r = requests.get(url, allow_redirects=True)
a = os.environ["ProgramFiles"]
open('{essa}\\4RDM\\4rdmlogo.ico'.format(essa = a), 'wb').write(r.content)

te = ["ProgramFiles"]
print(os.getenv("SystemDrive"))
desktop = winshell.desktop()
path = os.path.join(desktop, "4RDM.CF.url")
target = "fivem://connect/4rdm.cf"
shortcut = open(path, 'w')
shortcut.write('[InternetShortcut]\n')
a = ('URL=fivem://connect/4rdm.cf\nIconIndex=0\nHotKey=0\nIDList=\nIconFile={drive}\\Program Files\\4RDM\\4rdmlogo.ico\n').format(drive = os.getenv("SystemDrive"))
shortcut.write(a)
shortcut.close()