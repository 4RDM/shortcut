import os, winshell
import requests

try:
    os.mkdir(os.path.join(os.environ["ProgramFiles"], "4RDM"))
except:
    pass

r = requests.get('https://4rdm.cf/4rdmlogo.ico', allow_redirects=True)
open('{essa}\\4RDM\\4rdmlogo.ico'.format(essa = os.environ["ProgramFiles"]), 'wb').write(r.content)

desktop = winshell.desktop()
path = os.path.join(desktop, "4RDM.CF.url")
target = "fivem://connect/4rdm.cf"
shortcut = open(path, 'w')
shortcut.write('[InternetShortcut]\n')
a = ('URL=fivem://connect/4rdm.cf\nIconIndex=0\nHotKey=0\nIDList=\nIconFile={a}\\4RDM\\4rdmlogo.ico\n').format(a = os.environ["ProgramFiles"])
shortcut.write(a)
shortcut.close()
