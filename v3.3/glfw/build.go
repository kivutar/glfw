package glfw

/*
// Windows Build Tags
// ----------------
// GLFW Options:
#cgo windows CFLAGS: -D_GLFW_WIN32 -Iglfw/deps/mingw

// Linker Options:
#cgo windows LDFLAGS: -lopengl32 -lgdi32


// Darwin Build Tags
// ----------------
// GLFW Options:
#cgo darwin CFLAGS: -D_GLFW_COCOA -D_GLFW_USE_RETINA -Wno-deprecated-declarations -x objective-c

// Linker Options:
#cgo darwin LDFLAGS: -framework Cocoa -framework OpenGL -framework IOKit -framework IOSurface -framework QuartzCore -framework Metal -lc++


// Linux Build Tags
// ----------------
// GLFW Options:
#cgo linux,!wayland CFLAGS: -D_GLFW_X11 -D_GNU_SOURCE
#cgo linux,wayland CFLAGS: -D_GLFW_WAYLAND -D_GNU_SOURCE

// Linker Options:
#cgo linux,!wayland LDFLAGS: -lGL -lX11 -lXrandr -lXxf86vm -lXi -lXcursor -lm -lXinerama -ldl -lrt
#cgo linux,wayland LDFLAGS: -lGL -lwayland-client -lwayland-cursor -lwayland-egl -lxkbcommon -lm -ldl -lrt


// FreeBSD Build Tags
// ----------------
// GLFW Options:
#cgo freebsd,!wayland CFLAGS: -D_GLFW_X11 -D_GLFW_HAS_GLXGETPROCADDRESSARB -D_GLFW_HAS_DLOPEN
#cgo freebsd,wayland CFLAGS: -D_GLFW_WAYLAND -D_GLFW_HAS_DLOPEN

// Linker Options:
#cgo freebsd,!wayland LDFLAGS: -lGL -lX11 -lXrandr -lXxf86vm -lXi -lXcursor -lm -lXinerama
#cgo freebsd,wayland LDFLAGS: -lGL -lwayland-client -lwayland-cursor -lwayland-egl -lxkbcommon -lm
*/
import "C"
