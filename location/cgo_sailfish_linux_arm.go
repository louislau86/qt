// +build sailfish

package location

/*
#cgo CFLAGS: -pipe -O2 -g -pipe -Wall -Wp,-D_FORTIFY_SOURCE=2 -fexceptions -fstack-protector --param=ssp-buffer-size=4 -Wformat -Wformat-security -fmessage-length=0 -march=armv7-a -mfloat-abi=hard -mfpu=neon -mthumb -Wno-psabi -fPIC -fvisibility=hidden -fvisibility-inlines-hidden -Wall -W -D_REENTRANT -fPIE
#cgo CXXFLAGS: -pipe -O2 -g -pipe -Wall -Wp,-D_FORTIFY_SOURCE=2 -fexceptions -fstack-protector --param=ssp-buffer-size=4 -Wformat -Wformat-security -fmessage-length=0 -march=armv7-a -mfloat-abi=hard -mfpu=neon -mthumb -Wno-psabi -fPIC -fvisibility=hidden -fvisibility-inlines-hidden -Wall -W -D_REENTRANT -fPIE
#cgo CXXFLAGS: -DQT_NO_DEBUG -DQT_LOCATION_LIB -DQT_POSITIONING_LIB -DQT_QUICK_LIB -DQT_GUI_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -I/srv/mer/targets/SailfishOS-armv7hl/usr/share/qt5/mkspecs/linux-g++ -I/srv/mer/targets/SailfishOS-armv7hl/usr/include -I/srv/mer/targets/SailfishOS-armv7hl/usr/include/sailfishapp -I/srv/mer/targets/SailfishOS-armv7hl/usr/include/mdeclarativecache5 -I/srv/mer/targets/SailfishOS-armv7hl/usr/include/qt5
#cgo CXXFLAGS: -I/srv/mer/targets/SailfishOS-armv7hl/usr/include/qt5/QtLocation -I/srv/mer/targets/SailfishOS-armv7hl/usr/include/qt5/QtPositioning -I/srv/mer/targets/SailfishOS-armv7hl/usr/include/qt5/QtQuick -I/srv/mer/targets/SailfishOS-armv7hl/usr/include/qt5/QtGui -I/srv/mer/targets/SailfishOS-armv7hl/usr/include/qt5/QtCore

#cgo LDFLAGS: -Wl,-O1 -Wl,-rpath-link,/srv/mer/targets/SailfishOS-armv7hl/usr/lib -Wl,-rpath-link,/srv/mer/targets/SailfishOS-armv7hl/lib
#cgo LDFLAGS: -rdynamic -L/srv/mer/targets/SailfishOS-armv7hl/usr/lib -L/srv/mer/targets/SailfishOS-armv7hl/lib -lsailfishapp -lmdeclarativecache5 -lQt5Positioning -lQt5Quick -lQt5Gui -lQt5Core -lGLESv2 -lpthread
*/
import "C"
