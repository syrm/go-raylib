// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Thu, 10 Sep 2020 17:34:47 CST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package raygui

/*
#include "../lib/raygui/src/raygui.h"
#include "../lib/raygui/src/ricons.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// GuiStyleProp as declared in src/raygui.h:261
type gGuiStyleProp struct {
	gControlId     uint16
	gPropertyId    uint16
	gPropertyValue int32
	ref8760bcee    *C.GuiStyleProp
	allocs8760bcee interface{}
}
type GuiStyleProp struct {
	ControlId     uint16
	PropertyId    uint16
	PropertyValue int32
}

// Vector2 as declared in src/raylib.h:176
type gVector2 struct {
	gX             float32
	gY             float32
	ref29ca61a5    *C.Vector2
	allocs29ca61a5 interface{}
}
type Vector2 struct {
	X float32
	Y float32
}

// Vector3 as declared in src/raylib.h:183
type gVector3 struct {
	gX             float32
	gY             float32
	gZ             float32
	ref5ecd5133    *C.Vector3
	allocs5ecd5133 interface{}
}
type Vector3 struct {
	X float32
	Y float32
	Z float32
}

// Color as declared in src/raylib.h:210
type gColor struct {
	gR             byte
	gG             byte
	gB             byte
	gA             byte
	refa79767ed    *C.Color
	allocsa79767ed interface{}
}
type Color struct {
	R byte
	G byte
	B byte
	A byte
}

// Rectangle as declared in src/raylib.h:218
type gRectangle struct {
	gX             float32
	gY             float32
	gWidth         float32
	gHeight        float32
	refcee8783a    *C.Rectangle
	allocscee8783a interface{}
}
type Rectangle struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

// Image as declared in src/raylib.h:228
type gImage struct {
	gData         unsafe.Pointer
	gWidth        int32
	gHeight       int32
	gMipmaps      int32
	gFormat       int32
	ref4fc2b5b    *C.Image
	allocs4fc2b5b interface{}
}
type Image struct {
	Data    unsafe.Pointer
	Width   int32
	Height  int32
	Mipmaps int32
	Format  int32
}

// Texture2D as declared in src/raylib.h:238
type gTexture2D struct {
	gId            uint32
	gWidth         int32
	gHeight        int32
	gMipmaps       int32
	gFormat        int32
	ref3c51a40b    *C.Texture2D
	allocs3c51a40b interface{}
}
type Texture2D struct {
	Id      uint32
	Width   int32
	Height  int32
	Mipmaps int32
	Format  int32
}

// CharInfo as declared in src/raylib.h:274
type gCharInfo struct {
	gValue         int32
	gOffsetX       int32
	gOffsetY       int32
	gAdvanceX      int32
	gImage         gImage
	ref702c36c0    *C.CharInfo
	allocs702c36c0 interface{}
}
type CharInfo struct {
	Value    int32
	OffsetX  int32
	OffsetY  int32
	AdvanceX int32
	Image    Image
}

// Font as declared in src/raylib.h:283
type gFont struct {
	gBaseSize      int32
	gCharsCount    int32
	gTexture       gTexture2D
	gRecs          []gRectangle
	gChars         []gCharInfo
	ref70a6a7ec    *C.Font
	allocs70a6a7ec interface{}
}
type Font struct {
	BaseSize   int32
	CharsCount int32
	Texture    Texture2D
	Recs       *Rectangle
	Chars      *CharInfo
}

// MultiText as declared in src/raylib.h:467
type gMultiText struct {
	gText          []string
	refdf1ec495    *C.MultiText
	allocsdf1ec495 interface{}
}
type MultiText struct {
	Text *string
}
