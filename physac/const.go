// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Tue, 28 Jul 2020 16:18:41 CST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package physac

/*
#include "../lib/raylib/src/physac.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// C_FOR_GO_DISABLE as defined in go-raylib/<predefine>:24
	C_FOR_GO_DISABLE = 1
	// PHYSAC_IMPLEMENTATION as defined in go-raylib/<predefine>:25
	PHYSAC_IMPLEMENTATION = 1
	// PHYSAC_NO_THREADS as defined in go-raylib/<predefine>:26
	PHYSAC_NO_THREADS = 1
	// PHYSAC_STANDALONE as defined in go-raylib/<predefine>:27
	PHYSAC_STANDALONE = 1
	// PHYSAC_STATIC as defined in go-raylib/<predefine>:28
	PHYSAC_STATIC = 1
	// C_FOR_GO as defined in go-raylib/<predefine>:29
	C_FOR_GO = 1
	// PHYSACDEF as defined in src/physac.h:78
	PHYSACDEF = 0
	// PHYSAC_MAX_BODIES as defined in src/physac.h:98
	PHYSAC_MAX_BODIES = 64
	// PHYSAC_MAX_MANIFOLDS as defined in src/physac.h:99
	PHYSAC_MAX_MANIFOLDS = 4096
	// PHYSAC_MAX_VERTICES as defined in src/physac.h:100
	PHYSAC_MAX_VERTICES = 24
	// PHYSAC_CIRCLE_VERTICES as defined in src/physac.h:101
	PHYSAC_CIRCLE_VERTICES = 24
	// PHYSAC_COLLISION_ITERATIONS as defined in src/physac.h:103
	PHYSAC_COLLISION_ITERATIONS = 100
	// PHYSAC_PENETRATION_ALLOWANCE as defined in src/physac.h:104
	PHYSAC_PENETRATION_ALLOWANCE = 0.05
	// PHYSAC_PENETRATION_CORRECTION as defined in src/physac.h:105
	PHYSAC_PENETRATION_CORRECTION = 0.4
	// PHYSAC_PI as defined in src/physac.h:107
	PHYSAC_PI = 3.14159265358979323846
	// PHYSAC_DEG2RAD as defined in src/physac.h:108
	PHYSAC_DEG2RAD = (PHYSAC_PI / 180.0)
	// PHYSAC_FLT_MAX as defined in src/physac.h:289
	PHYSAC_FLT_MAX = 3.402823466e+38
	// PHYSAC_EPSILON as defined in src/physac.h:290
	PHYSAC_EPSILON = 0.000001
	// PHYSAC_K as defined in src/physac.h:291
	PHYSAC_K = 1.0 / 3.0
	// Bool as defined in include/stdbool.h:31
	Bool = 0
	// True as defined in include/stdbool.h:32
	True = 1
	// False as defined in include/stdbool.h:33
	False = 0
)

// PhysicsShapeType as declared in src/physac.h:129
type PhysicsShapeType int32

// PhysicsShapeType enumeration from src/physac.h:129
const (
	PHYSICS_CIRCLE  PhysicsShapeType = iota
	PHYSICS_POLYGON PhysicsShapeType = 1
)
