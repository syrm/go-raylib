# go-raylib


### Bindings
Header     | Supported          |
---------  | ------------------ |
raylib.h   | :heavy_check_mark: |
raymath.h  | :heavy_check_mark: |
raygui.h   | :heavy_check_mark: |
ricons.h   | :heavy_check_mark: |
physac.h   | :heavy_check_mark: |

### Platforms
OS         | Supported          |
---------  | ------------------ |
Mac        | :heavy_check_mark: |
Windows    | :heavy_check_mark: |

You can write your own cross-platform code, see 
[raylib_darwin.go](https://github.com/chunqian/go-raylib/blob/master/raylib/raylib_darwin.go) 
[raylib_windows.go](https://github.com/chunqian/go-raylib/blob/master/raylib/raylib_windows.go)

### Version
go-raylib binding [raylib](https://github.com/raysan5/raylib) C version is 3.5


### Performance
High performance, same with the raylib C version. Enjoy it!.

### Development Tools
I use [sublime text](https://www.sublimetext.com) and customize tools.

[c-for-go](https://github.com/chunqian/c-for-go) automatic C-Go bindings generator for raylib C version.

[LSP](https://github.com/chunqian/LSP) use [mistune](https://github.com/lepture/mistune) instead [mdpopups](https://github.com/facelessuser/sublime-markdown-popups).

[language-formatter](https://github.com/chunqian/Language-Formatter) generic code format tool.

[gopls](https://github.com/chunqian/golang-tools) fix go-raylib code autocomplete slow.


Memory
------
For example
```go
multext := rl.NewMultiText([]string{"Hello World!"})
```
The method will check if memory is requested through cgo.

If it detects, panic **Cgo memory alloced, please use func AllocMultiText.**

Rewrite.
```go
multext, men := rl.AllocMultiText([]string{"Hello World!"})
multext.GC(mem)
```
Don't forget, call GC() for register, it can be automated management.


Difference
----------
There are some differences between the processing in Go and C.

In C
```c
char multiTextBoxText[256] = "Multi text box";
```
In Go
```go
multiTextBoxText := rg.NewBytes("Multi text box", 256)
```

In C
```c
const char *listViewExList[8] = { "This", "is", "a", "list view", "with", "disable", "elements", "amazing!" };
```
In Go
```go
listViewExList, mem := rg.AllocMultiText([]string{"This", "is", "a", "list view", "with", "disable", "elements", "amazing!"})
listViewExList.GC(mem)
```

In C
```c
int dropsCount = 0;
char **droppedFiles = GetDroppedFiles(&dropsCount);
const char *droppedFilePath = droppedFiles[0];
```
In Go
```go
dropsCount := int32(0)
droppedFiles := rl.GetDroppedFiles(&dropsCount)
droppedFilePath := rl.ToString(droppedFiles, 0)
```

In C
```c
Texture2D texture = LoadTexture("resources/cubicmap_atlas.png");
model.materials[0].maps[MAP_DIFFUSE].texture = texture;
```
In Go
```go
texture := rl.LoadTexture("../models/resources/cubicmap_atlas.png")
model.Materialser(0).Mapser(rl.MAP_DIFFUSE).Texture = texture
```


Build
-----
Mac
```bash
$ git clone https://github.com/chunqian/go-raylib.git
$ cd go-raylib/lib/raylib/src
$ export MACOSX_DEPLOYMENT_TARGET=10.9
$ make PLATFORM=PLATFORM_DESKTOP PLATFORM_OS=OSX
$ cd ../../..
$ mkdir examples/bin
$ go build -o examples/bin/models_material_pbr examples/models/material_pbr/*.go
```
Run
```bash
$ cd examples/bin
$ ./models_material_pbr
```
Windows
```bash
$ git clone https://github.com/chunqian/go-raylib.git
$ cd go-raylib/lib/raylib/src
$ make CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ AR=i686-w64-mingw32-ar PLATFORM=PLATFORM_DESKTOP PLATFORM_OS=WINDOWS LANG=en_US
$ cd ../../..
$ mkdir examples/bin
$ CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/models_material_pbr.exe examples/models/material_pbr/*.go
```
Run
```bash
$ cd examples/bin
$ ./models_material_pbr.exe
```

Examples
--------
Category   | Example                | Supported          |
---------  | ---------------------- | ------------------ |
audio      | module_playing         | :heavy_check_mark: |
audio      | multichannel_sound     | :heavy_check_mark: |
audio      | music_stream           | :heavy_check_mark: |
audio      | raw_stream             | :heavy_check_mark: |
audio      | sound_loading          | :heavy_check_mark: |
core       | 2d_camera              | :heavy_check_mark: |
core       | 2d_camera_platformer   | :heavy_check_mark: |
core       | 3d_camera_first_person | :heavy_check_mark: |
core       | 3d_camera_free         | :heavy_check_mark: |
core       | 3d_camera_mode         | :heavy_check_mark: |
core       | 3d_picking             | :heavy_check_mark: |
core       | basic_window           | :heavy_check_mark: |
core       | drop_files             | :heavy_check_mark: |
core       | input_gestures         | :heavy_check_mark: |
core       | input_keys             | :heavy_check_mark: |
core       | input_mouse            | :heavy_check_mark: |
core       | input_mouse_wheel      | :heavy_check_mark: |
core       | input_multitouch       | :heavy_check_mark: |
core       | random_values          | :heavy_check_mark: |
core       | scissor                | :heavy_check_mark: |
core       | storage_values         | :heavy_check_mark: |
core       | vr_simulator           | :heavy_check_mark: |
core       | window_letterbox       | :heavy_check_mark: |
core       | world_screen           | :heavy_check_mark: |
gui        | controls_test_suite    | :heavy_check_mark: |
gui        | scroll_panel           | :heavy_check_mark: |
models     | animation              | :heavy_check_mark: |
models     | billboard              | :heavy_check_mark: |
models     | box_collisions         | :heavy_check_mark: |
models     | cubicmap               | :heavy_check_mark: |
models     | first_person_maze      | :heavy_check_mark: |
models     | loading                | :heavy_check_mark: |
models     | material_pbr           | :heavy_check_mark: |
models     | mesh_generation        | :heavy_check_mark: |
models     | mesh_picking           | :heavy_check_mark: |
models     | orthographic_projection| :heavy_check_mark: |
models     | skybox                 | :heavy_check_mark: |
models     | waving_cubes           | :heavy_check_mark: |
models     | yaw_pitch_roll         | :heavy_check_mark: |
physac     | demo                   | :heavy_check_mark: |
shaders    | postprocessing         | :heavy_check_mark: |
shaders    | basic_lighting         | :heavy_check_mark: |
shaders    | eratosthenes           | :heavy_check_mark: |
shaders    | fog                    | :heavy_check_mark: |
shaders    | julia_set              | :heavy_check_mark: |
shaders    | model_shader           | :heavy_check_mark: |
shaders    | palette_switch         | :heavy_check_mark: |
text       | font_filters           | :heavy_check_mark: |
text       | font_loading           | :heavy_check_mark: |
text       | font_sdf               | :heavy_check_mark: |
text       | font_spritefont        | :heavy_check_mark: |
text       | format_text            | :heavy_check_mark: |
text       | input_box              | :heavy_check_mark: |
text       | raylib_fonts           | :heavy_check_mark: |
text       | rectangle_bounds       | :heavy_check_mark: |
text       | unicode                | :heavy_check_mark: |
text       | writing_anim           | :heavy_check_mark: |
textures   | bunnymark              | :heavy_check_mark: |


License
-------
go-raylib is licensed under an unmodified zlib/libpng license, which is an OSI-certified, BSD-like license that allows static linking with closed source software. Check [LICENSE](LICENSE) for further details.
