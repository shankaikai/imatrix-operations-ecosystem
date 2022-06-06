LOCATION_KEY = "Menu_Key"
MAINMENU_NAME = "Main Menu"
SOSMENU_NAME = "SOS Menu"
REPORTMENU_NAME = "Report Menu"
ATTENDANCEMENU_NAME = "Attendance Menu"
HELPMENU_NAME = "Help Menu"
REPORTMAKERMENU_NAME = "Report Maker"


class MenuLocation:
    def __init__(self, parent = None, children = [], name = ""):
        self.parent = parent
        self.children = children
        self.name = name
    

LOC_MAINMENU = MenuLocation(parent = None, name = MAINMENU_NAME)

LOC_SOSMENU = MenuLocation(parent = LOC_MAINMENU, name = SOSMENU_NAME)
LOC_ATTENDANCEMENU = MenuLocation(parent = LOC_MAINMENU, name = ATTENDANCEMENU_NAME)
LOC_REPORTMENU = MenuLocation(parent = LOC_MAINMENU, name = REPORTMENU_NAME)
LOC_HELPMENU = MenuLocation(parent = LOC_MAINMENU, name = HELPMENU_NAME)
LOC_MAINMENU.children = [LOC_SOSMENU, LOC_ATTENDANCEMENU, LOC_REPORTMENU, LOC_HELPMENU]


def getLocation(menuName: str):
    if menuName == MAINMENU_NAME:
        return LOC_MAINMENU
    elif menuName == SOSMENU_NAME:
        return LOC_SOSMENU
    elif menuName == ATTENDANCEMENU_NAME:
        return LOC_ATTENDANCEMENU
    elif menuName == REPORTMENU_NAME:
        return LOC_REPORTMENU
    elif menuName == HELPMENU_NAME:
        return LOC_HELPMENU